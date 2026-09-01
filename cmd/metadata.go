/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"reflect"

	"github.com/outscale/goutils/sdk/metadata"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/output"
	"github.com/outscale/octl/pkg/output/read"
	"github.com/spf13/cobra"
)

// metadataCmd represents the metadata command
var metadataCmd = &cobra.Command{
	GroupID: "services",
	Use:     "metadata",
	Aliases: []string{"meta"},
	Short:   "Metadata server queries",
	Long:    `Query the metadata server`,
}

var metas = map[string]string{
	"region":            "GetRegion",
	"availability_zone": "GetSubregion",
	"instance_id":       "GetInstanceID",
	"instance_type":     "GetInstanceType",
	"device_mapping":    "GetDeviceMappings",
	"mac":               "GetMAC",
	"tags":              "GetTags",
	"placement_cluster": "GetPlacementCluster",
	"placement_server":  "GetPlacementServer",
}

func init() {
	rootCmd.AddCommand(metadataCmd)
	metadataCmd.AddCommand(&cobra.Command{
		Use:   "all",
		Short: "query all metadata from the metadata server",
		Run:   runMetadata("Fetch"),
	})
	for cmd, call := range metas {
		metadataCmd.AddCommand(&cobra.Command{
			Use:   cmd,
			Short: "query " + cmd + " from the metadata server",
			Run:   runMetadata(call),
		})
	}
}

func runMetadata(fn string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		_, out, err := output.NewFromFlags(cmd.Flags(), "yaml", "", nil, false, false)
		if err == nil {
			call := read.FetchPage{
				Method: reflect.ValueOf(metadata.DefaultService).MethodByName(fn),
				Args:   []reflect.Value{reflect.ValueOf(cmd.Context())},
			}
			err = out.Output(cmd.Context(), call)
		}
		if err != nil {
			messages.ExitErr(err)
		}
	}
}
