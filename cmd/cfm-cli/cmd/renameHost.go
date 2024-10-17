// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"

	"github.com/spf13/cobra"
)

var renameHostCmd = &cobra.Command{
	Use:   `host [--serv-ip | -i] [--serv-net-port | -p] <--host-id | -H> <--new-id | -N>`,
	Short: `Rename a specific cxl host to a new ID`,
	Long:  `Rename a specific cxl host to a new ID.`,
	Example: `
	cfm rename host --host-id hostId --new-id newId --serv-ip 127.0.0.1 --serv-net-port 8080

	cfm rename host -H hostId -N newId -a 127.0.0.1 -p 8080`,
	Args: cobra.MatchAll(cobra.NoArgs),
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

	renameHostCmd.Flags().StringP(flags.HOST_ID, flags.HOST_ID_SH, flags.ID_DFLT, "Current CXL host ID")
	renameHostCmd.MarkFlagRequired(flags.HOST_ID)
	renameHostCmd.Flags().String(flags.NEW_ID, flags.NEW_ID_SH, "New CXL host ID")
	renameHostCmd.MarkFlagRequired(flags.NEW_ID)

	//Add command to parent
	resyncCmd.AddCommand(renameHostCmd)
}
