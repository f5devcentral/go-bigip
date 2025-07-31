package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

// testUDPProfile tests CRUD operations for UDP profiles
func testUDPProfile(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing UDP Profile CRUD Operations ===")

	// Test profile configuration
	udpProfile := &bigip.UdpProfile{
		Name:             "test-udp-profile",
		DefaultsFrom:     "/Common/udp",
		IdleTimeout:      "60",
		BufferMaxBytes:   65535,
		BufferMaxPackets: 0,
		IPTosToClient:    "pass-through",
		NoChecksum:       "disabled",
	}

	// CREATE: Add UDP profile
	fmt.Println("Creating UDP profile...")
	err := f5.AddUDPProfile(udpProfile)
	if err != nil {
		log.Printf("Error creating UDP profile: %v", err)
		return
	}
	fmt.Println("UDP profile created successfully")

	// READ: Get all UDP profiles
	fmt.Println("Getting all UDP profiles...")
	profiles, err := f5.UDPProfiles()
	if err != nil {
		log.Printf("Error getting UDP profiles: %v", err)
		return
	}
	fmt.Printf("Found %d UDP profiles\n", len(profiles.UdpProfiles))

	// READ: Get specific UDP profile
	fmt.Println("Getting specific UDP profile...")
	profile, err := f5.GetUDPProfile("test-udp-profile")
	if err != nil {
		log.Printf("Error getting UDP profile: %v", err)
		return
	}
	if profile != nil {
		fmt.Printf("Retrieved UDP profile: %s with idle timeout: %s\n", profile.Name, profile.IdleTimeout)
	}

	// UPDATE: Modify UDP profile
	fmt.Println("Updating UDP profile...")
	udpProfile.IdleTimeout = "120"
	udpProfile.BufferMaxBytes = 32768
	err = f5.ModifyUDPProfile("test-udp-profile", udpProfile)
	if err != nil {
		log.Printf("Error updating UDP profile: %v", err)
		return
	}
	fmt.Println("UDP profile updated successfully")

	// Verify update
	updatedProfile, err := f5.GetUDPProfile("test-udp-profile")
	if err != nil {
		log.Printf("Error getting updated UDP profile: %v", err)
		return
	}
	if updatedProfile != nil && updatedProfile.IdleTimeout == "120" && updatedProfile.BufferMaxBytes == 32768 {
		fmt.Printf("Updated UDP profile idle timeout: %s\n", updatedProfile.IdleTimeout)
	}

	// DELETE: Remove UDP profile
	fmt.Println("Deleting UDP profile...")
	err = f5.DeleteUDPProfile("test-udp-profile")
	if err != nil {
		log.Printf("Error deleting UDP profile: %v", err)
		return
	}
	fmt.Println("UDP profile deleted successfully")

	// Verify deletion
	_, err = f5.GetUDPProfile("test-udp-profile")
	if err != nil && err.Error() == "01020036:3: The requested UDP Profile (/Common/test-udp-profile) was not found." {
		fmt.Println("UDP profile deletion confirmed")
	}
}

// testWebsocketProfile tests CRUD operations for Websocket profiles
func testWebsocketProfile(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Websocket Profile CRUD Operations ===")

	// Test profile configuration
	websocketProfile := &bigip.WebsocketProfile{
		Name:         "test-websocket-profile",
		DefaultsFrom: "/Common/websocket",
		Compression:  "disabled",
		Masking:      "selective",
		NoDelay:      "enabled",
		WindowBits:   15,
	}

	// CREATE: Add Websocket profile
	fmt.Println("Creating Websocket profile...")
	err := f5.AddWebsocketProfile(websocketProfile)
	if err != nil {
		log.Printf("Error creating Websocket profile: %v", err)
		return
	}
	fmt.Println("Websocket profile created successfully")

	// READ: Get all Websocket profiles
	fmt.Println("Getting all Websocket profiles...")
	profiles, err := f5.WebsocketProfiles()
	if err != nil {
		log.Printf("Error getting Websocket profiles: %v", err)
		return
	}
	fmt.Printf("Found %d Websocket profiles\n", len(profiles.WebsocketProfiles))

	// READ: Get specific Websocket profile
	fmt.Println("Getting specific Websocket profile...")
	profile, err := f5.GetWebsocketProfile("test-websocket-profile")
	if err != nil {
		log.Printf("Error getting Websocket profile: %v", err)
		return
	}
	if profile != nil {
		fmt.Printf("Retrieved Websocket profile: %s with compression: %s\n", profile.Name, profile.Compression)
	}

	// UPDATE: Modify Websocket profile
	fmt.Println("Updating Websocket profile...")
	websocketProfile.Compression = "enabled"
	websocketProfile.WindowBits = 10
	err = f5.ModifyWebsocketProfile("test-websocket-profile", websocketProfile)
	if err != nil {
		log.Printf("Error updating Websocket profile: %v", err)
		return
	}
	fmt.Println("Websocket profile updated successfully")

	// Verify update
	updatedProfile, err := f5.GetWebsocketProfile("test-websocket-profile")
	if err != nil {
		log.Printf("Error getting updated Websocket profile: %v", err)
		return
	}
	if updatedProfile != nil && updatedProfile.Compression == "enabled" && updatedProfile.WindowBits == 10 {
		fmt.Printf("Updated Websocket profile compression: %s\n", updatedProfile.Compression)
	}

	// DELETE: Remove Websocket profile
	fmt.Println("Deleting Websocket profile...")
	err = f5.DeleteWebsocketProfile("test-websocket-profile")
	if err != nil {
		log.Printf("Error deleting Websocket profile: %v", err)
		return
	}
	fmt.Println("Websocket profile deleted successfully")

	// Verify deletion
	_, err = f5.GetWebsocketProfile("test-websocket-profile")
	if err != nil && err.Error() == "01020036:3: The requested Websocket Profile (/Common/test-websocket-profile) was not found." {
		fmt.Println("Websocket profile deletion confirmed")
	}
}

// testHTMLProfile tests CRUD operations for HTML profiles
func testHTMLProfile(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HTML Profile CRUD Operations ===")

	// Test profile configuration
	htmlProfile := &bigip.HTMLProfile{
		Name:             "test-html-profile",
		DefaultsFrom:     "/Common/html",
		ContentDetection: "enabled",
		ContentSelection: []string{"text/html", "text/xhtml"},
	}

	// CREATE: Add HTML profile
	fmt.Println("Creating HTML profile...")
	err := f5.AddHTMLProfile(htmlProfile)
	if err != nil {
		log.Printf("Error creating HTML profile: %v", err)
		return
	}
	fmt.Println("HTML profile created successfully")

	// READ: Get all HTML profiles
	fmt.Println("Getting all HTML profiles...")
	profiles, err := f5.HTMLProfiles()
	if err != nil {
		log.Printf("Error getting HTML profiles: %v", err)
		return
	}
	fmt.Printf("Found %d HTML profiles\n", len(profiles.HTMLProfiles))

	// READ: Get specific HTML profile
	fmt.Println("Getting specific HTML profile...")
	profile, err := f5.GetHTMLProfile("test-html-profile")
	if err != nil {
		log.Printf("Error getting HTML profile: %v", err)
		return
	}
	if profile != nil {
		fmt.Printf("Retrieved HTML profile: %s with content detection: %s\n", profile.Name, profile.ContentDetection)
	}

	// UPDATE: Modify HTML profile
	fmt.Println("Updating HTML profile...")
	htmlProfile.ContentDetection = "disabled"
	htmlProfile.ContentSelection = []string{"text/html"}
	err = f5.ModifyHTMLProfile("test-html-profile", htmlProfile)
	if err != nil {
		log.Printf("Error updating HTML profile: %v", err)
		return
	}
	fmt.Println("HTML profile updated successfully")

	// Verify update
	updatedProfile, err := f5.GetHTMLProfile("test-html-profile")
	if err != nil {
		log.Printf("Error getting updated HTML profile: %v", err)
		return
	}
	if updatedProfile != nil && updatedProfile.ContentDetection == "disabled" {
		fmt.Printf("Updated HTML profile content detection: %s\n", updatedProfile.ContentDetection)
	}

	// DELETE: Remove HTML profile
	fmt.Println("Deleting HTML profile...")
	err = f5.DeleteHTMLProfile("test-html-profile")
	if err != nil {
		log.Printf("Error deleting HTML profile: %v", err)
		return
	}
	fmt.Println("HTML profile deleted successfully")

	// Verify deletion
	_, err = f5.GetHTMLProfile("test-html-profile")
	if err != nil && err.Error() == "01020036:3: The requested HTML Profile (/Common/test-html-profile) was not found." {
		fmt.Println("HTML profile deletion confirmed")
	}
}

// testAnalyticsProfile tests CRUD operations for Analytics profiles
func testAnalyticsProfile(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Analytics Profile CRUD Operations ===")

	// Test profile configuration
	analyticsProfile := &bigip.AnalyticsProfile{
		Name:                 "test-analytics-profile",
		DefaultsFrom:         "/Common/analytics",
		CollectGeo:           "enabled",
		CollectIp:            "enabled",
		CollectUrl:           "enabled",
		CollectMethods:       "enabled",
		CollectResponseCodes: "enabled",
		CollectUserAgent:     "disabled",
		NotificationByEmail:  "disabled",
		NotificationBySnmp:   "disabled",
		NotificationBySyslog: "disabled",
	}

	// CREATE: Add Analytics profile
	fmt.Println("Creating Analytics profile...")
	err := f5.AddAnalyticsProfile(analyticsProfile)
	if err != nil {
		log.Printf("Error creating Analytics profile: %v", err)
		return
	}
	fmt.Println("Analytics profile created successfully")

	// READ: Get all Analytics profiles
	fmt.Println("Getting all Analytics profiles...")
	profiles, err := f5.AnalyticsProfiles()
	if err != nil {
		log.Printf("Error getting Analytics profiles: %v", err)
		return
	}
	fmt.Printf("Found %d Analytics profiles\n", len(profiles.AnalyticsProfiles))

	// READ: Get specific Analytics profile
	fmt.Println("Getting specific Analytics profile...")
	profile, err := f5.GetAnalyticsProfile("test-analytics-profile")
	if err != nil {
		log.Printf("Error getting Analytics profile: %v", err)
		return
	}
	if profile != nil {
		fmt.Printf("Retrieved Analytics profile: %s with collect geo: %s\n", profile.Name, profile.CollectGeo)
	}

	// UPDATE: Modify Analytics profile
	fmt.Println("Updating Analytics profile...")
	analyticsProfile.CollectGeo = "disabled"
	analyticsProfile.CollectUserAgent = "enabled"
	err = f5.ModifyAnalyticsProfile("test-analytics-profile", analyticsProfile)
	if err != nil {
		log.Printf("Error updating Analytics profile: %v", err)
		return
	}
	fmt.Println("Analytics profile updated successfully")

	// Verify update
	updatedProfile, err := f5.GetAnalyticsProfile("test-analytics-profile")
	if err != nil {
		log.Printf("Error getting updated Analytics profile: %v", err)
		return
	}
	if updatedProfile != nil && updatedProfile.CollectGeo == "disabled" && updatedProfile.CollectUserAgent == "enabled" {
		fmt.Printf("Updated Analytics profile collect geo: %s\n", updatedProfile.CollectGeo)
	}

	// DELETE: Remove Analytics profile
	fmt.Println("Deleting Analytics profile...")
	err = f5.DeleteAnalyticsProfile("test-analytics-profile")
	if err != nil {
		log.Printf("Error deleting Analytics profile: %v", err)
		return
	}
	fmt.Println("Analytics profile deleted successfully")

	// Verify deletion
	_, err = f5.GetAnalyticsProfile("test-analytics-profile")
	if err != nil && err.Error() == "01020036:3: The requested Analytics Profile (/Common/test-analytics-profile) was not found." {
		fmt.Println("Analytics profile deletion confirmed")
	}
}

// testRequestAdaptProfile tests CRUD operations for Request Adapt profiles
func testRequestAdaptProfile(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Request Adapt Profile CRUD Operations ===")

	// Test profile configuration
	requestAdaptProfile := &bigip.RequestAdaptProfile{
		Name:              "test-request-adapt-profile",
		DefaultsFrom:      "/Common/requestadapt",
		Enabled:           "yes",
		PreviewSize:       1024,
		ServiceDownAction: "ignore",
		Timeout:           30,
	}

	// CREATE: Add Request Adapt profile
	fmt.Println("Creating Request Adapt profile...")
	err := f5.AddRequestAdaptProfile(requestAdaptProfile)
	if err != nil {
		log.Printf("Error creating Request Adapt profile: %v", err)
		return
	}
	fmt.Println("Request Adapt profile created successfully")

	// READ: Get all Request Adapt profiles
	fmt.Println("Getting all Request Adapt profiles...")
	profiles, err := f5.RequestAdaptProfiles()
	if err != nil {
		log.Printf("Error getting Request Adapt profiles: %v", err)
		return
	}
	fmt.Printf("Found %d Request Adapt profiles\n", len(profiles.RequestAdaptProfiles))

	// READ: Get specific Request Adapt profile
	fmt.Println("Getting specific Request Adapt profile...")
	profile, err := f5.GetRequestAdaptProfile("test-request-adapt-profile")
	if err != nil {
		log.Printf("Error getting Request Adapt profile: %v", err)
		return
	}
	if profile != nil {
		fmt.Printf("Retrieved Request Adapt profile: %s with enabled: %s\n", profile.Name, profile.Enabled)
	}

	// UPDATE: Modify Request Adapt profile
	fmt.Println("Updating Request Adapt profile...")
	requestAdaptProfile.PreviewSize = 2048
	requestAdaptProfile.Timeout = 60
	err = f5.ModifyRequestAdaptProfile("test-request-adapt-profile", requestAdaptProfile)
	if err != nil {
		log.Printf("Error updating Request Adapt profile: %v", err)
		return
	}
	fmt.Println("Request Adapt profile updated successfully")

	// Verify update
	updatedProfile, err := f5.GetRequestAdaptProfile("test-request-adapt-profile")
	if err != nil {
		log.Printf("Error getting updated Request Adapt profile: %v", err)
		return
	}
	if updatedProfile != nil && updatedProfile.PreviewSize == 2048 && updatedProfile.Timeout == 60 {
		fmt.Printf("Updated Request Adapt profile preview size: %d\n", updatedProfile.PreviewSize)
	}

	// DELETE: Remove Request Adapt profile
	fmt.Println("Deleting Request Adapt profile...")
	err = f5.DeleteRequestAdaptProfile("test-request-adapt-profile")
	if err != nil {
		log.Printf("Error deleting Request Adapt profile: %v", err)
		return
	}
	fmt.Println("Request Adapt profile deleted successfully")

	// Verify deletion
	_, err = f5.GetRequestAdaptProfile("test-request-adapt-profile")
	if err != nil && err.Error() == "01020036:3: The requested Request Adapt Profile (/Common/test-request-adapt-profile) was not found." {
		fmt.Println("Request Adapt profile deletion confirmed")
	}
}

// testResponseAdaptProfile tests CRUD operations for Response Adapt profiles
func testResponseAdaptProfile(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Response Adapt Profile CRUD Operations ===")

	// Test profile configuration
	responseAdaptProfile := &bigip.ResponseAdaptProfile{
		Name:              "test-response-adapt-profile",
		DefaultsFrom:      "/Common/responseadapt",
		Enabled:           "yes",
		PreviewSize:       1024,
		ServiceDownAction: "ignore",
		Timeout:           30,
	}

	// CREATE: Add Response Adapt profile
	fmt.Println("Creating Response Adapt profile...")
	err := f5.AddResponseAdaptProfile(responseAdaptProfile)
	if err != nil {
		log.Printf("Error creating Response Adapt profile: %v", err)
		return
	}
	fmt.Println("Response Adapt profile created successfully")

	// READ: Get all Response Adapt profiles
	fmt.Println("Getting all Response Adapt profiles...")
	profiles, err := f5.ResponseAdaptProfiles()
	if err != nil {
		log.Printf("Error getting Response Adapt profiles: %v", err)
		return
	}
	fmt.Printf("Found %d Response Adapt profiles\n", len(profiles.ResponseAdaptProfiles))

	// READ: Get specific Response Adapt profile
	fmt.Println("Getting specific Response Adapt profile...")
	profile, err := f5.GetResponseAdaptProfile("test-response-adapt-profile")
	if err != nil {
		log.Printf("Error getting Response Adapt profile: %v", err)
		return
	}
	if profile != nil {
		fmt.Printf("Retrieved Response Adapt profile: %s with enabled: %s\n", profile.Name, profile.Enabled)
	}

	// UPDATE: Modify Response Adapt profile
	fmt.Println("Updating Response Adapt profile...")
	responseAdaptProfile.PreviewSize = 2048
	responseAdaptProfile.Timeout = 60
	err = f5.ModifyResponseAdaptProfile("test-response-adapt-profile", responseAdaptProfile)
	if err != nil {
		log.Printf("Error updating Response Adapt profile: %v", err)
		return
	}
	fmt.Println("Response Adapt profile updated successfully")

	// Verify update
	updatedProfile, err := f5.GetResponseAdaptProfile("test-response-adapt-profile")
	if err != nil {
		log.Printf("Error getting updated Response Adapt profile: %v", err)
		return
	}
	if updatedProfile != nil && updatedProfile.PreviewSize == 2048 && updatedProfile.Timeout == 60 {
		fmt.Printf("Updated Response Adapt profile preview size: %d\n", updatedProfile.PreviewSize)
	}

	// DELETE: Remove Response Adapt profile
	fmt.Println("Deleting Response Adapt profile...")
	err = f5.DeleteResponseAdaptProfile("test-response-adapt-profile")
	if err != nil {
		log.Printf("Error deleting Response Adapt profile: %v", err)
		return
	}
	fmt.Println("Response Adapt profile deleted successfully")

	// Verify deletion
	_, err = f5.GetResponseAdaptProfile("test-response-adapt-profile")
	if err != nil && err.Error() == "01020036:3: The requested Response Adapt Profile (/Common/test-response-adapt-profile) was not found." {
		fmt.Println("Response Adapt profile deletion confirmed")
	}
}
