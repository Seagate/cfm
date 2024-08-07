// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package services

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"k8s.io/klog/v2"
)

const (
	PRIMARY_WEBUI_DIST_PATH   = "./webui/dist"
	SECONDARY_WEBUI_DIST_PATH = "../../webui/dist"
)

// StartWebService: Launch the Vue.js web service using local distribution files, if they are present.
func StartWebUIService(ctx context.Context, webuiPort *string, servicePort *string, webuiDistPath *string, hostIpOverride *string) {
	logger := klog.FromContext(ctx)

	var hostIp string

	// obtain host IP
	if *hostIpOverride == "" {
		hostIp = GetHostIp(ctx)
	} else {
		hostIp = *hostIpOverride
	}

	webuiSocket := fmt.Sprintf("%s:%s", hostIp, *webuiPort)
	serviceSocket := fmt.Sprintf("%s:%s", hostIp, *servicePort)

	// 9 comes from minimum possible length required for socket to have a valid IPv4 and port specified (X.X.X.X:X)
	// Assumes ip defaults to an empty string (which is then used to disable the webui service here)
	if len(webuiSocket) < 9 {
		logger.V(1).Info("[WEBUI] cfm-service's webui service NOT starred.  Incomplete socket specified.", "socket", webuiSocket)
		return
	}

	// overwrite base_path for the webui disto
	err := UpdateBasePath(ctx, &serviceSocket, webuiDistPath)
	if err != nil {
		logger.V(1).Info("[WEBUI] fail to update webui service base path", "socket", serviceSocket, "err", err)
		return
	}

	r := mux.NewRouter()

	// Serve up distribution of cfm-webui
	// Example: "./services/webui/dist"
	fs := http.FileServer(http.Dir(*webuiDistPath))
	r.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fullPath := filepath.Join(*webuiDistPath, path)

		// Check if file exists
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// If the file doesn't exist, serve index.html
			http.ServeFile(w, r, filepath.Join(*webuiDistPath, "index.html"))
			return
		}

		// Otherwise, serve the requested file
		fs.ServeHTTP(w, r)
	}))
	http.Handle("/", r)

	// Log and start the web UI service
	logger.V(1).Info("[WEBUI] Start cfm-service's webui service", "socket", webuiSocket)

	// Start the server
	err = http.ListenAndServe(webuiSocket, nil)
	if err != nil {
		logger.Error(err, ", [WEBUI] unable to start cfm-service's webui service", "socket", webuiSocket)
	}
}

// FindWebUIDistPath - Find the relative path to the webui distro.  Returns nil if not found.
//
// Locates the "services" pkg relative to the current working directory
// This is necessary because of the way the vscode debugger runs and how the project is structured.
// If using "make local" and then start the service using "./cfm-service", os.Getwd() == cfm-service/ (the project root folder)
// If start the service using vscode debugger, os.Getwd() == cfm-service/cmd/cfm-service (starting at the project root folder)
// So the "services" package is not in a consistent relative location to the current working directory.  Need to determine which one.
func FindWebUIDistPath(ctx context.Context) (*string, error) {
	logger := klog.FromContext(ctx)

	var webuiDistPath string

	_, err := os.Open(PRIMARY_WEBUI_DIST_PATH)
	if !errors.Is(err, os.ErrNotExist) {
		webuiDistPath = PRIMARY_WEBUI_DIST_PATH
		logger.V(4).Info("[WEBUI] found webui service", "webuiDistPath", webuiDistPath)
		return &webuiDistPath, nil
	}

	// This one is only used in special debug cases (where a webui distro has been manually added to the local cfm-service project AND the project is being run via vscode debugger)
	_, err = os.Open(SECONDARY_WEBUI_DIST_PATH)
	if !errors.Is(err, os.ErrNotExist) {
		webuiDistPath = SECONDARY_WEBUI_DIST_PATH
		logger.V(4).Info("[WEBUI] found webui service", "webuiDistPath", webuiDistPath)
		return &webuiDistPath, nil
	}

	return nil, fmt.Errorf("webui service NOT found within cfm-service")
}

// GetHostIp: get host ip address from hostname
func GetHostIp(ctx context.Context) string {
	logger := klog.FromContext(ctx)

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.Error(err, ", [WEBUI] unable to retrive cfm-service's ip address")
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	logger.V(2).Info("[WEBUI] found webui service", "ip addr", localAddr.IP)

	return fmt.Sprintf("%s", localAddr.IP[:]) // deep copy
}

// UpdateBasePath: Replace the base address in the webui distro file
func UpdateBasePath(ctx context.Context, serviceSocket *string, webuiDistPath *string) error {
	// logger := klog.FromContext(ctx)

	rawString := "http://127.0.0.1:8080"
	dirPath := *webuiDistPath + "/assets/"

	f, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return err
	}

	for _, file := range fileInfo {
		// only match suffix ".js"
		if file.Name()[len(file.Name())-3:] == ".js" {
			read, err := os.ReadFile(dirPath + file.Name())
			if err != nil {
				return err
			}
			newContents := strings.Replace(string(read), rawString, "http://"+*serviceSocket, -1)
			err = os.WriteFile(dirPath+file.Name(), []byte(newContents), 0)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
