// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var renameApplianceCmd = &cobra.Command{
	Use:   `appliance [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L> <--new-id | -N>`,
	Short: "Rename a specific composable memory appliance (CMA) to a new ID",
	Long:  `Rename a specific composable memory appliance (CMA) to a new ID.`,
	Example: `
	cfm rename appliance --appliance-id applId --new-id newId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm rename appliance -L applId -N newId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestRenameAppliance(cmd)
		appliance, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsApplianceAction(appliance, "rename")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	renameApplianceCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(renameApplianceCmd)

	renameApplianceCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "Current ID of composable memory appliance (CMA)")
	renameApplianceCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	renameApplianceCmd.Flags().StringP(flags.NEW_ID, flags.NEW_ID_SH, flags.ID_DFLT, "New ID of composable memory appliance (CMA)")
	renameApplianceCmd.MarkFlagRequired(flags.NEW_ID)

	//Add command to parent
	renameCmd.AddCommand(renameApplianceCmd)
}
