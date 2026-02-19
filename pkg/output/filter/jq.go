/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package filter

import (
	"context"
	"encoding/json"
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
		return JQ{}, fmt.Errorf("parse jq filter: %w", err)
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
			buf, err := json.Marshal(v.Ok)
			if err != nil {
				_ = yield(result.Result{Error: fmt.Errorf("jq to JSON: %w", err)})
				return
			}
			var raw any
			err = json.Unmarshal(buf, &raw)
			if err != nil {
				_ = yield(result.Result{Error: fmt.Errorf("jq from JSON: %w", err)})
				return
			}
			iter := j.query.RunWithContext(ctx, raw)
			for {
				v, ok := iter.Next()
				if !ok {
					break
				}
				if err, ok := v.(error); ok {
					if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
						break
					}
					_ = yield(result.Result{Error: fmt.Errorf("jq error: %w", err)})
					return
				}
				if !yield(result.Result{Ok: v}) {
					return
				}
			}
		}
	}
}

var _ Interface = JQ{}
