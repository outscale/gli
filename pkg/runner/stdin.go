/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package runner

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/mattn/go-isatty"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
)

var (
	stdinChecked bool
	stdin        []byte
)

func Stdin() ([]byte, bool) {
	return stdin, stdinChecked && len(stdin) > 0
}

func CheckStdin() error {
	stdinChecked = true
	if isatty.IsTerminal(os.Stdin.Fd()) {
		debug.Println("terminal, skipping stdin")
		return nil
	}
	debug.Println("reading stdin")
	var err error
	stdin, err = io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read stdin: %w", err)
	}
	if !slices.ContainsFunc(os.Args, func(arg string) bool {
		return strings.HasPrefix(arg, "{{")
	}) {
		return nil
	}
	debug.Println("templating args")
	var input map[string]any
	err = json.Unmarshal(stdin, &input)
	if err != nil {
		return fmt.Errorf("input is not a JSON object: %w", err)
	}
	for i, arg := range os.Args {
		if !strings.HasPrefix(arg, "{{") {
			continue
		}
		tmpl, err := template.New("tpl").Parse(arg)
		if err != nil {
			messages.Warn("unable to parse argument", arg)
			continue
		}
		w := &strings.Builder{}
		err = tmpl.Execute(w, input)
		if err != nil {
			messages.Warn("unable to get value for argument", arg)
			continue
		}
		debug.Println("replacing", arg, "with", w.String())
		os.Args[i] = w.String()
	}
	debug.Println("new args", os.Args)
	return nil
}
