// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"cfm/cli/pkg/serviceLib/serviceRequests"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listBladeResourcesCmd = &cobra.Command{
	Use:   GetCmdUsageListBladesResources(),
	Short: "List all available blade memory resources",
	Long: `Queries the cfm-service for existing memory resources.
	Outputs a detailed summary (including composition state) of those resources to stdout.
	Note: For any given ID option:
			If the option is included, ONLY THAT ID is searched.
			If the option is omitted, ALL POSSIBLE IDs (within cfm-service) are searched.`,
	Example: GetCmdExampleListBladesResources(),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initLogging(cmd)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := serviceRequests.NewServiceRequestListResources(cmd)
		results, err := request.Execute()
		if err != nil {
			cobra.CheckErr(err)
		}

		request.OutputResults(results)
	},
}

// Here you will define your flags and configuration settings.
func init() {
	listBladeResourcesCmd.DisableFlagsInUseLine = true

	listBladeResourcesCmd.Flags().StringP(flags.RESOURCE_ID, flags.RESOURCE_ID_SH, flags.ID_DFLT, "ID of a specific resource block\n (default \"all resource blocks listed\")")

	initCommonBladeListCmdFlags(listBladeResourcesCmd)

	//Add command to parent
	listBladesCmd.AddCommand(listBladeResourcesCmd)
}

// GetCmdUsageListBladeResources - Generates the command usage string for the cobra.Command.Use field.
func GetCmdUsageListBladesResources() string {
	return fmt.Sprintf("%s %s %s %s %s",
		flags.RESOURCES, // Note: The first word in the Command.Use string is how Cobra defines the "name" of this "command".
		flags.GetOptionUsageGroupServiceTcp(false),
		flags.GetOptionUsageApplianceId(true),
		flags.GetOptionUsageBladeId(true),
		flags.GetOptionUsageResourceId(true))
}

// GetCmdExampleListBladesResources - Generates the command example string for the cobra.Command.Example field.
func GetCmdExampleListBladesResources() string {
	baseCmd := fmt.Sprintf("cfm list %s %s", flags.BLADES, flags.RESOURCES)

	baseCmdLoopSh := fmt.Sprintf("%s %s",
		baseCmd,
		flags.GetOptionExampleShGroupServiceTcp())

	shIdExamplesMap := map[string]string{
		"applianceId": " " + flags.GetOptionExampleShApplianceId(),
		"bladeId":     " " + flags.GetOptionExampleShBladeId(),
		"resourceId":  " " + flags.GetOptionExampleShResourceId(),
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
			s += shIdExamplesMap["resourceId"]
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
		"resourceId":  " " + flags.GetOptionExampleLhResourceId(),
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
			s += lhIdExamplesMap["resourceId"]
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
