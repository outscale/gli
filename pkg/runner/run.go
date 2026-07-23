/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package runner

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/output"
	"github.com/outscale/octl/pkg/output/read"
	"github.com/outscale/octl/pkg/style"
	"github.com/outscale/octl/pkg/watch"
	"github.com/spf13/cobra"
)

func Run[Client any, Error error](cmd *cobra.Command, args []string, cl Client, cfg config.Config, hooks ...Hook) error {
	if cmd.Flags().Lookup("waitfor").Changed {
		return waitfor[Client, Error](cmd, args, cl, cfg, hooks...)
	}
	return doRun[Client, Error](cmd, args, cl, cfg, hooks...)
}

func doRun[Client any, Error error](cmd *cobra.Command, args []string, cl Client, cfg config.Config, hooks ...Hook) error {
	clt := reflect.TypeFor[Client]()
	m, _ := clt.MethodByName(cmd.Name())

	ctx := cmd.Context()
	callArgs := []reflect.Value{
		reflect.ValueOf(ctx),
	}

	argsIndex := 0
	// for methods, the first parameter is the receiver, we need to start at 2 (after context)
	firstIdx := 2
	// for interfaces, the method has no receiver, we need to start at 1 (after context)
	if clt.Kind() == reflect.Interface {
		firstIdx = 1
	}
	injected := false
	for j := firstIdx; j < m.Type.NumIn(); j++ {
		argType := m.Type.In(j)
		if m.Type.IsVariadic() && j == m.Type.NumIn()-1 {
			continue
		}
		if argType.Kind() == reflect.Struct || argType.Kind() == reflect.Pointer {
			arg := reflect.New(argType).Elem()
			if argType.Kind() == reflect.Pointer {
				debug.Println("allocating arg")
				arg.Set(reflect.New(argType.Elem()))
			}
			if !injected {
				err := ToStruct(cmd, arg, "", hooks...)
				switch {
				case err != nil:
					return err
				default:
					injected = true
				}
			}
			callArgs = append(callArgs, arg)
			continue
		}

		if argsIndex >= len(args) {
			return fmt.Errorf("not enough arguments for %s", cmd.Name())
		}

		callArgs = append(callArgs, reflect.ValueOf(args[argsIndex]))
		argsIndex++
	}

	if len(args) > argsIndex {
		return fmt.Errorf("too many arguments for %s", cmd.Name())
	}

	c := cfg.Calls[cmd.Name()]
	debug.Println("call", cmd.Name())
	e := cfg.Entities[c.Entity]
	debug.Println("entity", c.Entity)
	fmter, out, err := output.NewFromFlags(cmd.Flags(), "", c.Content, e.Columns, e.Explode, e.Sort)
	if err != nil {
		return err
	}

	dryRun, _ := cmd.Flags().GetBool("dry-run")
	if dryRun {
		arg := callArgs[len(callArgs)-1]
		if !arg.CanInterface() {
			messages.ExitErr(errors.New("payload cannot be displayed"))
		}
		return fmter.Format(cmd.Context(), os.Stdout, arg.Interface())
	}

	doWatch, _ := cmd.Flags().GetBool("watch")
	if doWatch {
		interval, _ := cmd.Flags().GetDuration("interval")
		err = watch.Run[Error](cmd.Context(), fmter, func(ctx context.Context) error {
			return runSingle[Client, Error](ctx, cl, cmd.Name(), callArgs, out)
		}, interval)
	} else {
		err = runSingle[Client, Error](cmd.Context(), cl, cmd.Name(), callArgs, out)
	}

	if appErr, ok := errors.AsType[Error](err); ok {
		_, _ = fmt.Fprintln(os.Stderr, style.Error.Render("The server returned an error"))
		_ = out.Error(ctx, appErr)
		os.Exit(1)
	}
	return err
}

func runSingle[Client any, Error error](ctx context.Context, cl Client, callName string, callArgs []reflect.Value, out output.Outputter) error {
	call := read.FetchPage{
		Method: reflect.ValueOf(cl).MethodByName(callName),
		Args:   callArgs,
	}
	return out.Output(ctx, call)
}
