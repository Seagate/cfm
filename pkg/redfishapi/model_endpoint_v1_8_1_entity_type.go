/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type EndpointV181EntityType string

// List of EndpointV181EntityType
const (
	ENDPOINTV181ENTITYTYPE_STORAGE_INITIATOR     EndpointV181EntityType = "StorageInitiator"
	ENDPOINTV181ENTITYTYPE_ROOT_COMPLEX          EndpointV181EntityType = "RootComplex"
	ENDPOINTV181ENTITYTYPE_NETWORK_CONTROLLER    EndpointV181EntityType = "NetworkController"
	ENDPOINTV181ENTITYTYPE_DRIVE                 EndpointV181EntityType = "Drive"
	ENDPOINTV181ENTITYTYPE_STORAGE_EXPANDER      EndpointV181EntityType = "StorageExpander"
	ENDPOINTV181ENTITYTYPE_DISPLAY_CONTROLLER    EndpointV181EntityType = "DisplayController"
	ENDPOINTV181ENTITYTYPE_BRIDGE                EndpointV181EntityType = "Bridge"
	ENDPOINTV181ENTITYTYPE_PROCESSOR             EndpointV181EntityType = "Processor"
	ENDPOINTV181ENTITYTYPE_VOLUME                EndpointV181EntityType = "Volume"
	ENDPOINTV181ENTITYTYPE_ACCELERATION_FUNCTION EndpointV181EntityType = "AccelerationFunction"
	ENDPOINTV181ENTITYTYPE_MEDIA_CONTROLLER      EndpointV181EntityType = "MediaController"
	ENDPOINTV181ENTITYTYPE_MEMORY_CHUNK          EndpointV181EntityType = "MemoryChunk"
	ENDPOINTV181ENTITYTYPE_SWITCH                EndpointV181EntityType = "Switch"
	ENDPOINTV181ENTITYTYPE_FABRIC_BRIDGE         EndpointV181EntityType = "FabricBridge"
	ENDPOINTV181ENTITYTYPE_MANAGER               EndpointV181EntityType = "Manager"
	ENDPOINTV181ENTITYTYPE_STORAGE_SUBSYSTEM     EndpointV181EntityType = "StorageSubsystem"
	ENDPOINTV181ENTITYTYPE_MEMORY                EndpointV181EntityType = "Memory"
	ENDPOINTV181ENTITYTYPE_CXL_DEVICE            EndpointV181EntityType = "CXLDevice"
)

// AssertEndpointV181EntityTypeRequired checks if the required fields are not zero-ed
func AssertEndpointV181EntityTypeRequired(obj EndpointV181EntityType) error {
	return nil
}

// AssertEndpointV181EntityTypeConstraints checks if the values respects the defined constraints
func AssertEndpointV181EntityTypeConstraints(obj EndpointV181EntityType) error {
	return nil
}
