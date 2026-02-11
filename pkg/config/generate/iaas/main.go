/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"fmt"
	"maps"
	"os"
	"reflect"
	"slices"
	"strings"

	"dario.cat/mergo"
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
		if strings.HasSuffix(m.Name, "Raw") || strings.HasSuffix(m.Name, "WithBody") || m.Type.NumOut() != 2 {
			continue
		}
		if strings.HasPrefix(m.Name, "Read") {
			_, err := buildReadCommand(&base, m)
			if err != nil {
				errors.ExitErr(err)
			}
		}
	}
	fd, err := os.Create(dst) //nolint:gosec
	if err != nil {
		errors.ExitErr(err)
	}
	err = yaml.NewEncoder(fd, yaml.UseSingleQuote(true)).Encode(base)
	if err != nil {
		errors.ExitErr(err)
	}
}

func buildReadCommand(cfg *config.Config, m reflect.Method) (bool, error) {
	typesName := strings.TrimPrefix(m.Name, "Read")
	typeName := strings.TrimSuffix(typesName, "s")
	if strings.HasSuffix(typesName, "ies") {
		typeName = strings.TrimSuffix(typesName, "ies") + "y"
	}
	fmt.Println("***", typeName)
	entity := strings.ToLower(typeName)
	typeNameList := []string{typeName}

	// Entity & columns, based on response
	resp := m.Type.Out(0).Elem()
	respContent, found := resp.FieldByName(typesName)
	if !found {
		for i := range resp.NumField() {
			field := resp.Field(i)
			if field.Name == "NextPageToken" || field.Name == "ResponseContext" {
				continue
			}
			respContent = field
			found = true
			break
		}
		if !found {
			fmt.Println("*** no response found")
			return false, nil
		}
		fmt.Println("guessed response", m.Name, respContent.Name)
		typeNameList = append(typeNameList, strings.TrimSuffix(respContent.Name, "s"))
		fmt.Println("type names", typeNameList)
	}
	respContentType := respContent.Type
	if respContentType.Kind() == reflect.Ptr {
		respContentType = respContentType.Elem()
	}
	if respContentType.Kind() == reflect.Slice {
		respContentType = respContentType.Elem()
	}
	if respContentType.Kind() != reflect.Struct {
		fmt.Println("*** invalid response type", respContentType.Kind())
		return false, nil
	}
	e := config.Entity{}
	for _, typeName := range typeNameList {
		if f, found := respContentType.FieldByName(typeName + "Id"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "ID",
				Content: f.Name,
			})
			break
		}
	}
	if f, found := respContentType.FieldByName(typeName + "Name"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Name",
			Content: f.Name,
		})
	} else if _, found := respContentType.FieldByName("Tags"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Name",
			Content: `find(Tags, #?.Key == "Name")?.Value`,
		})
	}
	if f, found := respContentType.FieldByName(typeName + "Type"); found {
		e.Columns = append(e.Columns, config.Column{
			Title:   "Type",
			Content: f.Name,
		})
	}
	for _, name := range priorityFields {
		if slices.ContainsFunc(e.Columns, func(c config.Column) bool { return c.Content == name }) {
			continue
		}
		if f, found := respContentType.FieldByName(name); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   f.Name,
				Content: f.Name,
			})
		}
	}
	ecfg := cfg.Entities[entity]
	err := mergo.Merge(&ecfg, e)
	if err != nil {
		return true, err
	}
	cfg.Entities[entity] = ecfg

	// Store response content field
	cfg.Calls[m.Name] = config.Call{
		Entity:  entity,
		Content: respContent.Name,
	}

	// list
	req := m.Type.In(2)
	flags := make(map[string]string)
	filters, found := req.FieldByName("Filters")
	var fType reflect.Type
	var prefix string
	if found {
		fType = filters.Type
		prefix = "Filters."
	} else {
		fType = req
	}
	if fType.Kind() == reflect.Pointer {
		fType = fType.Elem()
	}

	for i := range fType.NumField() {
		filt := fType.Field(i)
		flag := strings.TrimPrefix(strcase.ToSnake(filt.Name), entity+"_")
		flags[flag] = prefix + filt.Name
	}
	format := "table"
	if len(ecfg.Columns) == 0 {
		format = "yaml"
	}
	fmt.Println("list", typeName, "filters", slices.Collect(maps.Keys(flags)))
	cfg.Aliases = append(cfg.Aliases, config.Alias{
		Entity:  entity,
		Use:     "list",
		Aliases: []string{"ls"},
		Group:   entity,
		Short:   "alias for api " + m.Name,
		Command: []string{
			"api",
			m.Name,
			"-o", format,
		},
		Flags: flags,
	})

	// describe
	if fType != nil {
		// Guess id filter
		fids, found := fType.FieldByName(typeName + "Ids")
		if !found {
			fids, found = fType.FieldByName(typeName + "Names")
		}
		if !found {
			for i := range fType.NumField() {
				field := fType.Field(i)
				if strings.HasSuffix(field.Name, "Ids") && strings.HasSuffix(typeName, strings.TrimSuffix(field.Name, "Ids")) {
					fids, found = field, true
					typeNameList = append(typeNameList, strings.TrimSuffix(field.Name, "Ids"))
					fmt.Println("type names", typeNameList)
				}
			}
		}
		if found {
			fmt.Println("describe", typeName, "Filters."+fids.Name)
			cfg.Aliases = append(cfg.Aliases, config.Alias{
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
			})
		}
	}

	return true, nil
}
