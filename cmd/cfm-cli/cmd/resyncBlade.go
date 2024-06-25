// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var resyncBladeCmd = &cobra.Command{
	Use:   `blade  [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L> <--blade-id | -B>`,
	Short: "Resynchronize the cfm service to a single blade on a specific composable memory appliance (CMA)",
	Long:  `Resynchronize the cfm service to a single blade on a specific composable memory appliance (CMA).`,
	Example: `
	cfm resync blade --appliance-id applId --blade-id bladeId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm resync blade -L applId -B bladeId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestResyncBlade(cmd)
		blade, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsBladeAction(blade, "resync")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	resyncBladeCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(resyncBladeCmd)

	resyncBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of blade's appliance")
	resyncBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of blade")

	//Add command to parent
	resyncCmd.AddCommand(resyncBladeCmd)
}
