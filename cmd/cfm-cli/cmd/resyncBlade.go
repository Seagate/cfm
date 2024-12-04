// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"

	"github.com/spf13/cobra"
)

var resyncBladeCmd = &cobra.Command{
	Use:     GetCmdUsageResyncBlade(),
	Short:   "Resynchronize the cfm service to a single blade on a specific composable memory appliance (CMA)",
	Long:    `Resynchronize the cfm service to a single blade on a specific composable memory appliance (CMA).`,
	Example: GetCmdExampleResyncBlade(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestResyncBlade(cmd)
		blade, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		serviceRequests.OutputResultsBladeAction(blade, "resync")
	},
}

// Here you will define your flags and configuration settings.
func init() {
	resyncBladeCmd.DisableFlagsInUseLine = true

	initCommonPersistentFlags(resyncBladeCmd)

	resyncBladeCmd.Flags().StringP(flags.APPLIANCE_ID, flags.APPLIANCE_ID_SH, flags.ID_DFLT, "ID of blade's appliance\n")
	resyncBladeCmd.MarkFlagRequired(flags.APPLIANCE_ID)
	resyncBladeCmd.Flags().StringP(flags.BLADE_ID, flags.BLADE_ID_SH, flags.ID_DFLT, "ID of blade\n")
	resyncBladeCmd.MarkFlagRequired(flags.BLADE_ID)

	//Add command to parent
	resyncCmd.AddCommand(resyncBladeCmd)
}

// GetCmdUsageResyncBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageResyncBlade() string {
	return fmt.Sprintf("%s %s %s %s",
		flags.BLADE, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(false),
		flags.GetOptionUsageBladeId(false))
}

// GetCmdExampleResyncBlade - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleResyncBlade() string {
	baseCmd := fmt.Sprintf("cfm resync %s", flags.BLADE)

	shorthandFormat := fmt.Sprintf("%s %s %s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp(),
		flags.GetOptionExampleShApplianceId(),
		flags.GetOptionExampleShBladeId())

	longhandFormat := fmt.Sprintf("%s %s %s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp(),
		flags.GetOptionExampleLhApplianceId(),
		flags.GetOptionExampleLhBladeId())

	return fmt.Sprintf(`
	%s

	%s`, shorthandFormat, longhandFormat)
}
