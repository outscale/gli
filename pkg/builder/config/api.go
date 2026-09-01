package configbuilder

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"slices"
	"strings"
	"time"

	"dario.cat/mergo"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/descriptions"
	"github.com/outscale/octl/pkg/flags"
	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
)

type APIBuilder struct {
	cfg Config

	build *config.Config

	m reflect.Method

	typeName string

	call config.APICall
}

func NewAPIBuilder(cfg Config, build *config.Config, m reflect.Method) *APIBuilder {
	typeName := cfg.getTypeNameFor(m.Name)
	call := build.API[m.Name]
	return &APIBuilder{
		cfg:      cfg,
		build:    build,
		m:        m,
		typeName: typeName,

		call: call,
	}
}

func (b *APIBuilder) Build(spec Spec) error {
	fmt.Println(" - building API call")
	if err := b.buildBase(spec); err != nil {
		return err
	}
	if err := b.buildContentEntity(); err != nil {
		return err
	}
	if err := b.buildArgsAndFlags(spec); err != nil {
		return err
	}
	return nil
}

func (b *APIBuilder) buildBase(spec Spec) error {
	if b.call.Use == "" {
		b.call.Use = b.m.Name
	}
	if b.call.Help == "" {
		b.call.Help = spec.ForCall(b.m.Name).Help
	}
	if b.call.Short == "" && !strings.HasPrefix(descriptions.Summary(b.call.Help), "alias for") {
		b.call.Short = descriptions.Summary(b.call.Help)
	}
	if b.call.Group == "" {
		b.call.Group = spec.ForCall(b.m.Name).Group
	}
	return nil
}

func (b *APIBuilder) buildContentEntity() error {
	if b.call.Content != "" && b.call.Entity != "" {
		return nil
	}
	if b.m.Type.NumOut() == 1 {
		return nil
	}
	resp := b.m.Type.Out(0)
	if resp.Kind() == reflect.Pointer {
		resp = resp.Elem()
	}
	if b.call.Content == "" {
		var found bool
		var respContent reflect.StructField
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
			found = true
			break
		}
		if found {
			b.call.Content = respContent.Name
			fmt.Println("  -> guessed content field:", b.call.Content)
		}
	}
	if b.call.Entity == "" {
		b.call.Entity = b.cfg.getEntityNameFor(b.m.Name)
		fmt.Println("  -> guessed entity name:", b.call.Entity)
	}
	return nil
}

func (b *APIBuilder) buildArgsAndFlags(spec Spec) error {
	parse := false
	for arg := range b.m.Type.Ins() {
		if arg.Implements(reflect.TypeFor[context.Context]()) {
			parse = true
			continue
		}
		if !parse {
			continue
		}
		if err := b.buildArgAndFlags(arg, spec); err != nil {
			return err
		}
	}
	return nil
}

func (b *APIBuilder) buildArgAndFlags(arg reflect.Type, spec Spec) error {
	switch arg.Kind() {
	case reflect.Struct:
		return b.buildFlags(arg, spec)
	case reflect.Pointer:
		arg = arg.Elem()
		switch arg.Kind() {
		case reflect.Struct:
			return b.buildFlags(arg, spec)
		default:
			fmt.Println("### unsupported type for command flags", arg.Kind())
		}
	case reflect.String:
		b.call.Use += " id"
	default:
		fmt.Println("### unsupported type for command flags", arg.Kind())
	}
	return nil
}

func (b *APIBuilder) buildFlags(arg reflect.Type, spec Spec) error {
	if b.call.Flags == nil {
		b.call.Flags = config.FlagSet{}
	}
	return b.buildFlagSet(&b.call.Flags, arg, "", true, spec)
}

func (b *APIBuilder) buildFlagSet(fs *config.FlagSet, arg reflect.Type, prefix string, allowRequired bool, spec Spec) error {
	if arg.Kind() == reflect.Pointer {
		arg = arg.Elem()
	}
	typeName := arg.Name()
	skipPrefixes := b.cfg.skipFlagsPrefixes(b.call.Entity)
	for f := range arg.Fields() {
		ot := f.Type
		t := ot

		required := false
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		} else {
			required = b.cfg.RequiredFromFieldPointer
		}

		var container reflect.Kind
		switch t.Kind() {
		case reflect.Slice:
			container = reflect.Slice
			ot = t.Elem()
			t = ot
			if t.Kind() == reflect.Pointer {
				t = t.Elem()
			}
			// a slice without pointer is not necessarily required
			required = false
		case reflect.Map:
			container = reflect.Map
			if t.Key().Kind() != reflect.String || t.Elem().Kind() != reflect.String {
				fmt.Println("### only map[string]string are supported", t)
				continue
			}
			ot = t.Elem()
			t = ot
			if t.Kind() == reflect.Pointer {
				t = t.Elem()
			}
			required = false
		}

		flagName := prefix + f.Name
		attrSpec := spec.ForAttribute(typeName, f.Name)
		if attrSpec.Required {
			required = attrSpec.Required
		}
		required = required && allowRequired
		help := attrSpec.Help
		if required {
			help = RequiredHelp(help)
		}
		var f config.Flag
		switch t.Kind() {
		case reflect.Bool, reflect.String, reflect.Int, reflect.Int32, reflect.Int64:
			f = config.Flag{
				Name:          flagName,
				Kind:          t.Kind(),
				Help:          help,
				ContainerKind: container,
			}
			if t.Implements(reflect.TypeFor[Enum]()) {
				f.AllowedValues = reflect.New(t).Interface().(Enum).Values()
			}
		case reflect.Interface:
			if t == reflect.TypeFor[io.Reader]() {
				f = config.Flag{
					Name:          flagName,
					Kind:          reflect.String,
					Help:          help,
					ContainerKind: container,
					CustomValue:   flags.Reader,
				}
			}
		case reflect.Struct:
			switch {
			case t == reflect.TypeFor[iso8601.Time]() || t == reflect.TypeFor[time.Time]() || t == reflect.TypeFor[openapi_types.Date]():
				f = config.Flag{
					Name:          flagName,
					Kind:          reflect.String,
					Help:          help,
					ContainerKind: container,
					CustomValue:   flags.Time,
				}
			case ot.Implements(reflect.TypeFor[json.Marshaler]()):
				f = config.Flag{
					Name:          flagName,
					Kind:          reflect.String,
					Help:          help,
					ContainerKind: container,
					CustomValue:   flags.FileOrJSON,
				}
			default:
				var err error
				if container == reflect.Slice {
					err = b.buildFlagSet(fs, t, flagName+".#.", required, spec)
				} else {
					err = b.buildFlagSet(fs, t, flagName+".", required, spec)
				}
				if err != nil {
					return err
				}
				continue
			}
		default:
			fmt.Println("### unsupported Kind", t.Kind())
		}
		if slices.ContainsFunc(skipPrefixes, func(prefix string) bool {
			return strings.HasPrefix(f.Name, prefix)
		}) {
			continue
		}
		if required {
			f.Required = &required
		}
		nf := b.cfg.FlagOverrides[f.Name]
		err := mergo.Merge(&nf, f)
		if err != nil {
			return err
		}
		fs.Add(nf)
	}
	return nil
}

const required = "[REQUIRED]"

func RequiredHelp(help string) string {
	if strings.HasPrefix(help, required) {
		return help
	}
	return required + " " + help
}

type Enum interface {
	Values() []string
}

func (b *APIBuilder) Commit() {
	b.build.API[b.m.Name] = b.call
}

var _ MethodBuilder = (*APIBuilder)(nil)
