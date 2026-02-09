/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package config

import (
	"strings"
)

type Column struct {
	Title   string `yaml:"title"`
	Content string `yaml:"content"`
}

func (c Column) String() string {
	return c.Title + ":" + c.Content
}

type Columns []Column

func ParseColumns(s string) Columns {
	ss := strings.Split(s, ",")
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
	Columns Columns `yaml:"columns"`
}

type Alias struct {
	Entity  string            `yaml:"entity"`
	Group   string            `yaml:"group"`
	Use     string            `yaml:"use"`
	Aliases []string          `yaml:"aliases"`
	Short   string            `yaml:"short"`
	Command []string          `yaml:"command"`
	Flags   map[string]string `yaml:"flags,omitzero"`
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
	Flags          []FlagConfig      `yaml:"flags,omitempty"`
}

type Configs map[string]Config

func For(provider string) Config {
	return Defaults()[provider]
}
