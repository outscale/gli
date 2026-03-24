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
		RequiredFromFieldPointer: true,
	}

	sb := builder.NewSpecBuilder(cfg)
	sb.BuildSpec(&base, "github.com/outscale/goutils/oks/apis/oks.dev/v1beta2")

	// Aliases are too specific, we do not generate them automatically
	// b := builder.NewClientBuilder[oksv1beta2.NodePoolInterface](cfg)
	// b.BuildFor(&base)

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
