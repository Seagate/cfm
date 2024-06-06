package services

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"k8s.io/klog/v2"
)

const (
	PRIMARY_WEBUI_DIST_PATH   = "./webui/dist"
	SECONDARY_WEBUI_DIST_PATH = "../../webui/dist"
)

// StartWebService: Launch the Vue.js web service using local distribution files, if they are present.
func StartWebUIService(ctx context.Context, socket *string, webuiDistPath *string) {
	logger := klog.FromContext(ctx)

	// 9 comes from minimum possible length required for socket to have a valid IPv4 and port specified (X.X.X.X:X)
	// Assumes ip defaults to an empty string (which is then used to disable the webui service here)
	if len(*socket) < 9 {
		logger.V(1).Info("[WEBUI] cfm-service's webui service NOT starred.  Incomplete socket specified.", "socket", socket)
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
	logger.V(1).Info("[WEBUI] Start cfm-service's webui service", "socket", socket)

	// Start the server
	err := http.ListenAndServe(*socket, nil)
	if err != nil {
		logger.Error(err, ", [WEBUI] unable to start cfm-service's webui service", "socket", socket)
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
