// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cmd

import (
	"cfm/cli/pkg/serviceLib/flags"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "cfm <command> <device> [<flags>]",
	Version: "1.1.0",
	Short:   "CLI for managing memory appliance(s) via cfm-service.",
	Long:    `A command line interface that connects to a cfm-service instance.  Once connected, can use CLI to to manage one or more memory appliances.`,
	Args:    cobra.MatchAll(cobra.NoArgs),
	// Each child function will run this section of code.
	// Validates environment variables and flags.
	// On error, prints the underlying error.
	// PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
	// 	return nil
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("Welcome to cfm (the CLI for cfm-service)")

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().Int(flags.VERBOSITY, flags.VERBOSITY_DFLT, "set cfm logging verbosity")
}

// Initialize logging functionality via klog
func initLogging(cmd *cobra.Command) {

	// Use contextual logging by pulling the logger from the context, and structured logging using Info or Error
	klog.EnableContextualLogging(true)

	verbosity, err := cmd.Flags().GetInt(flags.VERBOSITY)
	if err != nil {
		klog.ErrorS(err, "Missing verbosity flag. Default setting required.")
		cobra.CheckErr(err)
	}

	// Set global logging level
	var lvl klog.Level
	lvl.Set(strconv.Itoa(verbosity))
	klog.V(1).InfoS("logging verbosity level", "verbosity", verbosity)
}

// Add set of flags used by all subcommands
// Kept separate because of cobra framework quirk
// Need to ONLY add them in the subcommand init() func so these flags ONLY show up in the help output when they can actually be used.
func initCommonPersistentFlags(cmd *cobra.Command) {
	//Add globally required cfm-service TCPIP connection flags
	cmd.PersistentFlags().StringP(flags.SERVICE_NET_IP, flags.SERVICE_NET_IP_SH, flags.SERVICE_NET_IP_DFLT, "cfm-service network IP address")
	cmd.PersistentFlags().Uint16P(flags.SERVICE_NET_PORT, flags.SERVICE_NET_PORT_SH, flags.SERVICE_NET_PORT_DFLT, "cfm-service network port")

	//Currently unused but need default values downstream
	cmd.PersistentFlags().Bool(flags.SERVICE_INSECURE, flags.SERVICE_INSECURE_DFLT, "cfm-service insecure connection flag")
	cmd.PersistentFlags().MarkHidden(flags.SERVICE_INSECURE)
	cmd.PersistentFlags().String(flags.SERVICE_PROTOCOL, flags.SERVICE_PROTOCOL_DFLT, "cfm-service network connection protocol (http/https)")
	cmd.PersistentFlags().MarkHidden(flags.SERVICE_PROTOCOL)
}
