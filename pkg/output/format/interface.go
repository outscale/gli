/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"context"
)

type Interface interface {
	Format(ctx context.Context, v any) error
	Error(ctx context.Context, v any) error
}
