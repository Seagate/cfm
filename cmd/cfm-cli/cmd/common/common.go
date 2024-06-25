// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package common

import "fmt"

const (
	WARNING_CXL_HOST_POWER_DOWN = "\nWARNING: Any connected CXL-Host MUST be powered down BEFORE (un)assigning memory to\\from a memory appliance port.\n"
)

func PromptYesNo(message string) error {

	input := "y" // Default to yes

	fmt.Println(message)

	for {
		fmt.Printf("Continue? [y\\N]: ")
		fmt.Scanln(&input)
		if input != "y" && input != "N" {
			fmt.Printf("Invalid input: ")
			continue
		}

		break
	}

	if input == "y" {
		fmt.Println("continuing...")
		return nil
	} else {
		fmt.Println("aborting...")
		return fmt.Errorf("aborting")
	}
}
