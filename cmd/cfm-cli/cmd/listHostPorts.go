// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var listHostsPortsCmd = &cobra.Command{
	Use:   `ports [--serv-ip | -a] [--serv-net-port | -p] [-host-id | -H] [--port-id | -o]`,
	Short: "List all available port(s) accessible to the host(s).",
	Long: `Queries the cfm-service for port(s) accessible to the host(s).
	Outputs a detailed summary of those ports to stdout.
	Note that, for any given item ID, if it is omitted, ALL items are collected\searched.`,
	Example: `
	cfm list hosts ports --serv-ip 127.0.0.1 --serv-net-port 8080 --host-id hostId --port-id portId
	cfm list hosts ports --serv-ip 127.0.0.1 --serv-net-port 8080 --host-id hostId
	cfm list hosts ports --serv-ip 127.0.0.1 --serv-net-port 8080 --port-id portId
	cfm list hosts ports --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list hosts ports -a 127.0.0.1 -p 8080 -B hostId -o portId
	cfm list hosts ports -a 127.0.0.1 -p 8080 -B hostId
	cfm list hosts ports -a 127.0.0.1 -p 8080 -o portId
	cfm list hosts ports -a 127.0.0.1 -p 8080 `,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListHostPorts(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummary(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listHostsPortsCmd.DisableFlagsInUseLine = true

	initCommonHostListCmdFlags(listHostsPortsCmd)

	listHostsPortsCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of a specific port. (default \"all ports returned.\")")

	//Add command to parent
	listHostsCmd.AddCommand(listHostsPortsCmd)
}
