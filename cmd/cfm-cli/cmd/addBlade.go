// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var addBladeCmd = &cobra.Command{
	Use:     GetCmdUsageAddBlade(),
	Short:   "Add a memory appliance blade connection to cfm-service",
	Long:    `Adds a netowrk connection from the cfm-service to an external memory appliance blade.`,
	Example: GetCmdExampleAddBlade(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestAddBlade(cmd)
		addedBlade, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsBladeAction(addedBlade, "add")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	addBladeCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(addBladeCmd)

	addBladeCmd.Flags().StringP(flags.BLADE_USERNAME, flags.BLADE_USERNAME_SH, flags.BLADE_USERNAME_DFLT, "Blade username for authentication\n")
	// addBladeCmd.MarkFlagRequired(flags.BLADE_USERNAME)
	addBladeCmd.Flags().StringP(flags.BLADE_PASSWORD, flags.BLADE_PASSWORD_SH, flags.BLADE_PASSWORD_DFLT, "Blade password for authentication\n")
	// addBladeCmd.MarkFlagRequired(flags.BLADE_PASSWORD)

	addBladeCmd.Flags().StringP(flags.BLADE_NET_IP, flags.BLADE_NET_IP_SH, flags.BLADE_NET_IP_DFLT, "Blade network IP address\n")
	addBladeCmd.Flags().Uint16P(flags.BLADE_NET_PORT, flags.BLADE_NET_PORT_SH, flags.BLADE_NET_PORT_DFLT, "Blade network port\n")
	addBladeCmd.Flags().BoolP(flags.BLADE_INSECURE, flags.BLADE_INSECURE_SH, flags.BLADE_INSECURE_DFLT, "Blade insecure connection flag\n (default false)")
	addBladeCmd.Flags().StringP(flags.BLADE_PROTOCOL, flags.BLADE_PROTOCOL_SH, flags.BLADE_PROTOCOL_DFLT, "Blade network connection protocol (http/https)\n")

	addBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "Appliance ID of target blade\n")
	addBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	addBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "User-defined ID for target blade (Optional)\n (default: random w\\ format: blade-XXXX)")

	//Add command to parent
	addCmd.AddCommand(addBladeCmd)
}

// GetCmdUsageAddBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageAddBlade() string {
	return fmt.Sprintf("%s %s %s %s %s %s %s",
		flags.BLADE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageBladeId(true),
		flags.GetOptionUsageGroupBladeTcp(false),
		flags.GetOptionUsageBladeUsername(false),
		flags.GetOptionUsageBladePassword(false))
}

// GetCmdExampleAddBlade - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleAddBlade() string {
	baseCmd := fmt.Sprintf("cfm add %s", flags.BLADE)

	shorthandFormat := fmt.Sprintf("%s %s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShBladeId(),
		flags.GetOptionExampleShGroupBladeTcp(),
		flags.GetOptionExampleShBladeUsername(),
		flags.GetOptionExampleShBladePassword())

	longhandFormat := fmt.Sprintf("%s %s %s %s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId(),
		flags.GetOptionExampleLhBladeId(),
		flags.GetOptionExampleLhGroupBladeTcp(),
		flags.GetOptionExampleLhBladeUsername(),
		flags.GetOptionExampleLhBladePassword())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
