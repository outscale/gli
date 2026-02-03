/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"context"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/outscale/gli/pkg/errors"
	"github.com/outscale/gli/pkg/update"
	"github.com/spf13/cobra"
)

var (
	a = color.RGB(174, 86, 207).Sprint
	b = color.RGB(35, 48, 182).Sprint
	c = color.RGB(23, 167, 238).Sprint
	d = color.RGB(24, 215, 89).Sprint
	e = color.RGB(241, 241, 57).Sprint
	// f = color.RGB(246, 104, 60).Sprint
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gli",
	Short: "An eperimental CLI for Outscale services",
	Long: a(`   ______   __       ____
  / ____/  / /      /  _/
 / / __   / /       / /
/ /_/ /  / /___   _/ /
\____/  /_____/  /___/

`) + b(`One CLI to rule them all,
`) + c(`one CLI to find them,
`) + d(`one CLI to bring them all
`) + e(`and in the terminal bind them.`),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !isatty.IsTerminal(os.Stdout.Fd()) {
			return
		}
		ctx := context.Background()
		ghCtx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		latest, err := update.LatestRelease(ghCtx)
		if err != nil {
			errors.Warn("❌ Unable to fetch latest release: %v\n", err)
			return
		}
		if latest == "" || latest == Version {
			return
		}
		_, _ = color.New(color.FgGreen).Add(color.Bold).Printf("⬆️ New version %s detected - call gli update to update\n", latest)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().String("jq", "", "jq-like output filter")
}
