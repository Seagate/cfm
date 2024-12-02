// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listHostsPortsCmd = &cobra.Command{
	Use:   GetCmdUsageListHostsPorts(),
	Short: "List all available port(s) accessible to the host(s).",
	Long: `Queries the cfm-service for port(s) accessible to the host(s).
	Outputs a detailed summary of those ports to stdout.
	Note: For any given ID option:
			If the option is included, ONLY THAT ID is searched.
			If the option is omitted, ALL POSSIBLE IDs (within cfm-service) are searched.`,
	Example: GetCmdExampleListHostsPorts(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListHostPorts(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummary(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listHostsPortsCmd.DisableFlagsInUseLine = true

	initCommonHostListCmdFlags(listHostsPortsCmd)

	listHostsPortsCmd.Flags().StringP(flags.PORT_ID, flags.PORT_ID_SH, flags.ID_DFLT, "ID of a specific port\n (default \"all ports listed.\")")

	//Add command to parent
	listHostsCmd.AddCommand(listHostsPortsCmd)
}

// GetCmdUsageListHostsPorts - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageListHostsPorts() string {
	return fmt.Sprintf("%s %s %s %s",
		flags.PORTS, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageHostId(true),
		flags.GetOptionUsagePortId(true))
}

// GetCmdExampleListHostsPorts - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleListHostsPorts() string {
	baseCmd := fmt.Sprintf("cfm list %s %s", flags.HOSTS, flags.PORTS)

	baseCmdLoopSh := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp())

	shIdExamplesMap := map[string]string{
		"hostId": " " + flags.GetOptionExampleShHostId(),
		"portId": " " + flags.GetOptionExampleShPortId(),
	}

	shExampleLines := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		s := baseCmdLoopSh
		if i&1 != 0 {
			s += shIdExamplesMap["hostId"]
		}
		if i&2 != 0 {
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
		"hostId": " " + flags.GetOptionExampleLhHostId(),
		"portId": " " + flags.GetOptionExampleLhPortId(),
	}

	lhExampleLines := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		s := baseCmdLoopLh
		if i&1 != 0 {
			s += lhIdExamplesMap["hostId"]
		}
		if i&2 != 0 {
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
