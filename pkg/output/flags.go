/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"slices"
	"strings"

	"github.com/outscale/octl/pkg/config"
	"github.com/spf13/pflag"
)

func NewFromFlags(fs *pflag.FlagSet, c config.Call, e config.Entity) (Output, error) {
	jq, _ := fs.GetString("jq")
	if jq != "" {
		return NewJQ(jq)
	}
	out, _ := fs.GetString("output")
	out, param, _ := strings.Cut(out, ",")
	switch strings.ToLower(out) {
	case "none":
		return None{}, nil
	case "json":
		return content{content: c.Content, output: JSON{}, single: param == "single"}, nil
	case "yaml":
		return content{content: c.Content, output: YAML{}, single: param == "single"}, nil
	case "table":
		var cols config.Columns
		fcols, _ := fs.GetString("columns")
		if fcols != "" {
			cols = config.ParseColumns(fcols)
		} else {
			cols = slices.Clone(e.Columns)
		}
		if len(cols) == 0 {
			return content{content: c.Content, output: YAML{}, single: param == "single"}, nil
		}
		return content{content: c.Content, output: Table{Columns: cols}, single: param == "single"}, nil
	default:
		return Default{}, nil
	}
}
