// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteBladeCmd = &cobra.Command{
	Use:     GetCmdUsageDeleteBlade(),
	Short:   "Delete a memory appliance blade connection from cfm-service",
	Long:    `Deletes a netowrk connection from the cfm-service to an external memory appliance blade.`,
	Example: GetCmdExampleDeleteBlade(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestDeleteBlade(cmd)
		deletedBlade, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsBladeAction(deletedBlade, "delete")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	deleteBladeCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(deleteBladeCmd)

	deleteBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of blade's appliance\n")
	deleteBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	deleteBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of blade to delete\n")
	deleteBladeCmd.MarkFlagRequired(flags.BLADE_ID)

	//Add command to parent
	deleteCmd.AddCommand(deleteBladeCmd)
}

// GetCmdUsageDeleteBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageDeleteBlade() string {
	return fmt.Sprintf("%s %s %s %s",
		flags.BLADE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageBladeId(false))
}

// GetCmdExampleDeleteBlade - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleDeleteBlade() string {
	baseCmd := fmt.Sprintf("cfm delete %s", flags.BLADE)

	shorthandFormat := fmt.Sprintf("%s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShBladeId())

	longhandFormat := fmt.Sprintf("%s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId(),
		flags.GetOptionExampleLhBladeId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
