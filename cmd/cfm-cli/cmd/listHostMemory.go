// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listHostsMemoryCmd = &cobra.Command{
	Use:   GetCmdUsageListHostsMemory(),
	Short: "List all available composed memory region(s) accessible to the host(s)",
	Long: `Queries the cfm-service for composed memory region(s) accessible to the host(s).
	Outputs a detailed summary of those memory regions to stdout.
	Note: For any given ID option:
			If the option is included, ONLY THAT ID is searched.
			If the option is omitted, ALL POSSIBLE IDs (within cfm-service) are searched.`,
	Example: GetCmdExampleListHostsMemory(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListHostMemoryRegions(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummaryListMemory(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listHostsMemoryCmd.DisableFlagsInUseLine = true

	initCommonHostListCmdFlags(listHostsMemoryCmd)

	listHostsMemoryCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of a specific memory region\n (default \"all memory regions listed\")")

	//Add command to parent
	listHostsCmd.AddCommand(listHostsMemoryCmd)
}

// GetCmdUsageListHostMemory - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageListHostsMemory() string {
	return fmt.Sprintf("%s %s %s %s",
		flags.MEMORY, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageHostId(true),
		flags.GetOptionUsageMemoryId(true))
}

// GetCmdExampleListHostsMemory - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleListHostsMemory() string {
	baseCmd := fmt.Sprintf("cfm list %s %s", flags.HOSTS, flags.MEMORY)

	baseCmdLoopSh := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp())

	shIdExamplesMap := map[string]string{
		"hostId":   " " + flags.GetOptionExampleShHostId(),
		"memoryId": " " + flags.GetOptionExampleShMemoryId(),
	}

	shExampleLines := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		s := baseCmdLoopSh
		if i&1 != 0 {
			s += shIdExamplesMap["hostId"]
		}
		if i&2 != 0 {
			s += shIdExamplesMap["memoryId"]
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
		"hostId":   " " + flags.GetOptionExampleShHostId(),
		"memoryId": " " + flags.GetOptionExampleLhMemoryId(),
	}

	lhExampleLines := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		s := baseCmdLoopLh
		if i&1 != 0 {
			s += shIdExamplesMap["hostId"]
		}
		if i&2 != 0 {
			s += lhIdExamplesMap["memoryId"]
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
