package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var ippoolCmd = &cobra.Command{
	GroupID: "service",
	Use:     "ippool",
	Short:   "ippool commands",
}

func init() {
	buildKubeAPI("kubeclient_ippool", ippoolCmd, oksCmd, func(client *clientset.Clientset) oksv1beta.IPPoolInterface {
		return client.OksV1beta().IPPools()
	})
}
