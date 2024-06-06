/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var composeCmd = &cobra.Command{
	Use:   "compose <device> [<flags>]",
	Short: "Creates (composes) a new memory region on a memory appliance blade.",
	Long:  `Creates (composes) a new memory region on a memory appliance blade.  The memory composition can occur via a memory appliance blade OR a cxl-host.`,
	Args:  cobra.MatchAll(cobra.NoArgs),
}

func init() {
	composeCmd.DisableFlagsInUseLine = true

	//Add command to parent
	rootCmd.AddCommand(composeCmd)
}
