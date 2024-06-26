// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"github.com/spf13/cobra"
)

var unassignCmd = &cobra.Command{
	Use:   "unassign <device> [<flags>]",
	Short: "Unassign an existing blade memory region from a port.",
	Long:  `Unassign an existing blade memory region from a port.`,
	Args:  cobra.MatchAll(cobra.NoArgs),
}

func init() {
	unassignCmd.DisableFlagsInUseLine = true

	//Add command to parent
	rootCmd.AddCommand(unassignCmd)
}
