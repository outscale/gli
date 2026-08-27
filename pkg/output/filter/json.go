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

	"github.com/outscale/octl/pkg/output/result"
)

type JSON struct{}

func (j JSON) Filter(ctx context.Context, seq iter.Seq[result.Result]) iter.Seq[result.Result] {
	return func(yield func(result.Result) bool) {
		for v := range seq {
			if v.Error != nil {
				_ = yield(v)
				return
			}
			buf, err := json.Marshal(v.Ok)
			if err != nil {
				_ = yield(result.Result{Error: fmt.Errorf("to JSON: %w", err)})
				return
			}
			var raw any
			err = json.Unmarshal(buf, &raw)
			if err != nil {
				_ = yield(result.Result{Error: fmt.Errorf("from JSON: %w", err)})
				return
			}
			if !yield(result.New(v, raw)) {
				return
			}
		}
	}
}

var _ Interface = JSON{}
