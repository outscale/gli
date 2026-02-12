/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package style

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	Green  = lipgloss.NewStyle().Foreground(lipgloss.Color("113"))
	Yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("184"))
	Red    = lipgloss.NewStyle().Foreground(lipgloss.Color("202"))

	Faint = lipgloss.NewStyle().Faint(true)
	Error = lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Bold(true)
)

func Theme() *huh.Theme {
	t := huh.ThemeBase()
	t.Focused.Title = t.Focused.Title.Foreground(lipgloss.Color("184"))
	return t
}
