/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
	"slices"

	"github.com/outscale/octl/pkg/output/filter"
	"github.com/outscale/octl/pkg/output/format"
	"github.com/outscale/octl/pkg/output/read"
	"github.com/outscale/octl/pkg/output/result"
	"github.com/samber/lo"
)

type Paginated struct {
	Read    read.Interface
	Format  format.Interface
	Filters []filter.Interface
}

func (p *Paginated) Output(ctx context.Context, fetch read.FetchPage) error {
	seq := p.Read.Read(ctx, fetch)
	for _, f := range p.Filters {
		seq = f.Filter(ctx, seq)
	}
	res := slices.Collect(seq)
	errRes, found := lo.Find(res, func(r result.Result) bool {
		return r.Error != nil
	})
	if found {
		return errRes.Error
	}
	if len(res) == 1 && res[0].SingleEntry {
		return p.Format.Format(ctx, res[0].Ok)
	}
	return p.Format.Format(ctx, lo.Map(res, func(r result.Result, _ int) any { return r.Ok }))
}

func (p *Paginated) Error(ctx context.Context, v any) error {
	return p.Format.Error(ctx, v)
}

var _ Outputter = (*Paginated)(nil)
