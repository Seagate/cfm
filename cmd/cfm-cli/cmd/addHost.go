// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

// addHostCmd represents the addHost command
var addHostCmd = &cobra.Command{
	Use: `host [--dev-username | -R] [--dev-password | -W] [--serv-ip | -a] [--serv-net-port | -p]
			[--host-id | -H] [--dev-ip | -A] [--dev-net-port | -P] [--dev-insecure | -S] [--dev-protocol | -T]`,
	Short: "Add a cxl host connection to cfm-service",
	Long:  `Adds a netowrk connection from the cfm-service to an external cxl host.`,
	Example: `
	cfm add host --serv-ip 127.0.0.1 --serv-net-port 8080 --host-id userDefinedId --dev-username user --dev-password pswd --dev-ip 127.0.0.1 --dev-net-port 7443 --dev-insecure --dev-protocol https

	cfm add host -a 127.0.0.1 -p 8080 -H userDefinedId -R user -W pswd -A 127.0.0.1 -P 7443 -S -T https`,
	Args: cobra.MatchAll(cobra.NoArgs),
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

	addHostCmd.Flags().StringP(flags.DEVICE_USERNAME, flags.DEVICE_USERNAME_SH, flags.HOST_USERNAME_DFLT, "Host username for authentication")
	// addHostCmd.MarkFlagRequired(flags.DEVICE_USERNAME)
	addHostCmd.Flags().StringP(flags.DEVICE_PASSWORD, flags.DEVICE_PASSWORD_SH, flags.HOST_PASSWORD_DFLT, "Host password for authentication")
	// addHostCmd.MarkFlagRequired(flags.DEVICE_PASSWORD)

	addHostCmd.Flags().StringP(flags.DEVICE_NET_IP, flags.DEVICE_NET_IP_SH, flags.HOST_NET_IP_DFLT, "Host network IP address")
	addHostCmd.Flags().Uint16P(flags.DEVICE_NET_PORT, flags.DEVICE_NET_PORT_SH, flags.HOST_NET_PORT_DFLT, "Host network port ")
	addHostCmd.Flags().BoolP(flags.DEVICE_INSECURE, flags.DEVICE_INSECURE_SH, flags.HOST_INSECURE_DFLT, "Host insecure connection flag")
	addHostCmd.Flags().StringP(flags.DEVICE_PROTOCOL, flags.DEVICE_PROTOCOL_SH, flags.HOST_PROTOCOL_DFLT, "Host network connection protocol (http/https)")

	addHostCmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "User-defined ID for target host")

	//Add command to parent
	addCmd.AddCommand(addHostCmd)
}
