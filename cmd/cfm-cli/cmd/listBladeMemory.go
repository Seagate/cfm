// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var listBladesMemoryCmd = &cobra.Command{
	Use:   `memory [--serv-ip | -a] [--serv-net-port | -p] [--appliance-id | -L] [--blade-id | -B] [--memory-id | -m]`,
	Short: "List all available blade composed and\\or provisioned memory regions",
	Long: `Queries the cfm-service for available composed\\provisioned memory regions.
	Outputs a detailed summary of those memory regions to stdout.
	Note that, for any given item ID, if it is omitted, ALL items are collected\searched.`,
	Example: `
	cfm list blades memory --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId --memory-id memId
	cfm list blades memory --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId
	cfm list blades memory --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --memory-id memId
	cfm list blades memory --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId
	cfm list blades memory --serv-ip 127.0.0.1 --serv-net-port 8080 --blade-id bladeId --memory-id memId
	cfm list blades memory --serv-ip 127.0.0.1 --serv-net-port 8080 --blade-id bladeId
	cfm list blades memory --serv-ip 127.0.0.1 --serv-net-port 8080 --memory-id memId
	cfm list blades memory --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm list blades memory -a 127.0.0.1 -p 8080 -L applId -B bladeId -m memId
	cfm list blades memory -a 127.0.0.1 -p 8080 -L applId -B bladeId
	cfm list blades memory -a 127.0.0.1 -p 8080 -L applId -m memId
	cfm list blades memory -a 127.0.0.1 -p 8080 -L applId
	cfm list blades memory -a 127.0.0.1 -p 8080 -B bladeId -m memId
	cfm list blades memory -a 127.0.0.1 -p 8080 -B bladeId
	cfm list blades memory -a 127.0.0.1 -p 8080 -m memId
	cfm list blades memory -a 127.0.0.1 -p 8080 `,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListMemoryRegions(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummaryListMemory(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listBladesMemoryCmd.DisableFlagsInUseLine = true

	listBladesMemoryCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of a specific memory region. (default \"all memory regions returned\")")

	initCommonBladeListCmdFlags(listBladesMemoryCmd)

	//Add command to parent
	listBladesCmd.AddCommand(listBladesMemoryCmd)
}

func testHelp(cmd *cobra.Command, test []string) {
	fmt.Println("A new help line")
	fmt.Println("Another new help line")
}
