// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package main

import (
	cxl_host "cfm/cmd/cxl-host/service"
	"cfm/pkg/accounts"
	"cfm/pkg/redfishapi"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/godbus/dbus/v5"
	"github.com/holoplot/go-avahi"
	"k8s.io/klog/v2"
)

// This variable is filled in during the linker step - -ldflags "-X main.buildTime=`date -u '+%Y-%m-%dT%H:%M:%S'`"
var buildTime = ""

func main() {

	// Extract settings and initialize context using command line args, env, config file, or defaults
	settings := cxl_host.Settings{}
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
	logger.Info("[] cxl-host", "version", cxl_host.Version, "build", buildTime)
	args := strings.Join(os.Args[1:], " ")
	logger.V(1).Info("cxl-host", "args", args)
	logger.V(2).Info("cxl-host", "settings", settings)

	if settings.Version {
		os.Exit(0)
	}

	accounts.AccountsHandler().InitLogger(logger)
	accounts.AccountsHandler().Restore()

	// avahi publish
	err = AvahiPublish(ctx, settings.Port)
	if err != nil {
		fmt.Printf("ERROR: avahi publish failed, err=%v\nContinue...\n", err)
	}
	DefaultApiService := cxl_host.NewCxlHostApiService()
	DefaultApiController := redfishapi.NewDefaultAPIController(DefaultApiService)
	OverrideAPIController := cxl_host.NewOverrideAPIController(DefaultApiService)
	router := cxl_host.NewCxlHostRouter(ctx, OverrideAPIController, DefaultApiController)
	log.Fatal(http.ListenAndServe(":"+settings.Port, router))
}

// AvahiPublish: Publish the service with avahi
func AvahiPublish(ctx context.Context, port string) error {
	txt := [][]byte{}
	txt = append(txt, []byte("cxl-host=true"))

	conn, err := dbus.SystemBus()
	if err != nil {
		return fmt.Errorf("avahi: Cannot get system bus: %v", err)
	}

	a, err := avahi.ServerNew(conn)
	if err != nil {
		return fmt.Errorf("avahi: Avahi new failed: %v", err)
	}

	eg, err := a.EntryGroupNew()
	if err != nil {
		return fmt.Errorf("avahi: EntryGroupNew() failed: %v", err)
	}

	hostname, err := a.GetHostName()
	if err != nil {
		return fmt.Errorf("avahi: GetHostName() failed: %v", err)
	}

	fqdn, err := a.GetHostNameFqdn()
	if err != nil {
		return fmt.Errorf("avahi: GetHostNameFqdn() failed: %v", err)
	}

	p, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("avahi: Cannot convert port to Integer: %v", err)
	}
	err = eg.AddService(avahi.InterfaceUnspec, avahi.ProtoInet, 0, hostname, "_obmc_redfish._tcp", "local", fqdn, uint16(p), txt)
	if err != nil {
		return fmt.Errorf("avahi: AddService() failed: %v", err)
	}

	err = eg.Commit()
	if err != nil {
		return fmt.Errorf("avahi: Commit() failed: %v", err)
	}

	return nil
}
