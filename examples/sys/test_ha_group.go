package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////       HA Group Testing                           /////////////////
/////////////////////////////////////////////////////////////////////////////////

func testHaGroupRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HA Group Read Operation ===")

	fmt.Println("Getting all HA Groups...")
	haGroups, err := f5.GetHaGroups()
	if err != nil {
		log.Printf("Error getting HA Groups: %v", err)
		return
	}
	fmt.Printf("Found %d HA Groups\n", len(haGroups.HaGroups))

	fmt.Println("\nGetting specific HA Groups...")
	for i := 0; i < len(haGroups.HaGroups); i++ {
		haGroup, err := f5.GetHaGroup(haGroups.HaGroups[i].Name)
		if err != nil {
			log.Printf("Error getting HA Group: %v", err)
			continue
		}
		if haGroup != nil {
			fmt.Printf("Retrieved HA Group:\n")
			fmt.Printf("  Name: %s\n", haGroup.Name)
			fmt.Printf("  Active Bonus: %d\n", haGroup.ActiveBonus)
			fmt.Printf("  Enabled: %t\n", haGroup.Enabled)
			fmt.Printf("  Description: %s\n", haGroup.Description)
			fmt.Printf("  Pools: %d\n", len(haGroup.Pools))
			fmt.Printf("  Clusters: %d\n", len(haGroup.Clusters))
			fmt.Printf("  Trunks: %d\n", len(haGroup.Trunks))
			if len(haGroup.Pools) > 0 {
				for _, pool := range haGroup.Pools {
					fmt.Printf("    Pool: %s (Weight: %d, MinThreshold: %d, SufficientThreshold: %s)\n",
						pool.Name, pool.Weight, pool.MinimumThreshold, pool.SufficientThreshold)
				}
			}
			fmt.Println()
		}
	}
}

func testHaGroupCreation(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HA Group Creation Operation ===")

	haGroup := &bigip.HaGroup{
		Name:        "test_ha_group",
		ActiveBonus: 10,
		Description: "Test HA Group created via API",
		Pools: []bigip.HaGroupPool{
			{
				Name:                "/Common/Pool_1",
				Attribute:           "percent-up-members",
				MinimumThreshold:    1,
				SufficientThreshold: "all",
				Weight:              10,
			},
		},
	}

	fmt.Println("Creating HA Group...")
	err := f5.CreateHaGroup(haGroup)
	if err != nil {
		log.Printf("Error creating HA Group: %v", err)
		return
	}
	fmt.Printf("HA Group %s created successfully.\n", haGroup.Name)
}

func testHaGroupUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HA Group Update Operation ===")

	haGroup := &bigip.HaGroup{
		Name:        "test_ha_group",
		ActiveBonus: 20,
		Description: "Updated test HA Group - changed active bonus",
		Pools: []bigip.HaGroupPool{
			{
				Name:                "/Common/Pool_1",
				Attribute:           "percent-up-members",
				MinimumThreshold:    1,
				SufficientThreshold: "all",
				Weight:              15,
			},
			{
				Name:                "/Common/devdos_pool",
				Attribute:           "percent-up-members",
				MinimumThreshold:    1,
				SufficientThreshold: "all",
				Weight:              20,
			},
		},
	}

	fmt.Println("Updating HA Group...")
	err := f5.ModifyHaGroup(haGroup.Name, haGroup)
	if err != nil {
		log.Printf("Error updating HA Group: %v", err)
		return
	}
	fmt.Printf("HA Group %s updated successfully.\n", haGroup.Name)
}

func testHaGroupDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing HA Group Delete Operation ===")

	fmt.Println("Deleting HA Group test_ha_group...")
	err := f5.DeleteHaGroup("test_ha_group")
	if err != nil {
		log.Printf("Error deleting HA Group: %v", err)
		return
	}
	fmt.Printf("HA Group test_ha_group deleted successfully.\n")
}
