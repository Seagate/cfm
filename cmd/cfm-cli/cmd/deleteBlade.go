// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var deleteBladeCmd = &cobra.Command{
	Use:   `blade  [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L> <--blade-id | -B>`,
	Short: "Delete a memory appliance blade connection from cfm-service",
	Long:  `Deletes a netowrk connection from the cfm-service to an external memory appliance blade.`,
	Example: `
	cfm delete blade --appliance-id applId --blade-id bladeId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm delete blade -L applId -B bladeId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestDeleteBlade(cmd)
		deletedBlade, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsBladeAction(deletedBlade, "delete")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	deleteBladeCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(deleteBladeCmd)

	deleteBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of blade's appliance")
	deleteBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of blade to delete")

	//Add command to parent
	deleteCmd.AddCommand(deleteBladeCmd)
}
