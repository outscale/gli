/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package builder

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func runAlias(provider string, a config.Alias, cmd *cobra.Command) func(cmd *cobra.Command, args []string) {
	run := runFunc(provider, a.Command, a.Flags, cmd, false)
	if a.Prompt != nil {
		var display func(cmd *cobra.Command, args []string)
		if len(a.Prompt.DisplayCommand) > 0 {
			display = runFunc(provider, a.Prompt.DisplayCommand, nil, cmd, true)
		}
		return confirm(a.Prompt.Action, display, run)
	}
	return run
}

func saveFlags(fs *pflag.FlagSet) map[string]string {
	saved := map[string]string{}
	fs.VisitAll(func(f *pflag.Flag) {
		if f.Changed {
			saved[f.Name] = f.Value.String()
		}
	})
	return saved
}

func restoreFlags(fs *pflag.FlagSet, saved map[string]string) {
	fs.VisitAll(func(f *pflag.Flag) {
		val, found := saved[f.Name]
		switch {
		case found:
			_ = f.Value.Set(val)
			f.Changed = true
		case !found && f.Changed:
			_ = f.Value.Set(f.DefValue)
			f.Changed = false
		}
	})
}

func runFunc(provider string, command []string, flags map[string]string, cmd *cobra.Command, skipUserFlags bool) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		nargs := make([]string, 2, len(command)+2)
		nargs[0] = "octl"
		nargs[1] = provider
		saved := saveFlags(cmd.Flags())
		var userArgs []string
		if !skipUserFlags {
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				if f.Changed {
					newFlag := f.Name
					if flags[newFlag] != "" {
						newFlag = flags[newFlag]
					}
					userArgs = append(userArgs, "--"+newFlag+"="+f.Value.String())
				}
			})
		}
		skipnextvalue := false
		for _, arg := range command {
			if !strings.HasPrefix(arg, "%") {
				// skip flags already present in user flags
				isFlag := strings.HasPrefix(arg, "--")
				if isFlag && slices.ContainsFunc(
					userArgs, func(uf string) bool {
						return strings.HasPrefix(uf, arg+"=")
					}) {
					skipnextvalue = true
					continue
				}
				// skip value present after skipped flag
				if skipnextvalue && !isFlag {
					skipnextvalue = false
					continue
				}
				skipnextvalue = false
				nargs = append(nargs, arg)
				continue
			}
			idx, err := strconv.Atoi(arg[1:])
			if err != nil {
				errors.Warn(err.Error())
				continue
			}
			if idx >= len(args) {
				_ = cmd.Usage()
				return
			}
			nargs = append(nargs, args[idx])
		}
		nargs = append(nargs, userArgs...)
		errors.Info("Resolving alias to %v", nargs)
		// no need to check for an update a second time
		nargs = append(nargs, "--no-upgrade")
		os.Args = nargs
		err := cmd.Execute()
		if err != nil {
			errors.ExitErr(err)
		}
		restoreFlags(cmd.Flags(), saved)
	}
}
