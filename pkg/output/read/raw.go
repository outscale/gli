/*
SPDX-FileCopyrightText: 2025 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package read

import (
	"context"
	"iter"

	"github.com/outscale/octl/pkg/output/result"
)

type Raw struct{}

func NewRaw() *Raw {
	return &Raw{}
}

func (p *Raw) Read(ctx context.Context, fetch FetchPage) iter.Seq[result.Result] {
	return func(yield func(result.Result) bool) {
		vres := fetch.Call(ctx)
		if len(vres) == 0 {
			return
		}
		if err, ok := vres[len(vres)-1].Interface().(error); ok && err != nil {
			_ = yield(result.Result{Error: err})
			return
		}
		if len(vres) == 0 {
			_ = yield(result.Result{})
			return
		}
		_ = yield(result.Result{Ok: vres[0].Interface()})
	}
}

var _ Interface = (*Raw)(nil)
