/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package builder

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/outscale/gli/pkg/config"
	"github.com/outscale/gli/pkg/flags"
	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
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
		cfg:      config.For(provider),
	}
}

// Build builds the high-level API, must be run after BuildAPI as aliases might use API flags.
func (b *Builder[T]) Build(rootCmd *cobra.Command) {
	rootCmd.AddGroup(&cobra.Group{
		ID:    "service",
		Title: "service",
	})
	apiCmd, _ := lo.Find(rootCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "api" })
	for _, a := range b.cfg.Aliases {
		c, found := lo.Find(rootCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == a.Entity })
		if !found {
			e := b.cfg.Entities[a.Entity]
			c = &cobra.Command{
				GroupID: "service",
				Use:     a.Entity,
				Short:   a.Entity + " commands",
				Aliases: e.Aliases,
			}
			rootCmd.AddCommand(c)
		}
		cmd := &cobra.Command{
			Use:     a.Use,
			Aliases: a.Aliases,
			Short:   a.Short,
			Run:     runAlias(b.provider, a, rootCmd),
		}
		c.AddCommand(cmd)
		if apiCmd == nil || len(a.Command) < 2 {
			continue
		}
		callCmd, _ := lo.Find(apiCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == a.Command[1] })
		if callCmd == nil {
			continue
		}
		for f, fapi := range a.Flags {
			flag := callCmd.Flags().Lookup(fapi)
			if flag != nil {
				nflag := *flag
				nflag.Name = f
				cmd.Flags().AddFlag(&nflag)

				completion, found := callCmd.GetFlagCompletionFunc(fapi)
				if found {
					_ = cmd.RegisterFlagCompletionFunc(f, completion)
				}
			}
		}
	}
}

// BuildAPI builds the api command.
func (b *Builder[T]) BuildAPI(rootCmd *cobra.Command, methodFilter func(m reflect.Method) bool, run func(cmd *cobra.Command, args []string)) {
	rootCmd.AddGroup(&cobra.Group{
		ID:    "api",
		Title: "api",
	})
	var apiCmd = &cobra.Command{
		Use:     "api",
		GroupID: "api",
		Short:   rootCmd.Use + " api calls",
	}
	rootCmd.AddCommand(apiCmd)
	var client *T
	ct := reflect.TypeOf(client)
	for i := range ct.NumMethod() {
		m := ct.Method(i)
		if !methodFilter(m) {
			continue
		}
		short, help, group, _ := b.spec.SummaryForOperation(m.Name)
		if !apiCmd.ContainsGroup(group) {
			apiCmd.AddGroup(&cobra.Group{ID: group, Title: group})
		}
		cmd := &cobra.Command{
			Use:     m.Name,
			Short:   short,
			Long:    help,
			GroupID: group,
			Run:     run,
		}
		arg := m.Type.In(2)
		b.BuildFlags(cmd, arg, "")
		apiCmd.AddCommand(cmd)
	}
}

func (b *Builder[T]) BuildFlags(cmd *cobra.Command, arg reflect.Type, prefix string) {
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
				switch {
				case ot == reflect.TypeFor[iso8601.Time]() || ot == reflect.TypeFor[time.Time]():
					// TODO: add slice
					fs.Var(flags.NewTimeValue(), prefix+f.Name, help)
				case t.Elem().Implements(reflect.TypeFor[json.Marshaler]()):
					fs.StringSlice(prefix+f.Name, nil, help)
				default:
					for i := range NumEntriesInSlices {
						b.BuildFlags(cmd, t.Elem(), prefix+f.Name+"."+strconv.Itoa(i)+".")
					}
				}
			}
		case reflect.Struct:
			switch {
			case t == reflect.TypeFor[iso8601.Time]() || t == reflect.TypeFor[time.Time]():
				fs.Var(flags.NewTimeValue(), prefix+f.Name, help)
			case ot.Implements(reflect.TypeFor[json.Marshaler]()):
				fs.String(prefix+f.Name, "", help)
			default:
				b.BuildFlags(cmd, t, prefix+f.Name+".")
			}
		}
	}
}
