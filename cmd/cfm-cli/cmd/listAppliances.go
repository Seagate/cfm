/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var listAppliancesCmd = &cobra.Command{
	Use:   `appliances [--serv-ip | -i] [--serv-net-port | -p]`,
	Short: "List all recognized memory appliance(s)",
	Long:  `Queries the cfm-service for all recognized memory appliance(s) and outputs a summary to stdout.`,
	Example: `
	cfm list appliances --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list appliances -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListAppliances(cmd)
		appliances, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(appliances)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listAppliancesCmd.DisableFlagsInUseLine = true

	//Just need service flags here
	initCommonPersistentFlags(listAppliancesCmd)

	//Add command to parent
	listCmd.AddCommand(listAppliancesCmd)
}
