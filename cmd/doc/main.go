/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package main

import (
	"os"

	_ "github.com/outscale/octl/cmd/doc/notty" // this needs to be loaded first

	"github.com/outscale/octl/cmd"
	"github.com/outscale/octl/pkg/messages"
	"github.com/spf13/cobra/doc"
)

//go:generate go run main.go ../../docs/reference
func main() {
	path := os.Args[1]

	err := doc.GenMarkdownTree(cmd.Root(), path)
	if err != nil {
		messages.ExitErr(err)
	}
}
