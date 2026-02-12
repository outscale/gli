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
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/style"
	"github.com/samber/lo"
)

var (
	colors = map[string]map[string]lipgloss.Style{
		"State": map[string]lipgloss.Style{
			"running":   style.Green,
			"available": style.Green,
			"ACTIVE":    style.Green,
			"InService": style.Green,
			"attached":  style.Green,

			"pending":       style.Yellow,
			"stopping":      style.Yellow,
			"stopped":       style.Yellow,
			"deleting":      style.Yellow,
			"in-use":        style.Yellow,
			"INACTIVE":      style.Yellow,
			"disabled":      style.Yellow,
			"confirming":    style.Yellow,
			"allocated":     style.Yellow,
			"attaching":     style.Yellow,
			"detaching":     style.Yellow,
			"shutting-down": style.Yellow,

			"deleted":      style.Red,
			"OutOfService": style.Red,
			"rejected":     style.Red,
			"expired":      style.Red,
			"failed":       style.Red,
			"terminated":   style.Red,
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
