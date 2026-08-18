package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta2 "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta2"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var nodepoolCmd = &cobra.Command{
	GroupID: "service",
	Use:     "nodepool",
	Short:   "nodepool commands",
	Aliases: []string{"np"},
}

func init() {
	buildKubeAPI("kubeclient_nodepool", nodepoolCmd, oksCmd, func(client *clientset.Clientset) oksv1beta2.NodePoolInterface {
		return client.OksV1beta2().NodePools()
	})
}
