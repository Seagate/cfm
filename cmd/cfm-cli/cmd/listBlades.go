/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var listBladesCmd = &cobra.Command{
	Use:   `blades [--serv-ip | -a] [--serv-net-port | -p] [--appliance-id | -L] [--blade-id | -B]`,
	Short: "List some or all recognized appliance blades",
	Long:  `Queries the cfm-service for some or all recognized appliance blades and outputs a detailed summary of the discovered blades to stdout.`,
	Example: `
	cfm list blades --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId
	cfm list blades --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId
	cfm list blades --serv-ip 127.0.0.1 --serv-net-port 8080 --blade-id bladeId
	cfm list blades --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list blades -a 127.0.0.1 -p 8080 -L applId -B bladeId
	cfm list blades -a 127.0.0.1 -p 8080 -L applId
	cfm list blades -a 127.0.0.1 -p 8080 -B bladeId
	cfm list blades -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListBlades(cmd)
		bladesSummary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(bladesSummary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listBladesCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(listBladesCmd)
	listBladesCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate (default \"all appliances searched\")")
	listBladesCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of a specific appliance blade. (default \"all blades returned\")")

	//Add command to parent
	listCmd.AddCommand(listBladesCmd)
}
