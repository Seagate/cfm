/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var listBladesPortsCmd = &cobra.Command{
	Use:   `ports [--serv-ip | -a] [--serv-net-port | -p] [--appliance-id | -L] [--blade-id | -B] [--port-id | -o]`,
	Short: "List all available blade ports",
	Long: `Queries the cfm-service for available appliance blade ports.
	Outputs a detailed summary of those ports to stdout.
	Note that, for any given item ID, if it is omitted, ALL items are collected\searched.`,
	Example: `
	cfm list blades ports --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId --port-id portId
	cfm list blades ports --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId
	cfm list blades ports --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --port-id portId
	cfm list blades ports --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId
	cfm list blades ports --serv-ip 127.0.0.1 --serv-net-port 8080 --blade-id bladeId --port-id portId
	cfm list blades ports --serv-ip 127.0.0.1 --serv-net-port 8080 --blade-id bladeId
	cfm list blades ports --serv-ip 127.0.0.1 --serv-net-port 8080 --port-id portId
	cfm list blades ports --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list blades ports -a 127.0.0.1 -p 8080 -L applId -B bladeId -o portId
	cfm list blades ports -a 127.0.0.1 -p 8080 -L applId -B bladeId
	cfm list blades ports -a 127.0.0.1 -p 8080 -L applId -o portId
	cfm list blades ports -a 127.0.0.1 -p 8080 -L applId
	cfm list blades ports -a 127.0.0.1 -p 8080 -B bladeId -o portId
	cfm list blades ports -a 127.0.0.1 -p 8080 -B bladeId
	cfm list blades ports -a 127.0.0.1 -p 8080 -o portId
	cfm list blades ports -a 127.0.0.1 -p 8080 `,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListBladePorts(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummary(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listBladesPortsCmd.DisableFlagsInUseLine = true

	listBladesPortsCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of a specific port. (default \"all ports returned.\")")

	initCommonBladeListCmdFlags(listBladesPortsCmd)

	//Add command to parent
	listBladesCmd.AddCommand(listBladesPortsCmd)
}
