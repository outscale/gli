/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"context"
	"os"

	"github.com/outscale/octl/cmd"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/runner"
)

func main() {
	err := runner.CheckStdin()
	if err != nil {
		messages.ExitErr(err)
	}
	os.Args, err = runner.TemplateArgs(os.Args)
	if err != nil {
		messages.ExitErr(err)
	}
	ctx := context.Background()
	err = cmd.Root().ExecuteContext(ctx)
	if err != nil {
		messages.ExitErr(err)
	}
}
