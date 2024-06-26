// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

// listHostsCmd represents the listHosts command
var listHostsCmd = &cobra.Command{
	Use:   `hosts [--serv-ip | -i] [--serv-net-port | -p]`,
	Short: "List all recognized cxl host(s)",
	Long:  `Queries the cfm-service for all recognized cxl host(s) and outputs a summary to stdout.`,
	Example: `
	cfm list hosts --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list hosts -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListHosts(cmd)
		hosts, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(hosts)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listHostsCmd.DisableFlagsInUseLine = true

	//Just need service flags here
	initCommonPersistentFlags(listHostsCmd)

	//Add command to parent
	listCmd.AddCommand(listHostsCmd)
}
