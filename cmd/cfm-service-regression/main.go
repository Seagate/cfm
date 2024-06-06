// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cfm/pkg/regression"

	"github.com/onsi/ginkgo/v2"
	"k8s.io/klog/v2"
)

const (
	VERSION                  = "v0.0.2"
	DefaultConfigurationFile = "cfm-service-regression-config.yaml"
)

func stringVar(p *string, name string, usage string) {
	flag.StringVar(p, name, *p, usage)
}

func intVar(p *int, name string, usage string) {
	flag.IntVar(p, name, *p, usage)
}

type testing struct {
	result int
}

func (t *testing) Fail() {
	t.result = 1
}

func main() {

	// Context
	ctx := context.Background()

	// Flags for version and debug level
	version := flag.Bool("version", false, "print the version of this program")
	debug := flag.String("debug", "", "set the klog debug level")
	config := flag.String("config", DefaultConfigurationFile, "use this yaml file for configuration settings")

	// Parse flags
	flag.Parse()

	if *version {
		fmt.Printf("cfm-service-regression (%s)\n", VERSION)
		os.Exit(0)
	}

	if *debug != "" {
		fmt.Printf("setting debug level to (%s)\n", *debug)
		level := klog.Level(0)
		level.Set(*debug)
	}

	// Use the default configuration file, or a command line option filename
	configFilename := DefaultConfigurationFile
	if *config != "" {
		configFilename = *config
	}

	// Read config info from yaml file and set up configuration instance
	configFile, err := os.ReadFile(configFilename)
	if err != nil {
		fmt.Printf("ERROR: reading config file (%s), err: %v\n", configFilename, err)
		os.Exit(1)
	}
	cfmConfig, err := regression.NewTestConfig(ctx, configFile)
	if err != nil {
		fmt.Printf("ERROR: invalid config file (%s), err: %v\n", configFilename, err)
		os.Exit(2)
	}

	// Logging
	klog.EnableContextualLogging(true)
	klog.SetOutput(ginkgo.GinkgoWriter)
	logger := klog.FromContext(ctx)

	logger.V(4).Info("configuration", "filename", configFilename, "config", cfmConfig)

	// Create testing object
	t := testing{}

	// Run regression tests
	regression.Test(&t, cfmConfig, logger)

	os.Exit(t.result)
}
