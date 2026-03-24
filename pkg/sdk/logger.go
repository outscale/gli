/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type VerboseLogger struct{}

func (VerboseLogger) RequestHttp(ctx context.Context, req *http.Request) {
	fmt.Fprintf(os.Stderr, "- REQUEST -------------------\n\n%s %s\n\n", req.Method, req.URL)
	h := req.Header
	if h.Get("Secretkey") != "" {
		h = req.Header.Clone()
		h.Set("Secretkey", "[REDACTED]")
	}
	_ = h.Write(os.Stderr)
	fmt.Fprintln(os.Stderr)
	if req.GetBody != nil {
		bodyReader, err := req.GetBody()
		if err == nil {
			bodyBytes, _ := io.ReadAll(bodyReader)
			fmt.Fprintf(os.Stderr, "%s\n\n", string(bodyBytes))
		}
	}
	fmt.Fprint(os.Stderr, "- REQUEST -------------------\n\n")
}

func (VerboseLogger) ResponseHttp(ctx context.Context, resp *http.Response, d time.Duration) {
	fmt.Fprintf(os.Stderr, "- RESPONSE ------------------\n\n%s\n\n", resp.Status)
	_ = resp.Header.Write(os.Stderr)
	fmt.Fprintln(os.Stderr)
	bodyBytes, _ := io.ReadAll(resp.Body)
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	fmt.Fprintf(os.Stderr, "%s\n\n", string(bodyBytes))
	fmt.Fprint(os.Stderr, "- RESPONSE ------------------\n\n")
}

func (VerboseLogger) Request(ctx context.Context, req any) {}

func (VerboseLogger) Response(ctx context.Context, resp any) {}

func (VerboseLogger) Error(ctx context.Context, err error) {}
