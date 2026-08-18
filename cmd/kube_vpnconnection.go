package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var vpnconnectionCmd = &cobra.Command{
	GroupID: "service",
	Use:     "vpnconnection",
	Short:   "vpnconnection commands",
}

func init() {
	buildKubeAPI("kubeclient_vpnconnection", vpnconnectionCmd, oksCmd, func(client *clientset.Clientset) oksv1beta.VpnConnectionInterface {
		return client.OksV1beta().VpnConnections()
	})
}
