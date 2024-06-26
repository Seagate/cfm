// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var listBladeResourcesCmd = &cobra.Command{
	Use:   `resources [--serv-ip | -a] [--serv-net-port | -p] [--appliance-id | -L] [--blade-id | -B] [--resource-id | -i]`,
	Short: "List all available blade memory resources",
	Long: `Queries the cfm-service for existing memory resources.
	Outputs a detailed summary (including composition state) of those resources to stdout.
	Note that, for any given item ID, if it is omitted, ALL items are collected\searched.`,
	Example: `
	cfm list blades resources --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId --resource-id resId
	cfm list blades resources --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId
	cfm list blades resources --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --resource-id resId
	cfm list blades resources --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId
	cfm list blades resources --serv-ip 127.0.0.1 --serv-net-port 8080 --blade-id bladeId --resource-id resId
	cfm list blades resources --serv-ip 127.0.0.1 --serv-net-port 8080 --blade-id bladeId
	cfm list blades resources --serv-ip 127.0.0.1 --serv-net-port 8080 --resource-id resId
	cfm list blades resources --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list blades resources -a 127.0.0.1 -p 8080 -L applId -B bladeId -r resId
	cfm list blades resources -a 127.0.0.1 -p 8080 -L applId -B bladeId
	cfm list blades resources -a 127.0.0.1 -p 8080 -L applId -r resId
	cfm list blades resources -a 127.0.0.1 -p 8080 -L applId
	cfm list blades resources -a 127.0.0.1 -p 8080 -B bladeId -r resId
	cfm list blades resources -a 127.0.0.1 -p 8080 -B bladeId
	cfm list blades resources -a 127.0.0.1 -p 8080 -r resId
	cfm list blades resources -a 127.0.0.1 -p 8080 `,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListResources(cmd)
		results, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(results)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listBladeResourcesCmd.DisableFlagsInUseLine = true

	listBladeResourcesCmd.Flags().StringP(flags.RESOURCE_ID, flags.RESOURCE_ID_SH, flags.ID_DFLT, "ID of a specific resource block. (default \"all resource blocks returned\")")

	initCommonBladeListCmdFlags(listBladeResourcesCmd)

	//Add command to parent
	listBladesCmd.AddCommand(listBladeResourcesCmd)
}
