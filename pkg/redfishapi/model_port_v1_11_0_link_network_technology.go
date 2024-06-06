/*
 * Composer and Fabric Manager Redfish Service OpenAPI
 *
 * This API allows users to interact through the CFM Service with CXL Hosts and Memory Appliances. The main purpose of this interface is to allow the retrieval of information and the creation and mapping of memory from a Memory Appliance to a CXL host.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type PortV1110LinkNetworkTechnology string

// List of PortV1110LinkNetworkTechnology
const (
	PORTV1110LINKNETWORKTECHNOLOGY_ETHERNET      PortV1110LinkNetworkTechnology = "Ethernet"
	PORTV1110LINKNETWORKTECHNOLOGY_INFINI_BAND   PortV1110LinkNetworkTechnology = "InfiniBand"
	PORTV1110LINKNETWORKTECHNOLOGY_FIBRE_CHANNEL PortV1110LinkNetworkTechnology = "FibreChannel"
	PORTV1110LINKNETWORKTECHNOLOGY_GEN_Z         PortV1110LinkNetworkTechnology = "GenZ"
	PORTV1110LINKNETWORKTECHNOLOGY_PCIE          PortV1110LinkNetworkTechnology = "PCIe"
)

// AssertPortV1110LinkNetworkTechnologyRequired checks if the required fields are not zero-ed
func AssertPortV1110LinkNetworkTechnologyRequired(obj PortV1110LinkNetworkTechnology) error {
	return nil
}

// AssertPortV1110LinkNetworkTechnologyConstraints checks if the values respects the defined constraints
func AssertPortV1110LinkNetworkTechnologyConstraints(obj PortV1110LinkNetworkTechnology) error {
	return nil
}