/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list <device> [<flags>]",
	Short: "List all recognized devices or hw resources",
	Long: `Queries the cfm-service for the specified device or hardware resources and
outputs a summary of them to stdout.`,
	Args: cobra.MatchAll(cobra.NoArgs),
}

// Here you will define your flags and configuration settings.
func init() {
	listCmd.DisableFlagsInUseLine = true

	rootCmd.AddCommand(listCmd)
}

// Add set of flags used by all the "list" commands
// Kept separate because of cobra framework quirk
// Need to ONLY add them in the subcommands init() func (not in listCmd.init()) so flags ONLY show up as "local" flags in help, not "global"
func initCommonBladeListCmdFlags(cmd *cobra.Command) {

	initCommonPersistentFlags(cmd)

	cmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate")
	cmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of blade to interrogate")
}

// Add set of flags used by all the "list" commands
// Kept separate because of cobra framework quirk
// Need to ONLY add them in the subcommands init() func (not in listCmd.init()) so flags ONLY show up as "local" flags in help, not "global"
func initCommonHostListCmdFlags(cmd *cobra.Command) {

	initCommonPersistentFlags(cmd)

	cmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate")
}
