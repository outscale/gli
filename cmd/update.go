/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"fmt"

	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/update"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update octl to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		err := update.Update(cmd.Context())
		if err != nil {
			messages.ExitErr(fmt.Errorf("unable to update: %w", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
