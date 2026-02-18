/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package output

import (
	"fmt"
	"slices"
	"strings"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/output/format"
	"github.com/outscale/octl/pkg/output/read"
	"github.com/spf13/pflag"
)

func NewFromFlags(fs *pflag.FlagSet, out, contentField string, cols config.Columns, explode bool) (format.Interface, Outputter, error) {
	jq, _ := fs.GetString("jq")
	if jq != "" {
		format, err := format.NewJQ(jq)
		if err != nil {
			return nil, nil, err
		}
		return format, &Paginated{Read: &read.Raw{}, Format: format}, nil
	}
	fout, _ := fs.GetString("output")
	if fout != "" {
		out = fout
	}
	if out == "raw" || out == "" {
		out = "json,raw"
	}
	out, param, _ := strings.Cut(out, ",")

	var fmter format.Interface
	switch strings.ToLower(out) {
	case "none":
		fmter = format.None{}
	case "json":
		fmter = format.JSON{}
	case "yaml":
		fmter = format.YAML{}
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
			fmter = format.YAML{}
		} else {
			fmter = format.Table{Columns: cols, Explode: explode}
		}
	default:
		return nil, nil, fmt.Errorf("unknown format %q", out)
	}

	switch param {
	case "raw":
		return fmter, &Paginated{Read: read.NewRaw(), Format: fmter}, nil
	case "single":
		return fmter, &Single{Read: read.NewPaginated(contentField), Format: fmter}, nil
	case "":
		return fmter, &Paginated{Read: read.NewPaginated(contentField), Format: fmter}, nil
	default:
		return nil, nil, fmt.Errorf("unknown format option %q", param)
	}
}
