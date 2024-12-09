// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteHostCmd = &cobra.Command{
	Use:     GetCmdUsageDeleteHost(),
	Short:   `Delete a cxl host connection from cfm-service`,
	Long:    `Deletes a netowrk connection from the cfm-service to an external cxl host.`,
	Example: GetCmdExampleDeleteHost(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestDeleteHost(cmd)
		deletedHost, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsHostAction(deletedHost, "delete")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	deleteHostCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(deleteHostCmd)

	deleteHostCmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "ID of CXL host\n")
	deleteHostCmd.MarkFlagRequired(flags.HOST_ID)

	//Add command to parent
	deleteCmd.AddCommand(deleteHostCmd)
}

// GetCmdUsageDeleteHost - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageDeleteHost() string {
	return fmt.Sprintf("%s %s %s",
		flags.HOST, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageHostId(false))
}

// GetCmdExampleDeleteHost - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleDeleteHost() string {
	baseCmd := fmt.Sprintf("cfm delete %s", flags.HOST)

	shorthandFormat := fmt.Sprintf("%s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShHostId())

	longhandFormat := fmt.Sprintf("%s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhHostId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
