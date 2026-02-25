/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format

import (
	"context"
	"io"
	"os"
	"reflect"

	"github.com/gabriel-vasile/mimetype"
	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/structs"
)

type Body struct{}

func (Body) Format(ctx context.Context, v any) (err error) {
	vv, found := structs.FindFieldByType[io.ReadCloser](reflect.ValueOf(v))
	if !found {
		return YAML{}.Format(ctx, v)
	}
	r, ok := vv.Interface().(io.ReadCloser)
	if !ok {
		return YAML{}.Format(ctx, v)
	}
	defer func() {
		cerr := r.Close()
		if err == nil {
			err = cerr
		}
	}()
	// fetch the first 100 bytes to detect mime type
	buf := make([]byte, 100)
	_, err = r.Read(buf)
	if err != nil {
		return err
	}
	debug.Println("detected mime type", mimetype.Detect(buf).String(), "from", string(buf))
	if !mimetype.Detect(buf).Is("text/plain") && isatty.IsTerminal(os.Stdout.Fd()) {
		messages.Warn("not displaying binary data to terminal, you need to redirect output to a file")
		os.Exit(1)
	}
	// output first 10 bytes
	_, err = os.Stdout.Write(buf)
	if err == nil {
		// output remainder of body
		_, err = io.Copy(os.Stdout, r)
	}
	return err
}

func (Body) Error(ctx context.Context, v any) error {
	return YAML{}.Error(ctx, v)
}

var _ Interface = Body{}
