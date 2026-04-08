/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"fmt"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/outscale/octl/cmd/prerun"
	"github.com/outscale/octl/pkg/markdown"
	"github.com/outscale/octl/pkg/version"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	sdkversion "github.com/outscale/osc-sdk-go/v3/pkg/version"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var (
	a = lipgloss.NewStyle().Foreground(lipgloss.Color("#AE56CF")).Render
	b = lipgloss.NewStyle().Foreground(lipgloss.Color("#2330B6")).Render
	c = lipgloss.NewStyle().Foreground(lipgloss.Color("#17A7EE")).Render
	d = lipgloss.NewStyle().Foreground(lipgloss.Color("#18D759")).Render
	e = lipgloss.NewStyle().Foreground(lipgloss.Color("#F1F139")).Render
	// f = lipgloss.NewStyle().Foreground(lipgloss.Color("#F6683C")).Render
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "octl",
	Short: "A modern CLI for Outscale services",
	Long: "```" + a(`
                    __     __
  ____     _____   / /_   / /
 / __ \   / ___/  / __/  / /
/ /_/ /  / /__   / /_   / /
\____/   \___/   \__/  /_/`) + `
` + b(`One CLI to rule them all,`) + `
` + c(`one CLI to find them,`) + `
` + d(`one CLI to bring them all`) + `
` + e(`and in the terminal bind them.`) + `
` + "```",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		prerun.CheckFalse(cmd, args)
		prerun.CheckUpdate(cmd, args)
	},
	Run: func(cmd *cobra.Command, _ []string) {
		if b, _ := cmd.Flags().GetBool("version"); b {
			fmt.Printf("octl version %s\n", version.Version)
			fmt.Printf("based on Go SDK %s\n", sdkversion.Version)
			fmt.Printf("Providers:\n* iaas: %s\n", osc.Version)
			return
		}
		_ = cmd.Help()
	},
	SilenceErrors:     true, // do not display errors when an error occurred, we do it
	SilenceUsage:      true, // do not display usage when an error occurred, the user will need to call -h
	DisableAutoGenTag: true,
}

func Root() *cobra.Command {
	return rootCmd
}

func init() {
	md := markdown.NewRenderer()
	if long, err := md.Render(rootCmd.Long); err == nil {
		rootCmd.Long = long
	}
	rootCmd.AddGroup(&cobra.Group{ID: "services", Title: "OUTSCALE Services Commands"})
	rootCmd.AddGroup(&cobra.Group{ID: "config", Title: "Configuration Commands"})

	rootCmd.Flags().Bool("version", false, "Display version")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().String("config", "", "Path of profile file (by default, ~/.osc/config.json)")
	rootCmd.PersistentFlags().String("profile", "", fmt.Sprintf("Profile to use in profile file (by default, %q)", profile.DefaultProfile))

	rootCmd.PersistentFlags().String("template", "", "JSON template file for query body")
	rootCmd.PersistentFlags().String("template-root", "", "the root attribute for the template")
	_ = rootCmd.PersistentFlags().MarkHidden("template-root")
	rootCmd.PersistentFlags().String("payload", "", "JSON content for query body")

	rootCmd.PersistentFlags().String("jq", "", "jq filter")
	rootCmd.PersistentFlags().StringSlice("filter", nil, `comma separated list of filters for results - name:value,name:value, alias for jq filter 'select(.name | test("value"))'`)

	rootCmd.PersistentFlags().String("waitfor", "", `jq expression to wait for - octl will query every waitfor-interval until the expression returns 1/true or a non empty result`)
	rootCmd.PersistentFlags().Duration("waitfor-interval", 5*time.Second, `interval between two waitfor iterations`)
	rootCmd.PersistentFlags().Duration("waitfor-timeout", 10*time.Minute, `maximum duration of a wait`)

	rootCmd.PersistentFlags().StringP("columns", "c", "", "columns to display - [+]<title>:<jq query for content>||<title>:<jq query for content>")
	rootCmd.PersistentFlags().StringP("output", "o", "raw", "output format (raw, json, yaml, table, csv, none, base64, text)")
	rootCmd.PersistentFlags().StringP("out-file", "O", "", "redirect output to file")
	rootCmd.PersistentFlags().Bool("single", false, "convert single entry lists to a single object")

	rootCmd.PersistentFlags().Bool("no-upgrade", false, "do not check for new versions")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "answer yes to all prompts")

	rootCmd.PersistentFlags().StringSlice("hooks", nil, "")
	_ = rootCmd.PersistentFlags().MarkHidden("hooks")

	_ = rootCmd.RegisterFlagCompletionFunc("output", func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
		return []cobra.Completion{"raw", "json", "yaml", "table", "csv", "none", "base64", "text"}, cobra.ShellCompDirectiveDefault
	})

	_ = rootCmd.RegisterFlagCompletionFunc("profile", func(cmd *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
		cf, _ := loadConfig(cmd)
		return lo.Map(lo.Keys(cf.Profiles), func(k string, _ int) cobra.Completion { return cobra.Completion(k) }), cobra.ShellCompDirectiveDefault
	})
}
