package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////          SNMP Users Testing                      /////////////////
/////////////////////////////////////////////////////////////////////////////////

func testSnmpUsersRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Users Read Operation ===")

	fmt.Println("Getting all SNMP Users...")
	snmpUsers, err := f5.GetSnmpUsers()
	if err != nil {
		log.Printf("Error getting SNMP Users: %v", err)
		return
	}
	fmt.Printf("Found %d SNMP Users\n", len(snmpUsers.SnmpUsers))

	fmt.Println("\nGetting specific SNMP Users...")
	for i := 0; i < len(snmpUsers.SnmpUsers); i++ {
		snmpUser, err := f5.GetSnmpUser(snmpUsers.SnmpUsers[i].Name)
		if err != nil {
			log.Printf("Error getting SNMP User: %v", err)
			continue
		}
		if snmpUser != nil {
			fmt.Printf("Retrieved SNMP User:\n")
			fmt.Printf("  Name: %s\n", snmpUser.Name)
			fmt.Printf("  Username: %s\n", snmpUser.Username)
			fmt.Printf("  Access: %s\n", snmpUser.Access)
			fmt.Printf("  Auth Protocol: %s\n", snmpUser.AuthProtocol)
			fmt.Printf("  Privacy Protocol: %s\n", snmpUser.PrivacyProtocol)
			fmt.Printf("  Security Level: %s\n", snmpUser.SecurityLevel)
			fmt.Printf("  OID Subset: %s\n", snmpUser.OidSubset)
			fmt.Printf("  Description: %s\n", snmpUser.Description)
			fmt.Println()
		}
	}
}

func testSnmpUserCreation(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP User Creation Operation ===")

	snmpUser := &bigip.SnmpUser{
		Name:            "test_snmp_user",
		Username:        "testuser",
		Access:          "ro",
		AuthPassword:    "testpassword123",
		AuthProtocol:    "sha",
		PrivacyPassword: "testpassword123",
		PrivacyProtocol: "aes",
		SecurityLevel:   "auth-privacy",
		OidSubset:       ".1.3.6.1.4.1.3375.2.1.14.3.2",
		Description:     "Test SNMPv3 user created via API",
	}

	fmt.Println("Creating SNMP User...")
	err := f5.CreateSnmpUser(snmpUser)
	if err != nil {
		log.Printf("Error creating SNMP User: %v", err)
		return
	}
	fmt.Printf("SNMP User %s created successfully.\n", snmpUser.Name)
}

func testSnmpUserUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP User Update Operation ===")

	snmpUser := &bigip.SnmpUser{
		Name:            "test_snmp_user",
		Username:        "testuser_updated",
		Access:          "rw",
		AuthPassword:    "newpassword456",
		AuthProtocol:    "sha",
		PrivacyPassword: "newpassword456",
		PrivacyProtocol: "aes",
		SecurityLevel:   "auth-privacy",
		OidSubset:       ".1.3.6.1.4.1.3375.2.1.14",
		Description:     "Updated test SNMPv3 user - changed to read-write",
	}

	fmt.Println("Updating SNMP User...")
	err := f5.ModifySnmpUser(snmpUser.Name, snmpUser)
	if err != nil {
		log.Printf("Error updating SNMP User: %v", err)
		return
	}
	fmt.Printf("SNMP User %s updated successfully.\n", snmpUser.Name)
}

func testSnmpUserDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP User Delete Operation ===")

	fmt.Println("Deleting SNMP User test_snmp_user...")
	err := f5.DeleteSnmpUser("test_snmp_user")
	if err != nil {
		log.Printf("Error deleting SNMP User: %v", err)
		return
	}
	fmt.Printf("SNMP User test_snmp_user deleted successfully.\n")
}
