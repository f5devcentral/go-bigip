package main

import (
	"github.com/f5devcentral/go-bigip"
)

func main() {
	// Connect to the BIG-IP system.
	// Replace with your actual BIG-IP credentials and hostname
	config := &bigip.Config{
		Address:           "https://192.168.1.1",
		Username:          "admin",
		Password:          "admin",
		CertVerifyDisable: true, // Disable certificate verification for testing purposes
	}
	f5 := bigip.NewSession(config)
	testUDPProfile(f5)
	testWebsocketProfile(f5)
	testHTMLProfile(f5)
	testAnalyticsProfile(f5)
	testRequestAdaptProfile(f5)
	testResponseAdaptProfile(f5)
}
