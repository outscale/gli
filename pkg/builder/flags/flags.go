package flags

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/outscale/octl/pkg/builder/openapi"
	"github.com/outscale/octl/pkg/flags"
	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
	"github.com/spf13/pflag"
)

type FlagSet []Flag

func (fs *FlagSet) Add(f Flag) {
	*fs = append(*fs, f)
}

type Normalize func(string) string

type Flag struct {
	Name          string
	FieldPath     string
	Kind          reflect.Kind
	Slice         bool
	Help          string
	Required      bool
	AllowedValues []string
	FlagValue     pflag.Value
}

type Builder struct {
	Normalize Normalize
	spec      *openapi.Spec
}

type Option func(*Builder)

func WithNormalize(fn Normalize) Option {
	return func(b *Builder) {
		b.Normalize = fn
	}
}

func NewBuilder(spec *openapi.Spec, opts ...Option) *Builder {
	b := &Builder{
		Normalize: func(s string) string { return s },
		spec:      spec,
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func (b *Builder) Build(fs *FlagSet, arg reflect.Type, prefix string, allowRequired bool) {
	typeName := arg.Name()
	for i := range arg.NumField() {
		f := arg.Field(i)
		ot := f.Type
		t := ot

		slice := false
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() == reflect.Slice {
			slice = true
			ot = t.Elem()
			t = ot
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
		}

		help, required := b.spec.SummaryForAttribute(typeName, f.Name)
		flagName := prefix + f.Name
		switch t.Kind() {
		case reflect.Bool, reflect.String, reflect.Int:
			f := Flag{
				Name:      b.Normalize(flagName),
				FieldPath: flagName,
				Kind:      t.Kind(),
				Help:      help,
				Required:  required,
				Slice:     slice,
			}
			if t.Implements(reflect.TypeFor[openapi.Enum]()) {
				f.AllowedValues = reflect.New(t).Interface().(openapi.Enum).Values()
			}
			fs.Add(f)
		case reflect.Struct:
			switch {
			case t == reflect.TypeFor[iso8601.Time]() || t == reflect.TypeFor[time.Time]():
				f := Flag{
					Name:      b.Normalize(flagName),
					FieldPath: flagName,
					Kind:      reflect.String,
					Help:      help,
					Required:  required,
					Slice:     slice,
					FlagValue: flags.NewTimeValue(),
				}
				fs.Add(f)
			case ot.Implements(reflect.TypeFor[json.Marshaler]()):
				f := Flag{
					Name:      b.Normalize(flagName),
					FieldPath: flagName,
					Kind:      reflect.String,
					Help:      help,
					Required:  required,
					Slice:     slice,
				}
				fs.Add(f)
			default:
				if slice {
					for i := range NumEntriesInSlices {
						b.Build(fs, t, flagName+"."+strconv.Itoa(i)+".", required && allowRequired)
					}
				} else {
					b.Build(fs, t, flagName+".", required && allowRequired)
				}
			}
		}
	}
}
