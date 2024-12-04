// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteApplianceCmd = &cobra.Command{
	Use:     GetCmdUsageDeleteAppliance(),
	Short:   "Delete a memory appliance connection from cfm-service",
	Long:    `Deletes a netowrk connection from the cfm-service to an external memory appliance.`,
	Example: GetCmdExampleDeleteAppliance(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestDeleteAppliance(cmd)
		deletedAppliance, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsApplianceAction(deletedAppliance, "delete")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	deleteApplianceCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(deleteApplianceCmd)

	deleteApplianceCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of memory appliance\n")
	deleteApplianceCmd.MarkFlagRequired(flags.APPLIANCE_ID)

	//Add command to parent
	deleteCmd.AddCommand(deleteApplianceCmd)
}

// GetCmdUsageDeleteAppliance - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageDeleteAppliance() string {
	return fmt.Sprintf("%s %s %s",
		flags.APPLIANCE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false))
}

// GetCmdExampleDeleteAppliance - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleDeleteAppliance() string {
	baseCmd := fmt.Sprintf("cfm delete %s", flags.APPLIANCE)

	shorthandFormat := fmt.Sprintf("%s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId())

	longhandFormat := fmt.Sprintf("%s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
