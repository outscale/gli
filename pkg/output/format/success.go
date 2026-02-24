/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"context"

	"github.com/outscale/octl/pkg/messages"
)

type Success struct{}

func (Success) Format(ctx context.Context, v any) error {
	messages.Success("The operation completed successfully")
	return nil
}

func (Success) Error(ctx context.Context, v any) error {
	return YAML{}.Error(ctx, v)
}

var _ Interface = Success{}
