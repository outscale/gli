//go:build debug

/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package debug

import (
	"fmt"
	"os"

	"github.com/outscale/octl/pkg/style"
)

func Println(s ...any) {
	_, _ = fmt.Fprintln(os.Stderr, style.Faint.Render(s...))
}
