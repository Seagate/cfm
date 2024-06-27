// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var listServiceInfoCmd = &cobra.Command{
	Use:   `service [--serv-ip | -a] [--serv-net-port | -p]`,
	Short: "List general information on a specific cfm-service instance",
	Long:  `List general information on a specific cfm-service instance.`,
	Example: `
	cfm list service --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list service -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListServiceInfo(cmd)
		serviceInfo, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(serviceInfo)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listServiceInfoCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(listServiceInfoCmd)

	//Add command to parent
	listCmd.AddCommand(listServiceInfoCmd)
}
