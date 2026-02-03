/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package errors

import (
	"os"

	"github.com/fatih/color"
)

func Exit(code int, format string, a ...any) {
	_, _ = color.New(color.FgRed).Add(color.Bold).Fprintf(os.Stderr, format, a...)
	os.Exit(code)
}

func ExitErr(err error) {
	Exit(1, "an error occurred: %v", err)
}

func Warn(format string, a ...any) {
	_, _ = color.New(color.FgYellow).Add(color.Faint).Fprintf(os.Stderr, format, a...)
}
