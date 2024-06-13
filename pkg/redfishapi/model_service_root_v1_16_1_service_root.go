/*
 * Composable Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

// ServiceRootV1161ServiceRoot - The ServiceRoot schema describes the root of the Redfish service, located at the '/redfish/v1' URI.  All other resources accessible through the Redfish interface on this device are linked directly or indirectly from the service root.
type ServiceRootV1161ServiceRoot struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	AccountService OdataV4IdRef `json:"AccountService,omitempty"`

	AggregationService OdataV4IdRef `json:"AggregationService,omitempty"`

	Cables OdataV4IdRef `json:"Cables,omitempty"`

	CertificateService OdataV4IdRef `json:"CertificateService,omitempty"`

	Chassis OdataV4IdRef `json:"Chassis,omitempty"`

	ComponentIntegrity OdataV4IdRef `json:"ComponentIntegrity,omitempty"`

	CompositionService OdataV4IdRef `json:"CompositionService,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	EventService OdataV4IdRef `json:"EventService,omitempty"`

	Fabrics OdataV4IdRef `json:"Fabrics,omitempty"`

	Facilities OdataV4IdRef `json:"Facilities,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	JobService OdataV4IdRef `json:"JobService,omitempty"`

	JsonSchemas OdataV4IdRef `json:"JsonSchemas,omitempty"`

	KeyService OdataV4IdRef `json:"KeyService,omitempty"`

	LicenseService OdataV4IdRef `json:"LicenseService,omitempty"`

	Links ServiceRootV1161Links `json:"Links"`

	Managers OdataV4IdRef `json:"Managers,omitempty"`

	NVMeDomains OdataV4IdRef `json:"NVMeDomains,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PowerEquipment OdataV4IdRef `json:"PowerEquipment,omitempty"`

	// The product associated with this Redfish service.
	Product *string `json:"Product,omitempty"`

	ProtocolFeaturesSupported ServiceRootV1161ProtocolFeaturesSupported `json:"ProtocolFeaturesSupported,omitempty"`

	// The version of the Redfish service.
	RedfishVersion string `json:"RedfishVersion,omitempty"`

	RegisteredClients OdataV4IdRef `json:"RegisteredClients,omitempty"`

	Registries OdataV4IdRef `json:"Registries,omitempty"`

	ResourceBlocks OdataV4IdRef `json:"ResourceBlocks,omitempty"`

	ServiceConditions OdataV4IdRef `json:"ServiceConditions,omitempty"`

	// The vendor or user-provided product and service identifier.
	ServiceIdentification string `json:"ServiceIdentification,omitempty"`

	SessionService OdataV4IdRef `json:"SessionService,omitempty"`

	Storage OdataV4IdRef `json:"Storage,omitempty"`

	StorageServices OdataV4IdRef `json:"StorageServices,omitempty"`

	StorageSystems OdataV4IdRef `json:"StorageSystems,omitempty"`

	Systems OdataV4IdRef `json:"Systems,omitempty"`

	Tasks OdataV4IdRef `json:"Tasks,omitempty"`

	TelemetryService OdataV4IdRef `json:"TelemetryService,omitempty"`

	ThermalEquipment OdataV4IdRef `json:"ThermalEquipment,omitempty"`

	// Unique identifier for a service instance.  When SSDP is used, this value contains the same UUID returned in an HTTP `200 OK` response from an SSDP `M-SEARCH` request during discovery.
	UUID *string `json:"UUID,omitempty"`

	UpdateService OdataV4IdRef `json:"UpdateService,omitempty"`

	// The vendor or manufacturer associated with this Redfish service.
	Vendor *string `json:"Vendor,omitempty"`
}

// AssertServiceRootV1161ServiceRootRequired checks if the required fields are not zero-ed
func AssertServiceRootV1161ServiceRootRequired(obj ServiceRootV1161ServiceRoot) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Links":       obj.Links,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertOdataV4IdRefRequired(obj.AccountService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.AggregationService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Cables); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.CertificateService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Chassis); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ComponentIntegrity); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.CompositionService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.EventService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Fabrics); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Facilities); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.JobService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.JsonSchemas); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.KeyService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.LicenseService); err != nil {
		return err
	}
	if err := AssertServiceRootV1161LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Managers); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.NVMeDomains); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.PowerEquipment); err != nil {
		return err
	}
	if err := AssertServiceRootV1161ProtocolFeaturesSupportedRequired(obj.ProtocolFeaturesSupported); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.RegisteredClients); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Registries); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ResourceBlocks); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ServiceConditions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.SessionService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Storage); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.StorageServices); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.StorageSystems); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Systems); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Tasks); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.TelemetryService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ThermalEquipment); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.UpdateService); err != nil {
		return err
	}
	return nil
}

// AssertServiceRootV1161ServiceRootConstraints checks if the values respects the defined constraints
func AssertServiceRootV1161ServiceRootConstraints(obj ServiceRootV1161ServiceRoot) error {
	return nil
}
