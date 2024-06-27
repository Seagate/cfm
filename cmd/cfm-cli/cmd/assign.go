// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"github.com/spf13/cobra"
)

var assignCmd = &cobra.Command{
	Use:   "assign <device> [<flags>]",
	Short: "Assign an existing blade memory region to a port.",
	Long:  `Assign an existing blade memory region to a port.`,
	Args:  cobra.MatchAll(cobra.NoArgs),
}

func init() {
	assignCmd.DisableFlagsInUseLine = true

	//Add command to parent
	rootCmd.AddCommand(assignCmd)
}
