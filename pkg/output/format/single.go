/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"context"
	"reflect"
)

type Single struct {
	ForFormat Interface
}

func (s Single) Format(ctx context.Context, v any) error {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Slice && vv.Len() == 1 {
		return s.ForFormat.Format(ctx, vv.Index(0).Interface())
	}
	return s.ForFormat.Format(ctx, v)
}

func (s Single) Error(ctx context.Context, v any) error {
	return s.ForFormat.Error(ctx, v)
}

var _ Interface = Single{}
