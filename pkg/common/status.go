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
	StatusOK             StatusCodeType = iota
	StatusCreatedSuccess                //201

	StatusAppliancesExceedMaximum //422
	StatusBladesExceedMaximum     //422
	StatusHostsExceedMaximum      //422

	StatusInvalidBackend                   //400: StatusBadRequest
	StatusApplianceComposeMemoryRequestBad //400
	StatusApplianceIdMismatch              //400
	StatusResourceNotEnough                //400
	StatusEndpointNotConnected             //400
	StatusBackendInterfaceFailure          //500: StatusInternalServerError

	StatusComposeMemoryFailure              //409: StatusConflict
	StatusComposeMemoryByResourceFailure    //409
	StatusAssignMemoryFailure               //409
	StatusUnassignMemoryFailure             //409
	StatusHostGetMemoryFailure              //409
	StatusApplianceGetMemoryFailure         //409
	StatusApplianceGetMemoryByIdFailure     //409
	StatusHostGetMemoryByIdFailure          //409
	StatusApplianceFreeMemoryFailure        //409
	StatusHostFreeMemoryFailure             //409
	StatusGetMemoryResourceBlocksFailure    //409
	StatusGetMemoryResourceBlockByIdFailure //409
	StatusApplianceGetPortsFailure          //409
	StatusHostGetPortsFailure               //409
	StatusGetPortDetailsFailure             //409
	StatusGetMemoryDevicesFailure           //409
	StatusGetMemoryDevicesDetailsFailure    //409

	StatusApplianceNameDoesNotExist     //404: StatusNotFound
	StatusBladeIdDoesNotExist           //404
	StatusApplianceCreateSessionFailure //500
	StatusApplianceDeleteSessionFailure //500
	StatusApplianceIdDoesNotExist       //404

	StatusHostNameDoesNotExist             //404
	StatusHostCreateSessionFailure         //500
	StatusHostDeleteSessionFailure         //500
	StatusApplianceGetResourcesFailure     //500
	StatusApplianceGetResourceFailure      //500
	StatusApplianceGetEndpointsFailure     //500
	StatusApplianceGetOneEndpointFailure   //500
	StatusApplianceUnassginMemoryFailure   // 500
	StatusApplianceUnallocateMemoryFailure //500

	StatusHostIdDoesNotExist //404

	StatusMemoryIdDoesNotExist       //404
	StatusPortIdDoesNotExist         //404
	StatusResourceIdDoesNotExist     //404
	StatusMemoryDeviceIdDoesNotExist //404
	StatusNoCapacityInBlock          //404

	StatusEndpointDoesNotExist //404
	StatusSessionDoesNotExist  //404

	StatusManagerInitializationFailure //500

)

// Return a string representation of our StatusType. When this value is printed as a
// string (%s), the value is determined by this function.
func (e StatusCodeType) String() string {
	switch e {
	case StatusOK:
		return "Success"
	case StatusCreatedSuccess:
		return "Created Successfully"
	case StatusComposeMemoryFailure:
		return "Compose Memory Failed"
	case StatusComposeMemoryByResourceFailure:
		return "Compose Memory By Resource Failed"
	case StatusAssignMemoryFailure:
		return "Assign Memory Failed"
	case StatusUnassignMemoryFailure:
		return "Unassign Memory Failed"
	case StatusApplianceNameDoesNotExist:
		return "Appliance Name does not exist"
	case StatusHostNameDoesNotExist:
		return "Host Name does not exist"
	case StatusInvalidBackend:
		return "Backend interface does not exist"
	case StatusBackendInterfaceFailure:
		return "Backend Interface does not exist"
	case StatusApplianceCreateSessionFailure:
		return "Appliance Create Session Failure"
	case StatusApplianceDeleteSessionFailure:
		return "Appliance Delete Session Failure"
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
	case StatusApplianceGetMemoryFailure:
		return "Appliance Get Memory Failure"
	case StatusHostGetMemoryFailure:
		return "Host Get Memory Failure"
	case StatusApplianceGetMemoryByIdFailure:
		return "Appliance Get Memory By Id Failure"
	case StatusHostGetMemoryByIdFailure:
		return "Host Get Memory By Id Failure"
	case StatusApplianceFreeMemoryFailure:
		return "Appliance Free Memory Failure"
	case StatusHostFreeMemoryFailure:
		return "Host Free Memory Failure"
	case StatusGetMemoryResourceBlocksFailure:
		return "Appliance Get Memory Resource Blocks Failure"
	case StatusGetMemoryResourceBlockByIdFailure:
		return "Appliance Get Memory Resource Block By Id Failure"
	case StatusApplianceGetPortsFailure:
		return "Appliance Get Ports Failure"
	case StatusHostGetPortsFailure:
		return "Host Get Ports Failure"
	case StatusGetPortDetailsFailure:
		return "Get Port By Id Failure"
	case StatusGetMemoryDevicesFailure:
		return "Get Memory Devices Failure"
	case StatusGetMemoryDevicesDetailsFailure:
		return "Get Memory Devices Details Failure"
	case StatusResourceNotEnough:
		return "Resource Not Enough"
	case StatusSessionDoesNotExist:
		return "Session Does Not Exist"
	case StatusApplianceGetResourceFailure:
		return "Appliance Get A Specific Resource Failure"
	case StatusApplianceGetResourcesFailure:
		return "Appliance Get Resources Failure"
	case StatusNoCapacityInBlock:
		return "No Capacity In Block"
	case StatusEndpointDoesNotExist:
		return "Appliance Does Not Exist"
	case StatusApplianceGetEndpointsFailure:
		return "Appliance Get Endpoints Failure"
	case StatusApplianceGetOneEndpointFailure:
		return "Appliance Get One Endpoint Failure"
	case StatusApplianceIdMismatch:
		return "Appliance Id Mismatch"
	case StatusApplianceComposeMemoryRequestBad:
		return "Appliance Compose Memory Request Empty"
	case StatusEndpointNotConnected:
		return "Endpoint Not Connected"
	case StatusApplianceUnassginMemoryFailure:
		return "Appliance Unassgin Memory Failure"
	case StatusApplianceUnallocateMemoryFailure:
		return "Appliance Unallocate Memory Failure"
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
	}
	return "Unknown"

}

// Return the http status code of the the StatusType.
// return hardedcoded int value instead of defining from http module to reduce the import
func (e StatusCodeType) HttpStatusCode() int {
	switch e {
	case StatusOK:
		return http.StatusOK // 200
	case StatusCreatedSuccess:
		return http.StatusCreated // 201
	case StatusInvalidBackend,
		StatusApplianceComposeMemoryRequestBad,
		StatusApplianceIdMismatch,
		StatusResourceNotEnough,
		StatusEndpointNotConnected,
		StatusBladeIdDoesNotExist:
		return http.StatusBadRequest // 400
	case StatusApplianceNameDoesNotExist,
		StatusApplianceIdDoesNotExist,
		StatusHostNameDoesNotExist,
		StatusHostIdDoesNotExist,
		StatusMemoryIdDoesNotExist,
		StatusPortIdDoesNotExist,
		StatusResourceIdDoesNotExist,
		StatusEndpointDoesNotExist,
		StatusMemoryDeviceIdDoesNotExist,
		StatusNoCapacityInBlock,
		StatusSessionDoesNotExist:
		return http.StatusNotFound // 404
	case StatusComposeMemoryFailure,
		StatusComposeMemoryByResourceFailure,
		StatusAssignMemoryFailure,
		StatusUnassignMemoryFailure,
		StatusHostGetMemoryFailure,
		StatusApplianceGetMemoryFailure,
		StatusApplianceGetMemoryByIdFailure,
		StatusHostGetMemoryByIdFailure,
		StatusApplianceFreeMemoryFailure,
		StatusHostFreeMemoryFailure,
		StatusGetMemoryResourceBlocksFailure,
		StatusGetMemoryResourceBlockByIdFailure,
		StatusApplianceGetPortsFailure,
		StatusHostGetPortsFailure,
		StatusGetPortDetailsFailure,
		StatusGetMemoryDevicesFailure,
		StatusGetMemoryDevicesDetailsFailure:
		return http.StatusConflict // 409
	case StatusBackendInterfaceFailure,
		StatusApplianceCreateSessionFailure,
		StatusApplianceDeleteSessionFailure,
		StatusHostCreateSessionFailure,
		StatusHostDeleteSessionFailure,
		StatusApplianceGetResourcesFailure,
		StatusApplianceGetResourceFailure,
		StatusApplianceGetEndpointsFailure,
		StatusApplianceGetOneEndpointFailure,
		StatusApplianceUnassginMemoryFailure,
		StatusApplianceUnallocateMemoryFailure,
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
