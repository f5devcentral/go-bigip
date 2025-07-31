package main

import (
	"fmt"
	"github.com/f5devcentral/go-bigip"
	"log"
)

// test access-profile
func testAccessProfiles(f5 *bigip.BigIP) {
	// Example 1: List all existing Access Profiles
	fmt.Println("1. Listing all existing Access Profiles...")
	listAccessProfiles(f5)
}

// listAccessProfiles demonstrates how to retrieve all access profiles
func listAccessProfiles(f5 *bigip.BigIP) {
	profiles, err := f5.AccessProfiles()
	if err != nil {
		log.Printf("Error retrieving access profiles: %v", err)
		return
	}

	if len(profiles.AccessProfiles) == 0 {
		fmt.Println("No access profiles found.")
		return
	}

	fmt.Printf("Found %d access profile(s):\n", len(profiles.AccessProfiles))
	for i, profile := range profiles.AccessProfiles {
		fmt.Printf("  %d. Name: %s, Partition: %s, Type: %s\n",
			i+1, profile.Name, profile.Partition, profile.Type)
		if profile.Description != "" {
			fmt.Printf("     Description: %s\n", profile.Description)
		}
		fmt.Printf("     Full Path: %s\n", profile.FullPath)
	}
}
