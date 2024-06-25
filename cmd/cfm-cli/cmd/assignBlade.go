// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"cfm/cmd/cfm-cli/cmd/common"
	"fmt"

	"github.com/spf13/cobra"
)

var assignBladeCmd = &cobra.Command{
	Use:   "blade  [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L>  <--blade-id | -B> <--memory-id | -m> <--port-id | -o>",
	Short: "Assign an existing blade memory region to a blade port.",
	Long:  `Assign an existing blade memory region to a blade port.  A physical connection must already exist between the appliance blade and the target cxl-host.`,
	Example: `
	cfm assign blade --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId --memory-id memoryId --port-id portId

	cfm assign blade -a 127.0.0.1 -p 8080 -L applId  -B bladeId -m memoryId -o portId`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Assign Memory...")

		err := common.PromptYesNo(common.WARNING_CXL_HOST_POWER_DOWN)
		if err != nil {
			return
		}

		request := serviceRequests.NewServiceRequestBladesAssignMemory(cmd, "assign")
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummaryBladesAssignMemory(summary)
	},
}

func init() {
	assignBladeCmd.DisableFlagsInUseLine = true

	assignBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate")
	assignBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	assignBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of appliance blade to interrogate")
	assignBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	assignBladeCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of the appliance blade memory region to assign to the specified port.")
	assignBladeCmd.MarkFlagRequired(flags.MEMORY_ID)
	assignBladeCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of the appliance blade port to assign to the specified memory region.")
	assignBladeCmd.MarkFlagRequired(flags.PORT_ID)

	initCommonPersistentFlags(assignBladeCmd)

	//Add command to parent
	assignCmd.AddCommand(assignBladeCmd)
}
