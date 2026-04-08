/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package style

import (
	"os"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
)

var (
	Green  = lipgloss.NewStyle().Foreground(lipgloss.Color("113"))
	Yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("184"))
	Red    = lipgloss.NewStyle().Foreground(lipgloss.Color("202"))

	Faint = lipgloss.NewStyle().Faint(true)
	Error = lipgloss.NewStyle().Foreground(lipgloss.Color("202")).Bold(true)
)

func Theme() *huh.Styles {
	t := huh.ThemeBase(lipgloss.HasDarkBackground(os.Stdin, os.Stdout))
	t.Focused.Title = t.Focused.Title.Foreground(lipgloss.Color("184"))
	t.Focused.FocusedButton = t.Focused.FocusedButton.
		Border(lipgloss.NormalBorder()).
		BorderForeground(compat.AdaptiveColor{Light: lipgloss.Color("236"), Dark: lipgloss.Color("254")}).
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color("0"), Dark: lipgloss.Color("15")}).
		Background(lipgloss.NoColor{})
	t.Focused.BlurredButton = t.Focused.BlurredButton.
		Border(lipgloss.NormalBorder()).
		BorderForeground(compat.AdaptiveColor{Light: lipgloss.Color("254"), Dark: lipgloss.Color("236")}).
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color("239"), Dark: lipgloss.Color("248")}).
		Background(lipgloss.NoColor{})
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.PaddingRight(1)
	return t
}
