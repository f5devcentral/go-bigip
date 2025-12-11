package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

////////////////////////////////////////////////////////////////////////
//////////////          Management Route Testing          //////////////
////////////////////////////////////////////////////////////////////////

func testManagementRouteRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Management Route Read Operation ===")

	fmt.Println("Getting all Management Routes...")
	mgmtroutes, err := f5.GetManagementRoutes()
	if err != nil {
		log.Printf("Error getting Management Routes: %v", err)
		return
	}
	fmt.Printf("Found %d Management Routes\n", len(mgmtroutes.ManagementRoutes))

	fmt.Println("Getting specific Management Route...")
	for i := 0; i < len(mgmtroutes.ManagementRoutes); i++ {
		mgmtroute, err := f5.GetManagementRoute(mgmtroutes.ManagementRoutes[i].Name)
		if err != nil {
			log.Printf("Error getting Management Route: %v", err)
			return
		}
		if mgmtroute != nil {
			fmt.Printf("Retrieved Management Route:\n%s - %s via %s   (MTU: %d)\n", mgmtroute.Name, mgmtroute.Network, mgmtroute.Gateway, mgmtroute.MTU)
		}
	}

}

func testManagementRouteCreation(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Management Route Creation Operation ===")

	mgmtRoute := &bigip.ManagementRoute{
		Name:    "/Common/test-route",
		Gateway: "10.171.125.61",
		//MTU:     1500,
		Network: "32.50.33.0/24",
	}

	fmt.Println("Creating Management Route...")
	err := f5.CreateManagementRoute(mgmtRoute)
	if err != nil {
		log.Printf("Error creating Management Route: %v", err)
		return
	}
	fmt.Printf("Management Route %s created successfully.\n", mgmtRoute.Name)

}

func testManagementRouteUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Management Route Update Operation ===")

	mgmtRoute := &bigip.ManagementRoute{
		Name:    "/Common/test-route",
		Gateway: "10.171.125.45",
		//MTU:     1500,
		Network: "32.50.33.0/24",
	}

	fmt.Println("Updating Management Route...")
	err := f5.ModifyManagementRoute(mgmtRoute.Name, mgmtRoute)
	if err != nil {
		log.Printf("Error updating Management Route: %v", err)
		return
	}
	fmt.Printf("Management Route %s updated successfully.\n", mgmtRoute.Name)

}

func testManagementRouteDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Management Route Delete Operation ===")

	mgmtRoute := &bigip.ManagementRoute{
		Name: "/Common/test-route",
	}

	fmt.Println("Deleting Management Route...")
	err := f5.DeleteManagementRoute(mgmtRoute.Name)
	if err != nil {
		log.Printf("Error deleting Management Route: %v", err)
		return
	}
	fmt.Printf("Management Route %s deleted successfully.\n", mgmtRoute.Name)
}
