// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package cxl_host

import (
	openapi "cfm/pkg/redfishapi"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// NewCxlHostRouter creates a new router for any number of api routers
func NewCxlHostRouter(ctx context.Context, routers ...openapi.Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for name, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc
			handler = ApiValidateAndLog(ctx, handler, name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(name).
				Handler(handler)
		}
	}

	return router
}

// OverrideAPIController binds http requests to an api service and writes the service results to the http response
type OverrideAPIController struct {
	service      openapi.DefaultAPIServicer
	errorHandler openapi.ErrorHandler
}

// NewDefaultAPIController creates a override api controller
func NewOverrideAPIController(s openapi.DefaultAPIServicer) openapi.Router {
	controller := &OverrideAPIController{
		service:      s,
		errorHandler: openapi.DefaultErrorHandler,
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *OverrideAPIController) Routes() openapi.Routes {
	return openapi.Routes{
		"RedfishV1SessionServiceSessionsPost": openapi.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/redfish/v1/SessionService/Sessions",
			HandlerFunc: c.RedfishV1SessionServiceSessionsPost,
		},
	}
}

// RedfishV1SessionServiceSessionsPost -
func (c *OverrideAPIController) RedfishV1SessionServiceSessionsPost(w http.ResponseWriter, r *http.Request) {
	sessionV171SessionParam := openapi.SessionV171Session{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&sessionV171SessionParam); err != nil {
		c.errorHandler(w, r, &openapi.ParsingError{Err: err}, nil)
		return
	}
	// if err := AssertSessionV171SessionRequired(sessionV171SessionParam); err != nil {
	// 	c.errorHandler(w, r, err, nil)
	// 	return
	// }
	// if err := AssertSessionV171SessionConstraints(sessionV171SessionParam); err != nil {
	// 	c.errorHandler(w, r, err, nil)
	// 	return
	// }
	result, err := c.service.RedfishV1SessionServiceSessionsPost(r.Context(), sessionV171SessionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

	// Store the token in the HTTP header
	resource, ok := result.Body.(openapi.SessionV171Session)
	if ok {
		w.Header().Set("X-Auth-Token", *resource.Password)
		w.Header().Set("Location", resource.OdataId)
	}

	// If no error, encode the body and the result code
	openapi.EncodeJSONResponse(result.Body, &result.Code, w)
}
