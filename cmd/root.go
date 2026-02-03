/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gli",
	Short: "An eperimental CLI for Outscale services",
	Long: `   ______   __       ____
  / ____/  / /      /  _/
 / / __   / /       / /
/ /_/ /  / /___   _/ /
\____/  /_____/  /___/

One CLI to rule them all,
one CLI to find them,
one CLI to bring them all
and in the terminal bind them`,
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
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Auto accept all prompts")
	rootCmd.PersistentFlags().String("jq", "", "jq-like output filter")
}
