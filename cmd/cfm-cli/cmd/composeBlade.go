// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"cfm/cmd/cfm-cli/cmd/common"
	"fmt"

	"github.com/spf13/cobra"
)

var composeBladeCmd = &cobra.Command{
	Use:     GetCmdUsageComposeBlade(),
	Short:   "Compose a new memory region on the specified memory appliance blade.",
	Long:    `Compose a new memory region on the specified memory appliance blade.  The composed memory region can be optionally assigned to a specified memory appliance blade port for use by an external device (such as a cxl host).`,
	Example: GetCmdExampleComposeBlade(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Compose Memory...")

		request := serviceRequests.NewServiceRequestComposeMemory(cmd)

		if request.PortId.GetId() != "" {
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

	composeBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of appliance to interrogate\n")
	composeBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	composeBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of appliance blade to interrogate\n")
	composeBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	composeBladeCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of a specific appliance port\n")
	composeBladeCmd.Flags().StringP(flags.RESOURCE_SIZE, flags.RESOURCE_SIZE_SH, flags.SIZE_DFLT, "Total size of new, composed memory region(minimum of 1 GiB). Use 'G' or 'g' to specify GiB\n")
	composeBladeCmd.MarkFlagRequired(flags.RESOURCE_SIZE)
	composeBladeCmd.Flags().Int32P(flags.MEMORY_QOS, flags.MEMORY_QOS_SH, flags.MEMORY_QOS_DFLT, "Quality of Service level (ie: channel bandwidth)\n")
	composeBladeCmd.MarkFlagRequired(flags.MEMORY_QOS)

	initCommonPersistentFlags(composeBladeCmd)

	//Add command to parent
	composeCmd.AddCommand(composeBladeCmd)
}

// GetCmdUsageComposeBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageComposeBlade() string {
	return fmt.Sprintf("%s %s %s %s %s %s %s",
		flags.BLADE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageBladeId(false),
		flags.GetOptionUsagePortId(true),
		flags.GetOptionUsageResourceSize(false),
		flags.GetOptionUsageMemoryQos(false))
}

// GetCmdExampleComposeBlade - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleComposeBlade() string {
	baseCmd := fmt.Sprintf("cfm compose %s", flags.BLADE)

	shorthandFormat := fmt.Sprintf("%s %s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShBladeId(),
		flags.GetOptionExampleShPortId(),
		flags.GetOptionExampleShResourceSize(),
		flags.GetOptionExampleShMemoryQos())

	shorthandFormat2 := fmt.Sprintf("%s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShBladeId(),
		flags.GetOptionExampleShResourceSize(),
		flags.GetOptionExampleShMemoryQos())

	longhandFormat := fmt.Sprintf("%s %s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId(),
		flags.GetOptionExampleLhBladeId(),
		flags.GetOptionExampleLhPortId(),
		flags.GetOptionExampleLhResourceSize(),
		flags.GetOptionExampleLhMemoryQos())

	return fmt.Sprintf(`
	%s

	%s

	%s`, shorthandFormat, shorthandFormat2, longhandFormat)
}
