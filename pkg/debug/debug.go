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
	"github.com/samber/lo"
)

func Println(s ...any) {
	_, _ = fmt.Fprintln(os.Stderr, style.Faint.Render(lo.Map(s, func(s any, _ int) string { return fmt.Sprintf("%v", s) })...))
}
