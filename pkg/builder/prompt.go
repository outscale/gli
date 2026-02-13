/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package builder

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/errors"
	"github.com/outscale/octl/pkg/style"
	"github.com/spf13/cobra"
)

var (
	prompts = map[config.Action]string{
		config.ActionDelete: "Are you sure you want to delete this resource ?",
	}
	success = map[config.Action]string{
		config.ActionDelete: "The resource has been deleted",
	}
)

func confirm(action config.Action, display, run func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if yes, _ := cmd.Flags().GetBool("yes"); yes {
			run(cmd, args)
			return
		}
		if !isatty.IsTerminal(os.Stdout.Fd()) {
			errors.Exit(1, "unable to confirm action, aborting. Add -y to skip confirmation prompt")
		}
		if display != nil {
			display(cmd, args)
		}
		var yes bool
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title(prompts[action]).
					Affirmative("✅ Yes").
					Negative("❌ No").
					Value(&yes),
			),
		).WithTheme(style.Theme()).Run()
		if err != nil {
			errors.ExitErr(err)
		}
		if yes {
			run(cmd, args)
			_, _ = fmt.Fprint(os.Stderr, style.Green.Render(success[action]))
		}
	}
}
