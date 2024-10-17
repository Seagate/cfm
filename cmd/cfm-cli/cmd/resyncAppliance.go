// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var resyncApplianceCmd = &cobra.Command{
	Use:   `appliance [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L>`,
	Short: "Resynchronize the cfm service to all the added blades for a specific composable memory appliance (CMA)",
	Long:  `Resynchronize the cfm service to all the added blades for a specific composable memory appliance (CMA).`,
	Example: `
	cfm resync appliance --appliance-id applId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm resync appliance -L applId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestResyncAppliance(cmd)
		appliance, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsApplianceAction(appliance, "resync")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	resyncApplianceCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(resyncApplianceCmd)

	resyncApplianceCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of composable memory appliance (CMA)")
	resyncApplianceCmd.MarkFlagRequired(flags.APPLIANCE_ID)

	//Add command to parent
	resyncCmd.AddCommand(resyncApplianceCmd)
}
