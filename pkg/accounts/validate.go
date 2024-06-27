// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package accounts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/slices"
	"k8s.io/klog/v2"
)

type ResourceHealth string

// List of ResourceHealth
const (
	RESOURCEHEALTH_OK       ResourceHealth = "OK"
	RESOURCEHEALTH_WARNING  ResourceHealth = "Warning"
	RESOURCEHEALTH_CRITICAL ResourceHealth = "Critical"
)

// MessageV112Message - The message that the Redfish service returns.
type MessageV112Message struct {

	// The human-readable message.
	Message string `json:"Message,omitempty"`

	// An array of message arguments that are substituted for the arguments in the message when looked up in the message registry.
	MessageArgs []string `json:"MessageArgs,omitempty"`

	// The identifier for the message.
	MessageId string `json:"MessageId"`

	MessageSeverity ResourceHealth `json:"MessageSeverity,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// A set of properties described by the message.
	RelatedProperties []string `json:"RelatedProperties,omitempty"`

	// Used to provide suggestions on how to resolve the situation that caused the message.
	Resolution string `json:"Resolution,omitempty"`

	// The severity of the message.
	// Deprecated
	Severity string `json:"Severity,omitempty"`
}

// RedfishError - The error payload from a Redfish service.
type RedfishError struct {
	Error RedfishErrorError `json:"error"`
}

// RedfishErrorError - The properties that describe an error from a Redfish service.
type RedfishErrorError struct {

	// An array of messages describing one or more error messages.
	MessageExtendedInfo []MessageV112Message `json:"@Message.ExtendedInfo,omitempty"`

	// A string indicating a specific MessageId from a message registry.
	Code string `json:"code"`

	// A human-readable error message corresponding to the message in a message registry.
	Message string `json:"message"`
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

// ValidateAndLog: Validate the path before executing, and log an elapsed time
func ValidateAndLog(ctx context.Context, inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		doRequest := true

		// If this path requires validation, validate before proceeding
		if isValidationRequired(r.Method, r.RequestURI) {
			token := r.Header.Get("X-Auth-Token")
			if !IsSessionTokenActive(ctx, token) {
				doRequest = false

				// Return a Redfish Error
				redfishError := RedfishError{
					Error: RedfishErrorError{
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
		logger.V(1).Info("logger", "method", r.Method, "uri", r.RequestURI, "name", name, "duration", time.Since(start))
	})
}
