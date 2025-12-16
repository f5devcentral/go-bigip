package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////       SNMP Config Testing                        /////////////////
/////////////////////////////////////////////////////////////////////////////////

func testSnmpConfigRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Config Read Operation ===")

	fmt.Println("Getting SNMP Config...")
	snmpConfig, err := f5.GetSnmpConfig()
	if err != nil {
		log.Printf("Error getting SNMP Config: %v", err)
		return
	}

	if snmpConfig != nil {
		fmt.Printf("Retrieved SNMP Config:\n")
		fmt.Printf("  System Contact: %s\n", snmpConfig.SysContact)
		fmt.Printf("  System Location: %s\n", snmpConfig.SysLocation)
		fmt.Printf("  Allowed Addresses: %v\n", snmpConfig.AllowedAddresses)
		fmt.Println()
	}
}

func testSnmpConfigCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Config Create Operation ===")

	snmpConfig := &bigip.SnmpConfig{
		SysContact:       "Network Admin <netadmin@example.com>",
		SysLocation:      "Data Center 1",
		AllowedAddresses: []string{"10.0.0.0/8", "172.18.0.0/16"},
	}

	fmt.Println("Creating SNMP Config...")
	err := f5.CreateSnmpConfig(snmpConfig)
	if err != nil {
		log.Printf("Error creating SNMP Config: %v", err)
		return
	}
	fmt.Printf("SNMP Config created successfully.\n")
}

func testSnmpConfigUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Config Update Operation ===")

	snmpConfig := &bigip.SnmpConfig{
		SysContact:       "Updated Admin <updated@example.com>",
		SysLocation:      "Data Center 2",
		AllowedAddresses: []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"},
	}

	fmt.Println("Updating SNMP Config...")
	err := f5.ModifySnmpConfig(snmpConfig)
	if err != nil {
		log.Printf("Error updating SNMP Config: %v", err)
		return
	}
	fmt.Printf("SNMP Config updated successfully.\n")
}

func testSnmpConfigDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing SNMP Config Delete Operation ===")

	fmt.Println("Deleting (resetting) SNMP Config...")
	err := f5.DeleteSnmpConfig()
	if err != nil {
		log.Printf("Error deleting SNMP Config: %v", err)
		return
	}
	fmt.Printf("SNMP Config reset to defaults successfully.\n")
}
