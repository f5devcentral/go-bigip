package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////          Authentication Source Testing       /////////////////////
/////////////////////////////////////////////////////////////////////////////////

func testAuthSourceRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Authentication Source Read Operation ===")

	fmt.Println("Getting Authentication Source configuration...")
	authSource, err := f5.GetAuthSource()
	if err != nil {
		log.Printf("Error getting Authentication Source: %v", err)
		return
	}
	if authSource != nil {
		fmt.Printf("Retrieved Authentication Source:\n")
		fmt.Printf("  Type: %s\n", authSource.Type)
		fmt.Printf("  Fallback: %s\n", authSource.Fallback)
	}
}

func testAuthSourceCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Authentication Source Create Operation ===")

	authSource := &bigip.AuthSource{
		Type:     "ldap",
		Fallback: "false",
	}

	fmt.Println("Setting Authentication Source to LDAP...")
	err := f5.CreateAuthSource(authSource)
	if err != nil {
		log.Printf("Error creating Authentication Source: %v", err)
		return
	}
	fmt.Printf("Authentication Source set successfully:\n  Type: %s\n  Fallback: %s\n",
		authSource.Type, authSource.Fallback)
}

func testAuthSourceUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Authentication Source Update Operation ===")

	authSource := &bigip.AuthSource{
		Type:     "radius",
		Fallback: "false",
	}

	fmt.Println("Updating Authentication Source...")
	err := f5.ModifyAuthSource(authSource)
	if err != nil {
		log.Printf("Error updating Authentication Source: %v", err)
		return
	}
	fmt.Printf("Authentication Source updated successfully:\n  Type: %s\n  Fallback: %s\n",
		authSource.Type, authSource.Fallback)
}

func testAuthSourceDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Authentication Source Delete Operation ===")

	fmt.Println("Resetting Authentication Source to defaults...")
	err := f5.DeleteAuthSource()
	if err != nil {
		log.Printf("Error deleting Authentication Source: %v", err)
		return
	}
	fmt.Printf("Authentication Source reset to defaults:\n  Type: local\n  Fallback: false\n")
}
