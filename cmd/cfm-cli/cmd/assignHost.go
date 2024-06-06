/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var assignHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Not Implemented.",
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
	assignHostCmd.DisableFlagsInUseLine = true

	//TODO: Add flags

	initCommonPersistentFlags(assignHostCmd)

	//Add command to parent
	assignCmd.AddCommand(assignHostCmd)
}
