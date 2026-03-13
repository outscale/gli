/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/output/filter"
	"github.com/outscale/octl/pkg/output/format"
	"github.com/outscale/octl/pkg/output/read"
	"github.com/outscale/octl/pkg/output/result"
	"github.com/samber/lo"
)

var writeTo io.Writer = os.Stdout

func InjectOutput(w io.Writer) {
	writeTo = w
}

type Paginated struct {
	Read    read.Interface
	Format  format.Interface
	Filters []filter.Interface
	WriteTo string
}

func (p *Paginated) Output(ctx context.Context, fetch read.FetchPage) (err error) {
	writeTo := writeTo
	if p.WriteTo != "" {
		fd, err := os.Create(p.WriteTo)
		if err != nil {
			return fmt.Errorf("unable to write to %q: %w", p.WriteTo, err)
		}
		messages.Info("Writing output to %s", p.WriteTo)
		writeTo = fd
		defer func() {
			if wc, ok := writeTo.(io.Closer); ok {
				cerr := wc.Close()
				if cerr != nil && err == nil {
					err = fmt.Errorf("output error: %w", cerr)
				}
			}
		}()
	}
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
		return p.Format.Format(ctx, writeTo, res[0].Ok)
	}
	return p.Format.Format(ctx, writeTo, lo.Map(res, func(r result.Result, _ int) any { return r.Ok }))
}

func (p *Paginated) Error(ctx context.Context, v any) error {
	return p.Format.Error(ctx, v)
}

var _ Outputter = (*Paginated)(nil)
