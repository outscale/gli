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
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
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
		messages.Warn("Unable to load OpenAPI spec: %v", err)
	}
	b := builder.NewBuilder[osc.Client]("iaas", spec)
	b.BuildAPI(iaasCmd, func(m reflect.Method) bool {
		return m.Type.NumIn() == 4 && m.Type.NumOut() == 2 && !strings.HasSuffix(m.Name, "Raw")
	}, oapi)
	b.Build(iaasCmd)
}

func oapi(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")
	p := loadProfile(cmd)
	cl, err := osc.NewClient(p, sdkOptions(cmd)...)
	if err == nil {
		err = runner.Run[osc.Client, *osc.ErrorResponse](cmd, args, cl, config.For("iaas"))
	}
	if err != nil {
		messages.ExitErr(err)
	}
}
