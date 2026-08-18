package cmd

import (
	"github.com/outscale/goutils/oks/clientset"
	oksv1beta "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta"
	"github.com/spf13/cobra"
)

// oksCmd represents the kubecommand
var oosaccessCmd = &cobra.Command{
	GroupID: "service",
	Use:     "oosaccess",
	Short:   "oosaccess commands",
}

func init() {
	buildKubeAPI("kubeclient_oosaccess", oosaccessCmd, oksCmd, func(client *clientset.Clientset) oksv1beta.OOSAccessInterface {
		return client.OksV1beta().OOSAccesses()
	})
}
