package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta2 "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta2"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/runner"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

func kubeapi(provider string) func(cmd *cobra.Command, args []string) {
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

		err = runner.Run[oksv1beta2.NodePoolInterface, *osc.ErrorResponse](cmd, args, client.OksV1beta2().NodePools(), config.For(provider))
		if err != nil {
			messages.ExitErr(err)
		}
	}
}
