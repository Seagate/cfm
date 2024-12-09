// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var renameBladeCmd = &cobra.Command{
	Use:     GetCmdUsageRenameBlade(),
	Short:   "Rename a single blade ID on a specific composable memory appliance (CMA) to a new ID",
	Long:    `Rename a single blade ID on a specific composable memory appliance (CMA) to a new ID.`,
	Example: GetCmdExampleRenameBlade(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestRenameBlade(cmd)
		blade, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsBladeAction(blade, "rename")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	renameBladeCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(renameBladeCmd)

	renameBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of blade's composable memory appliance (CMA)\n")
	renameBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	renameBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "Current blade ID\n")
	renameBladeCmd.MarkFlagRequired(flags.BLADE_ID)
	renameBladeCmd.Flags().StringP(flags.NEW_ID, flags.NEW_ID_SH, flags.ID_DFLT, "New blade ID\n")
	renameBladeCmd.MarkFlagRequired(flags.NEW_ID)

	//Add command to parent
	renameCmd.AddCommand(renameBladeCmd)
}

// GetCmdUsageRenameBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageRenameBlade() string {
	return fmt.Sprintf("%s %s %s %s %s",
		flags.BLADE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageBladeId(false),
		flags.GetOptionUsageNewId(false))
}

// GetCmdExampleRenameBlade - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleRenameBlade() string {
	baseCmd := fmt.Sprintf("cfm rename %s", flags.BLADE)

	shorthandFormat := fmt.Sprintf("%s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShBladeId(),
		flags.GetOptionExampleShNewId())

	longhandFormat := fmt.Sprintf("%s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId(),
		flags.GetOptionExampleLhBladeId(),
		flags.GetOptionExampleLhNewId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
