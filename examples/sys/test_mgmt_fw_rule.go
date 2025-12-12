package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////          Management Firewall Rules Testing          //////////////
/////////////////////////////////////////////////////////////////////////////////

func testManagementFwRulesRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Management Firewall Rule Read Operation ===")

	fmt.Println("Getting all Management Firewall Rules...")
	mgmtiprules, err := f5.GetManagementFwRules()
	if err != nil {
		log.Printf("Error getting Management Firewall Rules: %v", err)
		return
	}
	fmt.Printf("Found %d Management Firewall Rules\n", len(mgmtiprules.MgmtFirewallRules))

	fmt.Println("Getting specific Management Firewall Rule...")
	for i := 0; i < len(mgmtiprules.MgmtFirewallRules); i++ {
		mgmtiprule, err := f5.GetManagementFwRule(mgmtiprules.MgmtFirewallRules[i].Name)
		if err != nil {
			log.Printf("Error getting Management Firewall Rule: %v", err)
			return
		}
		if mgmtiprule != nil {
			fmt.Printf("Retrieved Management Firewall Rule:\n%s - %s via %v\n", mgmtiprule.Name, mgmtiprule.Action, mgmtiprule.Destination)
		}
	}
}

func testManagementFwRuleCreation(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Management Firewall Rule Creation Operation ===")

	mgmtFwRule := &bigip.MgmtFirewallRule{
		Name:        "test-rule",
		IpProtocol:  "tcp",
		Action:      "accept",
		Log:         "no",
		Status:      "enabled",
		PlaceBefore: "first",
		Destination: bigip.MgmtFwRuleIpPortData{
			Addresses: []bigip.MgmtFwRuleAddress{
				{Name: "30.40.40.0/24"},
				{Name: "30.40.50.0/24"},
			},
			Ports: []bigip.MgmtFwRulePort{
				{Name: "555"},
			},
		},
		Source: bigip.MgmtFwRuleIpPortData{
			Addresses: []bigip.MgmtFwRuleAddress{
				{Name: "10.30.40.0/24"},
				{Name: "20.30.50.0/24"},
			},
			Ports: []bigip.MgmtFwRulePort{
				{Name: "8080"},
				{Name: "4443"},
			},
		},
	}

	fmt.Println("Creating Management Firewall Rule...")
	err := f5.CreateManagementFwRule(mgmtFwRule)
	if err != nil {
		log.Printf("Error creating Management Firewall Rule: %v", err)
		return
	}
	fmt.Printf("Management Firewall Rule %s created successfully.\n", mgmtFwRule.Name)
}

func testManagementFwRuleModify(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Management Firewall Rule Modify Operation ===")

	mgmtFwRule := &bigip.MgmtFirewallRule{
		Name:       "test-rule",
		IpProtocol: "tcp",
		Schedule:   "/Common/ssh_schedule",
		Status:     "scheduled",
		Destination: bigip.MgmtFwRuleIpPortData{
			Addresses: []bigip.MgmtFwRuleAddress{
				{Name: "30.30.70.0/24"},
				{Name: "30.30.80.0/24"},
			},
			Ports: []bigip.MgmtFwRulePort{},
		},
	}

	fmt.Println("Updating Management Firewall Rule...")
	err := f5.ModifyManagementFwRule(mgmtFwRule.Name, mgmtFwRule)
	if err != nil {
		log.Printf("Error updating Management Firewall Rule: %v", err)
		return
	}
	fmt.Printf("Management Firewall Rule %s updated successfully.\n", mgmtFwRule.Name)
}

func testManagementFwRuleDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Management IP Rules Delete Operation ===")

	mgmtfwRule := &bigip.MgmtFirewallRule{
		Name: "test-rule",
	}

	fmt.Println("Deleting Management Firewall Rule...")
	err := f5.DeleteManagementFwRule(mgmtfwRule.Name)
	if err != nil {
		log.Printf("Error deleting Management Firewall Rule: %v", err)
		return
	}
	fmt.Printf("Management Firewall Rule %s deleted successfully.\n", mgmtfwRule.Name)

}
