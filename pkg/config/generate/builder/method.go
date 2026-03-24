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
	call := build.Calls[m.Name]
	if call.Entity != "" {
		entity = call.Entity
	}
	return &MethodBuilder{
		cfg:          cfg,
		build:        build,
		m:            m,
		typeName:     typeName,
		typesName:    typesName,
		typeNameList: typeNameList,
		entityName:   entity,

		entity: build.Entities[entity],
		call:   call,
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
	case strings.HasPrefix(b.m.Name, "Read") || strings.HasPrefix(b.m.Name, "List") || strings.HasPrefix(b.m.Name, "Get"):
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
	if b.call.Content != "" && b.call.Entity != "" {
		return nil
	}
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
		for field := range resp.Fields() {
			if field.Anonymous || strings.HasSuffix(field.Name, "Context") || strings.HasSuffix(field.Name, "Metadata") || strings.HasSuffix(field.Name, "Pagination") {
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
	call := config.Call{}
	call.Entity = b.entityName
	if !found {
		fmt.Println("-> no response found")
		return nil
	}
	fmt.Println("guessed response", b.m.Name, respContent.Name)
	b.typeNameList = append(b.typeNameList, strings.TrimSuffix(respContent.Name, "s"))
	fmt.Println("type names", b.typeNameList)
	b.respContentType = respContentType
	call.Content = respContent.Name
	return mergo.Merge(&b.call, call)
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
		if slices.ContainsFunc(e.Columns, func(c config.Column) bool { return c.Content == "."+name }) {
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

func normalizeFlag(cfg Config, prefixes []string) func(s string) string {
	return func(s string) string {
		for _, prefix := range prefixes {
			s = strings.TrimPrefix(s, prefix)
		}
		words := lo.Words(s)
		words = lo.FilterMap(words, func(s string, i int) (string, bool) {
			if s == "0" {
				return "", false
			}
			return Singular(strings.ToLower(s)), true
		})
		s = strings.Join(words, "-")
		if len(cfg.FlagReplaces) > 0 {
			rep := strings.NewReplacer(cfg.FlagReplaces...)
			s = rep.Replace(s)
		}
		words = strings.Split(s, "-")
		if len(words) > 4 && lo.HasSuffix(words[:len(words)-2], words[len(words)-2:]) {
			fmt.Println("normalize reduce stutter", words, "=>", words[:len(words)-2])
			words = words[:len(words)-2]
		}
		if len(words) > 2 && lo.HasPrefix(words[1:], words[:1]) {
			fmt.Println("normalize reduce stutter", words, "=>", words[1:])
			words = words[1:]
		}
		return strings.Join(words, "-")
	}
}

func (b *MethodBuilder) buildFlags(t reflect.Type, prefixes []string, ignore []string) config.FlagSet {
	ignore = append(ignore, b.cfg.SkipFlagsPrefixes...)
	fs := flags.FlagSet{}
	fbs := flags.NewBuilder(*b.build, flags.WithNormalize(normalizeFlag(b.cfg, prefixes)), flags.RequiredFromPointer(b.cfg.RequiredFromFieldPointer))
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

func (b *MethodBuilder) requestIndex() int {
	reqIdx := 2
	if b.m.Type.In(0).Name() == "Context" {
		reqIdx = 1
	}
	return reqIdx
}

func (b *MethodBuilder) command(cmd ...string) []string {
	ncmd := slices.Clone(b.cfg.AliasRootPath)
	return append(ncmd, cmd...)
}

func (b *MethodBuilder) buildReadAliases() error {
	if b.entity.NoAliases {
		fmt.Println("-> no alias")
		return nil
	}

	reqIdx := b.requestIndex()
	if b.m.Type.NumIn() <= reqIdx {
		return nil
	}
	// list
	req := b.m.Type.In(reqIdx)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	var flags config.FlagSet
	if req.Kind() == reflect.Struct {
		flags = b.buildFlags(req, b.cfg.ReadFlagPrefixes, nil)
		fmt.Println("list", b.typeName, "flags", flags.Names())
		b.aliases = append(b.aliases, config.Alias{
			Entity:  b.entityName,
			Use:     "list",
			Aliases: []string{"ls"},
			AliasTo: b.m.Name,
			Short:   "alias for api " + b.m.Name,
			Command: b.command(
				"api",
				b.m.Name,
				"--output", "table",
			),
			Flags: flags,
		})
	}
	// describe
	var idFilter string
	desc := false
	switch {
	case req.Kind() != reflect.Struct:
		desc = true
	case len(flags) > 0:
		// Guess id filter
		if b.entity.Primary == "" {
			return nil
		}
		fids, found := lo.Find(flags, func(f config.Flag) bool {
			return slices.ContainsFunc(b.cfg.ReadFlagPrefixes, func(prefix string) bool { return strings.HasPrefix(f.AliasTo, prefix+b.entity.Primary) })
		})
		if found {
			idFilter = fids.AliasTo
			desc = true
		}
	}
	if !desc {
		fmt.Println("no describe", b.typeName)
		return nil
	}
	fmt.Println("describe", b.typeName, "idFilter", idFilter)
	id := lo.SnakeCase(b.entity.Primary)
	if id == "" {
		id = "id"
	}
	cmd := []string{
		"api",
		b.m.Name,
	}
	if idFilter != "" {
		idFilter = "--" + idFilter
		cmd = append(cmd, idFilter)
	}
	cmd = append(cmd, "%*",
		"--output", "yaml",
		"--single",
	)
	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     "describe " + id + " [" + id + "]...",
		Aliases: []string{"desc"},
		Short:   "alias for api " + b.m.Name + " " + idFilter + " " + id,
		Command: b.command(cmd...),
	})
	return nil
}

func (b *MethodBuilder) buildCreateAlias() error {
	if b.entity.NoAliases {
		fmt.Println("-> no alias")
		return nil
	}
	reqIdx := b.requestIndex()
	if b.m.Type.NumIn() <= reqIdx {
		return nil
	}
	req := b.m.Type.In(reqIdx)
	flags := b.buildFlags(req, b.cfg.CreateFlagPrefixes, nil)
	fmt.Println("create", b.typeName, "flags", flags.Names())
	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     "create",
		Aliases: []string{"add"},
		AliasTo: b.m.Name,
		Short:   "alias for api " + b.m.Name,
		Command: b.command(
			"api",
			b.m.Name,
			"--output", "yaml",
			"--single",
		),
		Flags: flags,
	})
	return nil
}

func (b *MethodBuilder) buildUpdateAlias(verb string) error {
	if b.entity.NoAliases {
		fmt.Println("-> no alias")
		return nil
	}
	idField, err := b.guessIDFilter()
	if err != nil {
		return err
	}
	verb = strings.ToLower(verb)
	reqIdx := b.requestIndex()
	if b.m.Type.NumIn() <= reqIdx {
		return nil
	}
	req := b.m.Type.In(reqIdx)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	if req.Kind() != reflect.Struct && b.m.Type.NumIn() >= reqIdx+2 {
		req = b.m.Type.In(reqIdx + 1)
	}
	var flags config.FlagSet
	if req.Kind() == reflect.Struct {
		flags = b.buildFlags(req, b.cfg.UpdateFlagPrefixes, []string{idField})
	}
	fmt.Println(verb, b.typeName, "idField", idField, "flags", flags.Names())
	id := lo.SnakeCase(b.entity.Primary)
	cmd := []string{
		"api",
		b.m.Name,
	}
	if idField != "" {
		idField = "--" + idField
		cmd = append(cmd, idField)
	}
	cmd = append(cmd, "%0", "--output", "yaml")

	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     verb + " " + id + " [" + id + "]...",
		AliasTo: b.m.Name,
		Short:   "alias for api " + b.m.Name + " " + idField + " " + id,
		Command: b.command(cmd...),
		Flags:   flags,
	})
	return nil
}

func (b *MethodBuilder) guessIDFilter() (string, error) {
	req := b.m.Type.In(2)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	if req.Kind() != reflect.Struct {
		return "", nil
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
		fmt.Println("-> no alias")
		return nil
	}
	idField, err := b.guessIDFilter()
	if err != nil {
		return err
	}
	reqIdx := b.requestIndex()
	if b.m.Type.NumIn() <= reqIdx {
		return nil
	}
	req := b.m.Type.In(reqIdx)
	if req.Kind() == reflect.Pointer {
		req = req.Elem()
	}
	if req.Kind() != reflect.Struct && b.m.Type.NumIn() >= reqIdx+2 {
		req = b.m.Type.In(reqIdx + 1)
	}
	var flags config.FlagSet
	if req.Kind() == reflect.Struct {
		flags = b.buildFlags(req, b.cfg.DeleteFlagPrefixes, []string{idField})
	}
	fmt.Println("delete", b.typeName, "idField", idField, "flags", flags.Names())
	var displayCmd []string
	var displayFlags config.FlagSet
	for _, a := range b.build.Aliases {
		if a.Entity == b.entityName && strings.HasPrefix(a.Use, "describe") && a.SubCommand == "" {
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
	cmd := []string{
		"api",
		b.m.Name,
	}
	if idField != "" {
		idField = "--" + idField
		cmd = append(cmd, idField)
	}
	cmd = append(cmd, "%0", "--output", "none")

	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     "delete " + id + " [" + id + "]...",
		Aliases: []string{"del", "rm"},
		AliasTo: b.m.Name,
		Short:   "alias for api " + b.m.Name + " " + idField + " " + id,
		Command: b.command(cmd...),
		Flags:   flags,
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
	newAliases := make([]config.Alias, 0, len(b.build.Aliases)+len(b.aliases))
	for _, a := range b.build.Aliases {
		// dropping existing aliases with a required flag where the new alias has no required flag
		if slices.ContainsFunc(b.aliases, func(aa config.Alias) bool {
			if a.Entity == aa.Entity && a.Use == aa.Use && a.HasRequiredFlag() && !aa.HasRequiredFlag() {
				fmt.Println("### duplicate", a, aa)
			}
			return a.Entity == aa.Entity && a.Use == aa.Use && a.HasRequiredFlag() && !aa.HasRequiredFlag()
		}) {
			continue
		}
		newAliases = append(newAliases, a)
	}
	for _, a := range b.aliases {
		if slices.ContainsFunc(newAliases, func(aa config.Alias) bool {
			if a.Entity == aa.Entity && a.Use == aa.Use {
				fmt.Println("### duplicate", a, aa)
			}
			return a.Entity == aa.Entity && a.Use == aa.Use
		}) {
			continue
		}
		newAliases = append(newAliases, a)
	}
	b.build.Aliases = newAliases
}
