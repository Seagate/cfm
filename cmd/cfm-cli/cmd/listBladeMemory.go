// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listBladesMemoryCmd = &cobra.Command{
	Use:   GetCmdUsageListBladesMemory(),
	Short: "List all available blade composed and\\or provisioned memory regions",
	Long: `Queries the cfm-service for available composed\\provisioned memory regions.
	Outputs a detailed summary of those memory regions to stdout.
	Note: For any given ID option:
			If the option is included, ONLY THAT ID is searched.
			If the option is omitted, ALL POSSIBLE IDs (within cfm-service) are searched.`,
	Example: GetCmdExampleListBladesMemory(),
	Args:    cobra.MatchAll(cobra.NoArgs),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListMemoryRegions(cmd)
		summary, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputSummaryListMemory(summary)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listBladesMemoryCmd.DisableFlagsInUseLine = true

	listBladesMemoryCmd.Flags().StringP(flags.MEMORY_ID, flags.MEMORY_ID_SH, flags.ID_DFLT, "ID of a specific memory region\n (default \"all memory regions listed\")")

	initCommonBladeListCmdFlags(listBladesMemoryCmd)

	//Add command to parent
	listBladesCmd.AddCommand(listBladesMemoryCmd)
}

// GetCmdUsageListBladeMemory - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageListBladesMemory() string {
	return fmt.Sprintf("%s %s %s %s %s",
		flags.MEMORY, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(true),
		flags.GetOptionUsageBladeId(true),
		flags.GetOptionUsageMemoryId(true))
}

// GetCmdExampleListBladesMemory - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleListBladesMemory() string {
	baseCmd := fmt.Sprintf("cfm list %s %s", flags.BLADES, flags.MEMORY)

	baseCmdLoopSh := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp())

	shIdExamplesMap := map[string]string{
		"applianceId": " " + flags.GetOptionExampleShApplianceId(),
		"bladeId":     " " + flags.GetOptionExampleShBladeId(),
		"memoryId":    " " + flags.GetOptionExampleShMemoryId(),
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
		"applianceId": " " + flags.GetOptionExampleLhApplianceId(),
		"bladeId":     " " + flags.GetOptionExampleLhBladeId(),
		"memoryId":    " " + flags.GetOptionExampleLhMemoryId(),
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
