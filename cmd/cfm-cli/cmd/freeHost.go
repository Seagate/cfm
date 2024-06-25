// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"github.com/spf13/cobra"
)

var freeHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Not Implemented.", // free memory on an appliance via a cxl-host
	//TODO: Add "Long" and "Examples"
	Args: cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr("Not Implemented")
	},
}

func init() {
	freeHostCmd.DisableFlagsInUseLine = true

	//TODO: Add flags

	initCommonPersistentFlags(freeHostCmd)

	//Add command to parent
	freeCmd.AddCommand(freeHostCmd)
}
