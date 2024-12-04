// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var renameHostCmd = &cobra.Command{
	Use:     GetCmdUsageRenameHost(),
	Short:   `Rename a specific cxl host to a new ID`,
	Long:    `Rename a specific cxl host to a new ID.`,
	Example: GetCmdExampleRenameHost(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestRenameHost(cmd)
		host, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsHostAction(host, "rename")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	renameHostCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(renameHostCmd)

	renameHostCmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "Current CXL host ID\n")
	renameHostCmd.MarkFlagRequired(flags.HOST_ID)
	renameHostCmd.Flags().StringP(flags.NEW_ID, flags.NEW_ID_SH, flags.ID_DFLT, "New CXL host ID\n")
	renameHostCmd.MarkFlagRequired(flags.NEW_ID)

	//Add command to parent
	renameCmd.AddCommand(renameHostCmd)
}

// GetCmdUsageRenameHost - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageRenameHost() string {
	return fmt.Sprintf("%s %s %s %s",
		flags.HOST, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageHostId(false),
		flags.GetOptionUsageNewId(false))
}

// GetCmdExampleRenameHost - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleRenameHost() string {
	baseCmd := fmt.Sprintf("cfm rename %s", flags.HOST)

	shorthandFormat := fmt.Sprintf("%s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShHostId(),
		flags.GetOptionExampleShNewId())

	longhandFormat := fmt.Sprintf("%s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhHostId(),
		flags.GetOptionExampleLhNewId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
