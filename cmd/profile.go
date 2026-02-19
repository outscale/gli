/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"cmp"
	"errors"
	"io/fs"
	"os"
	"slices"

	"github.com/charmbracelet/huh"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/output"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Profile file management",
	Long:  `Creates, updates profile from a config file`,
}

var profileListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all profiles from a config file",
	Run:   listProfiles,
}

var profileAddCmd = &cobra.Command{
	Use:   "add name",
	Short: "Add a profile to a config file",
	Run:   addProfile,
}

var profileUseCmd = &cobra.Command{
	Use:   "use name",
	Short: "Mark a profile as the default one",
	Run:   setDefaultProfile,
}

var profileDeleteCmd = &cobra.Command{
	Use:     "delete name",
	Aliases: []string{"del", "rm"},
	Short:   "Delete a profile from a config file",
	Run:     deleteProfile,
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.AddCommand(profileListCmd)
	profileCmd.AddCommand(profileAddCmd)
	profileCmd.AddCommand(profileUseCmd)
	profileCmd.AddCommand(profileDeleteCmd)

	profileAddCmd.Flags().String("ak", "", "Access Key")
	profileAddCmd.Flags().String("sk", "", "Secret Key - if not specified, you prompted to enter it")
	profileAddCmd.Flags().String("region", "eu-west-2", "Region")
	profileAddCmd.Flags().Bool("default", false, "Sets the new profile as the default")
	_ = cobra.MarkFlagRequired(profileAddCmd.Flags(), "ak")
	_ = profileAddCmd.RegisterFlagCompletionFunc("region", func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
		return []cobra.Completion{"eu-west-2", "us-west-1", "us-east-2", "cloudgouv-eu-west-1", "ap-northeast-1"}, cobra.ShellCompDirectiveDefault
	})
}

type profileEntry struct {
	Name    string
	Default bool
	profile.Profile
}

var profileColumns = config.Columns{{Title: "Name", Content: "Name"}, {Title: "Region", Content: "Region"}, {Title: "Default", Content: "Default"}}

func configPath(cmd *cobra.Command) string {
	path, _ := cmd.Flags().GetString("config")
	if path == "" {
		path = os.Getenv("OSC_CONFIG_FILE")
	}
	return path
}

func loadConfig(cmd *cobra.Command) (*profile.ConfigFile, error) {
	return profile.LoadConfigFile(configPath(cmd))
}

func listProfiles(cmd *cobra.Command, _ []string) {
	cf, err := loadConfig(cmd)
	if err != nil {
		messages.ExitErr(err)
	}
	out, _, err := output.NewFromFlags(cmd.Flags(), "table", "", profileColumns, false, true)
	if err != nil {
		messages.ExitErr(err)
	}
	def, _, err := cf.DefaultProfile()
	switch {
	case errors.Is(err, profile.ErrNoDefaultProfile):
	case err != nil:
		messages.ExitErr(err)
	}
	lst := lo.MapToSlice(cf.Profiles, func(k string, v profile.Profile) profileEntry {
		return profileEntry{Name: k, Profile: v, Default: k == def}
	})
	slices.SortFunc(lst,
		func(a, b profileEntry) int {
			return cmp.Compare(a.Name, b.Name)
		})
	_ = out.Format(cmd.Context(), lst)
}

func addProfile(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Usage()
		return
	}
	name := args[0]

	cf, err := loadConfig(cmd)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		cf = &profile.ConfigFile{
			Path:     configPath(cmd),
			Profiles: map[string]profile.Profile{},
		}
	case err != nil:
		messages.ExitErr(err)
	}
	if _, found := cf.Profiles[name]; found {
		messages.Exit(1, "Profile %q already exists", name)
	}

	ak, _ := cmd.Flags().GetString("ak")
	if ak == "" {
		messages.Exit(1, "Access key is required")
	}
	for name, eprof := range cf.Profiles {
		if eprof.AccessKey == ak {
			messages.Exit(1, "Profile %q already exists with the same access key", name)
		}
	}

	sk, _ := cmd.Flags().GetString("sk")
	if sk == "" {
		var err error
		sk, err = Prompt("Enter the secret Key:", huh.EchoModePassword)
		if err != nil {
			messages.ExitErr(err)
		}
	}
	if sk == "" {
		messages.Exit(1, "Secret key is required")
	}
	region, _ := cmd.Flags().GetString("region")
	if region == "" {
		messages.Exit(1, "Region is required")
	}
	newProfile := profile.Profile{
		AccessKey: ak,
		SecretKey: sk,
		Region:    region,
	}

	cf.Profiles[name] = newProfile
	err = cf.Save()
	if err != nil {
		messages.ExitErr(err)
	}
	messages.Success("Profile %q has been added", name)
}

func setDefaultProfile(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Usage()
		return
	}
	name := args[0]
	cf, err := loadConfig(cmd)
	if err != nil {
		messages.ExitErr(err)
	}
	err = cf.SetDefault(name)
	if err == nil {
		err = cf.Save()
	}
	if err != nil {
		messages.ExitErr(err)
	}
	messages.Success("Profile %q is now the default", name)
}

func deleteProfile(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = cmd.Usage()
		return
	}
	name := args[0]
	cf, err := loadConfig(cmd)
	if err != nil {
		messages.ExitErr(err)
	}
	if _, found := cf.Profiles[name]; !found {
		messages.Exit(1, "Profile %q does not exist", name)
	}
	delete(cf.Profiles, name)
	err = cf.Save()
	if err != nil {
		messages.ExitErr(err)
	}
	messages.Success("Profile %q has been deleted", name)
}
