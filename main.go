/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"github.com/outscale/octl/cmd"
	"github.com/outscale/octl/pkg/errors"
	"github.com/outscale/octl/pkg/runner"
)

func main() {
	err := runner.CheckStdin()
	if err != nil {
		errors.ExitErr(err)
	}
	cmd.Execute()
}
