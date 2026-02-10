/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/outscale/gli/pkg/config"
	"github.com/samber/lo"
)

var (
	green  = lipgloss.NewStyle().Foreground(lipgloss.Color("113"))
	yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("184"))
	red    = lipgloss.NewStyle().Foreground(lipgloss.Color("202"))
	colors = map[string]map[string]lipgloss.Style{
		"State": map[string]lipgloss.Style{
			"running":   green,
			"available": green,
			"ACTIVE":    green,
			"InService": green,
			"attached":  green,

			"pending":       yellow,
			"stopping":      yellow,
			"stopped":       yellow,
			"deleting":      yellow,
			"in-use":        yellow,
			"INACTIVE":      yellow,
			"disabled":      yellow,
			"confirming":    yellow,
			"allocated":     yellow,
			"attaching":     yellow,
			"detaching":     yellow,
			"shutting-down": yellow,

			"deleted":      red,
			"OutOfService": red,
			"rejected":     red,
			"expired":      red,
			"failed":       red,
			"terminated":   red,
		},
	}
)

type Table struct {
	Content string
	Columns config.Columns
}

func (t Table) Output(ctx context.Context, v any) error {
	headers := lo.Map(t.Columns, func(c config.Column, _ int) string {
		return c.Title
	})
	rows := [][]string{}
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Slice {
		for i := range vv.Len() {
			row, err := GetRow(vv.Index(i).Interface(), t.Columns)
			if err != nil {
				return err
			}
			rows = append(rows, row)
		}
	} else {
		row, err := GetRow(v, t.Columns)
		if err != nil {
			return err
		}
		rows = append(rows, row)
	}
	for r := range rows {
		for c := range rows[r] {
			styles, found := colors[headers[c]]
			if !found {
				continue
			}
			vstyle, found := styles[rows[r][c]]
			if found {
				rows[r][c] = vstyle.Render(rows[r][c])
			}
		}
	}
	cellStyle := lipgloss.NewStyle().Padding(0, 1)
	headerStyle := lipgloss.NewStyle().Bold(true).Align(lipgloss.Center)
	ot := table.New().
		Border(lipgloss.NormalBorder()).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return cellStyle
		}).Headers(headers...).
		Rows(rows...)

	_, err := fmt.Fprintln(os.Stdout, ot)
	return err
}
