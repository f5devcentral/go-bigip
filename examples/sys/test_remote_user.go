package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////            Remote User Testing               /////////////////////
/////////////////////////////////////////////////////////////////////////////////

func testRemoteUserRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Remote User Read Operation ===")

	fmt.Println("Getting Remote User configuration...")
	remoteUser, err := f5.GetRemoteUser()
	if err != nil {
		log.Printf("Error getting Remote User: %v", err)
		return
	}
	if remoteUser != nil {
		fmt.Printf("Retrieved Remote User:\n")
		fmt.Printf("  Default Partition: %s\n", remoteUser.DefaultPartition)
		fmt.Printf("  Default Role: %s\n", remoteUser.DefaultRole)
		fmt.Printf("  Remote Console Access: %s\n", remoteUser.RemoteConsoleAccess)
	}
}

func testRemoteUserCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Remote User Create Operation ===")

	remoteUser := &bigip.RemoteUser{
		DefaultPartition:    "Common",
		DefaultRole:         "guest",
		RemoteConsoleAccess: "tmsh",
	}

	fmt.Println("Setting Remote User configuration...")
	err := f5.CreateRemoteUser(remoteUser)
	if err != nil {
		log.Printf("Error creating Remote User: %v", err)
		return
	}
	fmt.Printf("Remote User set successfully:\n  DefaultPartition: %s\n  DefaultRole: %s\n  RemoteConsoleAccess: %s\n",
		remoteUser.DefaultPartition, remoteUser.DefaultRole, remoteUser.RemoteConsoleAccess)
}

func testRemoteUserUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Remote User Update Operation ===")

	remoteUser := &bigip.RemoteUser{
		DefaultPartition:    "all",
		DefaultRole:         "operator",
		RemoteConsoleAccess: "disabled",
	}

	fmt.Println("Updating Remote User...")
	err := f5.ModifyRemoteUser(remoteUser)
	if err != nil {
		log.Printf("Error updating Remote User: %v", err)
		return
	}
	fmt.Printf("Remote User updated successfully:\n  DefaultPartition: %s\n  DefaultRole: %s\n  RemoteConsoleAccess: %s\n",
		remoteUser.DefaultPartition, remoteUser.DefaultRole, remoteUser.RemoteConsoleAccess)
}

func testRemoteUserDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Remote User Delete Operation ===")

	fmt.Println("Resetting Remote User to defaults...")
	err := f5.DeleteRemoteUser()
	if err != nil {
		log.Printf("Error deleting Remote User: %v", err)
		return
	}
	fmt.Printf("Remote User reset to defaults:\n  DefaultPartition: all\n  DefaultRole: no-access\n  RemoteConsoleAccess: disabled\n")
}
