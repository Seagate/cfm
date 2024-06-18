// Copyright (c) 2022 Seagate Technology LLC and/or its Affiliates

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cfm/pkg/accounts"
	"cfm/pkg/api"
	"cfm/pkg/common"
	"cfm/pkg/openapi"

	"k8s.io/klog/v2"
)

// This variable is filled in during the linker step - -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`"
var buildTime = ""

func main() {

	// Extract settings and initialize context using command line args, env, config file, or defaults
	settings := common.Settings{}
	ctx := context.Background()
	var err error
	err, ctx = settings.InitContext(os.Args, ctx)

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

	// cxl-host banner
	logger.Info("[] cxl-host", "version", common.Version, "build", buildTime)
	args := strings.Join(os.Args[1:], " ")
	logger.V(1).Info("cxl-host", "args", args)
	logger.V(2).Info("cxl-host", "settings", settings)

	if settings.Version {
		os.Exit(0)
	}

	accounts.AccountsHandler().InitLogger(logger)
	accounts.AccountsHandler().Restore()

	DefaultApiService := api.NewCxlHostApiService()
	DefaultApiController := openapi.NewDefaultAPIController(DefaultApiService)
	OverrideAPIController := api.NewOverrideAPIController(DefaultApiService)
	router := api.NewCxlHostRouter(ctx, OverrideAPIController, DefaultApiController)
	log.Fatal(http.ListenAndServe(":"+settings.Port, router))
}
