package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////       SNMP Communities Testing                   /////////////////
/////////////////////////////////////////////////////////////////////////////////

func testSnmpCommunitiesRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Communities Read Operation ===")

	fmt.Println("Getting all SNMP Communities...")
	snmpCommunities, err := f5.GetSnmpCommunities()
	if err != nil {
		log.Printf("Error getting SNMP Communities: %v", err)
		return
	}
	fmt.Printf("Found %d SNMP Communities\n", len(snmpCommunities.SnmpCommunities))

	fmt.Println("\nGetting specific SNMP Communities...")
	for i := 0; i < len(snmpCommunities.SnmpCommunities); i++ {
		snmpCommunity, err := f5.GetSnmpCommunity(snmpCommunities.SnmpCommunities[i].Name)
		if err != nil {
			log.Printf("Error getting SNMP Community: %v", err)
			continue
		}
		if snmpCommunity != nil {
			fmt.Printf("Retrieved SNMP Community:\n")
			fmt.Printf("  Name: %s\n", snmpCommunity.Name)
			fmt.Printf("  Community Name: %s\n", snmpCommunity.CommunityName)
			fmt.Printf("  Access: %s\n", snmpCommunity.Access)
			fmt.Printf("  IPv6: %s\n", snmpCommunity.Ipv6)
			fmt.Printf("  OID Subset: %s\n", snmpCommunity.OidSubset)
			fmt.Printf("  Source: %s\n", snmpCommunity.Source)
			fmt.Printf("  Description: %s\n", snmpCommunity.Description)
			fmt.Println()
		}
	}
}

func testSnmpCommunityCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Community Creation Operation ===")

	snmpCommunity := &bigip.SnmpCommunity{
		Name:          "test_community",
		CommunityName: "test_public",
		Access:        "ro",
		Ipv6:          "disabled",
		OidSubset:     ".1.3.6.1.4.1.3375.2.1.14.3.2",
		Source:        "default",
		Description:   "Test SNMP community created via API",
	}

	fmt.Println("Creating SNMP Community...")
	err := f5.CreateSnmpCommunity(snmpCommunity)
	if err != nil {
		log.Printf("Error creating SNMP Community: %v", err)
		return
	}
	fmt.Printf("SNMP Community %s created successfully.\n", snmpCommunity.Name)
}

func testSnmpCommunityUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Community Update Operation ===")

	snmpCommunity := &bigip.SnmpCommunity{
		Name:          "test_community",
		CommunityName: "test_public_updated",
		Access:        "rw",
		OidSubset:     ".1.3.6.1",
		Source:        "all",
		Description:   "Updated test SNMP community",
	}

	fmt.Println("Updating SNMP Community...")
	err := f5.ModifySnmpCommunity(snmpCommunity.Name, snmpCommunity)
	if err != nil {
		log.Printf("Error updating SNMP Community: %v", err)
		return
	}
	fmt.Printf("SNMP Community %s updated successfully.\n", snmpCommunity.Name)
}

func testSnmpCommunityDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Community Delete Operation ===")

	fmt.Println("Deleting SNMP Community test_community...")
	err := f5.DeleteSnmpCommunity("test_community")
	if err != nil {
		log.Printf("Error deleting SNMP Community: %v", err)
		return
	}
	fmt.Printf("SNMP Community test_community deleted successfully.\n")
}
