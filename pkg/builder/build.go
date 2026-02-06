/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/package builder

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/outscale/gli/pkg/config"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

type Builder[T any] struct {
	provider string
	spec     *Spec
	cfg      config.Config
}

func NewBuilder[T any](provider string, spec *openapi3.T) *Builder[T] {
	return &Builder[T]{
		provider: provider,
		spec:     NewSpec(spec),
		cfg:      config.For("oapi"),
	}
}

func (b *Builder[T]) Build(rootCmd *cobra.Command, methodFilter func(m reflect.Method) bool, run func(cmd *cobra.Command, args []string)) {
	var client *T
	ct := reflect.TypeOf(client)
	for i := range ct.NumMethod() {
		m := ct.Method(i)
		if !methodFilter(m) {
			continue
		}
		short, help, group, _ := b.spec.SummaryForOperation(m.Name)
		if !rootCmd.ContainsGroup(group) {
			rootCmd.AddGroup(&cobra.Group{ID: group, Title: group})
		}
		cmd := &cobra.Command{
			Use:     m.Name,
			Short:   short,
			Long:    help,
			GroupID: group,
			Run:     run,
		}
		arg := m.Type.In(2)
		b.BuildArg(cmd, arg, "")
		rootCmd.AddCommand(cmd)
	}
	for _, a := range b.cfg.Aliases {
		cmd := &cobra.Command{
			Use:     a.Use,
			Short:   a.Short,
			GroupID: a.Group,
			Run:     runAlias(b.provider, a, rootCmd),
		}
		rootCmd.AddCommand(cmd)
	}
}

func (b *Builder[T]) BuildArg(cmd *cobra.Command, arg reflect.Type, prefix string) {
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
					for i := range NumEntriesInSlices {
						b.BuildArg(cmd, t.Elem(), prefix+f.Name+"."+strconv.Itoa(i)+".")
					}
				}
			}
		case reflect.Struct:
			if ot.Implements(reflect.TypeFor[json.Marshaler]()) {
				fs.String(prefix+f.Name, "", help)
			} else {
				b.BuildArg(cmd, t, prefix+f.Name+".")
			}
		}
	}
}
