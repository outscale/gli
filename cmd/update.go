/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/outscale/gli/pkg/errors"
	"github.com/outscale/gli/pkg/update"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gli to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		err := update.Update(context.Background())
		if err != nil {
			errors.ExitErr(fmt.Errorf("unable to update: %w", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
