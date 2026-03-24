/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package config

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/itchyny/gojq"
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
	Skip      bool     `yaml:"skip,omitempty"`
	NoAliases bool     `yaml:"no_aliases,omitempty"`
	Explode   bool     `yaml:"explode,omitempty"`
	Sort      bool     `yaml:"sort,omitempty"`
	Aliases   []string `yaml:"aliases,omitempty"`
	Columns   Columns  `yaml:"columns,omitempty"`
	Primary   string   `yaml:"primary,omitempty"`
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

type Flag struct {
	Name     string `yaml:"name"`
	AliasTo  string `yaml:"alias_to,omitempty"`
	Required bool   `yaml:"required,omitempty"`
	Type     string `yaml:"type,omitempty"`
	Usage    string `yaml:"usage,omitempty"`
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
	AliasTo    string   `yaml:"alias_to,omitempty"`
	Aliases    []string `yaml:"aliases,omitempty"`
	Short      string   `yaml:"short,omitempty"`
	Command    []string `yaml:"command"`
	Flags      FlagSet  `yaml:"flags,omitempty"`
	Prompt     *Prompt  `yaml:"prompt,omitempty"`
}

func (a *Alias) HasRequiredFlag() bool {
	return slices.ContainsFunc(a.Flags, func(f Flag) bool { return f.Required })
}

type FlagConfig struct {
	AppliesTo  string `yaml:"applies_to"`
	TrimPrefix string `yaml:"trim_prefix"`
}

type Call struct {
	Content string  `yaml:"content,omitempty"`
	Entity  string  `yaml:"entity,omitempty"`
	Group   string  `yaml:"group,omitempty"`
	AliasTo string  `yaml:"alias_to,omitempty"`
	Short   string  `yaml:"short,omitempty"`
	Flags   FlagSet `yaml:"flags,omitempty"`
}

type SpecCall struct {
	Help string `yaml:"help,omitempty"`
}

type SpecAttribute struct {
	Help     string `yaml:"help,omitempty"`
	Required bool   `yaml:"required,omitempty"`
}
type Spec struct {
	Calls      map[string]SpecCall      `yaml:"calls,omitempty"`
	Attributes map[string]SpecAttribute `yaml:"attributes,omitempty"`
}

func (s Spec) ForCall(name string) SpecCall {
	if s.Calls == nil {
		return SpecCall{}
	}
	return s.Calls[name]
}

func (s Spec) ForAttribute(call, name string) SpecAttribute {
	if s.Attributes == nil {
		return SpecAttribute{}
	}
	return s.Attributes[call+"."+name]
}

func (s *Spec) SetAttribute(call, name string, spec SpecAttribute) {
	s.Attributes[call+"."+name] = spec
}

type Config struct {
	DefaultContent string            `yaml:"default_content,omitempty"`
	Calls          map[string]Call   `yaml:"contents,omitempty"`
	Entities       map[string]Entity `yaml:"entities,omitempty"`
	Aliases        []Alias           `yaml:"aliases,omitempty"`
	Spec           Spec              `yaml:"spec,omitzero"`
}

type Configs map[string]Config

func For(provider string) Config {
	return Defaults()[provider]
}
