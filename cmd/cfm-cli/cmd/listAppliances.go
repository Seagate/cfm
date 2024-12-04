// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var listAppliancesCmd = &cobra.Command{
	Use:     GetCmdUsageListAppliances(),
	Short:   "List all recognized memory appliance(s)",
	Long:    `Queries the cfm-service for all recognized memory appliance(s) and outputs a summary to stdout.`,
	Example: GetCmdExampleListAppliances(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListAppliances(cmd)
		appliances, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(appliances)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listAppliancesCmd.DisableFlagsInUseLine = true

	//Just need service flags here
	initCommonPersistentFlags(listAppliancesCmd)

	//Add command to parent
	listCmd.AddCommand(listAppliancesCmd)
}

// GetCmdUsageListAppliances - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageListAppliances() string {
	return fmt.Sprintf("%s %s",
		flags.APPLIANCES, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false))
}

// GetCmdExampleListAppliances - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleListAppliances() string {
	baseCmd := fmt.Sprintf("cfm list %s", flags.APPLIANCES)

	shorthandFormat := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp())

	longhandFormat := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
