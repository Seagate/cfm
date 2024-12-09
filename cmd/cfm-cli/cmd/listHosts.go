// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

// listHostsCmd represents the listHosts command
var listHostsCmd = &cobra.Command{
	Use:     GetCmdUsageListHosts(),
	Short:   "List all recognized cxl host(s)",
	Long:    `Queries the cfm-service for all recognized cxl host(s) and outputs a summary to stdout.`,
	Example: GetCmdExampleListHosts(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListHosts(cmd)
		hosts, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(hosts)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listHostsCmd.DisableFlagsInUseLine = true

	//Just need service flags here
	initCommonPersistentFlags(listHostsCmd)

	//Add command to parent
	listCmd.AddCommand(listHostsCmd)
}

// GetCmdUsageListHost - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageListHosts() string {
	return fmt.Sprintf("%s %s",
		flags.HOSTS, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false))
}

// GetCmdExampleListHosts - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleListHosts() string {
	baseCmd := fmt.Sprintf("cfm list %s", flags.HOSTS)

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
