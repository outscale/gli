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

	"github.com/charmbracelet/huh/spinner"
	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/style"
)

type FetchPage struct {
	Method reflect.Value
	Args   []reflect.Value
}

func (f *FetchPage) Call(ctx context.Context) []reflect.Value {
	// display a spinner if API call lasts more than 200ms
	stopSpinner := func() {}
	if isatty.IsTerminal(os.Stderr.Fd()) {
		lctx, cancel := context.WithCancel(ctx)
		spinnerDone := make(chan struct{})
		t := time.AfterFunc(200*time.Millisecond, func() {
			_ = spinner.New().
				Title("Waiting for server...").
				Context(lctx).
				Output(os.Stderr).
				Style(style.Yellow).
				TitleStyle(style.Faint).
				Run()
			close(spinnerDone)
		})
		defer t.Stop()
		stopSpinner = func() {
			// kill the spinner
			cancel()
			// wait for the spinner to stop and clear it's display
			<-spinnerDone
		}
	}
	// call api
	res := f.Method.Call(f.Args)
	// stop the spinner
	stopSpinner()

	return res
}
