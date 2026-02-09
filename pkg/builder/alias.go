/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package builder

import (
	"os"
	"strconv"
	"strings"

	"github.com/outscale/gli/pkg/config"
	"github.com/outscale/gli/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func runAlias(provider string, a config.Alias, rootCmd *cobra.Command) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		nargs := make([]string, 2, len(a.Command)+2)
		nargs[0] = "gli"
		nargs[1] = provider
		for _, arg := range a.Command {
			if !strings.HasPrefix(arg, "%") {
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
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				newFlag := f.Name
				if a.Flags[newFlag] != "" {
					newFlag = a.Flags[newFlag]
				}
				nargs = append(nargs, "--"+newFlag+"="+f.Value.String())
			}
		})
		errors.Info("Resolving alias to %v", nargs)
		os.Args = nargs
		err := rootCmd.Execute()
		if err != nil {
			errors.ExitErr(err)
		}
	}
}
