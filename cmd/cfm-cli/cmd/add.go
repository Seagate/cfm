/*
Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

*/

package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <device> [<flags>]",
	Short: "Add an Appliance, Blade or Host to the service",
	Long:  `Add an Appliance, Blade or Host to the service`,
	Args:  cobra.MatchAll(cobra.NoArgs),
}

// Here you will define your flags and configuration settings.
func init() {
	addCmd.DisableFlagsInUseLine = true

	rootCmd.AddCommand(addCmd)
}
