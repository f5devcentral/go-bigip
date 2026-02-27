package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////          Route Domain Extended Testing          //////////////////
/////////////////////////////////////////////////////////////////////////////////

func testRouteDomainRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Route Domain Read Operation ===")

	// READ: Get all Route Domains
	fmt.Println("Getting all Route Domains...")
	routedomains, err := f5.GetRouteDomains()
	if err != nil {
		log.Printf("Error getting Route Domains: %v", err)
		return
	}
	fmt.Printf("Found %d Route Domains\n", len(routedomains.RouteDomains))

	// READ: Get specific Route Domain
	fmt.Println("\nGetting specific Route Domains...")
	for i := 0; i < len(routedomains.RouteDomains); i++ {
		routedomain, err := f5.GetRouteDomain(routedomains.RouteDomains[i].Name)
		if err != nil {
			log.Printf("Error getting Route Domain: %v", err)
			continue
		}
		if routedomain != nil {
			fmt.Printf("Retrieved Route Domain:\n")
			fmt.Printf("  Name: %s\n", routedomain.Name)
			fmt.Printf("  ID: %d\n", routedomain.ID)
			fmt.Printf("  Strict: %s\n", routedomain.Strict)
			fmt.Printf("  Parent: %s\n", routedomain.Parent)
			fmt.Printf("  VLANs: %v\n", routedomain.Vlans)
			fmt.Printf("  Description: %s\n", routedomain.Description)
			fmt.Printf("  Connection Limit: %d\n", routedomain.ConnectionLimit)
			fmt.Println()
		}
	}
}

func testRouteDomainCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Route Domain Creation Operation ===")

	routeDomain := &bigip.RouteDomain{
		Name:        "rd_test",
		ID:          300,
		Strict:      "enabled",
		Description: "Test route domain created via API",
		//ConnectionLimit: 0,
		//Vlans:           []string{"/Common/vlan100", "/Common/vlan101"},
		//RoutingProtocol: []string{},
		//Parent:          "",
	}

	fmt.Println("Creating Route Domain...")
	err := f5.CreateRouteDomain(routeDomain)
	if err != nil {
		log.Printf("Error creating Route Domain: %v", err)
		return
	}
	fmt.Printf("Route Domain %s (ID: %d) created successfully.\n", routeDomain.Name, routeDomain.ID)
}

func testRouteDomainUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Route Domain Update Operation ===")

	routeDomain := &bigip.RouteDomain{
		Name: "rd_test",
		//ID:              100,
		Strict:          "disabled",
		Description:     "Updated test route domain - strict mode disabled",
		ConnectionLimit: 5000,
		//Vlans:           []string{"/Common/vlan100", "/Common/vlan101", "/Common/vlan102"},
	}

	// MODIFY: Update Route Domain
	fmt.Println("Updating Route Domain...")
	err := f5.ModifyRouteDomain(routeDomain.Name, routeDomain)
	if err != nil {
		log.Printf("Error updating Route Domain: %v", err)
		return
	}
	fmt.Printf("Route Domain %s updated successfully.\n", routeDomain.Name)
}

func testRouteDomainDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Route Domain Delete Operation ===")

	// DELETE: Remove Route Domain
	fmt.Println("Deleting Route Domain rd_test_100...")
	err := f5.DeleteRouteDomain("rd_test")
	if err != nil {
		log.Printf("Error deleting Route Domain: %v", err)
		return
	}
	fmt.Printf("Route Domain rd_test deleted successfully.\n")
}
