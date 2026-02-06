/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"github.com/outscale/gli/cmd"
	"github.com/outscale/gli/pkg/errors"
	"github.com/outscale/gli/pkg/runner"
)

func main() {
	err := runner.CheckStdin()
	if err != nil {
		errors.ExitErr(err)
	}
	cmd.Execute()
}
