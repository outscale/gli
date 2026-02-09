/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/iancoleman/strcase"
	"github.com/outscale/gli/pkg/config"
	"github.com/outscale/gli/pkg/errors"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
)

var priorityFields = []string{
	"State",
	"PublicIp",
	"PrivateIp",
	"NetId",
	"SubnetId",
	"IpRange",
	"SubregionName",
	"SubregionNames",
	"Subregion",
	"Subregions",
	"Size",
	"Iops",
}

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	var base config.Config
	data, err := os.ReadFile(src) //nolint:gosec
	if err != nil {
		errors.ExitErr(err)
	}
	err = yaml.Unmarshal(data, &base)
	if err != nil {
		errors.ExitErr(err)
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
		if !strings.HasPrefix(m.Name, "Read") || strings.HasSuffix(m.Name, "Raw") || strings.HasSuffix(m.Name, "WithBody") || m.Type.NumOut() != 2 {
			continue
		}
		typesName := strings.TrimPrefix(m.Name, "Read")
		typeName := strings.TrimSuffix(typesName, "s")
		entity := strings.ToLower(typeName)

		// Call
		base.Calls[m.Name] = config.Call{
			Entity:  entity,
			Content: typesName,
		}

		// Aliases, based on request
		req := m.Type.In(2)
		filters, found := req.FieldByName("Filters")
		if !found {
			continue
		}
		fType := filters.Type
		if fType.Kind() == reflect.Pointer {
			fType = fType.Elem()
		}
		fids, found := fType.FieldByName(typeName + "Ids")
		if !found {
			fids, found = fType.FieldByName(typeName + "Names")
		}
		if !found {
			continue
		}
		flags := make(map[string]string, fType.NumField())
		for i := range fType.NumField() {
			filt := fType.Field(i)
			flag := strings.TrimPrefix(strcase.ToSnake(filt.Name), entity+"_")
			flags[flag] = "Filters." + filt.Name
		}
		fmt.Println(m.Name, "->", typeName, "Filters."+fids.Name)
		base.Aliases = append(base.Aliases, config.Alias{
			Entity:  entity,
			Use:     "describe " + entity + "_id",
			Aliases: []string{"desc"},
			Group:   entity,
			Short:   "alias for api " + m.Name + " --Filters." + fids.Name + " " + entity + "_id",
			Command: []string{
				"api",
				m.Name,
				"--Filters." + fids.Name, "%0",
				"-o", "yaml,single",
			},
		}, config.Alias{
			Entity:  entity,
			Use:     "list",
			Aliases: []string{"l"},
			Group:   entity,
			Short:   "alias for api " + m.Name,
			Command: []string{
				"api",
				m.Name,
				"-o", "table",
			},
			Flags: flags,
		})

		// Entity & columns, based on response
		resp := m.Type.Out(0).Elem()
		obj, found := resp.FieldByName(typesName)
		if !found {
			continue
		}
		e := config.Entity{}
		tObj := obj.Type
		if tObj.Kind() == reflect.Ptr {
			tObj = tObj.Elem()
		}
		tObj = tObj.Elem()
		if f, found := tObj.FieldByName(typeName + "Id"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "ID",
				Content: f.Name,
			})
		}
		if f, found := tObj.FieldByName(typeName + "Name"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "Name",
				Content: f.Name,
			})
		}
		if f, found := tObj.FieldByName(typeName + "Type"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "Type",
				Content: f.Name,
			})
		}
		for _, name := range priorityFields {
			if slices.ContainsFunc(e.Columns, func(c config.Column) bool { return c.Content == name }) {
				continue
			}
			if f, found := tObj.FieldByName(name); found {
				e.Columns = append(e.Columns, config.Column{
					Title:   f.Name,
					Content: f.Name,
				})
			}
		}
		base.Entities[entity] = e
	}
	fd, err := os.Create(dst) //nolint:gosec
	if err != nil {
		errors.ExitErr(err)
	}
	err = yaml.NewEncoder(fd).Encode(base)
	if err != nil {
		errors.ExitErr(err)
	}
}
