package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////       Device Name Testing                        /////////////////
/////////////////////////////////////////////////////////////////////////////////

func testDeviceNameRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Device Name Read Operation ===")

	fmt.Println("Getting Device Name...")
	name, err := f5.GetSelfDeviceName()
	if err != nil {
		log.Printf("Error getting Device Name: %v", err)
		return
	}

	fmt.Printf("  Current Device Name: %s\n", name)
}

func testDeviceNameCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Device Name Create Operation ===")

	targetName := "ltm02.f5lab.com"

	fmt.Printf("Renaming device to %s...\n", targetName)
	err := f5.CreateDeviceName(targetName)
	if err != nil {
		log.Printf("Error creating Device Name: %v", err)
		return
	}
	fmt.Printf("Device renamed successfully.\n")
}

func testDeviceNameUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Device Name Update Operation ===")

	targetName := "ltm02-updated.f5lab.com"

	fmt.Printf("Renaming device to %s...\n", targetName)
	err := f5.ModifyDeviceName(targetName)
	if err != nil {
		log.Printf("Error updating Device Name: %v", err)
		return
	}
	fmt.Printf("Device renamed successfully.\n")
}

func testDeviceNameDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Device Name Delete Operation ===")

	fmt.Println("Resetting device name to default (bigip1)...")
	err := f5.DeleteDeviceName()
	if err != nil {
		log.Printf("Error deleting Device Name: %v", err)
		return
	}
	fmt.Printf("Device name reset to default successfully.\n")

	// Verify default name after reset
	fmt.Println("\nVerifying default name after reset...")
	name, err := f5.GetSelfDeviceName()
	if err != nil {
		log.Printf("Error getting Device Name after reset: %v", err)
		return
	}

	fmt.Printf("  Device Name after reset: %s (expected: bigip1)\n", name)
	if name == "bigip1" {
		fmt.Println("Default name verified successfully.")
	} else {
		log.Println("Warning: Device name does not match expected default.")
	}
}
