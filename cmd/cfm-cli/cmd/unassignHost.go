// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"github.com/spf13/cobra"
)

var unassignHostCmd = &cobra.Command{
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
	unassignHostCmd.DisableFlagsInUseLine = true

	//TODO: Add flags

	initCommonPersistentFlags(unassignHostCmd)

	//Add command to parent
	unassignCmd.AddCommand(unassignHostCmd)
}
