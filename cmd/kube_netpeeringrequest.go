package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var netpeeringrequestCmd = &cobra.Command{
	Use:   "request",
	Short: "Manage Netpeering Request resources",
}

func init() {
	buildKubeAPI("kubeclient_netpeering_request", netpeeringrequestCmd, netpeeringCmd, func(client *clientset.Clientset) oksv1beta.NetPeeringRequestInterface {
		return client.OksV1beta().NetPeeringRequests()
	})
}
