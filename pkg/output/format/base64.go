/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"reflect"

	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
)

type Base64 struct{}

func (b Base64) Format(ctx context.Context, v any) error {
	debug.Println("base64 decode", reflect.TypeOf(v).Name(), "kind", reflect.ValueOf(v).Kind())
	switch v := v.(type) {
	case []any:
		for _, s := range v {
			err := b.Format(ctx, s)
			if err != nil {
				messages.Info("Unable to decode base64 array, switching to YAML...")
			}
		}
	case []string:
		for _, s := range v {
			err := b.Format(ctx, s)
			if err != nil {
				messages.Info("Unable to decode base64 array, switching to YAML...")
			}
		}
	case string:
		buf, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return fmt.Errorf("base64: %w", err)
		}
		_, _ = fmt.Fprintln(os.Stdout, string(buf))
	default:
		messages.Info("Unable to decode base64, switching to YAML...")
		return YAML{}.Format(ctx, v)
	}
	return nil
}

func (Base64) Error(ctx context.Context, v any) error {
	return YAML{}.Format(ctx, v)
}

var _ Interface = Base64{}
