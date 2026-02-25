/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"os"

	"github.com/goccy/go-yaml"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/config/generate/builder"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
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
	if base.Calls == nil {
		base.Calls = map[string]config.Call{}
	}
	if base.Entities == nil {
		base.Entities = map[string]config.Entity{}
	}
	if base.Spec.Calls == nil {
		base.Spec.Calls = map[string]config.SpecCall{}
	}
	if base.Spec.Attributes == nil {
		base.Spec.Attributes = map[string]config.SpecAttribute{}
	}

	cfg := builder.Config{
		InputStructSuffix: "Request",
		ReadFlagPrefix:    "Filters.",
		SkipFlags:         []string{"DryRun", "NextPageToken", "ResultsPerPage"},
		PriorityFields: []string{
			"State",
			"PublicIp",
			"PrivateIp",
			"NetId",
			"SubnetId",
			"IpRange",
			"SubregionName",
			"SubregionNames",
			"Subregion",
			"Subregions",
			"Size",
			"Iops",
			"Email",
		},
		FlagOverrides: map[string]config.Flag{
			"public-key": {
				Type:  "base64File",
				Usage: "The file storing the public key to import in your account, if you are importing an existing keypair.",
			},
			"user-data": {
				Type:  "base64File",
				Usage: "The file storing the data or script used to add a specific configuration to the VM (max size 500 KiB).",
			},
			"policy-document": {
				Type:  "file",
				Usage: "The file storing the policy document, corresponding to a JSON string that contains the policy.",
			},
			"document": {
				Type:  "file",
				Usage: "The file storing the policy document, corresponding to a JSON string that contains the policy.",
			},
		},
		FlagReplaces: []string{
			"block-device-mapping-bsu", "volume",
			"block-device-mapping", "volume",
		},
		RequiredFromFieldPointer: true,
	}

	sb := builder.NewSpecBuilder(cfg)
	sb.BuildSpec(&base, "github.com/outscale/osc-sdk-go/v3/pkg/osc")

	var client *osc.Client
	b := builder.NewClientBuilder(cfg)
	b.BuildFor(&base, client)

	fd, err := os.Create(dst) //nolint:gosec
	if err != nil {
		messages.ExitErr(err)
	}
	err = yaml.NewEncoder(fd, yaml.UseSingleQuote(true), yaml.UseLiteralStyleIfMultiline(true)).Encode(base)
	if err != nil {
		messages.ExitErr(err)
	}
}
