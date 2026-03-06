/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"cmp"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"slices"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/x/term"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
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

type TabularFormatter interface {
	Format(ctx context.Context, w io.Writer, header []string, data [][]string) error
}

type Tabular struct {
	Explode, Sort bool
	Columns       config.Columns

	Formatter TabularFormatter
}

func validForTable(v any) bool {
	vv := reflect.ValueOf(v)
	for vv.Kind() == reflect.Interface || vv.Kind() == reflect.Pointer {
		vv = vv.Elem()
	}
	switch vv.Kind() {
	case reflect.Struct:
		return true
	case reflect.Slice:
	default:
		return false
	}
	if vv.Len() == 0 {
		return true
	}
	vv = vv.Index(0)
	for vv.Kind() == reflect.Interface || vv.Kind() == reflect.Pointer {
		vv = vv.Elem()
	}
	return vv.Kind() == reflect.Struct || vv.Kind() == reflect.Map
}

func (t Tabular) Format(ctx context.Context, w io.Writer, v any) error {
	if !validForTable(v) {
		messages.Info("Unable to format as a table, switching to YAML...")
		return YAML{}.Format(ctx, w, v)
	}
	headers := lo.Map(t.Columns, func(c config.Column, _ int) string {
		return c.Title
	})
	// build table
	rows := [][]string{}
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Slice {
		for i := range vv.Len() {
			add, err := GetRows(vv.Index(i).Interface(), t.Columns, t.Explode)
			if err != nil {
				return err
			}
			rows = append(rows, add...)
		}
	} else {
		add, err := GetRows(v, t.Columns, t.Explode)
		if err != nil {
			return err
		}
		rows = append(rows, add...)
	}
	// styling
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
	// sort
	if t.Sort {
		slices.SortFunc(rows, func(a, b []string) int {
			for i, ai := range a {
				var bi string
				if len(b) >= i+1 {
					bi = b[i]
				}
				c := cmp.Compare(ai, bi)
				if c != 0 {
					return c
				}
			}
			return 0
		})
	}

	return t.Formatter.Format(ctx, w, headers, rows)
}

func (Tabular) Error(ctx context.Context, v any) error {
	return YAML{}.Error(ctx, v)
}

type TableFormatter struct{}

func (f TableFormatter) Format(ctx context.Context, w io.Writer, headers []string, data [][]string) error {
	// build table
	cellStyle := lipgloss.NewStyle().Padding(0, 1)
	headerStyle := lipgloss.NewStyle().Bold(true).Align(lipgloss.Center)
	ot := table.New().
		Border(lipgloss.NormalBorder()).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return cellStyle
		}).Headers(headers...)
	// compute table width
	width := lo.Sum(
		lo.Map(headers, func(_ string, i int) int {
			// max is the row having the longest value for col i
			max := lo.MaxBy(data, func(a []string, b []string) bool {
				if len(a) <= i {
					return false
				}
				if len(b) <= i {
					return true
				}
				return len(a[i]) > len(b[i])
			})
			if len(max) <= i {
				return 3
			}
			debug.Println("col", i, "max", max[i], "len", len(max[i])+3)
			return len(max[i]) + 3
		}))
	if fd, ok := w.(*os.File); ok && IsTerminal(w) {
		// resize if width greater than term width
		termWidth, _, _ := term.GetSize(fd.Fd())
		if termWidth > 40 && width > termWidth-1 {
			debug.Println("reducing width from", width, "to", termWidth)
			ot.Width(termWidth)
		}
	}
	// render !
	ot.Rows(data...)

	_, err := fmt.Fprintln(w, ot)
	return err
}

type CSVFormatter struct{}

func (f CSVFormatter) Format(ctx context.Context, w io.Writer, headers []string, data [][]string) error {
	cw := csv.NewWriter(w)
	defer func() {
		cw.Flush()
	}()

	err := cw.Write(headers)
	if err != nil {
		return err
	}
	for _, row := range data {
		err := cw.Write(row)
		if err != nil {
			return err
		}
	}
	return nil
}
