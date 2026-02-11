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
	"github.com/outscale/gli/pkg/debug"
	"github.com/outscale/gli/pkg/flags"
	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

const requiredAnnotation = "gli/required"

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
				if _, found := nflag.Annotations[requiredAnnotation]; found {
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

		for j := 2; j < m.Type.NumIn()-1; j++ {
			arg := m.Type.In(j)
			b.BuildFlags(cmd, arg)
		}

		apiCmd.AddCommand(cmd)
	}
}

func (b *Builder[T]) BuildFlags(
	cmd *cobra.Command,
	arg reflect.Type,
) {
	switch arg.Kind() {
	case reflect.Struct:
		
		b.buildFlagsFromStruct(cmd, arg, "", true)
	case reflect.Pointer:
		b.BuildFlags(cmd, arg.Elem())
	case reflect.String:
		// projectId, ClusterId
		cmd.Flags().String("Id", "", "Id of the object")
	default:
		debug.Println("unsupported type for command flags: %v", arg.Kind())
	}
}

func (b *Builder[T]) buildFlagsFromStruct(cmd *cobra.Command, arg reflect.Type, prefix string, allowRequired bool) {
	typeName := arg.Name()
	fs := cmd.Flags()
	for i := range arg.NumField() {
		f := arg.Field(i)
		t := f.Type
		ot := t
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		help, required := b.spec.SummaryForAttribute(typeName, f.Name)
		flagName := prefix + f.Name
		switch t.Kind() {
		case reflect.Bool:
			fs.Bool(flagName, false, help)
		case reflect.String:
			fs.String(flagName, "", help)
			if t.Implements(reflect.TypeFor[enum]()) {
				values := reflect.New(t).Interface().(enum).Values()
				_ = cmd.RegisterFlagCompletionFunc(flagName, func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
					return lo.Map(values, func(v string, _ int) cobra.Completion { return cobra.Completion(v) }), cobra.ShellCompDirectiveDefault
				})
			}
		case reflect.Int:
			fs.Int(flagName, 0, help)
		case reflect.Slice:
			switch t.Elem().Kind() {
			case reflect.Bool:
				fs.BoolSlice(flagName, nil, help)
			case reflect.String:
				fs.StringSlice(flagName, nil, help)
				if t.Elem().Implements(reflect.TypeFor[enum]()) {
					values := reflect.New(t.Elem()).Interface().(enum).Values()
					_ = cmd.RegisterFlagCompletionFunc(flagName, func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
						return lo.Map(values, func(v string, _ int) cobra.Completion { return cobra.Completion(v) }), cobra.ShellCompDirectiveDefault
					})
				}
			case reflect.Int:
				fs.IntSlice(flagName, nil, help)
			case reflect.Struct:
				switch {
				case ot == reflect.TypeFor[iso8601.Time]() || ot == reflect.TypeFor[time.Time]():
					// TODO: add slice
					fs.Var(flags.NewTimeValue(), flagName, help)
				case t.Elem().Implements(reflect.TypeFor[json.Marshaler]()):
					fs.StringSlice(flagName, nil, help)
				default:
					for i := range NumEntriesInSlices {
						b.buildFlagsFromStruct(cmd, t.Elem(), flagName+"."+strconv.Itoa(i)+".", required && allowRequired)
					}
				}
			}
		case reflect.Struct:
			switch {
			case t == reflect.TypeFor[iso8601.Time]() || t == reflect.TypeFor[time.Time]():
				fs.Var(flags.NewTimeValue(), flagName, help)
			case ot.Implements(reflect.TypeFor[json.Marshaler]()):
				fs.String(flagName, "", help)
			default:
				b.buildFlagsFromStruct(cmd, t, flagName+".", required && allowRequired)
			}
		}
		// we do not mark the flag as reuired as it breaks templating (the value might come from the template)
		// an annotation is set in order to mark the flag as required in the high level cli.
		if required && allowRequired {
			debug.Println(flagName, "is required")
			_ = fs.SetAnnotation(flagName, requiredAnnotation, []string{"true"})
		}
	}
}
