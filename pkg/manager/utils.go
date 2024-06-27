// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"fmt"
	"strings"
)

// convertPortGCxlIdToSn - convert a cxl global id to a cxl port serial number (but without the cxl logical devices on the end).
// Example: "cb-7b-6a-39-22-df-e1-00" to "0xcb7b6a3922dfe100"
func ConvertPortGCxlIdToSn(gCxlId *string) *string {
	var sn string

	id := strings.Split(*gCxlId, ":")[0]
	tokens := strings.Split(id, "-")
	if len(tokens) > 0 {
		sn = fmt.Sprintf("0x%s", strings.Join(tokens, ""))
	}
	return &sn
}

// convertPortSnToGCxlId - convert a cxl port serial number to a cxl global id
// Example: "0xcb7b6a3922dfe100" to "cb-7b-6a-39-22-df-e1-00"
func ConvertPortSnToGCxlId(sn *string) *string {
	var ret string

	trimmed := strings.TrimPrefix(*sn, "0x")
	for i, chr := range strings.Split(trimmed, "") {
		ret += chr
		if i != 0 && i%2 == 0 {
			ret += "-"
		}
	}

	return &ret
}

func GetIdFromUriByName(uri, name string) (string, error) {
	target := -1

	elements := strings.Split(uri, "/")
	for i, elem := range elements {
		if elem == name {
			target = i
		}
	}

	if target < 0 {
		return "", fmt.Errorf("uri [%s] does not contain '%s'", uri, name)
	}

	if len(elements) < target+2 {
		return "", fmt.Errorf("uri [%s] does not contain an ID for '%s'", uri, name)
	}

	return elements[target+1], nil
}

// // Experiment to see what the code looks like to scan all the real\external devices for this information
// func GetExternHostPortByCxlSn(ctx context.Context, cxlSn string) (*CxlHostPort, error) {
// 	//Search for a linked host port by SN

// 	// CACHE: Get
// 	for hostId, host := range deviceCache.GetHosts() {
// 		// CACHE: Get
// 		//		Do I want to go look at the backend here, or, use the cached URIs?
// 		//			Right now, this woud be "locked" to using the cache
// 		hostPortUris, err := host.GetPorts(ctx)
// 		if err != nil {
// 			newErr := fmt.Errorf("failed to get host ports: %w", err)
// 			// logger.Error(newErr, "failure: get ports(host)")
// 			return nil, newErr
// 		}

// 		for _, uri := range hostPortUris {
// 			tokens := strings.Split(uri, "/")
// 			if len(tokens) == 0 {
// 				continue
// 			}
// 			pcieId := tokens[len(tokens)-1]
// 			hostPortId := "port" + pcieId

// 			reqHost := backend.GetHostPortSnByIdRequest{
// 				PortId:    hostPortId,
// 			}

// 			responseHost, err := host.BackendOps.GetHostPortSnById(ctx, &backend.ConfigurationSettings{}, &reqHost)
// 			if err != nil || responseHost == nil {
// 				newErr := fmt.Errorf("failed to get host [%s] port [%s] sn by id: %w", hostId, hostPortId, err)
// 				// logger.Error(newErr, "failure: get blade port details")
// 				return nil, &common.RequestError{StatusCode: common.StatusGetPortDetailsFailure, Err: newErr}
// 			}

// 			if responseHost.SerialNumber == cxlSn {
// 				//
// 			}
// 		}
// 	}

// 	return nil, nil
// }
