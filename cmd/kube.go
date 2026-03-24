/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
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
	GroupID: "services",
	Use:     "kube",
	Short:   "OUTSCALE Kubernetes as a Service (OKS) management",
	Aliases: []string{"oks"},
}

func init() {
	rootCmd.AddCommand(oksCmd)
	b := builder.NewBuilder[oks.Client]("kube", "https://docs.outscale.com/oks.html")
	b.BuildAPI(oksCmd, func(m reflect.Method) bool {
		return m.Type.NumIn() >= 3 && m.Type.NumOut() == 2 && !strings.HasSuffix(m.Name, "Raw") &&
			!strings.HasSuffix(m.Name, "WithBody")
	}, kube)
	b.Build(oksCmd, nil)

	oksCmd.AddCommand(kubectlCmd)
}

func kube(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")
	p := loadProfile(cmd)
	cl, err := oks.NewClient(p, sdkOptions(cmd)...)
	if err == nil {
		args = argNameToID(cmd, args, cl)
		flagNameToID(cmd, cl)
		err = runner.Run[*oks.Client, *oks.ErrorResponse](cmd, args, cl, config.For("kube"))
	}
	if err != nil {
		messages.ExitErr(err)
	}
}

func argNameToID(cmd *cobra.Command, args []string, cl *oks.Client) []string {
	if len(args) == 0 {
		debug.Println("no arg to replace")
		return args
	}
	if _, err := uuid.Parse(args[0]); err == nil {
		debug.Println("arg is an uuid")
		return args
	}
	var err error
	switch cmd.Name() {
	case "GetProject", "DeleteProject", "GetProjectNets", "GetProjectQuotas", "GetProjectPublicIps", "GetProjectSnapshots":
		args[0], err = projectNameToID(cmd.Context(), args[0], cl)
	case "GetCluster", "UpdateCluster", "DeleteCluster", "GetKubeconfig":
		args[0], err = clusterNameToID(cmd.Context(), args[0], cl)
	}
	if err != nil {
		messages.ExitErr(err)
	}
	return args
}

func flagNameToID(cmd *cobra.Command, cl *oks.Client) {
	f := cmd.Flags().Lookup("ProjectId")
	if f == nil {
		debug.Println("no flag to replace")
		return
	}
	id, err := projectNameToID(cmd.Context(), f.Value.String(), cl)
	if err != nil {
		messages.ExitErr(err)
	}
	_ = f.Value.Set(id)
}

func projectNameToID(ctx context.Context, name string, cl *oks.Client) (string, error) {
	pjs, err := cl.ListProjects(ctx, &oks.ListProjectsParams{Name: &name})
	if err != nil {
		return "", err
	}
	switch len(pjs.Projects) {
	case 0:
		return "", fmt.Errorf("project %q not found", name)
	default:
		debug.Println("replacing", name, "by", pjs.Projects[0].Id)
		return pjs.Projects[0].Id, nil
	}
}

func clusterNameToID(ctx context.Context, name string, cl *oks.Client) (string, error) {
	cs, err := cl.ListAllClusters(ctx, &oks.ListAllClustersParams{Name: &name})
	if err != nil {
		return "", err
	}
	switch len(cs.Clusters) {
	case 0:
		return "", fmt.Errorf("cluster %q not found", name)
	case 1:
		debug.Println("replacing", name, "by", cs.Clusters[0].Id)
		return cs.Clusters[0].Id, nil
	default:
		return "", fmt.Errorf("multiple clusters found with the name %q", name)
	}
}
