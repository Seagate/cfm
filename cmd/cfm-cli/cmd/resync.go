// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"github.com/spf13/cobra"
)

var resyncCmd = &cobra.Command{
	Use:   "resync <device> [<flags>]",
	Short: "Resynchronize the cfm service to the specified hardware",
	Long: `Resynchronize the cfm service to the specified hardware. (Required after any hardware power cycles)
	This command will NOT retrieve any data.  It'll just invalidate all cfm-service cached data, thus forcing any subsequent cfm-service data requests to interrogate the hardware.`,
	Args: cobra.MatchAll(cobra.NoArgs),
}

// Here you will define your flags and configuration settings.
func init() {
	resyncCmd.DisableFlagsInUseLine = true

	rootCmd.AddCommand(resyncCmd)
}
