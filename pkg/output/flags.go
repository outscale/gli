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
	"github.com/outscale/octl/pkg/watch"
	"github.com/spf13/pflag"
)

func NewFromFlags(fs *pflag.FlagSet, out, contentField string, cols config.Columns, explode, sort bool) (format.Interface, Outputter, error) {
	fout, _ := fs.GetString("output")
	if fout != "" {
		out = fout
	}
	out = strings.ToLower(out)
	doWatch, _ := fs.GetBool("watch")
	elapsed, _ := fs.GetBool("elapsed")
	switch {
	case doWatch && out != "table" && out != "csv" && out != "text":
		messages.Info("Switching to table...")
		out = "table"
	case out == "":
		out = "raw"
	}

	var filters []filter.Interface
	if decode, _ := fs.GetBool("decode"); decode {
		filters = append(filters, filter.NewDecode())
	}
	if split, _ := fs.GetBool("split"); split {
		filters = append(filters, filter.NewSplit())
	}
	if skip, _ := fs.GetInt("skip"); skip > 0 {
		filters = append(filters, filter.NewSkip(skip))
	}
	if doWatch && elapsed {
		filters = append(filters, filter.NewElapsed())
	}
	filts, _ := fs.GetStringSlice("filter")
	for _, filt := range filts {
		name, value, found := strings.Cut(filt, ":")
		var jqstr string
		if found {
			jqstr = fmt.Sprintf(`select(.%s | tostring | test("%s"))`, name, value)
		} else {
			jqstr = fmt.Sprintf(`select(. | tostring | test("%s"))`, name)
		}
		jqf, err := filter.NewJQ(jqstr)
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
	reverse, _ := fs.GetBool("reverse")
	if reverse {
		filters = append(filters, filter.Reverse{})
	}
	if len(filters) > 0 {
		filters = slices.Insert(filters, 0, filter.Interface(filter.JSON{}))
	}

	var fmter format.Interface
	switch out {
	case "none":
		fmter = format.None{}
	case "json", "raw":
		fmter = format.JSON{}
	case "yaml":
		fmter = format.YAML{}
	case "success":
		fmter = format.Success{}
	case "body":
		fmter = format.Body{}
	case "text":
		fmter = format.Text{}
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
		if doWatch && elapsed {
			cols = append(cols, config.Column{Title: "Elapsed", Content: "._elapsed"})
		}
		dryRun, _ := fs.GetBool("dry-run")
		switch {
		case dryRun:
			messages.Info("--dry-run is incompatible with the table output format, switching to YAML...")
			fmter = format.YAML{}
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

	switch {
	case doWatch && out == "table":
		fmter = watch.NewFormat(fmter)
	case doWatch:
		filters = append(filters, filter.NewDedup(true))
	default: // single breaks --watch
		single, _ := fs.GetBool("single")
		if single {
			fmter = format.Single{ForFormat: fmter}
		}
	}

	writeTo, _ := fs.GetString("out-file")
	if out == "raw" {
		return fmter, &Paginated{Read: read.NewRaw(), Format: fmter, Filters: filters, WriteTo: writeTo}, nil
	}
	maxPages, _ := fs.GetInt("max-pages")
	return fmter, &Paginated{Read: read.NewPaginated(contentField, maxPages), Format: fmter, Filters: filters, WriteTo: writeTo}, nil
}
