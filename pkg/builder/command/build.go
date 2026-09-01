/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package commandbuilder

import (
	"github.com/outscale/octl/pkg/config"
	"github.com/spf13/cobra"
)

// var md = markdown.NewRenderer()

type Builder struct {
	provider string
	cfg      config.Config
}

func NewBuilder(provider string, helpURL string) *Builder {
	return &Builder{
		provider: provider,
		cfg:      config.For(provider),
	}
}

func (b *Builder) Build(rootCmd *cobra.Command, runAPI func(cmd *cobra.Command, args []string)) {
	apiCmd := b.BuildAPI(rootCmd, runAPI)
	b.BuildAliases(rootCmd, apiCmd)
}
