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

	mgmtRouteTesting := false

	if mgmtRouteTesting {
		testManagementRouteCreation(f5)
		testManagementRouteRead(f5)
		testManagementRouteUpdate(f5)
		testManagementRouteRead(f5)
		testManagementRouteDelete(f5)
		testManagementRouteRead(f5)
	}

}
