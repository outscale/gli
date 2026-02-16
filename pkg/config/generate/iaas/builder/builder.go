package builder

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"dario.cat/mergo"
	"github.com/iancoleman/strcase"
	"github.com/outscale/octl/pkg/builder/flags"
	"github.com/outscale/octl/pkg/builder/openapi"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/samber/lo"
)

var priorityFields = []string{
	"State",
	"PublicIp",
	"PrivateIp",
	"NetId",
	"SubnetId",
	"IpRange",
	"SubregionName",
	"SubregionNames",
	"Subregion",
	"Subregions",
	"Size",
	"Iops",
}

var ErrCantBuild = errors.New("cannot build")

type Builder struct {
	cfg  *config.Config
	spec *openapi.Spec

	m reflect.Method

	typeName, typesName string
	typeNameList        []string
	entityName          string

	respContentType reflect.Type

	entity  config.Entity
	call    *config.Call
	aliases []config.Alias
}

func New(cfg *config.Config, m reflect.Method, prefix string) *Builder {
	typesName := strings.TrimPrefix(m.Name, prefix)
	typeName := strings.TrimSuffix(typesName, "s")
	if strings.HasSuffix(typesName, "ies") {
		typeName = strings.TrimSuffix(typesName, "ies") + "y"
	}
	entity := strings.ToLower(typeName)
	typeNameList := []string{typeName}
	spec, err := osc.GetSwagger()
	if err != nil {
		panic(err)
	}
	return &Builder{
		cfg:          cfg,
		spec:         openapi.NewSpec(spec),
		m:            m,
		typeName:     typeName,
		typesName:    typesName,
		typeNameList: typeNameList,
		entityName:   entity,

		entity: cfg.Entities[entity],
	}
}

func (b *Builder) Build() error {
	fmt.Println(b.m.Name)
	var err error
	switch {
	case strings.HasPrefix(b.m.Name, "Read"):
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
		err = b.buildUpdateAlias()
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

func (b *Builder) buildCall() error {
	resp := b.m.Type.Out(0).Elem()
	respContent, found := resp.FieldByName(b.typesName)
	if !found {
		for i := range resp.NumField() {
			field := resp.Field(i)
			if field.Name == "NextPageToken" || field.Name == "ResponseContext" {
				continue
			}
			respContent = field
			found = true
			break
		}
		if !found {
			fmt.Println("*** no response found")
			return ErrCantBuild
		}
		fmt.Println("guessed response", b.m.Name, respContent.Name)
		b.typeNameList = append(b.typeNameList, strings.TrimSuffix(respContent.Name, "s"))
		fmt.Println("type names", b.typeNameList)
	}
	respContentType := respContent.Type
	if respContentType.Kind() == reflect.Pointer {
		respContentType = respContentType.Elem()
	}
	if respContentType.Kind() == reflect.Slice {
		respContentType = respContentType.Elem()
	}
	if respContentType.Kind() != reflect.Struct {
		fmt.Println("*** invalid response type", respContentType.Kind())
		return ErrCantBuild
	}
	b.respContentType = respContentType
	b.call = &config.Call{
		Entity:  b.entityName,
		Content: respContent.Name,
	}
	return nil
}

func (b *Builder) buildEntity() error {
	e := config.Entity{}
	for _, typeName := range b.typeNameList {
		if f, found := b.respContentType.FieldByName(typeName + "Id"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "ID",
				Content: f.Name,
			})
			break
		}
	}
	if f, found := b.respContentType.FieldByName(b.typeName + "Name"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Name",
			Content: f.Name,
		})
	} else if _, found := b.respContentType.FieldByName("Tags"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Name",
			Content: `find(Tags, #?.Key == "Name")?.Value`,
		})
	}
	if f, found := b.respContentType.FieldByName(b.typeName + "Type"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Type",
			Content: f.Name,
		})
	}
	for _, name := range priorityFields {
		if slices.ContainsFunc(e.Columns, func(c config.Column) bool { return c.Content == name }) {
			continue
		}
		if f, found := b.respContentType.FieldByName(name); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   f.Name,
				Content: f.Name,
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

func (b *Builder) buildFlags(t reflect.Type, prefix string, ignore []string) config.FlagSet {
	ignore = append(ignore, "DryRun", "NextPageToken", "ResultsPerPage")
	fs := flags.FlagSet{}
	norm := func(s string) string {
		return strcase.ToKebab(strings.ReplaceAll(s, "s.0", ""))
	}
	fbs := flags.NewBuilder(b.spec, flags.WithNormalize(norm))
	fbs.Build(&fs, t, "", true)

	cfs := config.FlagSet{}
	for _, f := range fs {
		if slices.Contains(ignore, f.FieldPath) {
			continue
		}
		flag := strings.TrimPrefix(f.Name, prefix)
		stripped := strings.TrimPrefix(flag, b.entityName+"-")
		if _, found := cfs.Get(stripped); !found {
			flag = stripped
		}
		cfs = append(cfs, config.Flag{
			Name:     flag,
			AliasTo:  f.FieldPath,
			Required: f.Required,
		})
	}
	return cfs
}

func (b *Builder) buildReadAliases() error {
	// list
	req := b.m.Type.In(2)
	flags := b.buildFlags(req, "filters-", nil)
	fmt.Println("list", b.typeName, "filters", flags.Names())
	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     "list",
		Aliases: []string{"ls"},
		Group:   b.entityName,
		Short:   "alias for api " + b.m.Name,
		Command: []string{
			"api",
			b.m.Name,
			"--output", "table",
		},
		Flags: flags,
	})

	// describe
	filters, found := req.FieldByName("Filters")
	if found {
		fType := filters.Type.Elem()
		// Guess id filter
		fids, found := fType.FieldByName(b.typeName + "Ids")
		if !found {
			fids, found = fType.FieldByName(b.typeName + "Names")
		}
		if !found {
			for i := range fType.NumField() {
				field := fType.Field(i)
				if strings.HasSuffix(field.Name, "Ids") && strings.HasSuffix(b.typeName, strings.TrimSuffix(field.Name, "Ids")) {
					fids, found = field, true
					b.typeNameList = append(b.typeNameList, strings.TrimSuffix(field.Name, "Ids"))
					fmt.Println("type names", b.typeNameList)
				}
			}
		}
		if found {
			fmt.Println("describe", b.typeName, "Filters."+fids.Name)
			b.aliases = append(b.aliases, config.Alias{
				Entity:  b.entityName,
				Use:     "describe " + b.entityName + "_id",
				Aliases: []string{"desc"},
				Group:   b.entityName,
				Short:   "alias for api " + b.m.Name + " --Filters." + fids.Name + " " + b.entityName + "_id",
				Command: []string{
					"api",
					b.m.Name,
					"--Filters." + fids.Name, "%0",
					"--output", "yaml,single",
				},
			})
		}
	}
	return nil
}

func (b *Builder) buildCreateAlias() error {
	req := b.m.Type.In(2)
	flags := b.buildFlags(req, "", nil)
	fmt.Println("create", b.typeName, "flags", flags.Names())
	b.aliases = append(b.aliases, config.Alias{
		Entity: b.entityName,
		Use:    "create",
		Group:  b.entityName,
		Short:  "alias for api " + b.m.Name,
		Command: []string{
			"api",
			b.m.Name,
			"--output", "yaml",
		},
		Flags: flags,
	})
	return nil
}

func (b *Builder) buildUpdateAlias() error {
	idField, err := b.guessIDFilter()
	if err != nil {
		return err
	}
	req := b.m.Type.In(2)
	flags := b.buildFlags(req, "", []string{idField})
	fmt.Println("update", b.typeName, idField, "flags", flags.Names())
	b.aliases = append(b.aliases, config.Alias{
		Entity: b.entityName,
		Use:    "update " + b.entityName + "_id",
		Group:  b.entityName,
		Short:  "alias for api " + b.m.Name + " --" + idField + " " + b.entityName + "_id",
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

func (b *Builder) guessIDFilter() (string, error) {
	req := b.m.Type.In(2)
	fids, found := req.FieldByName(b.typeName + "Id")
	if !found {
		fids, found = req.FieldByName(b.typeName + "Ids")
	}
	if !found {
		fids, found = req.FieldByName(b.typeName + "Name")
	}
	if !found {
		fids, found = req.FieldByName(b.typeName + "Names")
	}
	if !found {
		return "", ErrCantBuild
	}
	return fids.Name, nil
}

func (b *Builder) buildDeleteAlias() error {
	idField, err := b.guessIDFilter()
	if err != nil {
		return err
	}
	// Guess id filter
	fmt.Println("delete", b.typeName, idField)
	var displayCmd []string
	for _, a := range b.cfg.Aliases {
		if a.Entity == b.entityName && strings.HasPrefix(a.Use, "describe") {
			displayCmd = lo.Map(a.Command, func(arg string, _ int) string {
				if arg == "yaml,single" {
					return "table,single"
				}
				return arg
			})
		}
	}
	b.aliases = append(b.aliases, config.Alias{
		Entity:  b.entityName,
		Use:     "delete " + b.entityName + "_id",
		Aliases: []string{"del", "rm"},
		Group:   b.entityName,
		Short:   "alias for api " + b.m.Name + " --" + idField + " " + b.entityName + "_id",
		Command: []string{
			"api",
			b.m.Name,
			"--" + idField, "%0",
			"--output", "none",
		},
		Prompt: &config.Prompt{
			Action:         config.ActionDelete,
			DisplayCommand: displayCmd,
		},
	})
	return nil
}

func (b *Builder) Commit() {
	if b.call != nil {
		b.cfg.Calls[b.m.Name] = *b.call
	}
	b.cfg.Entities[b.entityName] = b.entity
	b.cfg.Aliases = append(b.cfg.Aliases, b.aliases...)
}
