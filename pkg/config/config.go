/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/itchyny/gojq"
	"github.com/outscale/goutils/sdk/ptr"
	"github.com/outscale/octl/pkg/debug"
	"github.com/samber/lo"
)

type Column struct {
	Title   string `yaml:"title"`
	Content string `yaml:"content"`
	Primary bool   `yaml:"primary,omitempty"`
	query   *gojq.Query
}

func (c Column) String() string {
	return c.Title + ":" + c.Content
}

func (c *Column) compile() error {
	var err error
	c.query, err = gojq.Parse(c.Content)
	if err != nil {
		return fmt.Errorf("invalid expression %q: %w", c.Content, err)
	}
	return nil
}

func (c *Column) Get(v any) (any, error) {
	if c.query == nil {
		err := c.compile()
		if err != nil {
			return nil, fmt.Errorf("compile: %w", err)
		}
	}
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("jq to JSON: %w", err)
	}
	var raw any
	err = json.Unmarshal(buf, &raw)
	if err != nil {
		return nil, fmt.Errorf("jq from JSON: %w", err)
	}
	debug.Println(fmt.Sprintf("%+v", raw))
	iter := c.query.Run(raw)
	var vv []any
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil { //nolint
				break
			}
			return nil, fmt.Errorf("jq error: %w", err)
		}
		vv = append(vv, v)
	}
	switch len(vv) {
	case 0:
		return nil, nil
	case 1:
		return vv[0], nil
	default:
		return vv, nil
	}
}

type Columns []Column

func ParseColumns(s string) Columns {
	ss := strings.Split(s, "||")
	cs := make(Columns, 0, len(ss))
	for _, s := range ss {
		title, content, found := strings.Cut(s, ":")
		if !found {
			content = title
		}
		cs = append(cs, Column{Title: strings.TrimSpace(title), Content: strings.TrimSpace(content)})
	}
	return cs
}

type Entity struct {
	Title   string   `yaml:"title,omitempty"`
	Explode bool     `yaml:"explode,omitempty"`
	Sort    bool     `yaml:"sort,omitempty"`
	Aliases []string `yaml:"aliases,omitempty"`
	Columns Columns  `yaml:"columns,omitempty"`
	Primary string   `yaml:"primary,omitempty"`
}

func (e Entity) IsZero() bool {
	return e.Primary == "" && len(e.Columns) == 0 && len(e.Aliases) == 0 && e.Title == ""
}

type Action string

const (
	ActionDelete Action = "delete"
)

type FlagSet []Flag

func (fs FlagSet) Get(name string) (Flag, bool) {
	return lo.Find(fs, func(f Flag) bool {
		return f.Name == name
	})
}

func (fs FlagSet) Names() []string {
	return lo.Map(fs, func(f Flag, _ int) string { return f.Name })
}

func (fs *FlagSet) Add(f Flag) {
	*fs = append(*fs, f)
}

type Flag struct {
	Name          string       `yaml:"name"`
	AliasTo       string       `yaml:"alias_to,omitempty"`
	Required      *bool        `yaml:"required,omitempty"`
	Kind          reflect.Kind `yaml:"kind,omitempty"`
	ContainerKind reflect.Kind `yaml:"container,omitempty"`
	CustomValue   string       `yaml:"customValue,omitempty"`
	Help          string       `yaml:"help,omitempty"`
	Default       string       `yaml:"default,omitempty"`
	AllowedValues []string     `yaml:"values,omitzero"`
}

type Prompt struct {
	Action         Action   `yaml:"action"`
	DisplayCommand []string `yaml:"display,omitempty"`
	Flags          FlagSet  `yaml:"flags,omitempty"`
}

type Alias struct {
	Entity     string   `yaml:"entity"`
	SubCommand string   `yaml:"sub_command,omitempty"`
	Use        string   `yaml:"use"`
	Help       string   `yaml:"help,omitempty"`
	Short      string   `yaml:"short,omitempty"`
	AliasTo    string   `yaml:"alias_to,omitempty"`
	Aliases    []string `yaml:"aliases,omitempty"`
	AliasHelp  string   `yaml:"alias_help,omitempty"`
	Command    []string `yaml:"command"`
	Flags      FlagSet  `yaml:"flags,omitempty"`
	Prompt     *Prompt  `yaml:"prompt,omitempty"`
}

func (a *Alias) HasRequiredFlag() bool {
	return slices.ContainsFunc(a.Flags, func(f Flag) bool { return ptr.From(f.Required) })
}

type APICall struct {
	Use     string  `yaml:"use,omitempty"`
	Content string  `yaml:"content,omitempty"`
	Entity  string  `yaml:"entity,omitempty"`
	Group   string  `yaml:"group,omitempty"`
	AliasTo string  `yaml:"alias_to,omitempty"`
	Flags   FlagSet `yaml:"flags,omitempty"`
	Help    string  `yaml:"help,omitempty"`
	Short   string  `yaml:"short,omitempty"`
}

type Config struct {
	DefaultContent string             `yaml:"default_content,omitempty"`
	API            map[string]APICall `yaml:"api,omitzero"`
	Entities       map[string]Entity  `yaml:"entities,omitzero"`
	Aliases        []Alias            `yaml:"aliases,omitzero"`
}

type Configs map[string]Config

func For(provider string) Config {
	return Defaults()[provider]
}
