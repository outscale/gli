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
	"github.com/outscale/octl/cmd/prerun"
	"github.com/outscale/octl/pkg/builder"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/flags"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/preferences"
	"github.com/outscale/octl/pkg/runner"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var oksCmd = &cobra.Command{
	GroupID:           "services",
	Use:               "kube",
	Short:             "OUTSCALE Kubernetes as a Service (OKS) management",
	Aliases:           []string{"oks"},
	PersistentPreRunE: flagNamesToID,
}

var projectUseCmd = &cobra.Command{
	Use:   "use [id_or_name]",
	Short: "Set a default project for cluster commands, reset it without args",
	Run:   useProject,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(oksCmd)
	b := builder.NewBuilder[oks.Client]("kube", "https://docs.outscale.com/oks.html")
	b.BuildAPI(oksCmd, func(m reflect.Method) bool {
		return m.Type.NumIn() >= 3 && m.Type.NumOut() == 2 && !strings.HasSuffix(m.Name, "Raw") &&
			!strings.HasSuffix(m.Name, "WithBody")
	}, kube)
	b.Build(oksCmd, nil)

	// Add --project flag to kube cluster commands
	clusterCmd, _ := lo.Find(oksCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "cluster" })
	clusterCmd.PersistentFlags().String("project", "", "Name or ID of project")
	_ = flags.MarkAsNoForward(clusterCmd.PersistentFlags(), "project")
	clusterCmd.PersistentPreRunE = clusterArgToID
	clusterCreateCmd, _ := lo.Find(clusterCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "create" })
	_ = flags.SetDefault(clusterCreateCmd.Flags(), "project", "")

	// Remap <cluster ls --project name> to <project clusters name>
	projectCmd, _ := lo.Find(oksCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "project" })
	projectCmd.PersistentPreRunE = projectArgToID
	projectClustersCmd, _ := lo.Find(projectCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "clusters" })
	clusterListCmd, _ := lo.Find(clusterCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "list" })
	runClusterListCmd := clusterListCmd.Run
	clusterListCmd.Run = func(cmd *cobra.Command, args []string) {
		// remap flag to arg
		project, _ := cmd.Flags().GetString("project")
		if project != "" {
			cmd.Flag("project").Changed = false
			projectClustersCmd.Run(cmd, []string{project})
		} else {
			runClusterListCmd(cmd, nil)
		}
	}

	// cluster/project use
	projectCmd.AddCommand(projectUseCmd)
}

func kube(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")
	p := loadProfile(cmd)
	cl, err := oks.NewClient(p, sdkOptions(cmd)...)
	if err == nil {
		err = runner.Run[*oks.Client, *oks.ErrorResponse](cmd, args, cl, config.For("kube"))
	}
	if err != nil {
		messages.ExitErr(err)
	}
}

func clusterArgToID(cmd *cobra.Command, args []string) error {
	debug.Println("clusterArgToID")
	p := loadProfile(cmd)
	cl, err := oks.NewClient(p, sdkOptions(cmd)...)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		debug.Println("no arg to replace")
		return nil
	}
	if _, err := uuid.Parse(args[0]); err == nil {
		debug.Println("arg is an uuid")
		return nil
	}
	project, _ := cmd.Flags().GetString("project")
	args[0], err = clusterNameToID(cmd.Context(), args[0], project, cl)
	return err
}

func projectArgToID(cmd *cobra.Command, args []string) error {
	// project use should store the name, not the ID
	if cmd.Name() == "use" {
		return nil
	}

	debug.Println("projectArgToID")
	p := loadProfile(cmd)
	cl, err := oks.NewClient(p, sdkOptions(cmd)...)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		debug.Println("no arg to replace")
		return nil
	}
	if _, err := uuid.Parse(args[0]); err == nil {
		debug.Println("arg is an uuid")
		return nil
	}
	args[0], err = projectNameToID(cmd.Context(), args[0], cl)
	return err
}

func flagNamesToID(cmd *cobra.Command, args []string) error {
	debug.Println("flagNamesToID")
	p := loadProfile(cmd)
	cl, err := oks.NewClient(p, sdkOptions(cmd)...)
	if err != nil {
		return err
	}

	pf := cmd.Flags().Lookup("project")
	if pf == nil {
		debug.Println("no project flag to replace")
		return nil
	}
	pid, err := projectNameToID(cmd.Context(), pf.Value.String(), cl)
	if err != nil {
		return err
	}
	_ = pf.Value.Set(pid)

	cf := cmd.Flags().Lookup("cluster")
	if cf == nil {
		debug.Println("no cluster flag to replace")
		return nil
	}
	cid, err := clusterNameToID(cmd.Context(), cf.Value.String(), pid, cl)
	if err != nil {
		return err
	}
	_ = cf.Value.Set(cid)
	return nil
}

func projectNameToID(ctx context.Context, name string, cl *oks.Client) (string, error) {
	if name == "" {
		name = prerun.PreferencesFrom(ctx).Kube.DefaultProject
	}
	if name == "" {
		return "", nil
	}
	pjs, err := cl.ListProjects(ctx, &oks.ListProjectsParams{Name: &name})
	if err != nil {
		return "", err
	}
	switch len(pjs.Projects) {
	case 0:
		return "", fmt.Errorf("project %q not found", name)
	default:
		messages.Info("Resolving project name %q as ID %q", name, pjs.Projects[0].Id)
		return pjs.Projects[0].Id, nil
	}
}

func clusterNameToID(ctx context.Context, name, project string, cl *oks.Client) (string, error) {
	if name == "" {
		return "", nil
	}
	var (
		cs  *oks.ClusterResponseList
		err error
	)
	if project != "" {
		var pid string
		if _, err := uuid.Parse(project); err == nil {
			pid = project
			debug.Println("project is an uuid")
		} else {
			pid, err = projectNameToID(ctx, project, cl)
			if err != nil {
				return "", err
			}
		}
		cs, err = cl.ListClustersByProjectID(ctx, &oks.ListClustersByProjectIDParams{Name: &name, ProjectId: &pid})
	} else {
		cs, err = cl.ListAllClusters(ctx, &oks.ListAllClustersParams{Name: &name})
	}
	if err != nil {
		return "", err
	}
	clusters := cs.Clusters
	switch len(clusters) {
	case 0:
		return "", fmt.Errorf("cluster %q not found", name)
	case 1:
		messages.Info("Resolving cluster name %q as ID %q", name, clusters[0].Id)
		return clusters[0].Id, nil
	default:
		return "", fmt.Errorf("multiple clusters found with the name %q", name)
	}
}

func useProject(cmd *cobra.Command, args []string) {
	var def string
	if len(args) > 0 {
		def = args[0]
	}
	prof, _ := cmd.Flags().GetString("profile")
	if prof == "" {
		prof = profile.DefaultProfile
	}
	// retro compatibility
	_ = preferences.SetGlobal(func(prefs *preferences.Preferences) { prefs.Kube.DefaultProject = "" })
	err := preferences.Set(prof, func(prefs *preferences.Preferences) { prefs.Kube.DefaultProject = def })
	if err != nil {
		messages.ExitErr(err)
	}
	if def == "" {
		messages.Success("The default project has been reset for profile %q", prof)
	} else {
		messages.Success("%q is now the default project for profile %q", def, prof)
	}
}
