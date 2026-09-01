package configbuilder

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/outscale/octl/pkg/config"
	"github.com/samber/lo"
)

type EntityBuilder struct {
	cfg Config

	build *config.Config

	m reflect.Method

	typeName     string
	typeNameList []string
	entityName   string
	call         config.APICall

	entity config.Entity
}

func NewEntityBuilder(cfg Config, build *config.Config, m reflect.Method) *EntityBuilder {
	entity := cfg.getEntityNameFor(m.Name)
	typeName := cfg.getTypeNameFor(m.Name)
	typeNameList := []string{typeName, ""}
	if rawTypeName := cfg.getRawTypeNameFor(m.Name); rawTypeName != typeName {
		typeNameList = append(typeNameList, rawTypeName)
	}
	call := build.API[m.Name]
	if call.Entity != "" {
		entity = call.Entity
	}
	return &EntityBuilder{
		cfg:          cfg,
		build:        build,
		m:            m,
		typeName:     typeName,
		typeNameList: typeNameList,
		entityName:   entity,
		call:         call,

		entity: build.Entities[entity],
	}
}

func (b *EntityBuilder) Build(spec Spec) error {
	if !isReadMethod(b.m.Name) {
		return nil
	}
	if b.m.Type.NumOut() < 2 {
		fmt.Println(" - wrong numout, no entity")
		return nil
	}

	fmt.Println(" - building entity")

	if b.entity.Title == "" {
		fmt.Println("  -> title", b.typeName)
		b.entity.Title = b.typeName
	}
	if b.call.Content == "" {
		fmt.Println("  -> no content, no colums")
		return nil
	}

	resp := b.m.Type.Out(0)
	if resp.Kind() == reflect.Pointer {
		resp = resp.Elem()
	}
	respContent, found := resp.FieldByName(b.call.Content)
	if !found {
		return nil
	}
	respContentType := respContent.Type
	if respContentType.Kind() == reflect.Pointer {
		respContentType = respContentType.Elem()
	}
	if respContentType.Kind() == reflect.Slice {
		respContentType = respContentType.Elem()
	}
	if respContentType.Kind() != reflect.Struct {
		fmt.Println("  -> content is not a struct, no columns")
		return nil
	}

	e := config.Entity{}
	hasPrimary, hasName := false, false
	for _, typeName := range b.typeNameList {
		if f, found := respContentType.FieldByName(typeName + "Id"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "ID",
				Content: "." + f.Name,
			})
			e.Primary = f.Name
			fmt.Println("  -> primary", f.Name)
			hasPrimary = true
		}
	}
	for _, typeName := range b.typeNameList {
		if f, found := respContentType.FieldByName(typeName + "Name"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "Name",
				Content: "." + f.Name,
			})
			hasName = true
			if !hasPrimary {
				e.Primary = f.Name
				fmt.Println("  -> primary from name", f.Name)
				hasPrimary = true
			}
		}
	}
	if !hasName {
		if _, found := respContentType.FieldByName("Tags"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "Name",
				Content: `.Tags[] | select(.Key == "Name").Value`,
			})
		}
	}
	for _, typeName := range b.typeNameList {
		if f, found := respContentType.FieldByName(typeName + "Type"); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   "Type",
				Content: "." + f.Name,
			})
		}
	}
	for _, name := range b.cfg.PriorityFields {
		if slices.ContainsFunc(e.Columns, func(c config.Column) bool { return c.Content == "."+name }) {
			continue
		}
		if f, found := respContentType.FieldByName(name); found {
			e.Columns = append(e.Columns, config.Column{
				Title:   f.Name,
				Content: "." + f.Name,
			})
		}
	}
	if b.entity.Primary == "" {
		fmt.Println("  -> guessed primary", e.Primary)
		b.entity.Primary = e.Primary
	}
	if len(b.entity.Columns) == 0 {
		fmt.Println("  -> guessed columns", lo.Map(e.Columns, func(c config.Column, _ int) string { return c.Title }))
		b.entity.Columns = e.Columns
	}
	return nil
}

func (b *EntityBuilder) Commit() {
	if !b.entity.IsZero() {
		b.build.Entities[b.entityName] = b.entity
	}
}

var _ MethodBuilder = (*EntityBuilder)(nil)
