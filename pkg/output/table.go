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

type Table struct {
	Content string
	Columns []config.Column
}

func (t Table) Output(ctx context.Context, v any) error {
	headers := lo.Map(t.Columns, func(c config.Column, _ int) string {
		return c.Title
	})
	rows := [][]string{}
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Slice {
		for i := range vv.Len() {
			rows = append(rows, GetRow(vv.Index(i), t.Columns))
		}
	} else {
		rows = append(rows, GetRow(vv, t.Columns))
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
