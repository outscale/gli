/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package filter

import (
	"context"
	"iter"

	"github.com/outscale/octl/pkg/output/result"
)

type Skip struct {
	iter    int
	skip    int
	skipped int
}

func NewSkip(skip int) *Skip {
	return &Skip{
		skip: skip,
	}
}

func (e Skip) Filter(ctx context.Context, seq iter.Seq[result.Result]) iter.Seq[result.Result] {
	return func(yield func(result.Result) bool) {
		for v := range seq {
			if v.Error != nil {
				_ = yield(v)
				return
			}
			if v.Iter != e.iter {
				e.iter = v.Iter
				e.skipped = 0
			}
			if e.skipped < e.skip {
				e.skipped++
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}

var _ Interface = Skip{}
