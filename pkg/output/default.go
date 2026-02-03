/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/tidwall/pretty"
)

type Default struct{}

func (Default) Output(ctx context.Context, v any) error {
	buf, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}
	if isatty.IsTerminal(os.Stdout.Fd()) {
		buf = pretty.Color(buf, nil)
	}
	_, err = fmt.Fprintln(os.Stdout, string(buf))
	return err
}
