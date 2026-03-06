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
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/output/filter"
	"github.com/outscale/octl/pkg/output/format"
	"github.com/outscale/octl/pkg/output/read"
	"github.com/spf13/pflag"
)

func NewFromFlags(fs *pflag.FlagSet, out, contentField string, cols config.Columns, explode, sort bool) (format.Interface, Outputter, error) {
	fout, _ := fs.GetString("output")
	if fout != "" {
		out = fout
	}
	if out == "" {
		out = "raw"
	}
	out = strings.ToLower(out)

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
	switch out {
	case "none":
		fmter = format.None{}
	case "json", "raw":
		fmter = format.JSON{}
	case "yaml":
		fmter = format.YAML{}
	case "base64":
		fmter = format.Base64{}
	case "success":
		fmter = format.Success{}
	case "body":
		fmter = format.Body{}
	case "table", "csv":
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
		switch {
		case len(cols) == 0:
			messages.Info("No columns for table, switching to YAML...")
			fmter = format.YAML{}
		case out == "csv":
			fmter = format.Tabular{Columns: cols, Explode: explode, Sort: sort, Formatter: format.CSVFormatter{}}
		default:
			fmter = format.Tabular{Columns: cols, Explode: explode, Sort: sort, Formatter: format.TableFormatter{}}
		}
	default:
		return nil, nil, fmt.Errorf("unknown format %q", out)
	}

	writeTo, _ := fs.GetString("out-file")
	single, _ := fs.GetBool("single")
	switch {
	case out == "raw":
		return fmter, &Paginated{Read: read.NewRaw(), Format: fmter, Filters: filters, WriteTo: writeTo}, nil
	case single:
		return fmter, &Paginated{Read: read.NewPaginated(contentField), Format: format.Single{ForFormat: fmter}, Filters: filters, WriteTo: writeTo}, nil
	default:
		return fmter, &Paginated{Read: read.NewPaginated(contentField), Format: fmter, Filters: filters, WriteTo: writeTo}, nil
	}
}
