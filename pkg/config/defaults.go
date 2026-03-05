/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package config

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"strings"
	"sync"

	"github.com/goccy/go-yaml"
	"github.com/outscale/octl/pkg/messages"
)

//go:generate go run generate/storage/main.go generate/storage/defaults.yaml defaults_storage.yaml
//go:generate go run generate/iaas/main.go generate/iaas/defaults.yaml defaults_iaas.yaml
//go:generate go run generate/archive/main.go .
//go:embed defaults.zip
var defaults []byte

func Defaults() Configs {
	return sync.OnceValue(func() Configs {
		r, err := zip.NewReader(bytes.NewReader(defaults), int64(len(defaults)))
		if err != nil {
			messages.ExitErr(err)
		}

		defaults := Configs{}
		for _, f := range r.File {
			provider := strings.TrimPrefix(strings.TrimSuffix(f.Name, ".yaml"), "defaults_")
			rc, err := f.Open()
			if err != nil {
				messages.ExitErr(err)
			}
			var cfg Config
			err = yaml.NewDecoder(rc).Decode(&cfg)
			if err != nil {
				messages.ExitErr(err)
			}
			defaults[provider] = cfg
		}
		return defaults
	})()
}
