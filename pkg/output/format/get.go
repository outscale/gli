/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
)

func GetRows(v any, cols config.Columns, explode bool) ([][]string, error) {
	// compute result in a single line table
	raw := make([][]any, 1)
	raw[0] = make([]any, len(cols))
	lines := 1
	for i, c := range cols {
		val, err := c.Get(v)
		if err != nil {
			debug.Println("Get error: %v", err)
			// we continue with a nil value
		}
		val = deref(val)
		raw[0][i] = val
		if reflect.TypeOf(val).Kind() == reflect.Slice {
			lines = max(lines, reflect.ValueOf(val).Len())
		}
	}
	// if explode, explode each cell into multiple lines
	// slices are copied, values are repeated
	if explode && lines > 1 {
		debug.Println("exploding to", lines, "lines")
		for len(raw) < lines {
			raw = append(raw, make([]any, len(cols)))
		}
		for j, val := range raw[0] {
			rval := reflect.ValueOf(val)
			if rval.Kind() != reflect.Slice {
				sval := reflect.MakeSlice(reflect.SliceOf(reflect.TypeFor[any]()), lines, lines)
				for i := range lines {
					sval.Index(i).Set(rval)
				}
				rval = sval
			}
			for i := range rval.Len() {
				raw[i][j] = rval.Index(i).Interface()
			}
		}
	}
	// convert to table of strings
	rows := make([][]string, len(raw))
	for i := range raw {
		rows[i] = make([]string, len(cols))
		for j, val := range raw[i] {
			switch val := val.(type) {
			case iso8601.Time:
				rows[i][j] = val.String()
			case time.Time:
				rows[i][j] = val.Format(time.RFC3339)
			case float64:
				rows[i][j] = fmt.Sprintf("%.2f", val)
			default:
				rows[i][j] = strings.TrimSpace(fmt.Sprint(val))
			}
		}
	}
	return rows, nil
}

func deref(val any) any {
	rval := reflect.ValueOf(val)
	tval := reflect.TypeOf(val)
	switch {
	case val == nil || rval.IsZero():
		return ""
	case tval.Kind() == reflect.Ptr || tval.Kind() == reflect.Interface:
		return deref(rval.Elem().Interface())
	case tval.Kind() == reflect.Slice && rval.Len() == 0:
		return ""
	case tval.Kind() == reflect.Slice && rval.Len() == 1:
		return deref(rval.Index(0).Interface())
	case tval.Kind() == reflect.Slice:
		// convert slice to []any, with derefs values
		nval := make([]any, rval.Len())
		for i := range rval.Len() {
			nval[i] = deref(rval.Index(i).Interface())
		}
		return nval
	}
	return val
}
