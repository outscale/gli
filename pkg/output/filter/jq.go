/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package filter

import (
	"context"
	"fmt"
	"iter"

	"github.com/itchyny/gojq"
	"github.com/outscale/octl/pkg/output/result"
)

type JQ struct {
	query *gojq.Query
}

func NewJQ(s string) (JQ, error) {
	query, err := gojq.Parse(s)
	if err != nil {
		return JQ{}, fmt.Errorf("invalid jq filter: %w", err)
	}
	return JQ{query: query}, nil
}

func (j JQ) Filter(ctx context.Context, seq iter.Seq[result.Result]) iter.Seq[result.Result] {
	return func(yield func(result.Result) bool) {
		for v := range seq {
			if v.Error != nil {
				_ = yield(v)
				return
			}
			iter := j.query.RunWithContext(ctx, v.Ok)
			for {
				newOk, ok := iter.Next()
				if !ok {
					break
				}
				if err, ok := newOk.(error); ok {
					if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil { //nolint
						break
					}
					_ = yield(result.Result{Error: fmt.Errorf("jq error: %w", err)})
					return
				}
				if !yield(result.New(v, newOk)) {
					return
				}
			}
		}
	}
}

var _ Interface = JQ{}
