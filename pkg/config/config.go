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
	Explode bool     `yaml:"explode,omitempty"`
	Sort    bool     `yaml:"sort,omitempty"`
	Aliases []string `yaml:"aliases,omitempty"`
	Columns Columns  `yaml:"columns,omitempty"`
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
	AliasTo  string `yaml:"alias_to"`
	Required bool   `yaml:"required,omitempty"`
}

type Prompt struct {
	Action         Action   `yaml:"action"`
	DisplayCommand []string `yaml:"display"`
}

type Alias struct {
	Entity  string   `yaml:"entity"`
	Group   string   `yaml:"group"`
	Use     string   `yaml:"use"`
	Aliases []string `yaml:"aliases,omitempty"`
	Short   string   `yaml:"short"`
	Command []string `yaml:"command"`
	Flags   FlagSet  `yaml:"flags,omitempty"`
	Prompt  *Prompt  `yaml:"prompt,omitempty"`
}

type FlagConfig struct {
	AppliesTo  string `yaml:"applies_to"`
	TrimPrefix string `yaml:"trim_prefix"`
}

type Call struct {
	Content string `yaml:"content"`
	Entity  string `yaml:"entity"`
}

type Config struct {
	DefaultContent string            `yaml:"default_content,omitempty"`
	Calls          map[string]Call   `yaml:"contents,omitempty"`
	Entities       map[string]Entity `yaml:"entities,omitempty"`
	Aliases        []Alias           `yaml:"aliases,omitempty"`
}

type Configs map[string]Config

func For(provider string) Config {
	return Defaults()[provider]
}
