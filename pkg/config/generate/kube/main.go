/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"os"

	"github.com/goccy/go-yaml"
	configbuilder "github.com/outscale/octl/pkg/builder/config"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	var base config.Config
	data, err := os.ReadFile(src) //nolint:gosec
	if err != nil {
		messages.ExitErr(err)
	}
	err = yaml.Unmarshal(data, &base)
	if err != nil {
		messages.ExitErr(err)
	}
	if base.API == nil {
		base.API = map[string]config.APICall{}
	}
	if base.Entities == nil {
		base.Entities = map[string]config.Entity{}
	}

	cfg := configbuilder.Config{
		SkipSuffixes: []string{"Raw", "WithBody"},
		SkipAlias: []string{
			"GetCPSubregions",
			"GetClusterTemplate",
			"GetControlPlanePlans",
			"GetKubeconfig",
			"GetKubeconfigWithPubkeyNACL",
			"GetKubernetesVersions",
			"GetNetPeeringAcceptanceTemplate",
			"GetNetPeeringRequestTemplate",
			"GetNodepoolTemplate",
			"GetProjectNets",
			"GetProjectQuotas",
			"GetProjectTemplate",
			"GetQuotas",
		},
		AllowedNumOut:            []int{1, 2},
		InputStructSuffix:        "Request",
		RequiredFromFieldPointer: true,
		FlagOverrides: map[string]config.Flag{
			"tag": {
				Name: "tags",
				Help: "Tags (key=value,key=value)",
			},
		},
		IdFieldInUsage: "id_or_name",
	}

	sb := configbuilder.NewSpecBuilder(cfg)
	spec := sb.BuildSpec(&base, "github.com/outscale/osc-sdk-go/v3/pkg/oks")

	b := configbuilder.NewClientBuilder[*oks.Client](cfg)
	b.BuildFor(&base, spec)

	fd, err := os.Create(dst) //nolint:gosec
	if err != nil {
		messages.ExitErr(err)
	}
	err = yaml.NewEncoder(fd, yaml.UseSingleQuote(true), yaml.UseLiteralStyleIfMultiline(true)).Encode(base)
	if err != nil {
		messages.ExitErr(err)
	}
	err = fd.Close()
	if err != nil {
		messages.ExitErr(err)
	}
}
