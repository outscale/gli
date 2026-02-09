/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
	"reflect"

	"github.com/outscale/gli/pkg/debug"
)

type content struct {
	content string
	output  Output
	single  bool
}

func (sr content) Output(ctx context.Context, v any) error {
	vv := reflect.Indirect(reflect.ValueOf(v))
	if sr.content != "" {
		vv = reflect.Indirect(vv.FieldByName(sr.content))
	}
	if sr.single && vv.Kind() == reflect.Slice && vv.Len() == 1 {
		vv = reflect.Indirect(vv.Index(0))
	}
	debug.Println("content", sr.content, "type", vv.Type().Name(), "kind", vv.Kind())
	return sr.output.Output(ctx, vv.Interface())
}
