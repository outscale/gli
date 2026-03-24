/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"reflect"
	"strings"
	"time"

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
	GroupID: "services",
	Use:     "iaas",
	Short:   "OUTSCALE IaaS management",
}

func init() {
	rootCmd.AddCommand(iaasCmd)
	b := builder.NewBuilder[osc.Client]("iaas", "https://docs.outscale.com/api.html")
	b.BuildAPI(iaasCmd, func(m reflect.Method) bool {
		return m.Type.NumIn() == 4 && m.Type.NumOut() == 2 && !strings.HasSuffix(m.Name, "Raw")
	}, oapi)
	b.Build(iaasCmd, nil)

	cmd, _, err := iaasCmd.Find([]string{"net"})
	if err != nil {
		panic(err)
	}
	cmd.AddCommand(teardownCmd)
	teardownCmd.Flags().Duration("timeout", 10*time.Minute, "Timeout for a single resource deletion")
	teardownCmd.Flags().Bool("teardown-vms", false, "Tears down VM in net")
	cmd.AddCommand(depsCmd)
}

func oapi(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")
	p := loadProfile(cmd)
	cl, err := osc.NewClient(p, sdkOptions(cmd)...)
	if err == nil {
		err = runner.Run[*osc.Client, *osc.ErrorResponse](cmd, args, cl, config.For("iaas"))
	}
	if err != nil {
		messages.ExitErr(err)
	}
}
