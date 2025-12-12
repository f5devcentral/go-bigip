package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////          Remote Role Testing                     /////////////////
/////////////////////////////////////////////////////////////////////////////////

func testRemoteRoleRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Remote Role Read Operation ===")

	fmt.Println("Getting all Remote Roles...")
	remoteRoles, err := f5.GetRemoteRoles()
	if err != nil {
		log.Printf("Error getting Remote Roles: %v", err)
		return
	}
	fmt.Printf("Found %d Remote Roles\n", len(remoteRoles.RemoteRoles))

	fmt.Println("\nGetting specific Remote Roles...")
	for i := 0; i < len(remoteRoles.RemoteRoles); i++ {
		remoteRole, err := f5.GetRemoteRole(remoteRoles.RemoteRoles[i].Name)
		if err != nil {
			log.Printf("Error getting Remote Role: %v", err)
			continue
		}
		if remoteRole != nil {
			fmt.Printf("Retrieved Remote Role:\n")
			fmt.Printf("  Name: %s\n", remoteRole.Name)
			fmt.Printf("  Attribute: %s\n", remoteRole.Attribute)
			fmt.Printf("  Role: %s\n", remoteRole.Role)
			fmt.Printf("  Console: %s\n", remoteRole.Console)
			fmt.Printf("  User Partition: %s\n", remoteRole.UserPartition)
			fmt.Printf("  Line Order: %d\n", remoteRole.LineOrder)
			fmt.Printf("  Deny: %s\n", remoteRole.Deny)
			fmt.Printf("  Description: %s\n", remoteRole.Description)
			fmt.Println()
		}
	}
}

func testRemoteRoleCreation(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Remote Role Creation Operation ===")

	remoteRole := &bigip.RemoteRole{
		Name:          "lb_admins",
		Attribute:     "memberof=cn=lb_admins,ou=groups,dc=test,dc=com",
		Console:       "tmsh",
		Deny:          "disabled",
		LineOrder:     1000,
		Role:          "administrator",
		UserPartition: "All",
		Description:   "Test admin group created via API",
	}

	fmt.Println("Creating Remote Role...")
	err := f5.CreateRemoteRole(remoteRole)
	if err != nil {
		log.Printf("Error creating Remote Role: %v", err)
		return
	}
	fmt.Printf("Remote Role %s created successfully.\n", remoteRole.Name)
}

func testRemoteRoleUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Remote Role Update Operation ===")

	remoteRole := &bigip.RemoteRole{
		Name:          "lb_admins",
		Attribute:     "memberof=cn=lb_admins,ou=groups,dc=test,dc=com",
		Console:       "disabled",
		Deny:          "disabled",
		LineOrder:     1000,
		Role:          "guest",
		UserPartition: "Common",
		Description:   "Updated test admin group - changed to guest role",
	}

	fmt.Println("Updating Remote Role...")
	err := f5.ModifyRemoteRole(remoteRole.Name, remoteRole)
	if err != nil {
		log.Printf("Error updating Remote Role: %v", err)
		return
	}
	fmt.Printf("Remote Role %s updated successfully.\n", remoteRole.Name)
}

func testRemoteRoleDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Remote Role Delete Operation ===")

	fmt.Println("Deleting Remote Role test_admins...")
	err := f5.DeleteRemoteRole("lb_admins")
	if err != nil {
		log.Printf("Error deleting Remote Role: %v", err)
		return
	}
	fmt.Printf("Remote Role lb_admins deleted successfully.\n")
}
