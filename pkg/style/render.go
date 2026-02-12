/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package style

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Renderf(style lipgloss.Style, format string, args ...any) string {
	return style.Render(fmt.Sprintf(format, args...))
}
