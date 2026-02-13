/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
	"reflect"

	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
)

type content struct {
	contentField string
	output       Output
	single       bool
}

func (sr content) Content(ctx context.Context, v any) error {
	vv := reflect.Indirect(reflect.ValueOf(v))
	if sr.contentField != "" {
		vv = reflect.Indirect(vv.FieldByName(sr.contentField))
	}
	if sr.single && vv.Kind() == reflect.Slice {
		switch vv.Len() {
		case 0:
			messages.Exit(1, "no resource found")
		case 1:
			vv = reflect.Indirect(vv.Index(0))
		default:
			messages.Warn("dropping %d contents", vv.Len()-1)
			vv = reflect.Indirect(vv.Index(0))
		}
	}
	debug.Println("content", sr.contentField, "type", vv.Type().Name(), "kind", vv.Kind())
	return sr.output.Content(ctx, vv.Interface())
}

func (sr content) Error(ctx context.Context, v any) error {
	return sr.output.Error(ctx, v)
}
