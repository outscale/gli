/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package read

import (
	"context"
	"os"
	"reflect"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/spinner"
)

type FetchPage struct {
	Method reflect.Value
	Args   []reflect.Value
}

func (f *FetchPage) Clone() FetchPage {
	nArgs := make([]reflect.Value, len(f.Args))
	nArgs[0] = f.Args[0] // ctx
	for i := 1; i < len(f.Args); i++ {
		if f.Args[i].Kind() == reflect.Pointer {
			nval := reflect.New(f.Args[i].Type().Elem())
			nval.Elem().Set(f.Args[i].Elem())
			nArgs[i] = nval
		} else {
			nval := reflect.New(f.Args[i].Type())
			nval.Elem().Set(f.Args[i])
			nArgs[i] = nval.Elem()
		}
	}
	return FetchPage{
		Method: f.Method,
		Args:   nArgs,
	}
}

func (f *FetchPage) Call(ctx context.Context) []reflect.Value {
	// display a spinner if API call lasts more than 200ms
	stopSpinner := func() {}
	if isatty.IsTerminal(os.Stderr.Fd()) {
		t := time.AfterFunc(200*time.Millisecond, func() {
			stopSpinner = spinner.Run(ctx, "Waiting for server...", false)
		})
		defer t.Stop()
	}
	// call api
	res := f.Method.Call(f.Args)
	// stop the spinner
	stopSpinner()

	return res
}
