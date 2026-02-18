/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
	"slices"

	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/output/format"
	"github.com/outscale/octl/pkg/output/read"
	"github.com/outscale/octl/pkg/output/result"
	"github.com/samber/lo"
)

type Paginated struct {
	Read   read.Interface
	Format format.Interface
}

func (o *Paginated) Output(ctx context.Context, fetch read.FetchPage) error {
	seq := o.Read.Read(ctx, fetch)
	res := slices.Collect(seq)
	errRes, found := lo.Find(res, func(r result.Result) bool {
		return r.Error != nil
	})
	if found {
		return errRes.Error
	}
	return o.Format.Format(ctx, lo.Map(res, func(r result.Result, _ int) any { return r.Ok }))
}

func (o *Paginated) Error(ctx context.Context, v any) error {
	return o.Format.Error(ctx, v)
}

var _ Outputter = (*Paginated)(nil)

type Single struct {
	Read   read.Interface
	Format format.Interface
}

func (s *Single) Output(ctx context.Context, fetch read.FetchPage) error {
	seq := s.Read.Read(ctx, fetch)
	for v := range seq {
		if v.Error != nil {
			return v.Error
		}
		return s.Format.Format(ctx, v.Ok)
	}
	messages.Info("No results found")
	return nil
}

func (s *Single) Error(ctx context.Context, v any) error {
	return s.Format.Error(ctx, v)
}

var _ Outputter = (*Single)(nil)
