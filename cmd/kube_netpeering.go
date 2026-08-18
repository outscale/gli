package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var netpeeringCmd = &cobra.Command{
	GroupID: "service",
	Use:     "netpeering",
	Short:   "netpeering commands",
}

func init() {
	buildKubeAPI("kubeclient_netpeering", netpeeringCmd, oksCmd, func(client *clientset.Clientset) oksv1beta.NetPeeringInterface {
		return client.OksV1beta().NetPeerings()
	})
}
