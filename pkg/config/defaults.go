/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package config

import (
	"embed"

	"github.com/goccy/go-yaml"
	"github.com/outscale/gli/pkg/errors"
)

//go:generate go run generate/iaas/main.go generate/iaas/defaults.yaml defaults_iaas.yaml
//go:embed defaults_iaas.yaml
var f embed.FS

func Defaults() Configs {
	defaults := Configs{}
	for _, provider := range []string{"iaas"} {
		data, _ := f.ReadFile("defaults_" + provider + ".yaml")
		var cfg Config
		err := yaml.Unmarshal(data, &cfg)
		if err != nil {
			errors.ExitErr(err)
		}
		defaults[provider] = cfg
	}
	return defaults
}
