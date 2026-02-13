/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"reflect"
	"strings"

	"github.com/outscale/octl/pkg/builder"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/runner"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
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
		messages.Warn("Unable to load OpenAPI spec: %v", err)
	}
	b := builder.NewBuilder[oks.Client]("kube", spec)
	b.BuildAPI(oksCmd, func(m reflect.Method) bool {
		return m.Type.NumIn() >= 3 && m.Type.NumOut() == 2 && !strings.HasSuffix(m.Name, "Raw") &&
			!strings.HasSuffix(m.Name, "WithBody")
	}, kube)
	b.Build(oksCmd)
}

func kube(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")
	p := loadProfile(cmd)
	cl, err := oks.NewClient(p, sdkOptions(cmd)...)
	if err == nil {
		err = runner.Run[oks.Client, *oks.ErrorResponse](cmd, args, cl, config.For("kube"))
	}
	if err != nil {
		messages.ExitErr(err)
	}
}
