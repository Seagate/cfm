// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var addApplianceCmd = &cobra.Command{
	Use:   GetCmdUsageAddAppliance(),
	Short: "Add a memory appliance to cfm-service",
	Long: `Create a virtual memory appliance within the cfm-service.
           Use "cfm add blade" to add composable memory to the appliance.`,
	Example: GetCmdExampleAddAppliance(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestAddAppliance(cmd)
		addedAppliance, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsApplianceAction(addedAppliance, "add")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	addApplianceCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(addApplianceCmd)

	addApplianceCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "User-defined ID for target appliance (Optional)\n (default: random w\\ format: memory-appliance-XXXX)")

	// Unused by user, but values are required by cfm-service client frontend.
	addApplianceCmd.Flags().StringP(flags.APPLIANCE_USERNAME, flags.APPLIANCE_USERNAME_SH, flags.APPLIANCE_USERNAME_DFLT, "Appliance username for authentication\n")
	addApplianceCmd.Flags().StringP(flags.APPLIANCE_PASSWORD, flags.APPLIANCE_PASSWORD_SH, flags.APPLIANCE_PASSWORD_DFLT, "Appliance password for authentication\n")
	addApplianceCmd.Flags().MarkHidden(flags.APPLIANCE_USERNAME)
	addApplianceCmd.Flags().MarkHidden(flags.APPLIANCE_PASSWORD)

	addApplianceCmd.Flags().StringP(flags.APPLIANCE_NET_IP, flags.APPLIANCE_NET_IP_SH, flags.APPLIANCE_NET_IP_DFLT, "Appliance network IP address\n")
	addApplianceCmd.Flags().Uint16P(flags.APPLIANCE_NET_PORT, flags.APPLIANCE_NET_PORT_SH, flags.APPLIANCE_NET_PORT_DFLT, "Appliance network port\n")
	addApplianceCmd.Flags().BoolP(flags.APPLIANCE_INSECURE, flags.APPLIANCE_INSECURE_SH, flags.APPLIANCE_INSECURE_DFLT, "Appliance insecure connection flag\n (default false)")
	addApplianceCmd.Flags().StringP(flags.APPLIANCE_PROTOCOL, flags.APPLIANCE_PROTOCOL_SH, flags.APPLIANCE_PROTOCOL_DFLT, "Appliance network connection protocol (http/https)\n")
	addApplianceCmd.Flags().MarkHidden(flags.APPLIANCE_NET_IP)
	addApplianceCmd.Flags().MarkHidden(flags.APPLIANCE_NET_PORT)
	addApplianceCmd.Flags().MarkHidden(flags.APPLIANCE_INSECURE)
	addApplianceCmd.Flags().MarkHidden(flags.APPLIANCE_PROTOCOL)

	//Add command to parent
	addCmd.AddCommand(addApplianceCmd)
}

// GetCmdUsageAddAppliance - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageAddAppliance() string {
	return fmt.Sprintf("%s %s %s",
		flags.APPLIANCE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(true))
}

// GetCmdExampleAddAppliance - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleAddAppliance() string {
	baseCmd := fmt.Sprintf("cfm add %s", flags.APPLIANCE)

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
