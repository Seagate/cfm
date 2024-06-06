/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates
*/
package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"cfm/cmd/cfm-cli/cmd/common"
	"fmt"

	"github.com/spf13/cobra"
)

var unassignBladeCmd = &cobra.Command{
	Use:   "blade  [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L>  <--blade-id | -B> <--memory-id | -m> <--port-id | -o>",
	Short: "Unassign an existing blade memory region from a blade port.",
	Long:  `Unassign an existing blade memory region from a blade port.`,
	Example: `
	cfm unassign blade --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId --memory-id memoryId --port-id portId

	cfm unassign blade -a 127.0.0.1 -p 8080 -L applId  -B bladeId -m memoryId -o portId`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Unassign Memory...")

		err := common.PromptYesNo(common.WARNING_CXL_HOST_POWER_DOWN)
		if err != nil {
			cobra.CheckErr(err)
		}

		request := serviceRequests.NewServiceRequestBladesAssignMemory(cmd, "unassign")
		summary, err := request.Execute()
		if err != nil {
			return
		}

		request.OutputSummaryBladesAssignMemory(summary)
	},
}

func init() {
	unassignBladeCmd.DisableFlagsInUseLine = true

	unassignBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate")
	unassignBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	unassignBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of appliance blade to interrogate")
	unassignBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	unassignBladeCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of the appliance blade memory region to unassign from the specified port.")
	unassignBladeCmd.MarkFlagRequired(flags.MEMORY_ID)
	unassignBladeCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of the appliance blade port to unassign from the specified memory region.")
	unassignBladeCmd.MarkFlagRequired(flags.PORT_ID)

	initCommonPersistentFlags(unassignBladeCmd)

	//Add command to parent
	unassignCmd.AddCommand(unassignBladeCmd)
}
