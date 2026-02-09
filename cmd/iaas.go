/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/outscale/gli/pkg/builder"
	"github.com/outscale/gli/pkg/config"
	"github.com/outscale/gli/pkg/debug"
	"github.com/outscale/gli/pkg/errors"
	"github.com/outscale/gli/pkg/runner"
	"github.com/outscale/gli/pkg/sdk"
	"github.com/outscale/gli/pkg/version"
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	options "github.com/outscale/osc-sdk-go/v3/pkg/utils"
	"github.com/spf13/cobra"
)

// iaasCmd represents the iaascommand
var iaasCmd = &cobra.Command{
	Use:   "iaas",
	Short: "IaaS management",
}

func init() {
	rootCmd.AddCommand(iaasCmd)
	spec, err := osc.GetSwagger()
	if err != nil {
		errors.Warn(fmt.Sprintf("⚠️ unable to load OpenAPI spec: %v", err))
	}
	b := builder.NewBuilder[osc.Client]("iaas", spec)
	b.BuildAPI(iaasCmd, func(m reflect.Method) bool {
		return m.Type.NumIn() == 4 && m.Type.NumOut() == 2 && !strings.HasSuffix(m.Name, "Raw")
	}, func(cmd *cobra.Command, _ []string) {
		oapi(cmd)
	})
	b.Build(iaasCmd)
}

func oapi(cmd *cobra.Command) {
	debug.Println(cmd.Name() + " called")
	path, _ := cmd.Flags().GetString("config")
	prof, _ := cmd.Flags().GetString("profile")
	p, err := profile.NewFrom(path, prof)
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
	if err == nil {
		err = runner.Run[osc.Client, *osc.ErrorResponse](cmd, cl, config.For("iaas"))
	}
	if err != nil {
		errors.ExitErr(err)
	}
}
