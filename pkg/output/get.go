/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/outscale/gli/pkg/config"
	"github.com/outscale/gli/pkg/debug"
)

func Get(v reflect.Value, path string) (any, bool) {
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil, true
	}
	v = reflect.Indirect(v)
	if path == "" {
		if !v.IsValid() {
			return "", false
		}
		return v.Interface(), true
	}
	debug.Println("get", path, "on", v.Type().Name())
	before, after, _ := strings.Cut(path, ".")
	switch v.Kind() {
	case reflect.Struct:
		return Get(v.FieldByName(before), after)
	case reflect.Slice:
		idx, err := strconv.Atoi(before)
		if err != nil {
			return nil, false
		}
		if idx >= v.Len() {
			return nil, false
		}
		return Get(v.Index(idx), after)
	default:
		return nil, false
	}
}

func GetRow(v reflect.Value, cols []config.Column) []string {
	row := make([]string, len(cols))
	for i, c := range cols {
		val, found := Get(v, c.Content)
		if found && val != nil {
			row[i] = fmt.Sprint(val)
		}
	}
	return row
}
