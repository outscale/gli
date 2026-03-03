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

	"github.com/itchyny/gojq"
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
	messages.Info("Using standard input for command-line templating")
	var input any
	err = json.Unmarshal(stdin, &input)
	if err != nil {
		return fmt.Errorf("input is not a JSON object: %w", err)
	}
	flag := false
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			flag = true
			continue
		}
		if !strings.HasPrefix(arg, "{{") {
			continue
		}

		j, err := gojq.Parse(strings.TrimPrefix(strings.TrimSuffix(arg, "}}"), "{{"))
		if err != nil {
			messages.Warn("unable to parse argument", arg)
			continue
		}

		var strs []string
		iter := j.Run(input)
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}
			if err, ok := v.(error); ok {
				if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
					break
				}
			}
			switch v := v.(type) {
			case string:
				strs = append(strs, v)
			default:
				strs = append(strs, fmt.Sprintf("%v", v))
			}
		}
		newArgs := slices.Clone(os.Args[:i])
		if flag {
			replace := strings.Join(strs, ",")
			newArgs = append(newArgs, replace)
			debug.Println("replacing", arg, "with", replace)
		} else {
			debug.Println("replacing", arg, "with", strs)
			newArgs = append(newArgs, strs...)
		}
		newArgs = append(newArgs, os.Args[i+1:]...)
		os.Args = newArgs
		flag = false
	}
	debug.Println("new args", os.Args)
	return nil
}
