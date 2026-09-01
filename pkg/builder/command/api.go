package commandbuilder

import (
	"github.com/outscale/goutils/sdk/ptr"
	"github.com/outscale/octl/pkg/debug"
	"github.com/spf13/cobra"
)

// BuildAPI builds the low-level API.
func (b *Builder) BuildAPI(rootCmd *cobra.Command, run func(cmd *cobra.Command, args []string)) *cobra.Command {
	rootCmd.AddGroup(&cobra.Group{
		ID:    "api",
		Title: "API",
	})
	apiCmd := &cobra.Command{
		Use:     "api",
		GroupID: "api",
		Short:   "Call " + rootCmd.Use + " API",
	}
	rootCmd.AddCommand(apiCmd)

	for _, call := range b.cfg.API {
		if call.Use == "" {
			continue
		}
		if call.Group != "" && !apiCmd.ContainsGroup(call.Group) {
			apiCmd.AddGroup(&cobra.Group{ID: call.Group, Title: call.Group})
		}
		cmd := &cobra.Command{
			GroupID: call.Group,
			Use:     call.Use,
			Short:   call.Short,
			Long:    call.Help,
			Run:     run,
		}
		apiCmd.AddCommand(cmd)

		for _, f := range call.Flags {
			// Required flags are not configured as required
			// Templating (e.g. echo '{}' | octl foo bar) might set some required info without using the flag.
			if ptr.From(f.Required) {
				f.Required = new(false)
			}
			if err := b.buildFlag(cmd, f); err != nil {
				debug.Println(call.Entity, call.Use, "error building flag", f.Name, err)
			}
		}
	}
	return apiCmd
}
