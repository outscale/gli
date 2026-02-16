/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package builder

import (
	"reflect"

	"github.com/getkin/kin-openapi/openapi3"
	flagbuilder "github.com/outscale/octl/pkg/builder/flags"
	"github.com/outscale/octl/pkg/builder/openapi"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

type Builder[T any] struct {
	provider string
	spec     *openapi.Spec
	cfg      config.Config
}

func NewBuilder[T any](provider string, spec *openapi3.T) *Builder[T] {
	return &Builder[T]{
		provider: provider,
		spec:     openapi.NewSpec(spec),
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
		debug.Println(a.Entity, a.Use)
		for _, f := range a.Flags {
			flag := callCmd.Flags().Lookup(f.AliasTo)
			if flag != nil {
				nflag := *flag
				nflag.Name = f.Name
				cmd.Flags().AddFlag(&nflag)

				completion, found := callCmd.GetFlagCompletionFunc(f.AliasTo)
				if found {
					_ = cmd.RegisterFlagCompletionFunc(f.Name, completion)
				}
				if f.Required {
					_ = cmd.MarkFlagRequired(nflag.Name)
				}
			}
		}
	}
}

// BuildAPI builds the api command.
func (b *Builder[T]) BuildAPI(
	rootCmd *cobra.Command,
	methodFilter func(m reflect.Method) bool,
	run func(cmd *cobra.Command, args []string),
) {
	rootCmd.AddGroup(&cobra.Group{
		ID:    "api",
		Title: "api",
	})
	apiCmd := &cobra.Command{
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

		for j := 2; j < m.Type.NumIn()-1; j++ {
			arg := m.Type.In(j)
			b.BuildArgsAndFlags(cmd, arg)
		}

		apiCmd.AddCommand(cmd)
	}
}

func (b *Builder[T]) BuildArgsAndFlags(cmd *cobra.Command, arg reflect.Type) {
	switch arg.Kind() {
	case reflect.Struct:
		b.buildFlags(cmd, arg)
	case reflect.Pointer:
		arg = arg.Elem()
		switch arg.Kind() {
		case reflect.Struct:
			b.buildFlags(cmd, arg)
		default:
			debug.Println("unsupported type for command flags: *%v", arg.Kind())
		}
	case reflect.String:
		cmd.Use += " id"
	default:
		debug.Println("unsupported type for command flags: %v", arg.Kind())
	}
}

func (b *Builder[T]) buildFlags(cmd *cobra.Command, arg reflect.Type) {
	fbuilder := flagbuilder.NewBuilder(b.spec)
	fbs := flagbuilder.FlagSet{}
	fbuilder.Build(&fbs, arg, "", true)
	fs := cmd.Flags()
	for _, f := range fbs {
		switch f.Kind {
		case reflect.Bool:
			if f.Slice {
				fs.BoolSlice(f.Name, nil, f.Help)
			} else {
				fs.Bool(f.Name, false, f.Help)
			}
		case reflect.Int:
			if f.Slice {
				fs.IntSlice(f.Name, nil, f.Help)
			} else {
				fs.Int(f.Name, 0, f.Help)
			}
		case reflect.String:
			switch {
			case f.FlagValue != nil:
				fs.Var(f.FlagValue, f.Name, f.Help)
			case f.Slice:
				fs.StringSlice(f.Name, nil, f.Help)
			default:
				fs.String(f.Name, "", f.Help)
			}
			if len(f.AllowedValues) > 0 {
				_ = cmd.RegisterFlagCompletionFunc(f.Name, func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
					return lo.Map(f.AllowedValues, func(v string, _ int) cobra.Completion { return cobra.Completion(v) }), cobra.ShellCompDirectiveDefault
				})
			}
		}
	}
}
