// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listBladesCmd = &cobra.Command{
	Use:     GetCmdUsageListBlades(),
	Short:   "List some or all recognized appliance blades",
	Long:    `Queries the cfm-service for some or all recognized appliance blades and outputs a detailed summary of the discovered blades to stdout.`,
	Example: GetCmdExampleListBlades(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListBlades(cmd)
		bladesSummary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(bladesSummary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listBladesCmd.DisableFlagsInUseLine = true

	initCommonBladeListCmdFlags(listBladesCmd)

	//Add command to parent
	listCmd.AddCommand(listBladesCmd)
}

// GetCmdUsageListBlade - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageListBlades() string {
	return fmt.Sprintf("%s %s %s %s",
		flags.BLADES, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(true),
		flags.GetOptionUsageBladeId(true))
}

// GetCmdExampleListBlades - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleListBlades() string {
	baseCmd := fmt.Sprintf("cfm list %s", flags.BLADES)

	baseCmdLoopSh := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp())

	shIdExamplesMap := map[string]string{
		"applianceId": " " + flags.GetOptionExampleShApplianceId(),
		"bladeId":     " " + flags.GetOptionExampleShBladeId(),
	}

	shExampleLines := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		s := baseCmdLoopSh
		if i&1 != 0 {
			s += shIdExamplesMap["applianceId"]
		}
		if i&2 != 0 {
			s += shIdExamplesMap["bladeId"]
		}
		shExampleLines = append(shExampleLines, s)
	}

	var shorthandFormat strings.Builder
	for _, line := range shExampleLines {
		shorthandFormat.WriteString("\t" + line + "\n")
	}

	baseCmdLoopLh := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleLhGroupServiceTcp())

	lhIdExamplesMap := map[string]string{
		"applianceId": " " + flags.GetOptionExampleLhApplianceId(),
		"bladeId":     " " + flags.GetOptionExampleLhBladeId(),
	}

	lhExampleLines := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		s := baseCmdLoopLh
		if i&1 != 0 {
			s += lhIdExamplesMap["applianceId"]
		}
		if i&2 != 0 {
			s += lhIdExamplesMap["bladeId"]
		}
		lhExampleLines = append(lhExampleLines, s)
	}

	var longhandFormat strings.Builder
	for _, line := range lhExampleLines {
		longhandFormat.WriteString("\t" + line + "\n")
	}

	return fmt.Sprintf(`
%s

%s`, shorthandFormat.String(), longhandFormat.String())
}
