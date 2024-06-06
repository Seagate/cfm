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

var composeBladeCmd = &cobra.Command{
	Use:   "blade  [--serv-ip | -a] [--serv-net-port | -p] <--appliance-id | -L>  <--blade-id | -B> [--port-id | -i] <--resource-size | -z> [--memory-qos | -q]",
	Short: "Compose a new memory region on the specified memory appliance blade.",
	Long:  `Compose a new memory region on the specified memory appliance blade.  The composed memory region can be optionally assigned to a specified memory appliance blade port for use by an external device (such as a cxl host).`,
	Example: `
	cfm compose blade --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id bladeId --port-id portId --resource-size 32g --memory-qos 4

	cfm compose blade -a 127.0.0.1 -p 8080 -L applId  -B bladeId -o portId -z 32g -q 4`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Compose Memory...")

		request := serviceRequests.NewServiceRequestComposeMemory(cmd)

		if request.GetPortId() != "" {
			err := common.PromptYesNo(common.WARNING_CXL_HOST_POWER_DOWN)
			if err != nil {
				return
			}
		}

		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResultsComposedMemory(summary)
	},
}

func init() {
	composeBladeCmd.DisableFlagsInUseLine = true

	composeBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate")
	composeBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	composeBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of appliance blade to interrogate")
	composeBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	composeBladeCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of a specific appliance port.")
	composeBladeCmd.Flags().StringP(flags.RESOURCE_SIZE, flags.RESOURCE_SIZE_SH, flags.SIZE_DFLT, "Total size of new, composed memory region(minimum of 1 GiB). Use 'G' or 'g' to specific GiB.")
	composeBladeCmd.MarkFlagRequired(flags.RESOURCE_SIZE)
	composeBladeCmd.Flags().Int32P(flags.MEMORY_QOS, flags.MEMORY_QOS_SH, flags.MEMORY_QOS_DFLT, "NOT YET SUPPORTED, BUT, for now, cfm-service ***REQUIRES*** this to be 4: Quality of Service level.")

	initCommonPersistentFlags(composeBladeCmd)

	//Add command to parent
	composeCmd.AddCommand(composeBladeCmd)
}
