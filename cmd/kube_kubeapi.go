package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	commandbuilder "github.com/outscale/octl/pkg/builder/command"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/flags"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/runner"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

func buildKubeAPI[Client any](provider string, cmd, parent *cobra.Command, getclient func(client *clientset.Clientset) Client) {
	parent.AddCommand(cmd)
	b := commandbuilder.NewBuilder(provider, "https://docs.outscale.com/api.html")
	b.Build(cmd, runKubeAPI(func(cmd *cobra.Command, args []string, client *clientset.Clientset) error {
		return runner.Run[Client, *osc.ErrorResponse](cmd, args, getclient(client), config.For(provider))
	}))
	// Add --cluster flag to api commands (no --project as there is no name to id conversion here)
	apiCmd, _ := lo.Find(cmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "api" })
	apiCmd.PersistentFlags().String("cluster", "", "[REQUIRED] ID of cluster")
	_ = apiCmd.MarkPersistentFlagRequired("cluster")
	// Add --cluster and --project to high level commands
	for _, child := range cmd.Commands() {
		if child.Name() != "api" {
			child.Flags().String("cluster", "", "[REQUIRED] Name or ID of cluster")
			child.Flags().String("project", "", "Name or ID of project")
			_ = child.MarkFlagRequired("cluster")
			_ = flags.MarkAsNoForward(child.Flags(), "project")
		}
	}
}

func runKubeAPI(fn func(cmd *cobra.Command, args []string, client *clientset.Clientset) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		p := loadProfile(cmd)
		cl, err := oks.NewClient(p, sdkOptions(cmd)...)
		if err != nil {
			messages.ExitErr(err)
		}
		cluster, _ := cmd.Flags().GetString("cluster")
		kubeconfig, err := getKubeconfig(cmd.Context(), cluster, cl)
		if err != nil {
			messages.ExitErr(err)
		}

		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			messages.ExitErr(err)
		}
		client, err := clientset.NewForConfig(cfg)
		if err != nil {
			messages.ExitErr(err)
		}

		err = fn(cmd, args, client)
		if err != nil {
			messages.ExitErr(err)
		}
	}
}
