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
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	options "github.com/outscale/osc-sdk-go/v3/pkg/utils"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var oksCmd = &cobra.Command{
	Use:   "kube",
	Short: "Kubernetes as a Service (OKS) management",
}

func init() {
	rootCmd.AddCommand(oksCmd)
	spec, err := oks.GetSwagger()
	if err != nil {
		errors.Warn(fmt.Sprintf("⚠️ unable to load OpenAPI spec: %v", err))
	}
	b := builder.NewBuilder[oks.Client]("kube", spec)
	b.BuildAPI(oksCmd, func(m reflect.Method) bool {
		return m.Type.NumIn() >= 3 && m.Type.NumOut() == 2 && !strings.HasSuffix(m.Name, "Raw") &&
			!strings.HasSuffix(m.Name, "WithBody")
	}, func(cmd *cobra.Command, _ []string) {
		kube(cmd)
	})
	b.Build(oksCmd)
}

func kube(cmd *cobra.Command) {
	debug.Println(cmd.Name() + " called")
	path, _ := cmd.Flags().GetString("config")
	prof, _ := cmd.Flags().GetString("profile")
	p, err := profile.NewFrom(prof, path)
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
	cl, err := oks.NewClient(p, opts...)
	if err == nil {
		err = runner.Run[oks.Client, *oks.ErrorResponse](cmd, cl, config.For("kube"))
	}
	if err != nil {
		errors.ExitErr(err)
	}
}
