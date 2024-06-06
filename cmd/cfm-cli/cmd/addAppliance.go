/*
Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var addApplianceCmd = &cobra.Command{
	Use:   `appliance [--serv-ip | -a] [--serv-net-port | -p] [--appliance-id | -L]`,
	Short: "Add a memory appliance to cfm-service",
	Long: `Create a virtual memeory appliance within the cfm-service.
           Use "cfm add blade" to add composable memory to the appliance.`,
	Example: `
	cfm add appliance --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id userDefinedId

	cfm add appliance -a 127.0.0.1 -p 8080 -L userDefinedId`,
	Args: cobra.MatchAll(cobra.NoArgs),
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

		serviceRequests.OutputResultsAddDeleteAppliance(addedAppliance, "added")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	addApplianceCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(addApplianceCmd)

	// Unused by user, but some values are required by cfm-service frontend.
	addApplianceCmd.Flags().StringP(flags.DEVICE_USERNAME, flags.DEVICE_USERNAME_SH, flags.APPLIANCE_USERNAME_DFLT, "Appliance username for authentication")
	addApplianceCmd.Flags().StringP(flags.DEVICE_PASSWORD, flags.DEVICE_PASSWORD_SH, flags.APPLIANCE_PASSWORD_DFLT, "Appliance password for authentication")

	addApplianceCmd.Flags().StringP(flags.DEVICE_NET_IP, flags.DEVICE_NET_IP_SH, flags.APPLIANCE_NET_IP_DFLT, "Appliance network IP address")
	addApplianceCmd.Flags().Uint16P(flags.DEVICE_NET_PORT, flags.DEVICE_NET_PORT_SH, flags.APPLIANCE_NET_PORT_DFLT, "Appliance network port ")
	addApplianceCmd.Flags().BoolP(flags.DEVICE_INSECURE, flags.DEVICE_INSECURE_SH, flags.APPLIANCE_INSECURE_DFLT, "Appliance insecure connection flag")
	addApplianceCmd.Flags().StringP(flags.DEVICE_PROTOCOL, flags.DEVICE_PROTOCOL_SH, flags.APPLIANCE_PROTOCOL_DFLT, "Appliance network connection protocol (http/https)")

	addApplianceCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "User-defined ID for target appliance")

	//Add command to parent
	addCmd.AddCommand(addApplianceCmd)
}
