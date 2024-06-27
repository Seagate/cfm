// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cxl_host

import (
	"cfm/pkg/accounts"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/slices"
	"k8s.io/klog/v2"
)

// enableCORS: Allow browsers to access localhost
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// These paths do NOT require authentication
var noGetAuthentication []string = []string{"/redfish", "/redfish/v1", "/redfish/v1/odata", "/redfish/v1/$metadata", "/redfish/v1/openapi.yaml", "/redfish/v1/Schemas/cxl-host-schema"}
var noPostAuthentication []string = []string{"/redfish/v1/SessionService/Sessions"}

// isValidationRequired: Return true if the path requires authentication
func isValidationRequired(method string, path string) bool {
	validate := true

	if method == "GET" {
		if slices.Contains(noGetAuthentication, path) {
			validate = false
		}
	}

	if method == "POST" {
		if slices.Contains(noPostAuthentication, path) {
			validate = false
		}
	}

	return validate
}

// ApiValidateAndLog: Validate the path before executing, and log an elapsed time
func ApiValidateAndLog(ctx context.Context, inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		enableCORS(&w)
		doRequest := true

		// If this path requires validation, validate before proceeding
		if isValidationRequired(r.Method, r.RequestURI) {
			token := r.Header.Get("X-Auth-Token")
			if !accounts.IsSessionTokenActive(ctx, token) {
				doRequest = false

				// Return a Redfish Error
				redfishError := accounts.RedfishError{
					Error: accounts.RedfishErrorError{
						Code:    "Base.1.14.AccessDenied",
						Message: fmt.Sprintf("Session token (%s) is not valid or is not active", token),
					},
				}
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(redfishError)
			}
		}

		if doRequest {
			inner.ServeHTTP(w, r)
		}

		logger := klog.FromContext(r.Context())
		logger.V(1).Info(
			"logger",
			"method", r.Method,
			"uri", r.RequestURI,
			"name", name,
			"duration", time.Since(start))
	})
}
