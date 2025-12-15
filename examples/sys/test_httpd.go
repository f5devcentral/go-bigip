package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////              HTTPD Configuration Testing         /////////////////
/////////////////////////////////////////////////////////////////////////////////

func testHTTPDRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HTTPD Configuration Read Operation ===")

	fmt.Println("Getting HTTPD configuration...")
	httpdConfig, err := f5.GetHTTPDConfig()
	if err != nil {
		log.Printf("Error getting HTTPD configuration: %v", err)
		return
	}
	if httpdConfig != nil {
		fmt.Printf("Retrieved HTTPD Configuration:\n")
		fmt.Printf("  Allow: %v\n", httpdConfig.Allow)
		fmt.Printf("  AuthPamIdleTimeout: %d\n", httpdConfig.AuthPamIdleTimeout)
		fmt.Printf("  LogLevel: %s\n", httpdConfig.LogLevel)
		fmt.Printf("  MaxClients: %d\n", httpdConfig.MaxClients)
		fmt.Printf("  SslCertfile: %s\n", httpdConfig.SslCertfile)
		fmt.Printf("  SslCertkeyfile: %s\n", httpdConfig.SslCertkeyfile)
	}
}

func testHTTPDCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HTTPD Configuration Create Operation ===")

	httpdConfig := &bigip.HTTPDConfig{
		Allow:              []string{"10.1.10.0/24", "192.168.1.0/24"},
		AuthPamIdleTimeout: 1800,
		MaxClients:         20,
	}

	fmt.Println("Setting HTTPD configuration...")
	err := f5.CreateHTTPDConfig(httpdConfig)
	if err != nil {
		log.Printf("Error creating HTTPD configuration: %v", err)
		return
	}
	fmt.Printf("HTTPD configuration set successfully:\n")
	fmt.Printf("  Allow: %v\n  AuthPamIdleTimeout: %d\n  MaxClients: %d\n",
		httpdConfig.Allow, httpdConfig.AuthPamIdleTimeout, httpdConfig.MaxClients)
}

func testHTTPDUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HTTPD Configuration Update Operation ===")

	httpdConfig := &bigip.HTTPDConfig{
		AuthPamIdleTimeout: 2400,
		Allow:              []string{"172.18.0.0/16"},
	}

	fmt.Println("Updating HTTPD configuration...")
	err := f5.ModifyHTTPDConfig(httpdConfig)
	if err != nil {
		log.Printf("Error updating HTTPD configuration: %v", err)
		return
	}
	fmt.Printf("HTTPD configuration updated successfully:\n")
	fmt.Printf("  Allow: %v\n  AuthPamIdleTimeout: %d\n", httpdConfig.Allow, httpdConfig.AuthPamIdleTimeout)
}

func testHTTPDDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HTTPD Configuration Delete Operation ===")

	fmt.Println("Resetting HTTPD configuration to defaults...")
	err := f5.DeleteHTTPDConfig()
	if err != nil {
		log.Printf("Error deleting HTTPD configuration: %v", err)
		return
	}
	fmt.Printf("HTTPD configuration reset to defaults.\n")
}
