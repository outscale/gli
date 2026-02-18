/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"context"
)

type None struct{}

func (None) Format(ctx context.Context, v any) error {
	return nil
}

func (None) Error(ctx context.Context, v any) error {
	return nil
}

var _ Interface = None{}
