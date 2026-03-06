/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"os"

	"github.com/outscale/octl/cmd"
	"github.com/outscale/octl/pkg/messages"
	"github.com/spf13/cobra/doc"
)

func main() {
	path := os.Args[1]
	err := doc.GenMarkdownTree(cmd.Root(), path)
	if err != nil {
		messages.ExitErr(err)
	}
}
