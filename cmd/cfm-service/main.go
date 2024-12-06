// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/rs/cors"
	"k8s.io/klog/v2"

	"cfm/pkg/api"
	"cfm/pkg/common"
	"cfm/pkg/common/datastore"
	"cfm/pkg/openapi"
	"cfm/pkg/redfishapi"
	"cfm/services"
)

var Version = "1.x.x"

// This variable is filled in during the linker step - -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`"
var buildTime = ""

func main() {

	var err error
	var wg sync.WaitGroup

	// Extract settings and initialize context using command line args, env, config file, or defaults
	settings := common.Settings{}
	ctx := context.Background()
	ctx, err = settings.InitContext(os.Args, ctx)
	if err != nil {
		fmt.Printf("ERROR: parsing parameters, err=%v\n", err)
		os.Exit(1)
	}

	// Use contextual logging by pulling the logger from the context, and structured logging using Info or Error
	klog.EnableContextualLogging(true)

	// Set verbosity level according to the 'verbosity' flag
	var l klog.Level
	l.Set(settings.Verbosity)
	logger := klog.FromContext(ctx)

	// cfm-service banner
	logger.Info("[] cfm-service", "version", Version, "build", buildTime)
	args := strings.Join(os.Args[1:], " ")
	logger.V(0).Info("cfm-service", "args", args)
	logger.V(0).Info("cfm-service", "settings", settings)

	if settings.Version {
		os.Exit(0)
	}

	// Create a new router passing in context, which is passed on to every http request
	defaultApiService := api.NewCfmApiService(Version)
	defaultApiController := openapi.NewDefaultAPIController(defaultApiService)
	router := api.NewCfmRouter(ctx, defaultApiController)

	defaultRedfishService := api.NewRedfishApiService(Version)
	defaultRedfishController := redfishapi.NewDefaultAPIController(defaultRedfishService)
	api.AddRedfishRouter(ctx, router, defaultRedfishController)

	// Load datastore
	datastore.DStore().Restore()
	data := datastore.DStore().GetDataStore()

	// If there is any devices in the data store, skip discovery
	if len(data.ApplianceData) == 0 && len(data.HostData) == 0 {
		// Discover devices before loading datastore
		bladeDevices, errBlade := services.DiscoverDevices(ctx, defaultApiService, "blade")
		hostDevices, errHost := services.DiscoverDevices(ctx, defaultApiService, "cxl-host")
		// Add the discovered devices into datastore
		if errBlade == nil && errHost == nil {
			services.AddDiscoveredDevices(ctx, defaultApiService, bladeDevices, hostDevices)
		}
		// Update data
		data = datastore.DStore().GetDataStore()
	}

	datastore.ReloadDataStore(ctx, defaultApiService, data)

	// Set up CORS middleware (for webui)
	c := cors.AllowAll()
	handler := c.Handler(router)

	// Attempt to start cfm-service's webui service on a separate thread
	if settings.Webui {
		webuiDistPath, err := services.FindWebUIDistPath(ctx)
		if err != nil {
			logger.Error(err, ", [WEBUI] unable to locate cfm-service's webui service distro")
		} else {
			wg.Add(1)
			go services.StartWebUIService(ctx, &settings.WebuiPort, &settings.Port, webuiDistPath, &settings.HostIpOverride)
		}
	}

	// Start the main service
	logger.V(0).Info("cfm-service web server", "port", settings.Port)
	log.Fatal(http.ListenAndServe(":"+settings.Port, handler))
}
