/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/outscale/octl/cmd/prerun"
	"github.com/outscale/octl/pkg/errors"
	"github.com/outscale/octl/pkg/version"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	sdkversion "github.com/outscale/osc-sdk-go/v3/pkg/version"
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
	Short: "An eperimental CLI for Outscale services",
	Long: a(`                    __     __
  ____     _____   / /_   / /
 / __ \   / ___/  / __/  / / 
/ /_/ /  / /__   / /_   / /  
\____/   \___/   \__/  /_/`) + `   
` + b(`One CLI to rule them all,`) + `
` + c(`one CLI to find them,`) + `
` + d(`one CLI to bring them all`) + `
` + e(`and in the terminal bind them.`),
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
	SilenceErrors: true, // do not display errors when an error occurred, we do it
	SilenceUsage:  true, // do not display usage when an error occurred, the user will need to call -h
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		errors.ExitErr(err)
	}
}

func init() {
	rootCmd.Flags().Bool("version", false, "Display version")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().String("jq", "", "jq-like output filter")
	path, _ := profile.DefaultConfigPath()
	rootCmd.PersistentFlags().String("config", "", fmt.Sprintf("Path of profile file (by default, %q)", path))
	rootCmd.PersistentFlags().String("profile", "", fmt.Sprintf("Profile to use in profile file (by default, %q)", profile.DefaultProfile))
	rootCmd.PersistentFlags().String("template", "", "JSON template for query body")
	rootCmd.PersistentFlags().StringP("columns", "c", "", "columns to display - [+]title:content|title:content")
	rootCmd.PersistentFlags().StringP("output", "o", "", "output format (raw, json, yaml, table, none)")
	rootCmd.PersistentFlags().Bool("no-upgrade", false, "do not check for new versions")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "answer yes to all prompts")

	_ = rootCmd.RegisterFlagCompletionFunc("output", func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
		return []cobra.Completion{"raw", "json", "yaml", "table", "none"}, cobra.ShellCompDirectiveDefault
	})
}
