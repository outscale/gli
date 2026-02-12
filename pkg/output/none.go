/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
)

type None struct {
}

func (t None) Output(ctx context.Context, v any) error {
	return nil
}
