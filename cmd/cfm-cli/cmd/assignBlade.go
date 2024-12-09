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
	Use:     GetCmdUsageAssignBlade(),
	Short:   "Assign an existing blade memory region to a blade port.",
	Long:    `Assign an existing blade memory region to a blade port.  A physical connection must already exist between the appliance blade and the target cxl-host.`,
	Example: GetCmdExampleAssignBlade(),
	Args:    cobra.MatchAll(cobra.NoArgs),
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

	assignBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate\n")
	assignBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	assignBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of appliance blade to interrogate\n")
	assignBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	assignBladeCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of the appliance blade memory region to assign to the specified port\n")
	assignBladeCmd.MarkFlagRequired(flags.MEMORY_ID)
	assignBladeCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of the appliance blade port to assign to the specified memory region\n")
	assignBladeCmd.MarkFlagRequired(flags.PORT_ID)

	initCommonPersistentFlags(assignBladeCmd)

	//Add command to parent
	assignCmd.AddCommand(assignBladeCmd)
}

// GetCmdUsageAssignBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageAssignBlade() string {
	return fmt.Sprintf("%s %s %s %s %s %s",
		flags.BLADE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageBladeId(false),
		flags.GetOptionUsageMemoryId(false),
		flags.GetOptionUsagePortId(false))
}

// GetCmdExampleAssignBlade - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleAssignBlade() string {
	baseCmd := fmt.Sprintf("cfm assign %s", flags.BLADE)

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
