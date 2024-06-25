// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var listHostsMemoryDevicesCmd = &cobra.Command{
	Use:   `memory-devices [--serv-ip | -a] [--serv-net-port | -p] [--host-id | -H] [--memory-device-id | -d]`,
	Short: "List all available logical memory device(s) accessible to the host(s)",
	Long: `Queries the cfm-service for logical memory device(s) accessible to the host(s).
	Outputs a detailed summary of those memory regions to stdout.
	Note that, for any given item ID, if it is omitted, ALL items are collected\searched.`,
	Example: `
	cfm list hosts memory-devices --serv-ip 127.0.0.1 --serv-net-port 8080 --host-id hostId --memory-device-id memdevId
	cfm list hosts memory-devices --serv-ip 127.0.0.1 --serv-net-port 8080 --host-id hostId
	cfm list hosts memory-devices --serv-ip 127.0.0.1 --serv-net-port 8080 --memory-device-id memdevId
	cfm list hosts memory-devices --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list hosts memory-devices -a 127.0.0.1 -p 8080 -B hostId -d memdevId
	cfm list hosts memory-devices -a 127.0.0.1 -p 8080 -B hostId
	cfm list hosts memory-devices -a 127.0.0.1 -p 8080 -d memdevId
	cfm list hosts memory-devices -a 127.0.0.1 -p 8080 `,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListHostMemoryDevices(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummaryListMemoryDevices(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listHostsMemoryDevicesCmd.DisableFlagsInUseLine = true

	initCommonHostListCmdFlags(listHostsMemoryDevicesCmd)

	listHostsMemoryDevicesCmd.Flags().StringP(flags.MEMORY_DEVICE_ID, flags.MEMORY_DEVICE_ID_SH, flags.ID_DFLT, "ID of a specific memory device. (default \"all memory devices returned\")")

	//Add command to parent
	listHostsCmd.AddCommand(listHostsMemoryDevicesCmd)
}
