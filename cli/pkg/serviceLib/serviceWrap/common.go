/*
Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

*/

package serviceWrap

import (
	"fmt"
	"net/netip"
	"strings"

	service "cfm/pkg/client"

	"github.com/google/uuid"
)

func GetServiceClient(ip string, networkPort uint16) *service.APIClient {
	// Instantiate new configuration using openapi funciton.
	config := service.NewConfiguration()

	// Create, then pass, string for IP Address and Network port like "127.0.0.1:8080"
	config.Host = fmt.Sprintf("%s:%d", ip, networkPort)
	//TODO: Add this back in??  Check to see where this goes and if the service code is using it
	// // Pass debug value.
	// config.Debug = debug

	// This creates an API client, passing it the above configuration, and gathers a pointer to it.
	serviceClient := service.NewAPIClient(config)

	// Returns a pointer to the API client that will connect to the cfm-service.
	return serviceClient
}

// Returns error if provided network is not within the valid Linux range.
func ValidatePort(port uint16) error {
	if port == 0 {
		return fmt.Errorf("invalid port value")
	}
	return nil
}

// Returns error if provided IP is not a valid IP Address.
func ValidateIPAddress(ip string) error {
	parsedAddr, err := netip.ParseAddr(ip)
	// Check if there was an error parsing the string into an IPv4.
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// Check if the value in IPv4 is unspecified.  This disallows all zeros.
	if parsedAddr.IsUnspecified() {
		return fmt.Errorf("zero IP address is not allowed: %s", ip)
	}
	// Check if the value in IPv4 is a valid IP.  It should be at this point.
	// If the value in IPv4 is NOT valid, then we complain to the user.
	if !parsedAddr.IsValid() {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// Returns error if UUID is not a UUID.
func ValidateUUID(uuidStr string) error {
	if len(uuidStr) == 0 {
		return nil // Ignore the error if UUID is not provided
	}

	_, err := uuid.Parse(uuidStr)
	if err != nil {
		// // Check if the error is due to invalid UUID length
		// if strings.Contains(err.Error(), "invalid UUID length") {
		// 	return nil // Ignore the error if UUID length is 0
		// }
		return fmt.Errorf("invalid UUID format: %s", err)
	}

	return nil
}

// Check if a provided ID is included in a Service Collection.
func IsMember(c *service.Collection, id string) bool {
	// Confirm collection passed to function is not null.
	// If it is nill, then return false.
	if c == nil || c.GetMemberCount() == 0 {
		fmt.Println("empty collection")
		return false
	}

	// Iterate through the collection's membership, checking for an ID match.
	for _, mi := range c.Members {
		// Check if the ID passed to this function is in the collection URI.
		if strings.Contains(mi.GetUri(), id) {
			return true
		}
	}

	return false
}

// The URI embeds many useful values as the last item in the string.
// Convience function to extract the last string in the URI
func ReadLastItemFromUri(uri string) string {
	collectMembersURI := strings.Split(uri, "/")
	return collectMembersURI[len(collectMembersURI)-1]
}

// AskForConfirmation prompts the user for a confirmation (y/n) and returns true if confirmed, false otherwise
// func askForConfirmation(prompt string) bool {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print(prompt)
// 		answer, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("Error reading input.")
// 			return false
// 		}
// 		answer = strings.TrimSpace(answer)
// 		answer = strings.ToLower(answer)
// 		if answer == "y" || answer == "yes" {
// 			return true
// 		} else if answer == "n" || answer == "no" {
// 			return false
// 		} else {
// 			fmt.Println("Please enter either 'y' or 'n'")
// 		}
// 	}
// }

type ApplianceBladeKey struct {
	ApplianceId string
	BladeId     string
}

func NewApplianceBladeKey(applId, bladeId string) *ApplianceBladeKey {
	return &ApplianceBladeKey{
		ApplianceId: applId,
		BladeId:     bladeId,
	}
}
