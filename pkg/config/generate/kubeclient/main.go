/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	oksv1beta "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta"
	oksv1beta2 "github.com/outscale/goutils/oks/clientset/typed/oks.dev/v1beta2"
	configbuilder "github.com/outscale/octl/pkg/builder/config"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/messages"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	version := os.Args[3]
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
		AllowedNumOut: []int{1, 2},
		SkipMethods:   []string{"DeleteCollection", "Patch", "UpdateStatus", "Watch"},
		SkipAlias:     []string{"List", "Get", "Create", "Delete"},
		SkipFlagsPrefixes: map[string][]string{
			"*": {
				"TypeMeta",
				"ObjectMeta.ManagedFields", "ObjectMeta.OwnerReferences", "ObjectMeta.CreationTimestamp", "ObjectMeta.Generation",
				"Status", "Watch", "SendInitialEvents", "AllowWatchBookmarks", "Continue", "IgnoreStoreReadErrorWithClusterBreakingPotential",
			},
		},
		IdFieldInUsage: "id",
	}

	sb := configbuilder.NewSpecBuilder(cfg)
	spec := sb.BuildSpec(&base, "github.com/outscale/goutils/oks/apis/oks.dev/"+version)

	switch {
	case strings.Contains(src, "nodepool"):
		b := configbuilder.NewClientBuilder[oksv1beta2.NodePoolInterface](cfg)
		b.BuildFor(&base, spec)
	case strings.Contains(src, "netpeering_request"):
		b := configbuilder.NewClientBuilder[oksv1beta.NetPeeringRequestInterface](cfg)
		b.BuildFor(&base, spec)
	case strings.Contains(src, "netpeering_acceptance"):
		b := configbuilder.NewClientBuilder[oksv1beta.NetPeeringAcceptanceInterface](cfg)
		b.BuildFor(&base, spec)
	case strings.Contains(src, "netpeering"):
		b := configbuilder.NewClientBuilder[oksv1beta.NetPeeringInterface](cfg)
		b.BuildFor(&base, spec)
	case strings.Contains(src, "ippool"):
		b := configbuilder.NewClientBuilder[oksv1beta.IPPoolInterface](cfg)
		b.BuildFor(&base, spec)
	case strings.Contains(src, "oosaccess"):
		b := configbuilder.NewClientBuilder[oksv1beta.OOSAccessInterface](cfg)
		b.BuildFor(&base, spec)
	case strings.Contains(src, "vpnconnection"):
		b := configbuilder.NewClientBuilder[oksv1beta.VpnConnectionInterface](cfg)
		b.BuildFor(&base, spec)
	}
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
