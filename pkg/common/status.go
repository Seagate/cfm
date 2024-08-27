// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package common

import (
	"fmt"
	"net/http"
)

// Our new enhanced error
type RequestError struct {
	StatusCode StatusCodeType
	Err        error
}

// Our string representation of our enhanced error
func (r *RequestError) Error() string {
	return fmt.Sprintf("status [%d][%s] %v", int(r.StatusCode), r.StatusCode.String(), r.Err)
}

// A common error type to use throughout this application.
type StatusCodeType int

// A list of unique error codes returned by this service. These error codes
// can be used programmatically by clients to differentiate various errors returned.
// More error code will be added
const (
	StatusOK StatusCodeType = iota

	StatusComposePartialSuccess         //206
	StatusApplianceResyncPartialSuccess //206

	StatusAppliancesExceedMaximum //422
	StatusBladesExceedMaximum     //422
	StatusHostsExceedMaximum      //422

	StatusBackendInterfaceFailure //500: StatusInternalServerError

	StatusComposeMemoryByResourceFailure //409

	StatusAssignMemoryFailure   //409
	StatusUnassignMemoryFailure //409

	StatusHostGetMemoryFailure      //409
	StatusBladeGetMemoryFailure     //409
	StatusBladeGetMemoryByIdFailure //409
	StatusHostGetMemoryByIdFailure  //409
	StatusBladeFreeMemoryFailure    //409
	StatusHostFreeMemoryFailure     //409

	StatusBladeGetMemoryResourceBlocksFailure       //409
	StatusBladeGetMemoryResourceBlockDetailsFailure //409
	StatusBladeGetPortsFailure                      //409
	StatusHostGetPortsFailure                       //409
	StatusGetPortDetailsFailure                     //409
	StatusGetMemoryDevicesFailure                   //409
	StatusGetMemoryDevicesDetailsFailure            //409

	StatusApplianceResyncFailure //409

	StatusApplianceIdDuplicate //409
	StatusBladeIdDuplicate     //409
	StatusMemoryIdDuplicate    //409
	StatusPortIdDuplicate      //409
	StatusHostIdDuplicate      //409

	StatusApplianceCreateSessionFailure //500
	StatusApplianceDeleteSessionFailure //500
	StatusBladeCreateSessionFailure     //500
	StatusBladeDeleteSessionFailure     //500
	StatusHostCreateSessionFailure      //500
	StatusHostDeleteSessionFailure      //500

	StatusApplianceIdDoesNotExist    //404
	StatusBladeIdDoesNotExist        //404
	StatusHostIdDoesNotExist         //404
	StatusMemoryIdDoesNotExist       //404
	StatusPortIdDoesNotExist         //404
	StatusResourceIdDoesNotExist     //404
	StatusMemoryDeviceIdDoesNotExist //404

	StatusManagerInitializationFailure //500
)

// Return a string representation of our StatusType. When this value is printed as a
// string (%s), the value is determined by this function.
func (e StatusCodeType) String() string {
	switch e {
	case StatusOK:
		return "Success"
	case StatusComposePartialSuccess:
		return "Compose Partial Success"
	case StatusApplianceResyncPartialSuccess:
		return "Appliance Resync Partial Success"
	case StatusComposeMemoryByResourceFailure:
		return "Compose Memory By Resource Failed"
	case StatusAssignMemoryFailure:
		return "Assign Memory Failed"
	case StatusUnassignMemoryFailure:
		return "Unassign Memory Failed"
	case StatusBackendInterfaceFailure:
		return "Backend Interface does not exist"
	case StatusApplianceCreateSessionFailure:
		return "Appliance Create Session Failure"
	case StatusApplianceDeleteSessionFailure:
		return "Appliance Delete Session Failure"
	case StatusBladeCreateSessionFailure:
		return "Blade Create Session Failure"
	case StatusBladeDeleteSessionFailure:
		return "Blade Delete Session Failure"
	case StatusApplianceIdDoesNotExist:
		return "Appliance Id Does Not Exist"
	case StatusHostCreateSessionFailure:
		return "Host Create Session Failure"
	case StatusHostDeleteSessionFailure:
		return "Host Delete Session Failure"
	case StatusHostIdDoesNotExist:
		return "Host Id Does Not Exist"
	case StatusMemoryIdDoesNotExist:
		return "Memory Id Does Not Exist"
	case StatusPortIdDoesNotExist:
		return "Port Id Does Not Exist"
	case StatusResourceIdDoesNotExist:
		return "Resource Id Does Not Exist"
	case StatusBladeGetMemoryFailure:
		return "Blade Get Memory Failure"
	case StatusHostGetMemoryFailure:
		return "Host Get Memory Failure"
	case StatusBladeGetMemoryByIdFailure:
		return "Blade Get Memory By Id Failure"
	case StatusHostGetMemoryByIdFailure:
		return "Host Get Memory By Id Failure"
	case StatusBladeFreeMemoryFailure:
		return "Blade Free Memory Failure"
	case StatusHostFreeMemoryFailure:
		return "Host Free Memory Failure"
	case StatusBladeGetMemoryResourceBlocksFailure:
		return "Blade Get Memory Resource Blocks Failure"
	case StatusBladeGetMemoryResourceBlockDetailsFailure:
		return "Blade Get Memory Resource Block By Id Failure"
	case StatusBladeGetPortsFailure:
		return "Blade Get Ports Failure"
	case StatusHostGetPortsFailure:
		return "Host Get Ports Failure"
	case StatusGetPortDetailsFailure:
		return "Get Port By Id Failure"
	case StatusGetMemoryDevicesFailure:
		return "Get Memory Devices Failure"
	case StatusGetMemoryDevicesDetailsFailure:
		return "Get Memory Devices Details Failure"
	case StatusApplianceResyncFailure:
		return "Appliance Resync Failure"
	case StatusBladeIdDoesNotExist:
		return "Blade Id Does Not Exist"
	case StatusAppliancesExceedMaximum:
		return "Maximum Appliance count exceeded"
	case StatusBladesExceedMaximum:
		return "Maximum Blade count exceeded for Appliance"
	case StatusHostsExceedMaximum:
		return "Maximum Host count exceeded"
	case StatusManagerInitializationFailure:
		return "Manager Initialization Failure"
	case StatusApplianceIdDuplicate:
		return "Appliance Id Already Exists"
	case StatusBladeIdDuplicate:
		return "Blade Id Already Exists"
	case StatusMemoryIdDuplicate:
		return "Memory Id Already Exist"
	case StatusPortIdDuplicate:
		return "Port Id Already Exist"
	case StatusHostIdDuplicate:
		return "Host Id Already Exist"
	}
	return "Unknown"

}

// Return the http status code of the the StatusType.
// return hardedcoded int value instead of defining from http module to reduce the import
func (e StatusCodeType) HttpStatusCode() int {
	switch e {
	case StatusOK:
		return http.StatusOK // 200
	case StatusComposePartialSuccess,
		StatusApplianceResyncPartialSuccess:
		return http.StatusPartialContent // 206
	case StatusBladeIdDoesNotExist,
		StatusApplianceIdDoesNotExist,
		StatusHostIdDoesNotExist,
		StatusMemoryIdDoesNotExist,
		StatusPortIdDoesNotExist,
		StatusResourceIdDoesNotExist,
		StatusMemoryDeviceIdDoesNotExist:
		return http.StatusNotFound // 404
	case StatusComposeMemoryByResourceFailure,
		StatusAssignMemoryFailure,
		StatusUnassignMemoryFailure,
		StatusHostGetMemoryFailure,
		StatusBladeGetMemoryFailure,
		StatusBladeGetMemoryByIdFailure,
		StatusHostGetMemoryByIdFailure,
		StatusBladeFreeMemoryFailure,
		StatusHostFreeMemoryFailure,
		StatusBladeGetMemoryResourceBlocksFailure,
		StatusBladeGetMemoryResourceBlockDetailsFailure,
		StatusBladeGetPortsFailure,
		StatusHostGetPortsFailure,
		StatusGetPortDetailsFailure,
		StatusGetMemoryDevicesFailure,
		StatusGetMemoryDevicesDetailsFailure,
		StatusApplianceResyncFailure,
		StatusApplianceIdDuplicate,
		StatusBladeIdDuplicate,
		StatusPortIdDuplicate,
		StatusHostIdDuplicate:
		return http.StatusConflict // 409
	case StatusBackendInterfaceFailure,
		StatusBladeCreateSessionFailure,
		StatusBladeDeleteSessionFailure,
		StatusHostCreateSessionFailure,
		StatusHostDeleteSessionFailure,
		StatusApplianceCreateSessionFailure,
		StatusApplianceDeleteSessionFailure,
		StatusManagerInitializationFailure:
		return http.StatusInternalServerError // 500
	case StatusAppliancesExceedMaximum,
		StatusBladesExceedMaximum,
		StatusHostsExceedMaximum:
		return http.StatusUnprocessableEntity //422
	default:
		return 0
	}
}
