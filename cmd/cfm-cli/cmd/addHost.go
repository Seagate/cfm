// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

// addHostCmd represents the addHost command
var addHostCmd = &cobra.Command{
	Use:     GetCmdUsageAddHost(),
	Short:   "Add a cxl host connection to cfm-service",
	Long:    `Adds a netowrk connection from the cfm-service to an external cxl host.`,
	Example: GetCmdExampleAddHost(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestAddHost(cmd)
		addedHost, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsHostAction(addedHost, "add")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	addHostCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(addHostCmd)

	addHostCmd.Flags().StringP(flags.HOST_USERNAME, flags.HOST_USERNAME_SH, flags.HOST_USERNAME_DFLT, "Host(CXL) username for authentication\n")
	// addHostCmd.MarkFlagRequired(flags.HOST_USERNAME)
	addHostCmd.Flags().StringP(flags.HOST_PASSWORD, flags.HOST_PASSWORD_SH, flags.HOST_PASSWORD_DFLT, "Host(CXL) password for authentication\n")
	// addHostCmd.MarkFlagRequired(flags.HOST_PASSWORD)

	addHostCmd.Flags().StringP(flags.HOST_NET_IP, flags.HOST_NET_IP_SH, flags.HOST_NET_IP_DFLT, "Host(CXL) network IP address\n")
	addHostCmd.Flags().Uint16P(flags.HOST_NET_PORT, flags.HOST_NET_PORT_SH, flags.HOST_NET_PORT_DFLT, "Host(CXL) network port\n")
	addHostCmd.Flags().BoolP(flags.HOST_INSECURE, flags.HOST_INSECURE_SH, flags.HOST_INSECURE_DFLT, "Host(CXL) insecure connection flag\n (default false)")
	addHostCmd.Flags().StringP(flags.HOST_PROTOCOL, flags.HOST_PROTOCOL_SH, flags.HOST_PROTOCOL_DFLT, "Host(CXL) network connection protocol (http/https)\n")

	addHostCmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "User-defined ID for target host(CXL) (Optional)\n (default: random w\\ format: host-XXXX)")

	//Add command to parent
	addCmd.AddCommand(addHostCmd)
}

// GetCmdUsageAddHost - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageAddHost() string {
	return fmt.Sprintf("%s %s %s %s %s %s",
		flags.HOST, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageHostId(true),
		flags.GetOptionUsageGroupHostTcp(false),
		flags.GetOptionUsageHostUsername(false),
		flags.GetOptionUsageHostPassword(false))
}

// GetCmdExampleAddHost - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleAddHost() string {
	baseCmd := fmt.Sprintf("cfm add %s", flags.HOST)

	shorthandFormat := fmt.Sprintf("%s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShHostId(),
		flags.GetOptionExampleShGroupHostTcp(),
		flags.GetOptionExampleShHostUsername(),
		flags.GetOptionExampleShHostPassword())

	longhandFormat := fmt.Sprintf("%s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhHostId(),
		flags.GetOptionExampleLhGroupHostTcp(),
		flags.GetOptionExampleLhHostUsername(),
		flags.GetOptionExampleLhHostPassword())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
