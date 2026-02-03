/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"strings"

	"github.com/outscale/gli/pkg/debug"
	"github.com/outscale/gli/pkg/errors"
	"github.com/outscale/gli/pkg/flags"
	"github.com/outscale/gli/pkg/output"
	"github.com/outscale/gli/pkg/sdk"
	"github.com/outscale/gli/pkg/version"
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	options "github.com/outscale/osc-sdk-go/v3/pkg/utils"
	"github.com/spf13/cobra"
)

// oapiCmd represents the oapi command
var oapiCmd = &cobra.Command{
	Use:   "oapi",
	Short: "OAPI calls",
	Long:  `Send commands to the Outscale API.`,
}

func init() {
	rootCmd.AddCommand(oapiCmd)

	c := reflect.TypeOf(&osc.Client{})
	for i := range c.NumMethod() {
		m := c.Method(i)
		if m.Type.NumIn() != 4 || m.Type.NumOut() != 2 || strings.HasSuffix(m.Name, "Raw") {
			continue
		}
		cmd := &cobra.Command{
			Use: m.Name,
			Run: func(cmd *cobra.Command, args []string) {
				oapi(cmd)
			},
		}
		arg := m.Type.In(2)
		flags.FromStruct(cmd, arg, "")
		oapiCmd.AddCommand(cmd)
	}
}

func oapi(cmd *cobra.Command) {
	debug.Println(cmd.Name() + " called")

	p, err := profile.New()
	if err != nil {
		errors.ExitErr(err)
	}
	ua := "gli/" + version.Version
	opts := []middleware.MiddlewareChainOption{options.WithUseragent(ua)}
	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		opts = append(opts, options.WithLogging(sdk.VerboseLogger{}))
	} else {
		opts = append(opts, options.WithoutLogging())
	}
	cl, err := osc.NewClient(p, opts...)
	if err != nil {
		errors.ExitErr(err)
	}
	clt := reflect.TypeOf(cl)
	m, _ := clt.MethodByName(cmd.Name())
	argType := m.Type.In(2)
	arg := reflect.New(argType)
	err = flags.ToStruct(cmd, arg, "")
	if err != nil {
		errors.ExitErr(err)
	}
	ctx := context.Background()
	res := reflect.ValueOf(cl).MethodByName(cmd.Name()).Call([]reflect.Value{
		reflect.ValueOf(ctx),
		arg.Elem(),
	})
	debug.Println(len(res), "results")
	if !res[1].IsNil() {
		err = res[1].Interface().(error)
		if oerr := osc.AsErrorResponse(err); oerr != nil {
			enc := json.NewEncoder(os.Stderr)
			enc.SetIndent("", "  ")
			_ = enc.Encode(oerr)
			os.Exit(1)
		}
		errors.ExitErr(err)
	}
	out, err := output.NewFromFlags(cmd.Flags())
	if err != nil {
		errors.ExitErr(err)
	}
	_ = out.Output(ctx, res[0].Interface())
}
