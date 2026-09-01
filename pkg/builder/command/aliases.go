package commandbuilder

import (
	"strings"

	"dario.cat/mergo"
	"github.com/outscale/goutils/sdk/ptr"
	"github.com/outscale/octl/pkg/alias"
	"github.com/outscale/octl/pkg/debug"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// BuildAliases builds the high-level API, must be run after BuildAPI as aliases might use API flags.
func (b *Builder) BuildAliases(rootCmd, apiCmd *cobra.Command) {
	for _, a := range b.cfg.Aliases {
		if a.AliasTo == "" {
			debug.Println("empty alias", a)
			continue
		}
		var serviceCmd *cobra.Command
		if rootCmd.Name() == a.Entity {
			serviceCmd = rootCmd
		} else {
			var found bool
			serviceCmd, found = lo.Find(rootCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == a.Entity })
			if !found {
				e := b.cfg.Entities[a.Entity]
				serviceCmd = &cobra.Command{
					Use:     a.Entity,
					Short:   "Manage " + e.Title + " resources",
					Aliases: e.Aliases,
				}
				rootCmd.AddCommand(serviceCmd)
			}
		}
		if a.SubCommand != "" {
			subc, found := lo.Find(serviceCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == a.SubCommand })
			if !found {
				subc = &cobra.Command{
					Use:   a.SubCommand,
					Short: a.SubCommand + " commands",
				}
				serviceCmd.AddCommand(subc)
			}
			serviceCmd = subc
		}
		help := a.Help
		if help == "" {
			help = b.cfg.API[a.AliasTo].Help
		}
		short := a.Short
		if short == "" {
			short = b.cfg.API[a.AliasTo].Short
		}
		if a.AliasHelp != "" {
			help += "\n\n" + "> alias for " + a.AliasHelp
		}
		// help, _ = md.Render(help)
		rootPath := strings.Split(serviceCmd.CommandPath(), " ")[1]
		cmd := &cobra.Command{
			Use:     a.Use,
			Aliases: a.Aliases,
			Short:   short,
			Long:    help,
			Run:     alias.RunFunc(rootPath, a),
		}
		serviceCmd.AddCommand(cmd)
		if apiCmd == nil {
			continue
		}
		callCmd, _ := lo.Find(apiCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == a.AliasTo })
		if callCmd == nil {
			debug.Println("alias", a.AliasTo, "not found in", rootCmd.Use, apiCmd.Use)
			continue
		}
		for _, f := range a.Flags {
			if af, found := b.cfg.API[a.AliasTo].Flags.Get(f.AliasTo); found {
				// alias to local file

				// custom values are resolved by the API call
				if af.CustomValue != "" {
					af.CustomValue = ""
				}
				// merge with API flag (mergo.WithoutDereference to not overwrite required=&false with required=&true)
				err := mergo.Merge(&f, af, mergo.WithoutDereference)
				if err != nil {
					debug.Println(a.Entity, a.Use, "error merging flag", f.Name, err)
					continue
				}
				// If a default is present, no need to require the flag
				if ptr.From(f.Required) && f.Default != "" {
					f.Required = new(false)
				}
				if err := b.buildFlag(cmd, f); err != nil {
					debug.Println(a.Entity, a.Use, "error building flag", f.Name, err)
				}
			} else if af := callCmd.InheritedFlags().Lookup(f.AliasTo); af != nil {
				// probably a persistent flag
				naf := *af
				naf.Name = f.Name
				if f.Help != "" {
					naf.Usage = f.Help
				}
				cmd.Flags().AddFlag(&naf)
			} else {
				debug.Println("no source flag found for", f.Name, "in", a.Use)
			}
		}
	}
}
