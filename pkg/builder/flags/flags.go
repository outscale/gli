package flags

import (
	"encoding/json"
	"io"
	"reflect"
	"strconv"
	"time"

	"github.com/outscale/octl/pkg/config"
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
	cfg                 config.Config
	normalize           Normalize
	requiredFromPointer bool
}

type Option func(*Builder)

func WithNormalize(fn Normalize) Option {
	return func(b *Builder) {
		b.normalize = fn
	}
}

func RequiredFromPointer(required bool) Option {
	return func(b *Builder) {
		b.requiredFromPointer = required
	}
}

func NewBuilder(cfg config.Config, opts ...Option) *Builder {
	b := &Builder{
		cfg:       cfg,
		normalize: func(s string) string { return s },
	}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func (b *Builder) Build(fs *FlagSet, arg reflect.Type, prefix string, allowRequired bool) {
	if arg.Kind() == reflect.Pointer {
		arg = arg.Elem()
	}
	typeName := arg.Name()
	for i := range arg.NumField() {
		f := arg.Field(i)
		ot := f.Type
		t := ot

		slice := false
		required := false
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		} else {
			required = b.requiredFromPointer
		}
		if t.Kind() == reflect.Slice {
			slice = true
			ot = t.Elem()
			t = ot
			if t.Kind() == reflect.Pointer {
				t = t.Elem()
			}
			// a slice without pointer is not necessarily required
			required = false
		}

		flagName := prefix + f.Name
		required = required && allowRequired
		spec := b.cfg.Spec.ForAttribute(typeName, f.Name)
		if spec.Required {
			required = spec.Required
		}
		switch t.Kind() {
		case reflect.Bool, reflect.String, reflect.Int, reflect.Int32:
			f := Flag{
				Name:      b.normalize(flagName),
				FieldPath: flagName,
				Kind:      t.Kind(),
				Help:      spec.Help,
				Required:  required,
				Slice:     slice,
			}
			if t.Implements(reflect.TypeFor[Enum]()) {
				f.AllowedValues = reflect.New(t).Interface().(Enum).Values()
			}
			fs.Add(f)
		case reflect.Interface:
			if t == reflect.TypeFor[io.Reader]() {
				f := Flag{
					Name:      b.normalize(flagName),
					FieldPath: flagName,
					Kind:      reflect.String,
					Help:      spec.Help,
					Required:  required,
					Slice:     slice,
					FlagValue: flags.NewReaderValue(),
				}
				fs.Add(f)
			}
		case reflect.Struct:
			switch {
			case t == reflect.TypeFor[iso8601.Time]() || t == reflect.TypeFor[time.Time]():
				f := Flag{
					Name:      b.normalize(flagName),
					FieldPath: flagName,
					Kind:      reflect.String,
					Help:      spec.Help,
					Required:  required,
					Slice:     slice,
					FlagValue: flags.NewTimeValue(),
				}
				fs.Add(f)
			case ot.Implements(reflect.TypeFor[json.Marshaler]()):
				f := Flag{
					Name:      b.normalize(flagName),
					FieldPath: flagName,
					Kind:      reflect.String,
					Help:      spec.Help,
					Required:  required,
					Slice:     slice,
				}
				fs.Add(f)
			default:
				if slice {
					for i := range NumEntriesInSlices {
						b.Build(fs, t, flagName+"."+strconv.Itoa(i)+".", required)
					}
				} else {
					b.Build(fs, t, flagName+".", required)
				}
			}
		}
	}
}
