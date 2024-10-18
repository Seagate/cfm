// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var renameBladeCmd = &cobra.Command{
	Use:   `blade  [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L> <--blade-id | -B> <--new-id | -N>`,
	Short: "Rename a single blade ID on a specific composable memory appliance (CMA) to a new ID",
	Long:  `Rename a single blade ID on a specific composable memory appliance (CMA) to a new ID.`,
	Example: `
	cfm rename blade --appliance-id applId --blade-id bladeId --new-id newId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm rename blade -L applId -B bladeId -N newId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestRenameBlade(cmd)
		blade, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsBladeAction(blade, "rename")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	renameBladeCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(renameBladeCmd)

	renameBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of blade's composable memory appliance (CMA)")
	renameBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	renameBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "Current blade ID")
	renameBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	renameBladeCmd.Flags().StringP(flags.NEW_ID, flags.NEW_ID_SH, flags.ID_DFLT, "New blade ID")
	renameBladeCmd.MarkFlagRequired(flags.NEW_ID)

	//Add command to parent
	renameCmd.AddCommand(renameBladeCmd)
}
