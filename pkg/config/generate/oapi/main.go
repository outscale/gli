package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/outscale/gli/pkg/config"
	"github.com/outscale/gli/pkg/errors"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"go.yaml.in/yaml/v4"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	var base config.Config
	data, err := os.ReadFile(src)
	if err != nil {
		errors.ExitErr(err)
	}
	err = yaml.Unmarshal(data, &base)
	if err != nil {
		errors.ExitErr(err)
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
			continue
		}
		fmt.Println(m.Name, "->", typeName, "Filters."+fids.Name)
		base.Aliases = append(base.Aliases, config.Alias{
			Use:   "Read" + typeName,
			Group: typeName,
			Short: "alias for " + m.Name + " --Filters." + fids.Name,
			Command: []string{
				m.Name,
				"--Filters." + fids.Name,
				"%0",
			},
		})
	}
	fd, err := os.Create(dst)
	if err != nil {
		errors.ExitErr(err)
	}
	err = yaml.NewEncoder(fd).Encode(base)
	if err != nil {
		errors.ExitErr(err)
	}
}
