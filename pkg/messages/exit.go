/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package messages

import (
	"fmt"
	"os"

	"github.com/outscale/octl/pkg/style"
)

func Exit(code int, format string, a ...any) {
	_, _ = fmt.Fprintln(os.Stderr, style.Renderf(style.Error, "❌ "+format, a...))
	os.Exit(code)
}

func ExitErr(err error) {
	Exit(1, "an error occurred: %v", err)
}
