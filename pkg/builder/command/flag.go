package commandbuilder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/outscale/goutils/sdk/ptr"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/flags"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func (b *Builder) buildFlag(cmd *cobra.Command, f config.Flag) error {
	if !strings.Contains(f.Name, ".#.") {
		return b.buildSingleFlag(cmd, f)
	}
	return b.buildMultipleFlag(cmd, f)
}

func (b *Builder) buildMultipleFlag(cmd *cobra.Command, f config.Flag) error {
	if !strings.Contains(f.Name, ".#.") {
		return b.buildSingleFlag(cmd, f)
	}
	for _, flag := range b.buildFlagSlice(f.Name) {
		nf := f
		nf.Name = flag
		if err := b.buildSingleFlag(cmd, nf); err != nil {
			return err
		}
	}
	return nil
}

func (b *Builder) buildFlagSlice(flagName string) []string {
	before, after, found := strings.Cut(flagName, ".#.")
	if !found {
		return []string{flagName}
	}
	lst := []string{}
	for i := range NumEntriesInSlices(before) {
		lst = append(lst, b.buildFlagSlice(fmt.Sprintf("%s.%d.%s", before, i, after))...)
	}
	return lst
}

func (b *Builder) buildSingleFlag(cmd *cobra.Command, f config.Flag) error {
	fs := cmd.Flags()
	switch f.Kind {
	case reflect.Bool:
		if f.ContainerKind == reflect.Slice {
			fs.BoolSlice(f.Name, nil, f.Help)
		} else {
			fs.Bool(f.Name, false, f.Help)
		}
	case reflect.Int:
		if f.ContainerKind == reflect.Slice {
			fs.IntSlice(f.Name, nil, f.Help)
		} else {
			fs.Int(f.Name, 0, f.Help)
		}
	case reflect.Int32:
		if f.ContainerKind == reflect.Slice {
			fs.Int32Slice(f.Name, nil, f.Help)
		} else {
			fs.Int32(f.Name, 0, f.Help)
		}
	case reflect.Int64:
		if f.ContainerKind == reflect.Slice {
			fs.Int64Slice(f.Name, nil, f.Help)
		} else {
			fs.Int64(f.Name, 0, f.Help)
		}
	case reflect.String:
		switch {
		case f.CustomValue != "":
			switch f.CustomValue {
			case flags.Base64File:
				fs.Var(flags.NewBase64FileValue(), f.Name, f.Help)
			case flags.FileOrJSON:
				fs.Var(flags.NewFileOrJSONValue(), f.Name, f.Help)
			case flags.Reader:
				fs.Var(flags.NewReaderValue(), f.Name, f.Help)
			case flags.Time:
				fs.Var(flags.NewTimeValue(), f.Name, f.Help)
			default:
				debug.Println("missing custom value", f.CustomValue)
			}
		case f.ContainerKind == reflect.Slice:
			fs.StringSlice(f.Name, nil, f.Help)
		case f.ContainerKind == reflect.Map:
			fs.StringToString(f.Name, map[string]string{}, f.Help)
		default:
			fs.String(f.Name, "", f.Help)
		}
		if len(f.AllowedValues) > 0 {
			_ = cmd.RegisterFlagCompletionFunc(f.Name, func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
				return lo.Map(f.AllowedValues, func(v string, _ int) cobra.Completion { return cobra.Completion(v) }), cobra.ShellCompDirectiveDefault
			})
		}
	}
	if ptr.From(f.Required) {
		_ = cmd.MarkFlagRequired(f.Name)
	}
	if f.Default != "" {
		_ = flags.SetDefault(cmd.Flags(), f.Name, f.Default)
	}
	switch f.CustomValue {
	case flags.Base64File, flags.FileOrJSON:
		_ = cmd.MarkFlagFilename(f.Name)
	}
	return nil
}
