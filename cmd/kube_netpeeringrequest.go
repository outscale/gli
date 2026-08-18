package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var netpeeringrequestCmd = &cobra.Command{
	GroupID: "service",
	Use:     "request",
	Short:   "netpeering request commands",
}

func init() {
	buildKubeAPI("kubeclient_netpeeringrequests", netpeeringrequestCmd, netpeeringCmd, func(client *clientset.Clientset) oksv1beta.NetPeeringRequestInterface {
		return client.OksV1beta().NetPeeringRequests()
	})
}
