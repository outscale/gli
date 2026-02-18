/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"context"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type YAML struct{}

func (YAML) Format(ctx context.Context, v any) error {
	enc := yaml.NewEncoder(os.Stdout, yaml.OmitZero(), yaml.UseSingleQuote(true), yaml.Indent(2))
	err := enc.Encode(v)
	if err != nil {
		return fmt.Errorf("marshal yaml: %w", err)
	}
	return enc.Close()
}

func (y YAML) Error(ctx context.Context, v any) error {
	return y.Format(ctx, v)
}

var _ Interface = YAML{}
