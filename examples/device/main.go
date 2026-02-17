package main

import (
	"fmt"
	"os"

	"github.com/f5devcentral/go-bigip"
)

func main() {
	// Connect to the BIG-IP system.
	// Replace with your actual BIG-IP credentials and hostname
	config := &bigip.Config{
		//Address:           "https://192.168.1.1",
		Username: "admin",
		//Password:          "admin",
		CertVerifyDisable: true, // Disable certificate verification for testing purposes
	}

	config.Address = os.Getenv("BIGIP_ADDRESS")
	if config.Address != "" {
		fmt.Println("BIGIP_ADDRESS:", config.Address)
	} else {
		fmt.Println("BIGIP_ADDRESS is not set.")
	}

	config.Password = os.Getenv("BIGIP_PASSWORD")
	if config.Password != "" {
		fmt.Println("BIGIP_PASSWORD:", config.Password)
	} else {
		fmt.Println("BIGIP_PASSWORD is not set.")
	}

	f5 := bigip.NewSession(config)

	deviceSelfTesting := false
	deviceNameTesting := true

	if deviceSelfTesting {
		testDeviceSelfRead(f5)
		testDeviceSelfCreate(f5)
		testDeviceSelfRead(f5)
		testDeviceSelfUpdate(f5)
		testDeviceSelfRead(f5)
		testDeviceSelfDelete(f5)
		testDeviceSelfRead(f5)
	}
	if deviceNameTesting {
		testDeviceNameRead(f5)
		testDeviceNameCreate(f5)
		testDeviceNameRead(f5)
		testDeviceNameUpdate(f5)
		testDeviceNameRead(f5)
		testDeviceNameDelete(f5)
		testDeviceNameRead(f5)
	}

}
