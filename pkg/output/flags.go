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
	"github.com/outscale/octl/pkg/output/filter"
	"github.com/outscale/octl/pkg/output/format"
	"github.com/outscale/octl/pkg/output/read"
	"github.com/spf13/pflag"
)

func NewFromFlags(fs *pflag.FlagSet, out, contentField string, cols config.Columns, explode bool) (format.Interface, Outputter, error) {
	fout, _ := fs.GetString("output")
	if fout != "" {
		out = fout
	}
	if out == "raw" || out == "" {
		out = "json,raw"
	}
	out, param, _ := strings.Cut(out, ",")

	var filters []filter.Interface
	filts, _ := fs.GetStringSlice("filter")
	for _, filt := range filts {
		name, value, _ := strings.Cut(filt, ":")
		jqf, err := filter.NewJQ(fmt.Sprintf(`select(.%s | test("%s"))`, name, value))
		if err != nil {
			return nil, nil, err
		}
		filters = append(filters, jqf)
	}
	jq, _ := fs.GetString("jq")
	if jq != "" {
		jqf, err := filter.NewJQ(jq)
		if err != nil {
			return nil, nil, err
		}
		filters = append(filters, jqf)
	}

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

	if param == "single" {
		fmter = format.Single{ForFormat: fmter}
	}
	switch param {
	case "raw":
		return fmter, &Paginated{Read: read.NewRaw(), Format: fmter, Filters: filters}, nil
	case "", "single":
		return fmter, &Paginated{Read: read.NewPaginated(contentField), Format: fmter, Filters: filters}, nil
	default:
		return nil, nil, fmt.Errorf("unknown format option %q", param)
	}
}
