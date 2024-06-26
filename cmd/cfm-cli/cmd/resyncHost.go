// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var resyncHostCmd = &cobra.Command{
	Use:   `host [--serv-ip | -i] [--serv-net-port | -p] <--host-id | -H>`,
	Short: `Resynchronize the cfm service to a single cxl host`,
	Long:  `Resynchronize the cfm service to a single cxl host.`,
	Example: `
	cfm resync host --host-id hostId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm resync host -H hostId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
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

	resyncHostCmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "ID of CXL host")

	//Add command to parent
	resyncCmd.AddCommand(resyncHostCmd)
}
