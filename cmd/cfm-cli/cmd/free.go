// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package cmd

import (
	"github.com/spf13/cobra"
)

var freeCmd = &cobra.Command{
	Use:   "free <device> [<flags>]",
	Short: "Free an existing memory region on the specified memory appliance blade.",
	Long:  `Free an existing memory region on the specified memory appliance blade.  The blade port will be unassigned and the memory region's resource blocks will be deallocated.`,
	Args:  cobra.MatchAll(cobra.NoArgs),
}

func init() {
	freeCmd.DisableFlagsInUseLine = true

	//Add command to parent
	rootCmd.AddCommand(freeCmd)
}
