/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var deleteHostCmd = &cobra.Command{
	Use:   `host [--serv-ip | -i] [--serv-net-port | -p] <--host-id | -H>`,
	Short: `Delete a cxl host connection from cfm-service`,
	Long:  `Deletes a netowrk connection from the cfm-service to an external cxl host.`,
	Example: `
	cfm delete host --host-id hostId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm delete host -H hostId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestDeleteHost(cmd)
		deletedHost, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsAddDeleteHost(deletedHost, "deleted")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	deleteHostCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(deleteHostCmd)

	deleteHostCmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "ID of CXL host")

	//Add command to parent
	deleteCmd.AddCommand(deleteHostCmd)
}
