package cmd

import (
	"github.com/outscale/octl/pkg/output/format"
	"github.com/spf13/cobra"
)

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "alias management",
	Long:  `User defined aliases`,
}

var aliasListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all user-defined aliases",
	Run:   listAliases,
}

var aliasAddCmd = &cobra.Command{
	Use:   "add [alias]... -- [command]...",
	Short: "Creates a new user-defined alias",
	Run:   addAlias,
}

var aliasDeleteCmd = &cobra.Command{
	Use:   "delete [alias]...",
	Short: "Deletes an existing user-defined alias",
	Run:   deleteAlias,
}

func listAliases(cmd *cobra.Command, args []string) {
	format.YAML{}.Format(cmd.Context(), )
}

func addAlias(cmd *cobra.Command, args []string) {}

func deleteAlias(cmd *cobra.Command, args []string) {}
