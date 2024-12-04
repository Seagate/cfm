// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var resyncApplianceCmd = &cobra.Command{
	Use:     GetCmdUsageResyncAppliance(),
	Short:   "Resynchronize the cfm service to all the added blades for a specific composable memory appliance (CMA)",
	Long:    `Resynchronize the cfm service to all the added blades for a specific composable memory appliance (CMA).`,
	Example: GetCmdExampleResyncAppliance(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestResyncAppliance(cmd)
		appliance, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsApplianceAction(appliance, "resync")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	resyncApplianceCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(resyncApplianceCmd)

	resyncApplianceCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of composable memory appliance (CMA)\n")
	resyncApplianceCmd.MarkFlagRequired(flags.APPLIANCE_ID)

	//Add command to parent
	resyncCmd.AddCommand(resyncApplianceCmd)
}

// GetCmdUsageResyncAppliance - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageResyncAppliance() string {
	return fmt.Sprintf("%s %s %s",
		flags.APPLIANCE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false))
}

// GetCmdExampleResyncAppliance - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleResyncAppliance() string {
	baseCmd := fmt.Sprintf("cfm resync %s", flags.APPLIANCE)

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
