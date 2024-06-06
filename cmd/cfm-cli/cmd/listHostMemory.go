/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var listHostsMemoryCmd = &cobra.Command{
	Use:   `memory [--serv-ip | -a] [--serv-net-port | -p] [--host-id | -H] [--memory-id | -m]`,
	Short: "List all available composed memory region(s) accessible to the host(s)",
	Long: `Queries the cfm-service for composed memory region(s) accessible to the host(s).
	Outputs a detailed summary of those memory regions to stdout.
	Note that, for any given item ID, if it is omitted, ALL items are collected\searched.`,
	Example: `
	cfm list hosts memory --serv-ip 127.0.0.1 --serv-net-port 8080 --host-id hostId --memory-id memId
	cfm list hosts memory --serv-ip 127.0.0.1 --serv-net-port 8080 --host-id hostId
	cfm list hosts memory --serv-ip 127.0.0.1 --serv-net-port 8080 --memory-id memId
	cfm list hosts memory --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list hosts memory -a 127.0.0.1 -p 8080 -B hostId -m memId
	cfm list hosts memory -a 127.0.0.1 -p 8080 -B hostId
	cfm list hosts memory -a 127.0.0.1 -p 8080 -m memId
	cfm list hosts memory -a 127.0.0.1 -p 8080 `,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListHostMemoryRegions(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummaryListMemory(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listHostsMemoryCmd.DisableFlagsInUseLine = true

	initCommonHostListCmdFlags(listHostsMemoryCmd)

	listHostsMemoryCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of a specific memory region. (default \"all memory regions returned\")")

	//Add command to parent
	listHostsCmd.AddCommand(listHostsMemoryCmd)
}
