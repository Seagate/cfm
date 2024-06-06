/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <device> [<flags>]",
	Short: "Deletes a connection to an external device",
	Long:  `Closes a TCPIP connection from the cfm-service to an external device (a memory appliance -OR- a cxl host).`,
	Args:  cobra.MatchAll(cobra.NoArgs),
}

// Here you will define your flags and configuration settings.
func init() {
	deleteCmd.DisableFlagsInUseLine = true

	rootCmd.AddCommand(deleteCmd)
}
