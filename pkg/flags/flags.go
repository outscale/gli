/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package flags

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/outscale/gli/pkg/debug"
	"github.com/outscale/gli/pkg/openapi"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var numEntriesInSlices = 1

func init() {
	num := os.Getenv("NUM_ENTRIES")
	if num != "" {
		numEntriesInSlices, _ = strconv.Atoi(num)
	}
}

type Builder struct {
	spec *openapi.Spec
}

func NewBuilder(spec *openapi.Spec) *Builder {
	return &Builder{spec: spec}
}

func (b *Builder) FromStruct(cmd *cobra.Command, arg reflect.Type, prefix string) {
	typeName := arg.Name()
	fs := cmd.Flags()
	for i := range arg.NumField() {
		f := arg.Field(i)
		t := f.Type
		ot := t
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		help := b.spec.SummaryForAttribute(typeName, f.Name)
		switch t.Kind() {
		case reflect.Bool:
			fs.Bool(prefix+f.Name, false, help)
		case reflect.String:
			fs.String(prefix+f.Name, "", help)
			if t.Implements(reflect.TypeFor[enum]()) {
				values := reflect.New(t).Interface().(enum).Values()
				_ = cmd.RegisterFlagCompletionFunc(prefix+f.Name, func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
					return lo.Map(values, func(v string, _ int) cobra.Completion { return cobra.Completion(v) }), cobra.ShellCompDirectiveDefault
				})
			}
		case reflect.Int:
			fs.Int(prefix+f.Name, 0, help)
		case reflect.Slice:
			switch t.Elem().Kind() {
			case reflect.Bool:
				fs.BoolSlice(prefix+f.Name, nil, help)
			case reflect.String:
				fs.StringSlice(prefix+f.Name, nil, help)
				if t.Elem().Implements(reflect.TypeFor[enum]()) {
					values := reflect.New(t.Elem()).Interface().(enum).Values()
					_ = cmd.RegisterFlagCompletionFunc(prefix+f.Name, func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
						return lo.Map(values, func(v string, _ int) cobra.Completion { return cobra.Completion(v) }), cobra.ShellCompDirectiveDefault
					})
				}
			case reflect.Int:
				fs.IntSlice(prefix+f.Name, nil, help)
			case reflect.Struct:
				if t.Elem().Implements(reflect.TypeFor[json.Marshaler]()) {
					fs.StringSlice(prefix+f.Name, nil, help)
				} else {
					for i := range numEntriesInSlices {
						b.FromStruct(cmd, t.Elem(), prefix+f.Name+"."+strconv.Itoa(i)+".")
					}
				}
			}
		case reflect.Struct:
			if ot.Implements(reflect.TypeFor[json.Marshaler]()) {
				fs.String(prefix+f.Name, "", help)
			} else {
				b.FromStruct(cmd, t, prefix+f.Name+".")
			}
		}
	}
}

func ToStruct(cmd *cobra.Command, arg reflect.Value, prefix string) error {
	fs := cmd.Flags()
	debug.Println(reflect.Indirect(arg).Type().Name())
	var err error
	fs.VisitAll(func(f *pflag.Flag) {
		if f.Changed { // skipping default values
			debug.Println(f.Name, "=>", f.Value)
			if serr := set(arg, fs, f.Name, f.Name); serr != nil {
				err = serr
			}
		}
	})
	return err
}

func set(arg reflect.Value, fs *pflag.FlagSet, flag, path string) error {
	if arg.Kind() == reflect.Ptr && arg.IsNil() {
		debug.Println("allocating")
		arg.Set(reflect.New(arg.Type().Elem()))
	}
	arg = reflect.Indirect(arg)
	if path == "" {
		return setValue(arg, fs, flag)
	}
	debug.Println("set", path, "on", arg.Type().Name())
	before, after, _ := strings.Cut(path, ".")
	switch arg.Kind() {
	case reflect.Struct:
		return set(arg.FieldByName(before), fs, flag, after)
	case reflect.Slice:
		idx, err := strconv.Atoi(before)
		if err != nil {
			return fmt.Errorf("invalid array index %s in %s", before, path)
		}
		for arg.Len() < idx+1 {
			arg.Set(reflect.Append(arg, reflect.New(arg.Type().Elem()).Elem()))
		}
		return set(arg.Index(idx), fs, flag, after)
	default:
		return fmt.Errorf("invalid kind %v for %s in %s", arg.Kind(), before, flag)
	}
}

func setValue(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	debug.Println("setValue " + flag + " " + arg.Kind().String())
	switch arg.Kind() {
	case reflect.Slice:
		switch arg.Type().Elem().Kind() {
		case reflect.Bool:
			return setValueBoolSlice(arg, fs, flag)
		case reflect.Int:
			return setValueIntSlice(arg, fs, flag)
		case reflect.String:
			return setValueStringSlice(arg, fs, flag)
		case reflect.Struct:
			return setValueStructSlice(arg, fs, flag)
		default:
			debug.Println("invalid slice kind", arg.Elem().Kind())
		}
	case reflect.Struct:
		return setValueStruct(arg, fs, flag)
	case reflect.Bool:
		return setValueBool(arg, fs, flag)
	case reflect.Int:
		return setValueInt(arg, fs, flag)
	case reflect.String:
		return setValueString(arg, fs, flag)
	}
	return nil
}

func setValueBool(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	ss, err := fs.GetBool(flag)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	debug.Println("setValueBool", flag, ss)
	arg.Set(reflect.ValueOf(ss).Convert(arg.Type()))
	return nil
}

func setValueBoolSlice(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	ss, err := fs.GetBoolSlice(flag)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	debug.Println("setValueBoolSlice", flag, ss)
	narg := reflect.MakeSlice(arg.Type(), len(ss), len(ss))
	for i, s := range ss {
		narg.Index(i).Set(reflect.ValueOf(s).Convert(arg.Type().Elem()))
	}
	arg.Set(narg)
	return nil
}

func setValueInt(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	ss, err := fs.GetInt(flag)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	debug.Println("setValueInt", flag, ss)
	arg.Set(reflect.ValueOf(ss).Convert(arg.Type()))
	return nil
}

func setValueIntSlice(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	ss, err := fs.GetIntSlice(flag)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	debug.Println("setValueIntSlice", flag, ss)
	narg := reflect.MakeSlice(arg.Type(), len(ss), len(ss))
	for i, s := range ss {
		narg.Index(i).Set(reflect.ValueOf(s).Convert(arg.Type().Elem()))
	}
	arg.Set(narg)
	return nil
}

func setValueString(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	ss, err := fs.GetString(flag)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	debug.Println("setValueString", flag, ss)
	arg.Set(reflect.ValueOf(ss).Convert(arg.Type()))
	return nil
}

func setValueStringSlice(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	ss, err := fs.GetStringSlice(flag)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	debug.Println("setValueStringSlice", flag, ss)
	narg := reflect.MakeSlice(arg.Type(), len(ss), len(ss))
	for i, s := range ss {
		narg.Index(i).Set(reflect.ValueOf(s).Convert(arg.Type().Elem()))
	}
	arg.Set(narg)
	return nil
}

func setValueStruct(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	if !arg.Addr().Type().Implements(reflect.TypeFor[json.Unmarshaler]()) {
		return nil
	}
	ss, err := fs.GetString(flag)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	debug.Println("setValueStruct/JSON", flag, ss)
	v := reflect.New(arg.Type()).Interface()
	err = json.Unmarshal([]byte(`"`+ss+`"`), v)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	arg.Set(reflect.ValueOf(v).Elem())
	return nil
}

func setValueStructSlice(arg reflect.Value, fs *pflag.FlagSet, flag string) error {
	if !reflect.PointerTo(arg.Type().Elem()).Implements(reflect.TypeFor[json.Unmarshaler]()) {
		return nil
	}
	ss, err := fs.GetStringSlice(flag)
	if err != nil {
		return fmt.Errorf("invalid %s value: %w", flag, err)
	}
	debug.Println("setValueStructSlice/JSON", flag, ss)
	narg := reflect.MakeSlice(arg.Type(), len(ss), len(ss))
	for i, s := range ss {
		v := reflect.New(arg.Type().Elem()).Interface()
		err = json.Unmarshal([]byte(`"`+s+`"`), v)
		if err != nil {
			return fmt.Errorf("invalid %s value: %w", flag, err)
		}
		narg.Index(i).Set(reflect.ValueOf(v).Elem())
	}
	arg.Set(narg)
	return nil
}
