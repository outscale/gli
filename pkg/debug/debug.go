//go:build debug

/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package debug

import (
	"os"

	"github.com/fatih/color"
)

func Println(s ...any) {
	_, _ = color.New(color.Faint).Fprintln(os.Stderr, s...)
}
