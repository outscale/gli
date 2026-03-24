/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package alias

import (
	"bytes"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/output"
	"github.com/outscale/octl/pkg/runner"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func RunFunc(rootPath string, a config.Alias) func(cmd *cobra.Command, args []string) {
	if !slices.Contains(a.Command, "|") {
		return runAliasWithPrompt(rootPath, a)
	}
	first, last, _ := lo.Cut(a.Command, []string{"|"})
	var firstFlags, lastFlags config.FlagSet
	if slices.Contains(first, a.AliasTo) {
		firstFlags = a.Flags
	} else {
		lastFlags = a.Flags
	}
	runFirst := runFunc(rootPath, first, firstFlags, true)
	runLast := runFunc(rootPath, last, lastFlags, true)
	return func(cmd *cobra.Command, args []string) {
		buf := &bytes.Buffer{}
		output.InjectOutput(buf)
		runFirst(cmd, args)
		messages.Info("Piping output")
		runner.InjectStdin(buf.Bytes())
		output.InjectOutput(os.Stdout)
		runLast(cmd, args)
	}
}

func runAliasWithPrompt(rootPath string, a config.Alias) func(cmd *cobra.Command, args []string) {
	run := runFunc(rootPath, a.Command, a.Flags, false)
	if a.Prompt != nil {
		var display func(cmd *cobra.Command, args []string)
		if len(a.Prompt.DisplayCommand) > 0 {
			display = runFunc(rootPath, a.Prompt.DisplayCommand, a.Prompt.Flags, true)
		}
		return Confirm(a.Prompt.Action, display, run)
	}
	return run
}

func saveFlags(fs *pflag.FlagSet) map[string][]string {
	saved := map[string][]string{}
	fs.VisitAll(func(f *pflag.Flag) {
		if f.Changed {
			if svalue, ok := f.Value.(pflag.SliceValue); ok {
				saved[f.Name] = svalue.GetSlice()
				_ = svalue.Replace(nil)
			} else {
				saved[f.Name] = []string{f.Value.String()}
				_ = f.Value.Set(f.DefValue)
			}
			f.Changed = false
		}
	})
	return saved
}

func restoreFlags(fs *pflag.FlagSet, saved map[string][]string) {
	fs.VisitAll(func(f *pflag.Flag) {
		val, found := saved[f.Name]
		switch {
		case found:
			if svalue, ok := f.Value.(pflag.SliceValue); ok {
				_ = svalue.Replace(val)
			} else if len(val) > 0 {
				_ = f.Value.Set(val[0])
			}
			f.Changed = true
		case !found && f.Changed:
			if svalue, ok := f.Value.(pflag.SliceValue); ok {
				_ = svalue.Replace(nil)
			} else {
				_ = f.Value.Set(f.DefValue)
			}
			f.Changed = false
		}
	})
}

func once(fn func(cmd *cobra.Command, args []string) int) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		_ = fn(cmd, args)
	}
}

func iterate(fn func(cmd *cobra.Command, args []string) int) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		for {
			consumed := fn(cmd, args)
			debug.Println("consumed", consumed, "len", len(args))
			if consumed <= 0 || len(args) == consumed {
				break
			}
			args = args[consumed:]
		}
	}
}

func runFunc(rootPath string, command []string, flags config.FlagSet, skipUserFlags bool) func(cmd *cobra.Command, args []string) {
	exec := func(cmd *cobra.Command, args []string) int {
		nargs := make([]string, 2, len(command)+2)
		nargs[0] = "octl"
		nargs[1] = rootPath
		var userArgs []string
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				newFlag := f.Name
				nf, found := flags.Get(newFlag)
				switch {
				case !found && skipUserFlags:
					return
				case found:
					newFlag = nf.AliasTo
				}
				if svalue, ok := f.Value.(pflag.SliceValue); ok {
					userArgs = append(userArgs, "--"+newFlag+"="+strings.Join(svalue.GetSlice(), ","))
				} else {
					userArgs = append(userArgs, "--"+newFlag+"="+f.Value.String())
				}
			}
		})
		skipnextvalue := false
		consumed := -1
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
			if arg == "%*" {
				if len(args) == 0 {
					_ = cmd.Usage()
					os.Exit(1)
				}
				nargs = append(nargs, strings.Join(args, ","))
				consumed = len(args) - 1
				continue
			}
			idx, err := strconv.Atoi(arg[1:])
			if err != nil {
				messages.Warn(err.Error())
				continue
			}
			if idx >= len(args) {
				_ = cmd.Usage()
				os.Exit(1)
			}
			nargs = append(nargs, args[idx])
			consumed = max(consumed, idx)
		}
		nargs = append(nargs, userArgs...)
		nargs, err := runner.TemplateArgs(nargs)
		if err != nil {
			messages.ExitErr(err)
		}
		messages.Info("Resolving alias to %v", nargs)
		// no need to check for an update a second time
		nargs = append(nargs, "--no-upgrade")

		saved := saveFlags(cmd.Flags())
		os.Args = nargs
		err = cmd.Execute()
		if err != nil {
			messages.ExitErr(err)
		}
		restoreFlags(cmd.Flags(), saved)
		return consumed + 1
	}

	if lo.CountBy(command, func(arg string) bool { return strings.HasPrefix(arg, "%") }) == 1 {
		return iterate(exec)
	}
	return once(exec)
}
