/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"fmt"
	"reflect"
	"time"

	"github.com/outscale/gli/pkg/config"
	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
)

func GetRow(v any, cols config.Columns) ([]string, error) {
	row := make([]string, len(cols))
	for i, c := range cols {
		val, err := c.Get(v)
		if err != nil {
			return nil, err
		}
		val = deref(val)
		switch val := val.(type) {
		case iso8601.Time:
			row[i] = val.Format(time.RFC3339)
		case time.Time:
			row[i] = val.Format(time.RFC3339)
		default:
			row[i] = fmt.Sprint(val)
		}
	}
	return row, nil
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
