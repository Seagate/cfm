// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var renameApplianceCmd = &cobra.Command{
	Use:     GetCmdUsageRenameAppliance(),
	Short:   "Rename a specific composable memory appliance (CMA) to a new ID",
	Long:    `Rename a specific composable memory appliance (CMA) to a new ID.`,
	Example: GetCmdExampleRenameAppliance(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestRenameAppliance(cmd)
		appliance, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsApplianceAction(appliance, "rename")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	renameApplianceCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(renameApplianceCmd)

	renameApplianceCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "Current ID of composable memory appliance (CMA)\n")
	renameApplianceCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	renameApplianceCmd.Flags().StringP(flags.NEW_ID, flags.NEW_ID_SH, flags.ID_DFLT, "New ID of composable memory appliance (CMA)\n")
	renameApplianceCmd.MarkFlagRequired(flags.NEW_ID)

	//Add command to parent
	renameCmd.AddCommand(renameApplianceCmd)
}

// GetCmdUsageRenameAppliance - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageRenameAppliance() string {
	return fmt.Sprintf("%s %s %s %s",
		flags.APPLIANCE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageNewId(false))
}

// GetCmdExampleRenameAppliance - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleRenameAppliance() string {
	baseCmd := fmt.Sprintf("cfm rename %s", flags.APPLIANCE)

	shorthandFormat := fmt.Sprintf("%s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShNewId())

	longhandFormat := fmt.Sprintf("%s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId(),
		flags.GetOptionExampleLhNewId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
