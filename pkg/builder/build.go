/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package builder

import (
	"reflect"

	"github.com/outscale/octl/pkg/alias"
	flagbuilder "github.com/outscale/octl/pkg/builder/flags"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/descriptions"
	"github.com/outscale/octl/pkg/flags"
	"github.com/outscale/octl/pkg/markdown"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var md = markdown.NewRenderer()

type Builder[T any] struct {
	provider string
	cfg      config.Config
}

func NewBuilder[T any](provider string, helpURL string) *Builder[T] {
	return &Builder[T]{
		provider: provider,
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
		if a.SubCommand != "" {
			subc, found := lo.Find(c.Commands(), func(c *cobra.Command) bool { return c.Name() == a.SubCommand })
			if !found {
				subc = &cobra.Command{
					Use:   a.SubCommand,
					Short: a.SubCommand + " commands",
				}
				c.AddCommand(subc)
			}
			c = subc
		}
		spec := b.cfg.Spec.ForCall(a.AliasTo)
		help := "> *" + a.Short + "*\n\n" + spec.Help
		help, _ = md.Render(help)
		cmd := &cobra.Command{
			Use:     a.Use,
			Aliases: a.Aliases,
			Short:   a.Short,
			Long:    help,
			Run:     alias.RunFunc(b.provider, a, rootCmd),
		}
		c.AddCommand(cmd)
		if apiCmd == nil {
			continue
		}
		callCmd, _ := lo.Find(apiCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == a.AliasTo })
		if callCmd == nil {
			continue
		}
		for _, f := range a.Flags {
			// Flags does not return inherited flag during configuration
			flag := callCmd.Flags().Lookup(f.AliasTo)
			if flag == nil {
				flag = callCmd.InheritedFlags().Lookup(f.AliasTo)
			}
			if flag != nil {
				nflag := *flag
				nflag.Name = f.Name
				switch f.Type {
				case "base64File":
					debug.Println("overriding type for flag", f.Name, "to", f.Type)
					nflag.Value = flags.NewBase64FileValue()
				case "file":
					debug.Println("overriding type for flag", f.Name, "to", f.Type)
					nflag.Value = flags.NewFileValue()
				}
				if f.Usage != "" {
					nflag.Usage = f.Usage
				}
				cmd.Flags().AddFlag(&nflag)

				completion, found := callCmd.GetFlagCompletionFunc(f.AliasTo)
				if found {
					_ = cmd.RegisterFlagCompletionFunc(f.Name, completion)
				}
				if f.Required {
					_ = cmd.MarkFlagRequired(nflag.Name)
				}
				switch f.Type {
				case "base64File", "file":
					_ = cmd.MarkFlagFilename(f.Name)
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
	ct := reflect.TypeFor[*T]()
	for i := range ct.NumMethod() {
		m := ct.Method(i)
		if !methodFilter(m) {
			continue
		}
		spec := b.cfg.Spec.ForCall(m.Name)
		help := spec.Help
		short := descriptions.Summary(help)
		group := b.cfg.Calls[m.Name].Group
		if group != "" && !apiCmd.ContainsGroup(group) {
			apiCmd.AddGroup(&cobra.Group{ID: group, Title: group})
		}
		nhelp, err := md.Render(help)
		if err != nil {
			debug.Println("error rendering markdown", err)
		} else {
			help = nhelp
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
	fbuilder := flagbuilder.NewBuilder(b.cfg)
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
		case reflect.Int32:
			if f.Slice {
				fs.Int32Slice(f.Name, nil, f.Help)
			} else {
				fs.Int32(f.Name, 0, f.Help)
			}
		case reflect.Int64:
			if f.Slice {
				fs.Int64Slice(f.Name, nil, f.Help)
			} else {
				fs.Int64(f.Name, 0, f.Help)
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
