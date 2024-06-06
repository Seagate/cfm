/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var composeHostCmd = &cobra.Command{
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
	composeHostCmd.DisableFlagsInUseLine = true

	//TODO: Add flags

	initCommonPersistentFlags(composeHostCmd)

	//Add command to parent
	composeCmd.AddCommand(composeHostCmd)
}
