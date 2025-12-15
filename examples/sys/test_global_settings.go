package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////         Global Settings Testing             /////////////////////
/////////////////////////////////////////////////////////////////////////////////

func testGlobalSettingsRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Global Settings Read Operation ===")

	fmt.Println("Getting Global Settings configuration...")
	globalSettings, err := f5.GetGlobalSettings()
	if err != nil {
		log.Printf("Error getting Global Settings: %v", err)
		return
	}
	if globalSettings != nil {
		fmt.Printf("Retrieved Global Settings:\n")
		fmt.Printf("  Hostname: %s\n", globalSettings.Hostname)
		fmt.Printf("  GUI Security Banner: %s\n", globalSettings.GuiSecurityBanner)
		fmt.Printf("  GUI Security Banner Text: %s\n", globalSettings.GuiSecurityBannerText)
	}
}

func testGlobalSettingsCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Global Settings Create Operation ===")

	globalSettings := &bigip.GlobalSettings{
		Hostname:              "ltm02.f5lab.com",
		GuiSecurityBanner:     "enabled",
		GuiSecurityBannerText: "AUTHORIZED ACCESS ONLY\nUnauthorized access is prohibited.",
	}

	fmt.Println("Setting Global Settings configuration...")
	err := f5.CreateGlobalSettings(globalSettings)
	if err != nil {
		log.Printf("Error creating Global Settings: %v", err)
		return
	}
	fmt.Printf("Global Settings set successfully:\n")
	fmt.Printf("  Hostname: %s\n  GUI Security Banner: %s\n",
		globalSettings.Hostname, globalSettings.GuiSecurityBanner)
}

func testGlobalSettingsUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Global Settings Update Operation ===")

	globalSettings := &bigip.GlobalSettings{
		Hostname:              "mybigip02.lab.local",
		GuiSecurityBanner:     "disabled",
		GuiSecurityBannerText: "Disabled banner",
	}

	fmt.Println("Updating Global Settings...")
	err := f5.ModifyGlobalSettings(globalSettings)
	if err != nil {
		log.Printf("Error updating Global Settings: %v", err)
		return
	}
	fmt.Printf("Global Settings updated successfully:\n")
	fmt.Printf("  Hostname: %s\n  GUI Security Banner: %s\n",
		globalSettings.Hostname, globalSettings.GuiSecurityBanner)
}

func testGlobalSettingsDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Global Settings Delete Operation ===")

	fmt.Println("Resetting Global Settings to defaults...")
	err := f5.DeleteGlobalSettings()
	if err != nil {
		log.Printf("Error deleting Global Settings: %v", err)
		return
	}
	fmt.Printf("Global Settings reset to defaults.\n")
}
