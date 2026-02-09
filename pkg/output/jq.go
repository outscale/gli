/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/itchyny/gojq"
)

type JQ struct {
	query *gojq.Query
}

func NewJQ(s string) (*JQ, error) {
	query, err := gojq.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("parse jq filter: %w", err)
	}
	return &JQ{query: query}, nil
}

func (j *JQ) Output(ctx context.Context, v any) error {
	buf, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("to JSON: %w", err)
	}
	var raw map[string]any
	err = json.Unmarshal(buf, &raw)
	if err != nil {
		return fmt.Errorf("from JSON: %w", err)
	}
	iter := j.query.RunWithContext(ctx, raw)
	out := JSON{}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
				break
			}
			return fmt.Errorf("jq error: %w", err)
		}
		if err := out.Output(ctx, v); err != nil {
			return fmt.Errorf("output: %w", err)
		}
	}
	return nil
}
