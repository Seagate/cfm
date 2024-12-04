// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var resyncHostCmd = &cobra.Command{
	Use:     GetCmdUsageResyncHost(),
	Short:   `Resynchronize the cfm service to a single cxl host`,
	Long:    `Resynchronize the cfm service to a single cxl host.`,
	Example: GetCmdExampleResyncHost(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestResyncHost(cmd)
		host, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsHostAction(host, "resync")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	resyncHostCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(resyncHostCmd)

	resyncHostCmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "ID of CXL host\n")
	resyncHostCmd.MarkFlagRequired(flags.HOST_ID)

	//Add command to parent
	resyncCmd.AddCommand(resyncHostCmd)
}

// GetCmdUsageResyncHost - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageResyncHost() string {
	return fmt.Sprintf("%s %s %s",
		flags.HOST, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageHostId(false))
}

// GetCmdExampleResyncHost - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleResyncHost() string {
	baseCmd := fmt.Sprintf("cfm resync %s", flags.HOST)

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
