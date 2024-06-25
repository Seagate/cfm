// Copyright (c) 2022 Seagate Technology LLC and/or its Affiliates

package common

import (
	"context"
	"fmt"

	"github.com/namsral/flag"
	"k8s.io/klog/v2"
)

const (
	NumUuidCharsForId = 4 // Number of chars to strip from an interally generated uuid (starting from the right) for use in the internally generated ID's for appliance, blade and host
)
const (
	DefaultBackend   = "httpfish" // Default backend interface
	DefaultVerbosity = "0"        // Default log level
	DefaultPort      = "8080"     // Default cfm-service port
	DefaultWebui     = false      // Default mode for cfm-service's webui service.  This DISABLES the webui service.
	DefaultWebuiPort = "3000"     // Default port for cfm-service's webui service
)

var ValidBackends = []string{"httpfish"}

type Settings struct {
	Version   bool   // Print the version of this application and exit if true
	Verbosity string // The log level verbosity, where 0 is no longing and 4 is very verbose
	Backend   string // The backend interface to use, possible values are:  httpfish
	Port      string // The port that this service listens on
	Webui     bool   // The switch where cfm-service serves up its' webui service
	WebuiPort string // The port where cfm-service serves up its' webui service
}

const (
	KeyVerbosity = "cfmCtxVerbosity"
	KeyBackend   = "cfmCtxBackend"
	KeyUri       = "cfmCtxUri"
)

// Store a slice of context keys used in the cloning operation
var contextKeys = []string{
	KeyVerbosity,
	KeyBackend,
	KeyUri,
}

// InitFlags: initialize the configuration data using command line args, ENV, or a file
func (s *Settings) InitContext(args []string, ctx context.Context) (context.Context, error) {

	newContext := ctx

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.String(flag.DefaultConfigFlagname, "", "Path to config file")

	backendNote := fmt.Sprintf("Backend interface choice, options: %v", ValidBackends)

	var (
		version   = flags.Bool("version", false, "Display version and exit")
		verbosity = flags.String("verbosity", DefaultVerbosity, "Log level verbosity")
		backend   = flags.String("backend", DefaultBackend, backendNote)
		port      = flags.String("port", DefaultPort, "CFM service IP Address port")
		webui     = flags.Bool("webui", DefaultWebui, "Enable cfm-service's webui service")
		webuiPort = flags.String("webuiPort", DefaultWebuiPort, "Port for cfm-service's webui service")
	)

	// Parse 1) command line arguments, 2) env variables, 3) config file settings, and 4) defaults (in this order)
	err := flags.Parse(args[1:])
	if err != nil {
		return newContext, err
	}

	// Update the configuration object with the parsed values
	s.Version = *version
	s.Verbosity = *verbosity
	s.Backend = *backend
	s.Port = *port
	s.Webui = *webui
	s.WebuiPort = *webuiPort

	klog.V(4).InfoS("SetContextString", "KeyVerbosity", s.Verbosity, "KeyBackend", s.Backend)
	newContext = context.WithValue(newContext, KeyVerbosity, s.Verbosity)
	newContext = context.WithValue(newContext, KeyBackend, s.Backend)

	return newContext, nil
}

// GetContextString: return the value for a flag that is stored in a context
func GetContextString(ctx context.Context, key string) string {
	value := fmt.Sprintf("%v", ctx.Value(string(key)))
	klog.V(5).InfoS("GetContextString", "ctx", ctx)
	klog.V(5).InfoS("GetContextString", "key", key, "value", value)
	return value
}

// CloneContext: copy all context values from the incoming context to the new context
func CloneContext(mainContext, requestContext context.Context) context.Context {
	newContext := requestContext
	for _, key := range contextKeys {
		newContext = context.WithValue(newContext, key, GetContextString(mainContext, key))
	}
	return newContext
}
