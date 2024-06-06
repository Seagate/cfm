/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

Package main represents a command line utility for interacting with a cfm-service daemon to control a Seagate Memory Appliance.
*/

package main

import (
	"cfm/cmd/cfm-cli/cmd"

	"github.com/spf13/viper"
	"k8s.io/klog/v2"
)

func init() {
	//Setup Viper
	var V = viper.New()
	V.SetEnvPrefix("CFM")
}

func main() {
	// klog output is buffered and written periodically.  This flush forces the output to flush to the screen on exit.
	defer klog.Flush()

	// This calls the Cobra commands.
	// https://github.com/spf13/cobra-cli/blob/main/README.md
	cmd.Execute()
}
