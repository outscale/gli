package cmd

import (
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/sdk"
	"github.com/outscale/octl/pkg/version"
	"github.com/outscale/osc-sdk-go/v3/pkg/middleware"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	options "github.com/outscale/osc-sdk-go/v3/pkg/utils"
	"github.com/spf13/cobra"
)

func loadProfile(cmd *cobra.Command) *profile.Profile {
	path, _ := cmd.Flags().GetString("config")
	prof, _ := cmd.Flags().GetString("profile")
	p, err := profile.New(profile.FromFile(prof, path), profile.FromEnv())
	if err != nil {
		messages.ExitErr(err)
	}
	return p
}

func sdkOptions(cmd *cobra.Command) []middleware.MiddlewareChainOption {
	ua := "octl/" + version.Version
	opts := []middleware.MiddlewareChainOption{options.WithUseragent(ua)}
	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		opts = append(opts, options.WithLogging(sdk.VerboseLogger{}))
	} else {
		opts = append(opts, options.WithoutLogging())
	}
	return opts
}
