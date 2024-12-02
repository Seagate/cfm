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
	Use:     GetCmdUsageFreeBlade(),
	Short:   "Free an existing memory region on the specified memory appliance blade.",
	Long:    `Free an existing memory region on the specified memory appliance blade.  The blade port (if present) will be unassigned and the memory region's resource blocks will be deallocated.`,
	Example: GetCmdExampleFreeBlade(),
	Args:    cobra.MatchAll(cobra.NoArgs),
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

	freeBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate\n")
	freeBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	freeBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of appliance blade to interrogate\n")
	freeBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	freeBladeCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of a specific appliance blade memory region to free\n")
	freeBladeCmd.MarkFlagRequired(flags.MEMORY_ID)

	initCommonPersistentFlags(freeBladeCmd)

	//Add command to parent
	freeCmd.AddCommand(freeBladeCmd)
}

// GetCmdUsageFreeBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageFreeBlade() string {
	return fmt.Sprintf("%s %s %s %s %s",
		flags.BLADE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageBladeId(false),
		flags.GetOptionUsageMemoryId(false))
}

// GetCmdExampleFreeBlade - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleFreeBlade() string {
	baseCmd := fmt.Sprintf("cfm free %s", flags.BLADE)

	shorthandFormat := fmt.Sprintf("%s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShBladeId(),
		flags.GetOptionExampleShMemoryId())

	longhandFormat := fmt.Sprintf("%s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId(),
		flags.GetOptionExampleLhBladeId(),
		flags.GetOptionExampleLhMemoryId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
