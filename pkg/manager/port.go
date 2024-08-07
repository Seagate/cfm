// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates

package manager

import (
	"context"
	"fmt"
	"strings"

	"cfm/pkg/backend"
	"cfm/pkg/common"
	"cfm/pkg/openapi"

	"k8s.io/klog/v2"
)

const LINKED_PORT_NOT_FOUND = "NOT_FOUND"

//////////////////////////////////////////////////////////////////
///////////////////////// CxlBladePort ///////////////////////////
//////////////////////////////////////////////////////////////////

type CxlBladePort struct {
	Id          string
	Uri         string
	BladeId     string
	ApplianceId string
	// Status      string

	// Cached data
	cacheUpdated bool // Don't know what cache update mechanism looks like yet.  Start with this.
	details      openapi.PortInformation

	// Backend access data
	backendOps backend.BackendOperations
}

func NewCxlBladePortById(ctx context.Context, applianceId, bladeId, portId string, ops backend.BackendOperations) (*CxlBladePort, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewCxlBladePortById: ", "portId", portId, "bladeId", bladeId, "applianceId", applianceId, "backend", ops.GetBackendInfo(ctx).BackendName)

	p := CxlBladePort{
		Id:          portId,
		Uri:         GetCfmUriBladePortId(applianceId, bladeId, portId),
		BladeId:     bladeId,
		ApplianceId: applianceId,

		backendOps: ops,
	}

	logger.V(2).Info("success: new cxl blade port", "portId", p.Id, "bladeId", p.BladeId, "applianceId", p.ApplianceId)

	return &p, nil
}

func (p *CxlBladePort) GetDetails(ctx context.Context) (openapi.PortInformation, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetDetails: ", "portId", p.Id, "bladeId", p.BladeId, "applianceId", p.ApplianceId)

	if !p.cacheUpdated {
		req := backend.GetPortDetailsRequest{
			PortId: p.Id,
		}

		response, err := p.backendOps.GetPortDetails(ctx, &backend.ConfigurationSettings{}, &req)
		if err != nil || response == nil {
			newErr := fmt.Errorf("failed to get details (backend): appliance [%s] blade [%s] port [%s]: %w", p.ApplianceId, p.BladeId, p.Id, err)
			logger.Error(newErr, "failure: get blade port details")
			return openapi.PortInformation{}, &common.RequestError{StatusCode: common.StatusGetPortDetailsFailure, Err: newErr}
		}

		p.details = openapi.PortInformation{
			Id:               response.PortInformation.Id,
			GCxlId:           response.PortInformation.GCxlId,
			LinkedPortUri:    "", // Find below
			PortProtocol:     response.PortInformation.PortProtocol,
			PortMedium:       response.PortInformation.PortMedium,
			CurrentSpeedGbps: response.PortInformation.CurrentSpeedGbps,
			StatusHealth:     response.PortInformation.StatusHealth,
			StatusState:      response.PortInformation.StatusState,
			Width:            response.PortInformation.Width,
			LinkStatus:       response.PortInformation.LinkStatus,
			LinkState:        response.PortInformation.StatusState,
		}

		p.ValidateCache()
	}

	details := p.details // Use local copy to allow update of dynamic link outside of this objects cache

	cxlSn := GetCxlSnFromGCxlId(details.GCxlId)

OUT:
	for _, host := range deviceCache.GetHosts() {
		for _, port := range host.GetPorts(ctx) {

			if strings.EqualFold(port.CxlSn, cxlSn) { //NOTE: Do NOT call CxlHostPort.GetDetails here.  Creates inifinte loop.
				details.LinkedPortUri = GetCfmUriHostPortId(host.Id, port.Id)
				logger.V(2).Info("linked host port found", "uri", details.LinkedPortUri)
				break OUT
			}
		}
	}

	logger.V(2).Info("success: get cxl blade port details", "portId", p.Id, "bladeId", p.BladeId, "applianceId", p.ApplianceId)

	return details, nil
}

func (p *CxlBladePort) InvalidateCache() {
	p.cacheUpdated = false
}

func (p *CxlBladePort) ValidateCache() {
	p.cacheUpdated = true
}

func (p *CxlBladePort) init(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> init: ", "portId", p.Id, "bladeId", p.BladeId, "applianceId", p.ApplianceId)

	p.InvalidateCache()

	_, err := p.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("port [%s] init failed on blade [%s]: %w", p.Id, p.BladeId, err)
		logger.Error(newErr, "failure: init blade port")
		return newErr
	}

	logger.V(2).Info("success: init blade port", "portId", p.Id, "bladeId", p.BladeId, "applianceId", p.ApplianceId)

	return nil
}

//////////////////////////////////////////////////////////////////
///////////////////////// CxlHostPort ////////////////////////////
//////////////////////////////////////////////////////////////////

const (
	ID_PREFIX_HOST_PORT_DFLT string = "port"
)

type CxlHostPort struct {
	Id            string // "port19-00" (as input from client)
	BackendPortId string // "19-00" (only used internally)
	Uri           string
	HostId        string
	// Status        string

	// Cached data
	cacheUpdated bool // Don't know what cache update mechanism looks like yet.  Start with this.
	details      openapi.PortInformation
	CxlSn        string

	// Backend access data
	backendOps backend.BackendOperations
}

func NewCxlHostPortById(ctx context.Context, hostId, backendPortId string, ops backend.BackendOperations) (*CxlHostPort, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> NewCxlHostPortById: ", "backendPortId", backendPortId, "hostId", hostId, "backend", ops.GetBackendInfo(ctx).BackendName)

	portId := GenerateFrontendHostPortId(backendPortId)

	p := CxlHostPort{
		Id:            portId,
		BackendPortId: backendPortId,
		Uri:           GetCfmUriHostPortId(hostId, portId),
		HostId:        hostId,

		backendOps: ops,
	}

	logger.V(2).Info("success: new cxl host port", "portId", p.Id, "hostId", p.HostId)

	return &p, nil
}

// GetDetails: Gets the BLADE port details from the other side of the cxl cable connection.
// It first retrieves the connected HOST port serial number, then searches for that same SN\GCxlId on the blade side.
// From a hardware perspective, on the host, a PCIe device represents the physical port, which is identified by a "X9-00" string. (eg: 19-00, 29-00, etc)
// From a user perspective, the host port is identified by a "portX9-00" string. (to be similar to the nomenclature of the memory appliance port id's which are identified by a "portX" string. (eg: port0, port1, etc))
func (p *CxlHostPort) GetDetails(ctx context.Context) (openapi.PortInformation, error) {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> GetDetails: ", "portId", p.Id, "hostId", p.HostId)
	var err error

	if !p.cacheUpdated {
		// Get\Find SN from host port (ie host pcie device)
		reqHost := backend.GetHostPortSnByIdRequest{
			PortId: p.BackendPortId,
		}

		response, err := p.backendOps.GetHostPortSnById(ctx, &backend.ConfigurationSettings{}, &reqHost)
		if err != nil || response == nil {
			newErr := fmt.Errorf("failed to get host [%s] port [%s] sn by id: %w", p.HostId, p.Id, err)
			logger.Error(newErr, "failure: get host port details")
			return openapi.PortInformation{}, &common.RequestError{StatusCode: common.StatusGetPortDetailsFailure, Err: newErr}
		}

		p.CxlSn = response.SerialNumber

		p.details = openapi.PortInformation{
			Id:            p.Id,
			GCxlId:        GetGCxlIdFromCxlSn(p.CxlSn),
			LinkedPortUri: LINKED_PORT_NOT_FOUND,
		}

		p.ValidateCache()
	}

	details := p.details

OUT:
	// Find the linked blade port
	for _, appliance := range deviceCache.GetAppliances() {
		for _, blade := range appliance.GetBlades(ctx) {
			for _, port := range blade.GetPorts(ctx) {

				if strings.EqualFold(GetCxlSnFromGCxlId(port.details.GCxlId), p.CxlSn) {
					details, err = port.GetDetails(ctx)
					if err != nil {
						newErr := fmt.Errorf("failed to get appliance [%s] blade [%s] port [%s] details for host [%s] port [%s]: %w", appliance.Id, blade.Id, port.Id, p.HostId, p.Id, err)
						logger.Error(newErr, "failure: get host port details")
						return openapi.PortInformation{}, &common.RequestError{StatusCode: common.StatusGetPortDetailsFailure, Err: newErr}
					}

					details.Id = p.Id
					details.LinkedPortUri = port.Uri

					break OUT
				}
			}
		}
	}

	logger.V(4).Info("success: GetDetails: ", "portId", p.Id, "hostId", p.HostId)

	return details, nil
}

func (p *CxlHostPort) InvalidateCache() {
	p.cacheUpdated = false
}

func (p *CxlHostPort) ValidateCache() {
	// p.cacheUpdated = true
	p.cacheUpdated = false // Temporarily disable host cache usage
}

func (p *CxlHostPort) init(ctx context.Context) error {
	logger := klog.FromContext(ctx)
	logger.V(4).Info(">>>>>> init: ", "portId", p.Id, "host", p.HostId)

	p.InvalidateCache()

	_, err := p.GetDetails(ctx)
	if err != nil {
		newErr := fmt.Errorf("port [%s] init failed on host [%s]: %w", p.Id, p.HostId, err)
		logger.Error(newErr, "failure: init host port")
		return newErr
	}

	logger.V(2).Info("success: init host port", "portId", p.Id, "host", p.HostId)

	return nil
}

//////////////////////////////////////////////////////////////////
//////////////////////////// Helpers /////////////////////////////
//////////////////////////////////////////////////////////////////

func GenerateFrontendHostPortId(backendPortId string) string {
	// Current host port id string format: portXX-XX
	return fmt.Sprintf("port%s", backendPortId)
}
