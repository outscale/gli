package cmd

import (
	"reflect"
	"slices"

	oksv1beta2 "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta2"
	"github.com/outscale/octl/pkg/builder"
	"github.com/samber/lo"
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
	oksCmd.AddCommand(nodepoolCmd)
	b := builder.NewBuilder[oksv1beta2.NodePoolInterface]("kubeclient_nodepool", "https://docs.outscale.com/api.html")
	b.BuildAPI(nodepoolCmd, func(m reflect.Method) bool {
		return slices.Contains([]string{"List", "Get", "Create", "Update", "Delete"}, m.Name)
	}, kubeapi("kubeclient_nodepool"))
	apiCmd, _ := lo.Find(nodepoolCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "api" })
	apiCmd.PersistentFlags().String("cluster", "", "Name or ID of cluster")
	_ = apiCmd.MarkPersistentFlagRequired("cluster")
	// nodepool commands need to be added to the upper level, otherwise we will get kube nodepool nodepool
	b.Build(oksCmd, apiCmd)
}
