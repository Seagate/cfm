// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename <device> [<flags>]",
	Short: "Rename a cfm-service device with a new ID",
	Long:  `Rename a cfm-service device with a new ID.`,
	Args:  cobra.MatchAll(cobra.NoArgs),
}

// Here you will define your flags and configuration settings.
func init() {
	renameCmd.DisableFlagsInUseLine = true

	rootCmd.AddCommand(renameCmd)
}
