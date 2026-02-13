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

func NewFromFlags(fs *pflag.FlagSet, out, contentField string, cols config.Columns) (Output, error) {
	jq, _ := fs.GetString("jq")
	if jq != "" {
		return NewJQ(jq)
	}
	fout, _ := fs.GetString("output")
	if fout != "" {
		out = fout
	}
	out, param, _ := strings.Cut(out, ",")
	switch strings.ToLower(out) {
	case "none":
		return None{}, nil
	case "json":
		return content{contentField: contentField, output: JSON{}, single: param == "single"}, nil
	case "yaml":
		return content{contentField: contentField, output: YAML{}, single: param == "single"}, nil
	case "table":
		fcols, _ := fs.GetString("columns")
		if fcols != "" {
			add := strings.HasPrefix(fcols, "+")
			pfcols := config.ParseColumns(strings.TrimPrefix(fcols, "+"))
			if add {
				cols = append(slices.Clone(cols), pfcols...)
			} else {
				cols = pfcols
			}
		} else {
			cols = slices.Clone(cols)
		}
		if len(cols) == 0 {
			return content{contentField: contentField, output: YAML{}, single: param == "single"}, nil
		}
		return content{contentField: contentField, output: Table{Columns: cols}, single: param == "single"}, nil
	default:
		return Default{}, nil
	}
}
