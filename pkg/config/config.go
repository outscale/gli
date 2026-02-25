/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package config

import (
	"fmt"
	"strings"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/samber/lo"
)

type Column struct {
	Title    string `yaml:"title"`
	Content  string `yaml:"content"`
	Primary  bool   `yaml:"primary,omitempty"`
	compiled *vm.Program
}

func (c Column) String() string {
	return c.Title + ":" + c.Content
}

func (c *Column) compile(s any) error {
	var err error
	c.compiled, err = expr.Compile(c.Content, expr.Env(s))
	if err != nil {
		return fmt.Errorf("invalid expression %q: %w", c.Content, err)
	}
	return nil
}

func (c *Column) Get(v any) (any, error) {
	if c.compiled == nil {
		err := c.compile(v)
		if err != nil {
			return nil, fmt.Errorf("compile: %w", err)
		}
	}
	return expr.Run(c.compiled, v)
}

type Columns []Column

func ParseColumns(s string) Columns {
	ss := strings.Split(s, "|")
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
	Entity  string   `yaml:"entity"`
	Use     string   `yaml:"use"`
	AliasTo string   `yaml:"alias_to,omitempty"`
	Aliases []string `yaml:"aliases,omitempty"`
	Short   string   `yaml:"short,omitempty"`
	Command []string `yaml:"command"`
	Flags   FlagSet  `yaml:"flags,omitempty"`
	Prompt  *Prompt  `yaml:"prompt,omitempty"`
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
