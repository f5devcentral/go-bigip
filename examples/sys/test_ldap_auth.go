package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////          LDAP Authentication Testing           ///////////////////
/////////////////////////////////////////////////////////////////////////////////

func testLdapAuthRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing LDAP Authentication Read Operation ===")

	// READ: Get system LDAP Configuration
	fmt.Println("\nGetting System Auth LDAP Configurations...")

	ldapConfig, err := f5.GetLdapConfig("system-auth")
	if err != nil {
		log.Printf("Error getting LDAP Configuration: %v", err)
	}
	if ldapConfig != nil {
		fmt.Printf("Retrieved LDAP Configuration:\n")
		fmt.Printf("  Name: %s\n", ldapConfig.Name)
		fmt.Printf("  Servers: %v\n", ldapConfig.Servers)
		fmt.Printf("  Port: %d\n", ldapConfig.Port)
		fmt.Printf("  SearchBaseDn: %s\n", ldapConfig.SearchBaseDn)
		fmt.Printf("  BindDn: %s\n", ldapConfig.BindDn)
		fmt.Printf("  SSL: %s\n", ldapConfig.Ssl)
		fmt.Printf("  Referrals: %s\n", ldapConfig.Referrals)
		fmt.Println()
	} else {
		fmt.Printf("System LDAP Authentication is not configured.\n")
	}

}

func testLdapAuthCreation(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing LDAP Authentication Creation Operation ===")

	// Note: BIG-IP only allows "system-auth" as the LDAP configuration name
	ldapConfig := &bigip.LdapConfig{
		Name:            "system-auth",
		BindDn:          "cn=ldap_admin,dc=test,dc=com",
		BindPw:          "password",
		LoginAttribute:  "uid",
		CheckRolesGroup: "enabled",
		Debug:           "enabled",
		Port:            1389,
		Referrals:       "no",
		Scope:           "sub",
		SearchBaseDn:    "dc=test,dc=com",
		Servers:         []string{"192.168.252.11"},
	}

	fmt.Println("Creating LDAP Configuration...")
	err := f5.CreateLdapConfig(ldapConfig)
	if err != nil {
		log.Printf("Error creating LDAP Configuration: %v", err)
		return
	}
	fmt.Printf("LDAP Configuration %s created successfully.\n", ldapConfig.Name)
}

func testLdapAuthUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing LDAP Authentication Update Operation ===")

	ldapConfig := &bigip.LdapConfig{
		Name:            "system-auth",
		BindDn:          "cn=ldap_admin,dc=test,dc=com",
		BindPw:          "newpassword123",
		LoginAttribute:  "uid",
		CheckRolesGroup: "disabled",
		Debug:           "disabled",
		Port:            389,
		Referrals:       "yes",
		Scope:           "sub",
		SearchBaseDn:    "dc=test,dc=com",
	}

	fmt.Println("Updating LDAP Configuration...")
	err := f5.ModifyLdapConfig(ldapConfig.Name, ldapConfig)
	if err != nil {
		log.Printf("Error updating LDAP Configuration: %v", err)
		return
	}
	fmt.Printf("LDAP Configuration %s updated successfully.\n", ldapConfig.Name)
}

func testLdapAuthDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing LDAP Authentication Delete Operation ===")

	fmt.Println("Deleting LDAP Configuration system-auth...")
	err := f5.DeleteLdapConfig("system-auth")
	if err != nil {
		log.Printf("Error deleting LDAP Configuration: %v", err)
		return
	}
	fmt.Printf("LDAP Configuration system-auth deleted successfully.\n")
}
