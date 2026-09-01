package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var netpeeringacceptanceCmd = &cobra.Command{
	Use:   "acceptance",
	Short: "Manage Netpeering Acceptance resources",
}

func init() {
	buildKubeAPI("kubeclient_netpeering_acceptance", netpeeringacceptanceCmd, netpeeringCmd, func(client *clientset.Clientset) oksv1beta.NetPeeringAcceptanceInterface {
		return client.OksV1beta().NetPeeringAcceptances()
	})
}
