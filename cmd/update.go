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
		var opts []update.UpdateOption

		if cmd.Flag("ignore-signature-verification").Value.String() == "true" {
			opts = append(opts, update.WithIgnoreSignature())
		}
		if cmd.Flag("ignore-digest-verification").Value.String() == "true" {
			opts = append(opts, update.WithIgnoreDigest())
		}

		err := update.Update(cmd.Context(), opts...)
		if err != nil {
			messages.ExitErr(fmt.Errorf("unable to update: %w", err))
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().Bool("ignore-signature-verification", false, "Ignore signature verification for the update")
	updateCmd.Flags().Bool("ignore-digest-verification", false, "Ignore digest verification for the update")
}
