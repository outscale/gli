/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/structs"
)

type Body struct{}

func (Body) Format(ctx context.Context, w io.Writer, v any) (err error) {
	vv, found := structs.FindFieldByType[io.ReadCloser](reflect.ValueOf(v))
	if !found {
		return YAML{}.Format(ctx, w, v)
	}
	r, ok := vv.Interface().(io.ReadCloser)
	if !ok {
		return YAML{}.Format(ctx, w, v)
	}
	defer func() {
		cerr := r.Close()
		if err == nil {
			err = cerr
		}
	}()
	// fetch the first 100ish bytes to detect mime type
	wbuf := &bytes.Buffer{}
	_, err = io.Copy(wbuf, io.LimitReader(r, 100))
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}
	buf := wbuf.Bytes()
	debug.Println("detected mime type", mimetype.Detect(buf).String(), "from", string(buf))
	if !mimetype.Detect(buf).Is("text/plain") && IsTerminal(w) {
		_ = r.Close()
		messages.Warn("not displaying binary data to terminal, you need to redirect output to a file")
		os.Exit(1) //nolint
	}
	// output first 100ish bytes
	_, err = io.Copy(w, wbuf)
	if err == nil {
		// output remainder of body
		_, err = io.Copy(w, r)
	}
	return err
}

func (Body) Error(ctx context.Context, v any) error {
	return YAML{}.Error(ctx, v)
}

var _ Interface = Body{}
