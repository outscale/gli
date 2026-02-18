/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package read

import (
	"context"
	"iter"

	"github.com/outscale/octl/pkg/output/result"
)

type Interface interface {
	Read(ctx context.Context, fetch FetchPage) iter.Seq[result.Result]
}
