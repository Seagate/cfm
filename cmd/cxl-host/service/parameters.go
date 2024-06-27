// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cxl_host

import (
	"context"

	"github.com/namsral/flag"
)

const (
	DefaultVerbosity = "0"    // Default log level
	DefaultPort      = "8082" // Default service port
	RedfishVersion   = "/redfish/v1/"
	Version          = "1.1.0"
)

type Settings struct {
	Version   bool   // Print the version of this application and exit if true
	Verbosity string // The log level verbosity, where 0 is no longing and 4 is very verbose
	Port      string // The port that this Redfish service listens on
}

type ctxKey int

const (
	_ ctxKey = iota
	KeyVerbosity
	KeyPort
)

// InitFlags: initialize the configuration data using command line args, ENV, or a file
func (s *Settings) InitContext(args []string, ctx context.Context) (error, context.Context) {

	newContext := ctx

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.String(flag.DefaultConfigFlagname, "", "Path to config file")

	var (
		version   = flags.Bool("version", false, "Display version and exit")
		verbosity = flags.String("verbosity", DefaultVerbosity, "Log level verbosity")
		port      = flags.String("port", DefaultPort, "Service IP Address port")
	)

	// Parse 1) command line arguments, 2) env variables, 3) config file settings, and 4) defaults (in this order)
	err := flags.Parse(args[1:])
	if err != nil {
		return err, newContext
	}

	// Update the configuration object with the parsed values
	s.Version = *version
	s.Verbosity = *verbosity
	s.Port = *port

	newContext = context.WithValue(newContext, KeyVerbosity, s.Verbosity)
	newContext = context.WithValue(newContext, KeyPort, s.Port)

	return nil, newContext
}

// GetContextString: return the value for a flag that is stored in a context
func GetContextString(ctx context.Context, key ctxKey) string {
	value := ""
	if ctx != nil {
		value, _ = ctx.Value(key).(string)
	}
	return value
}

// GetContextInt: return the value for a flag that is stored in a context
func GetContextInt(ctx context.Context, key ctxKey) int {
	value := 0
	if ctx != nil {
		value, _ = ctx.Value(key).(int)
	}
	return value
}
