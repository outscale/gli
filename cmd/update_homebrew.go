//go:build homebrew

/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"github.com/outscale/octl/pkg/messages"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update octl to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		messages.Info(`Use "brew install octl" to upgrade octl.`)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
