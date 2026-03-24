/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"context"
	"fmt"
	"io"
	"os"
)

type Text struct{}

func (Text) Format(ctx context.Context, w io.Writer, v any) error {
	_, err := fmt.Fprint(w, v)
	return err
}

func (Text) Error(ctx context.Context, v any) error {
	return YAML{}.Format(ctx, os.Stderr, v)
}

var _ Interface = Text{}
