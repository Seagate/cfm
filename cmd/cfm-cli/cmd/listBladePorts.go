// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listBladesPortsCmd = &cobra.Command{
	Use:   GetCmdUsageListBladesPorts(),
	Short: "List all available blade ports",
	Long: `Queries the cfm-service for available appliance blade ports.
	Outputs a detailed summary of those ports to stdout.
	Note: For any given ID option:
			If the option is included, ONLY THAT ID is searched.
			If the option is omitted, ALL POSSIBLE IDs (within cfm-service) are searched.`,
	Example: GetCmdExampleListBladesPorts(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListBladePorts(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummary(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listBladesPortsCmd.DisableFlagsInUseLine = true

	listBladesPortsCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of a specific port\n (default \"all ports listed.\")")

	initCommonBladeListCmdFlags(listBladesPortsCmd)

	//Add command to parent
	listBladesCmd.AddCommand(listBladesPortsCmd)
}

// GetCmdUsageListBladesPorts - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageListBladesPorts() string {
	return fmt.Sprintf("%s %s %s %s %s",
		flags.PORTS, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(true),
		flags.GetOptionUsageBladeId(true),
		flags.GetOptionUsagePortId(true))
}

// GetCmdExampleListBladesPorts - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleListBladesPorts() string {
	baseCmd := fmt.Sprintf("cfm list %s %s", flags.BLADES, flags.PORTS)

	baseCmdLoopSh := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp())

	shIdExamplesMap := map[string]string{
		"applianceId": " " + flags.GetOptionExampleShApplianceId(),
		"bladeId":     " " + flags.GetOptionExampleShBladeId(),
		"portId":      " " + flags.GetOptionExampleShPortId(),
	}

	shExampleLines := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		s := baseCmdLoopSh
		if i&1 != 0 {
			s += shIdExamplesMap["applianceId"]
		}
		if i&2 != 0 {
			s += shIdExamplesMap["bladeId"]
		}
		if i&4 != 0 {
			s += shIdExamplesMap["portId"]
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
		"portId":      " " + flags.GetOptionExampleLhPortId(),
	}

	lhExampleLines := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		s := baseCmdLoopLh
		if i&1 != 0 {
			s += lhIdExamplesMap["applianceId"]
		}
		if i&2 != 0 {
			s += lhIdExamplesMap["bladeId"]
		}
		if i&4 != 0 {
			s += lhIdExamplesMap["portId"]
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
