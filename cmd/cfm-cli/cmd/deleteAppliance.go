/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var deleteApplianceCmd = &cobra.Command{
	Use:   `appliance [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L>`,
	Short: "Delete a memory appliance connection from cfm-service",
	Long:  `Deletes a netowrk connection from the cfm-service to an external memory appliance.`,
	Example: `
	cfm delete appliance --appliance-id applId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm delete appliance -L applId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestDeleteAppliance(cmd)
		deletedAppliance, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsAddDeleteAppliance(deletedAppliance, "deleted")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	deleteApplianceCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(deleteApplianceCmd)

	deleteApplianceCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of memory appliance")

	//Add command to parent
	deleteCmd.AddCommand(deleteApplianceCmd)
}
