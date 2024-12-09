// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"cfm/cmd/cfm-cli/cmd/common"
	"fmt"

	"github.com/spf13/cobra"
)

var unassignBladeCmd = &cobra.Command{
	Use:     GetCmdUsageUnassignBlade(),
	Short:   "Unassign an existing blade memory region from a blade port.",
	Long:    `Unassign an existing blade memory region from a blade port.`,
	Example: GetCmdExampleUnassignBlade(),
	Args:    cobra.MatchAll(cobra.NoArgs),
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

	unassignBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate\n")
	unassignBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	unassignBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of appliance blade to interrogate\n")
	unassignBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	unassignBladeCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of the appliance blade memory region to unassign from the specified port\n")
	unassignBladeCmd.MarkFlagRequired(flags.MEMORY_ID)
	unassignBladeCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of the appliance blade port to unassign from the specified memory region\n")
	unassignBladeCmd.MarkFlagRequired(flags.PORT_ID)

	initCommonPersistentFlags(unassignBladeCmd)

	//Add command to parent
	unassignCmd.AddCommand(unassignBladeCmd)
}

// GetCmdUsageUnassignBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageUnassignBlade() string {
	return fmt.Sprintf("%s %s %s %s %s %s",
		flags.BLADE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageBladeId(false),
		flags.GetOptionUsageMemoryId(false),
		flags.GetOptionUsagePortId(false))
}

// GetCmdExampleUnassignBlade - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleUnassignBlade() string {
	baseCmd := fmt.Sprintf("cfm unassign %s", flags.BLADE)

	shorthandFormat := fmt.Sprintf("%s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShBladeId(),
		flags.GetOptionExampleShMemoryId(),
		flags.GetOptionExampleShPortId())

	longhandFormat := fmt.Sprintf("%s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId(),
		flags.GetOptionExampleLhBladeId(),
		flags.GetOptionExampleLhMemoryId(),
		flags.GetOptionExampleLhPortId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
