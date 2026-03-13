package builder

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"dario.cat/mergo"
	"github.com/outscale/octl/pkg/builder/flags"
	"github.com/outscale/octl/pkg/config"
	"github.com/samber/lo"
)

var ErrCantBuild = errors.New("cannot build")

type MethodBuilder struct {
	cfg Config

	build *config.Config

	m reflect.Method

	typeName, typesName string
	typeNameList        []string
	entityName          string

	respContentType reflect.Type

	entity  config.Entity
	call    config.Call
	aliases []config.Alias
}

func NewMethodBuilder(cfg Config, build *config.Config, m reflect.Method) *MethodBuilder {
	typeName, typesName, entity := names(m.Name)
	typeNameList := []string{typeName}
	return &MethodBuilder{
		cfg:          cfg,
		build:        build,
		m:            m,
		typeName:     typeName,
		typesName:    typesName,
		typeNameList: typeNameList,
		entityName:   entity,

		entity: build.Entities[entity],
		call:   build.Calls[m.Name],
	}
}

func (b *MethodBuilder) Build() error {
	if b.entity.Skip {
		fmt.Println("***", b.m.Name, "(skipped)")
		return nil
	}
	fmt.Println("***", b.m.Name)
	var err error
	switch {
	case strings.HasPrefix(b.m.Name, "Read") || strings.HasPrefix(b.m.Name, "List"):
		err = b.buildCall()
		if err == nil {
			err = b.buildEntity()
		}
		if err == nil {
			err = b.buildReadAliases()
		}
	case strings.HasPrefix(b.m.Name, "Delete"):
		err = b.buildDeleteAlias()
	case strings.HasPrefix(b.m.Name, "Update"):
		err = b.buildCall()
		if err == nil {
			err = b.buildUpdateAlias("Update")
		}
	case strings.HasPrefix(b.m.Name, "Put"):
		err = b.buildCall()
		if err == nil {
			err = b.buildUpdateAlias("Put")
		}
	case strings.HasPrefix(b.m.Name, "Create"):
		err = b.buildCall()
		if err == nil {
			err = b.buildCreateAlias()
		}
	default:
		return ErrCantBuild
	}
	return err
}

func (b *MethodBuilder) buildCall() error {
	resp := b.m.Type.Out(0)
	if resp.Kind() == reflect.Pointer {
		resp = resp.Elem()
	}
	var found bool
	var respContent reflect.StructField
	var respContentType reflect.Type
	if b.call.Content != "" {
		respContent, found = resp.FieldByName(b.call.Content)
		respContentType = respContent.Type
	} else {
		for i := range resp.NumField() {
			field := resp.Field(i)
			if field.Anonymous || strings.HasSuffix(field.Name, "Context") || strings.HasSuffix(field.Name, "Metadata") {
				continue
			}
			t := field.Type
			respContent = field
			if t.Kind() == reflect.Pointer {
				t = t.Elem()
			}
			switch t.Kind() {
			case reflect.Struct:
			case reflect.Slice:
				t = t.Elem()
				if t.Kind() != reflect.Struct {
					continue
				}
			default:
				continue
			}
			respContentType = t
			found = true
			break
		}
	}
	b.call.Entity = b.entityName
	if !found {
		fmt.Println("*** no response found")
		return nil
	}
	fmt.Println("guessed response", b.m.Name, respContent.Name)
	b.typeNameList = append(b.typeNameList, strings.TrimSuffix(respContent.Name, "s"))
	fmt.Println("type names", b.typeNameList)
	b.respContentType = respContentType
	b.call.Content = respContent.Name
	return nil
}

func (b *MethodBuilder) buildEntity() error {
	if b.respContentType == nil || b.respContentType.Kind() != reflect.Struct {
		return nil
	}
	e := config.Entity{}
	hasPrimary := false
	for _, typeName := range b.typeNameList {
		if f, found := b.respContentType.FieldByName(typeName + "Id"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "ID",
				Content: "." + f.Name,
			})
			e.Primary = f.Name
			hasPrimary = true
			break
		}
	}
	if f, found := b.respContentType.FieldByName(b.typeName + "Name"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Name",
			Content: "." + f.Name,
		})
		if !hasPrimary {
			e.Primary = f.Name
			hasPrimary = true //nolint
		}
	} else if _, found := b.respContentType.FieldByName("Tags"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Name",
			Content: `.Tags[] | select(.Key == "Name").Value`,
		})
	}
	if f, found := b.respContentType.FieldByName(b.typeName + "Type"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Type",
			Content: "." + f.Name,
		})
	}
	for _, name := range b.cfg.PriorityFields {
		if slices.ContainsFunc(e.Columns, func(c config.Column) bool { return c.Content == name }) {
			continue
		}
		if f, found := b.respContentType.FieldByName(name); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   f.Name,
				Content: "." + f.Name,
			})
		}
	}
	ecfg := b.entity
	err := mergo.Merge(&ecfg, e)
	if err != nil {
		return err
	}
	b.entity = ecfg
	return nil
}

func normalizeFlag(cfg Config, prefix string) func(s string) string {
	return func(s string) string {
		items := lo.Words(strings.TrimPrefix(s, prefix))
		items = lo.FilterMap(items, func(s string, i int) (string, bool) {
			if s == "0" {
				return "", false
			}
			return Singular(strings.ToLower(s)), true
		})
		s = strings.Join(items, "-")
		if len(cfg.FlagReplaces) > 0 {
			rep := strings.NewReplacer(cfg.FlagReplaces...)
			s = rep.Replace(s)
		}
		items = strings.Split(s, "-")
		if len(items) > 4 && lo.HasSuffix(items[:len(items)-2], items[len(items)-2:]) {
			fmt.Println("normalize reduce stutter", items, "=>", items[:len(items)-2])
			items = items[:len(items)-2]
		}
		if len(items) > 2 && lo.HasPrefix(items[1:], items[:1]) {
			fmt.Println("normalize reduce stutter", items, "=>", items[1:])
			items = items[1:]
		}
		return strings.Join(items, "-")
	}
}

func (b *MethodBuilder) buildFlags(t reflect.Type, prefix string, ignore []string) config.FlagSet {
	ignore = append(ignore, b.cfg.SkipFlags...)
	fs := flags.FlagSet{}
	fbs := flags.NewBuilder(*b.build, flags.WithNormalize(normalizeFlag(b.cfg, prefix)), flags.RequiredFromPointer(b.cfg.RequiredFromFieldPointer))
	fbs.Build(&fs, t, "", true)

	cfs := config.FlagSet{}
	for _, f := range fs {
		if lo.ContainsBy(ignore, func(ignore string) bool { return strings.HasPrefix(f.FieldPath, ignore) }) {
			continue
		}
		flag := f.Name
		stripped := flag
		for _, name := range b.typeNameList {
			nstripped := strings.TrimPrefix(flag, lo.KebabCase(Singular(name))+"-")
			if len(nstripped) < len(stripped) {
				stripped = nstripped
			}
		}
		if _, found := cfs.Get(stripped); !found {
			flag = stripped
		}
		cf := b.cfg.FlagOverrides[f.Name]
		err := mergo.Merge(&cf, config.Flag{
			Name:     flag,
			AliasTo:  f.FieldPath,
			Required: f.Required,
		})
		if err != nil {
			panic(err)
		}
		cfs = append(cfs, cf)
	}
	return cfs
}

func (b *MethodBuilder) buildReadAliases() error {
	if b.entity.NoAliases {
		return nil
	}
	// list
	req := b.m.Type.In(2)
	flags := b.buildFlags(req, b.cfg.ReadFlagPrefix, nil)
	fmt.Println("list", b.typeName, "flags", flags.Names())
	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     "list",
		Aliases: []string{"ls"},
		AliasTo: b.m.Name,
		Short:   "alias for api " + b.m.Name,
		Command: []string{
			"api",
			b.m.Name,
			"--output", "table",
		},
		Flags: flags,
	})

	if b.entity.Primary == "" {
		return nil
	}
	// describe
	// Guess id filter
	fids, found := lo.Find(flags, func(f config.Flag) bool {
		return strings.HasPrefix(f.AliasTo, b.cfg.ReadFlagPrefix+b.entity.Primary)
	})
	if !found {
		return nil
	}
	if found {
		fmt.Println("describe", b.typeName, fids.AliasTo)
		id := lo.SnakeCase(b.entity.Primary)
		b.aliases = append(b.aliases, config.Alias{
			Entity:  b.entityName,
			Use:     "describe " + id + " [" + id + "]...",
			Aliases: []string{"desc"},
			Short:   "alias for api " + b.m.Name + " --" + fids.AliasTo + " " + id,
			Command: []string{
				"api",
				b.m.Name,
				"--" + fids.AliasTo, "%*",
				"--output", "yaml",
				"--single",
			},
		})
	}
	return nil
}

func (b *MethodBuilder) buildCreateAlias() error {
	if b.entity.NoAliases {
		return nil
	}
	req := b.m.Type.In(2)
	flags := b.buildFlags(req, "", nil)
	fmt.Println("create", b.typeName, "flags", flags.Names())
	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     "create",
		Aliases: []string{"add"},
		AliasTo: b.m.Name,
		Short:   "alias for api " + b.m.Name,
		Command: []string{
			"api",
			b.m.Name,
			"--output", "yaml",
			"--single",
		},
		Flags: flags,
	})
	return nil
}

func (b *MethodBuilder) buildUpdateAlias(verb string) error {
	if b.entity.NoAliases {
		return nil
	}
	idField, err := b.guessIDFilter()
	if err != nil {
		return err
	}
	verb = strings.ToLower(verb)
	req := b.m.Type.In(2)
	flags := b.buildFlags(req, "", []string{idField})
	fmt.Println(verb, b.typeName, idField, "flags", flags.Names())
	id := lo.SnakeCase(b.entity.Primary)
	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     verb + " " + id + " [" + id + "]...",
		AliasTo: b.m.Name,
		Short:   "alias for api " + b.m.Name + " --" + idField + " " + id,
		Command: []string{
			"api",
			b.m.Name,
			"--" + idField, "%0",
			"--output", "yaml",
		},
		Flags: flags,
	})
	return nil
}

func (b *MethodBuilder) guessIDFilter() (string, error) {
	req := b.m.Type.In(2)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	primary := b.build.Entities[b.entityName].Primary
	fids, found := req.FieldByName(primary)
	if !found {
		fids, found = req.FieldByName(primary + "s")
	}
	if !found {
		return "", ErrCantBuild
	}
	return fids.Name, nil
}

func (b *MethodBuilder) buildDeleteAlias() error {
	if b.entity.NoAliases {
		return nil
	}
	idField, err := b.guessIDFilter()
	if err != nil {
		return err
	}
	req := b.m.Type.In(2)
	flags := b.buildFlags(req, "", []string{idField})
	fmt.Println("delete", b.typeName, idField, "flags", flags.Names())
	var displayCmd []string
	var displayFlags config.FlagSet
	for _, a := range b.build.Aliases {
		if a.Entity == b.entityName && strings.HasPrefix(a.Use, "describe") {
			displayCmd = lo.Map(a.Command, func(arg string, i int) string {
				if i > 0 && (a.Command[i-1] == "-o" || a.Command[i-1] == "--output") {
					return "table"
				}
				return arg
			})
			displayFlags = lo.Filter(a.Flags, func(f config.Flag, _ int) bool {
				_, ok := flags.Get(f.Name)
				return ok
			})
		}
	}
	id := lo.SnakeCase(b.entity.Primary)
	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     "delete " + id + " [" + id + "]...",
		Aliases: []string{"del", "rm"},
		AliasTo: b.m.Name,
		Short:   "alias for api " + b.m.Name + " --" + idField + " " + id,
		Command: []string{
			"api",
			b.m.Name,
			"--" + idField, "%0",
			"--output", "none",
		},
		Flags: flags,
		Prompt: &config.Prompt{
			Action:         config.ActionDelete,
			DisplayCommand: displayCmd,
			Flags:          displayFlags,
		},
	})
	return nil
}

func (b *MethodBuilder) Commit() {
	b.build.Calls[b.m.Name] = b.call
	b.build.Entities[b.entityName] = b.entity
	b.build.Aliases = append(b.build.Aliases, b.aliases...)
}
