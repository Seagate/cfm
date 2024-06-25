/*
Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var addBladeCmd = &cobra.Command{
	Use: `blade <--appliance-id | -L> [--dev-username | -R] [--dev-password | -W] [--serv-ip | -a] [--serv-net-port | -p]
		[--blade-id | -B] [--dev-ip | -A] [--dev-net-port | -P] [--dev-insecure | -S] [--dev-protocol | -T]`,
	Short: "Add a memory appliance blade connection to cfm-service",
	Long:  `Adds a netowrk connection from the cfm-service to an external memory appliance blade.`,
	Example: `
	cfm add blade --serv-ip 127.0.0.1 --serv-net-port 8080 --appliance-id applId --blade-id userDefinedId --dev-username user --dev-password pswd --dev-ip 127.0.0.1 --dev-net-port 7443 --dev-insecure --dev-protocol https

	cfm add blade -a 127.0.0.1 -p 8080 -L applId -B userDefinedId -R user -W pswd -A 127.0.0.1 -P 7443 -S -T https`,
	Args: cobra.MatchAll(cobra.NoArgs),
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

	addBladeCmd.Flags().StringP(flags.DEVICE_USERNAME, flags.DEVICE_USERNAME_SH, flags.BLADE_USERNAME_DFLT, "Blade username for authentication")
	// addBladeCmd.MarkFlagRequired(flags.DEVICE_USERNAME)
	addBladeCmd.Flags().StringP(flags.DEVICE_PASSWORD, flags.DEVICE_PASSWORD_SH, flags.BLADE_PASSWORD_DFLT, "Blade password for authentication")
	// addBladeCmd.MarkFlagRequired(flags.DEVICE_PASSWORD)

	addBladeCmd.Flags().StringP(flags.DEVICE_NET_IP, flags.DEVICE_NET_IP_SH, flags.BLADE_NET_IP_DFLT, "Blade network IP address")
	addBladeCmd.Flags().Uint16P(flags.DEVICE_NET_PORT, flags.DEVICE_NET_PORT_SH, flags.BLADE_NET_PORT_DFLT, "Blade network port ")
	addBladeCmd.Flags().BoolP(flags.DEVICE_INSECURE, flags.DEVICE_INSECURE_SH, flags.BLADE_INSECURE_DFLT, "Blade insecure connection flag")
	addBladeCmd.Flags().StringP(flags.DEVICE_PROTOCOL, flags.DEVICE_PROTOCOL_SH, flags.BLADE_PROTOCOL_DFLT, "Blade network connection protocol (http/https)")

	addBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "Appliance ID of target blade")
	addBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	addBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "User-defined ID for target blade")

	//Add command to parent
	addCmd.AddCommand(addBladeCmd)
}
