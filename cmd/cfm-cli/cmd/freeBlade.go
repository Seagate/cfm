// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"cfm/cmd/cfm-cli/cmd/common"
	"fmt"

	"github.com/spf13/cobra"
)

var freeBladeCmd = &cobra.Command{
	Use:   "blade [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L> <--blade-id | -B> <--memory-id | -i>",
	Short: "Free an existing memory region on the specified memory appliance blade.",
	Long:  `Free an existing memory region on the specified memory appliance blade.  The blade port will be unassigned and the memory region's resource blocks will be deallocated.`,
	Example: `
	cfm free --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId --memory-id memoryId

	cfm free -a 127.0.0.1 -p 8080 -L applId -B bladeId -m memoryId`,
	Args: cobra.MatchAll(cobra.NoArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Free Memory...")

		request := serviceRequests.NewServiceRequestFreeMemory(cmd)

		err := common.PromptYesNo(common.WARNING_CXL_HOST_POWER_DOWN)
		if err != nil {
			return
		}

		region, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResultsFreeMemory(region)
	},
}

func init() {
	freeBladeCmd.DisableFlagsInUseLine = true

	freeBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate")
	freeBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	freeBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of appliance blade to interrogate")
	freeBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	freeBladeCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of a specific appliance blade memory region to free.")
	freeBladeCmd.MarkFlagRequired(flags.MEMORY_ID)

	initCommonPersistentFlags(freeBladeCmd)

	//Add command to parent
	freeCmd.AddCommand(freeBladeCmd)
}
