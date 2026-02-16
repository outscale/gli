/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/config/generate/iaas/builder"
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
	var client *osc.Client
	ct := reflect.TypeOf(client)
	for i := range ct.NumMethod() {
		m := ct.Method(i)
		if strings.HasSuffix(m.Name, "Raw") || strings.HasSuffix(m.Name, "WithBody") || m.Type.NumOut() != 2 {
			continue
		}
		if strings.HasPrefix(m.Name, "Read") {
			build(&base, m, "Read")
		}
	}
	for i := range ct.NumMethod() {
		m := ct.Method(i)
		if strings.HasSuffix(m.Name, "Raw") || strings.HasSuffix(m.Name, "WithBody") || m.Type.NumOut() != 2 {
			continue
		}
		if strings.HasPrefix(m.Name, "Delete") {
			build(&base, m, "Delete")
		}
		if strings.HasPrefix(m.Name, "Update") {
			build(&base, m, "Update")
		}
		if strings.HasPrefix(m.Name, "Create") {
			build(&base, m, "Create")
		}
	}
	fd, err := os.Create(dst) //nolint:gosec
	if err != nil {
		messages.ExitErr(err)
	}
	err = yaml.NewEncoder(fd, yaml.UseSingleQuote(true)).Encode(base)
	if err != nil {
		messages.ExitErr(err)
	}
}

func build(cfg *config.Config, m reflect.Method, prefix string) {
	b := builder.New(cfg, m, prefix)
	err := b.Build()
	switch {
	case errors.Is(err, builder.ErrCantBuild):
	case err != nil:
		messages.ExitErr(err)
	default:
		b.Commit()
	}
}
