package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////              SSHD Configuration Testing         //////////////////
/////////////////////////////////////////////////////////////////////////////////

func testSSHDRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SSHD Configuration Read Operation ===")

	fmt.Println("Getting SSHD configuration...")
	sshdConfig, err := f5.GetSSHDConfig()
	if err != nil {
		log.Printf("Error getting SSHD configuration: %v", err)
		return
	}
	if sshdConfig != nil {
		fmt.Printf("Retrieved SSHD Configuration:\n")
		fmt.Printf("  Port: %d\n", sshdConfig.Port)
		fmt.Printf("  Login: %s\n", sshdConfig.Login)
		fmt.Printf("  LogLevel: %s\n", sshdConfig.LogLevel)
		fmt.Printf("  Banner: %s\n", sshdConfig.Banner)
		fmt.Printf("  BannerText: %s\n", sshdConfig.BannerText)
		fmt.Printf("  InactivityTimeout: %d\n", sshdConfig.InactivityTimeout)
		fmt.Printf("  Allow: %v\n", sshdConfig.Allow)
	}
}

func testSSHDCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SSHD Configuration Create Operation ===")

	sshdConfig := &bigip.SSHDConfig{
		Banner:     "enabled",
		BannerText: "AUTHORIZED ACCESS ONLY\nUnauthorized access is prohibited and logged.",
		Allow:      []string{"10.1.10.0/24", "192.168.1.0/24"},
	}

	fmt.Println("Setting SSHD configuration...")
	err := f5.CreateSSHDConfig(sshdConfig)
	if err != nil {
		log.Printf("Error creating SSHD configuration: %v", err)
		return
	}
	fmt.Printf("SSHD configuration set successfully:\n")
	fmt.Printf("BannerText: %s\n Allow: %s\n ",
		sshdConfig.BannerText, sshdConfig.Allow)
}

func testSSHDUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SSHD Configuration Update Operation ===")

	sshdConfig := &bigip.SSHDConfig{
		BannerText: "Updated Banner",
		Allow:      []string{"172.18.20.0/24"},
	}

	fmt.Println("Updating SSHD configuration...")
	err := f5.ModifySSHDConfig(sshdConfig)
	if err != nil {
		log.Printf("Error updating SSHD configuration: %v", err)
		return
	}
	fmt.Printf("SSHD configuration updated successfully:\n")
	fmt.Printf("BannerText: %s\n Allow: %s\n ",
		sshdConfig.BannerText, sshdConfig.Allow)
}

func testSSHDDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SSHD Configuration Delete Operation ===")

	fmt.Println("Resetting SSHD configuration to defaults...")
	err := f5.DeleteSSHDConfig()
	if err != nil {
		log.Printf("Error deleting SSHD configuration: %v", err)
		return
	}
	fmt.Printf("SSHD configuration reset to defaults.\n")
}
