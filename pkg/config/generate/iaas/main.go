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
	"github.com/outscale/octl/pkg/flags"
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
	if base.API == nil {
		base.API = map[string]config.APICall{}
	}
	if base.Entities == nil {
		base.Entities = map[string]config.Entity{}
	}

	cfg := configbuilder.Config{
		SkipSuffixes: []string{"Raw", "WithBody"},
		SkipAlias: []string{
			"ReadConsoleOutput",
			"ReadFlexibleGpuCatalog",
			"ReadVmsHealth",
			"ReadVmsState",
			"ReadNetAccessPointServices",
			"DeletePolicy",
			"CreatePolicyVersion",
			"DeletePolicyVersion",
			"ReadPolicyVersion",
			"ReadPolicyVersions",
			"ReadUserGroup",
			"ReadVmsStopHistory",
		},
		AllowedNumOut: []int{1, 2},
		TypeName: map[string]string{
			"DhcpOption":        "DhcpOptionsSet",
			"UserGroupsPerUser": "UserGroup",
		},
		InputStructSuffix: "Request",
		ReadFlagPrefixes:  []string{"Filters."},
		SkipFlagsPrefixes: map[string][]string{
			"*": {
				"dry-run",
				"NextPageToken",
				"ResultsPerPage",
			},
		},
		StripFlagsPrefixes: map[string][]string{
			"securitygrouprule": {
				"security-",
			},
		},
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
			"PublicKey": {
				CustomValue: flags.Base64File,
				Help:        "The file storing the public key to import in your account, if you are importing an existing keypair.",
			},
			"UserData": {
				CustomValue: flags.Base64File,
				Help:        "The file storing the data or script used to add a specific configuration to the VM (max size 500 KiB).",
			},
			"PolicyDocument": {
				CustomValue: flags.FileOrJSON,
				Help:        "Either a file storing the policy document, or the policy document (in JSON format).",
			},
			"Document": {
				CustomValue: flags.FileOrJSON,
				Help:        "Either a file storing the policy document, or the policy document (in JSON format).",
			},
		},
		FlagReplaces: []string{
			"block-device-mapping-bsu", "volume",
			"block-device-mapping", "volume",
			"subregion-name", "subregion",
		},
		RequiredFromFieldPointer: true,
	}

	sb := configbuilder.NewSpecBuilder(cfg)
	spec := sb.BuildSpec(&base, "github.com/outscale/osc-sdk-go/v3/pkg/osc")

	b := configbuilder.NewClientBuilder[*osc.Client](cfg)
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
