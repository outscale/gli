/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"

	"github.com/outscale/octl/pkg/output/read"
)

type Outputter interface {
	Output(ctx context.Context, fetch read.FetchPage) error
	Error(ctx context.Context, v any) error
}
