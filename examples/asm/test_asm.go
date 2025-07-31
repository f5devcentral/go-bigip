package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

// testDosProfile tests CRUD operations for DOS profiles
func testDosProfile(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing DOS Profile CRUD Operations ===")

	// Test profile configuration
	dosProfile := &bigip.DOSProfile{
		Name:                 "test-dos-profile",
		ThresholdSensitivity: "medium",
	}

	// CREATE: Add DOS profile
	fmt.Println("Creating DOS profile...")
	err := f5.AddDOSProfile(dosProfile)
	if err != nil {
		log.Printf("Error creating DOS profile: %v", err)
		return
	}
	fmt.Println("DOS profile created successfully")

	// READ: Get all DOS profiles
	fmt.Println("Getting all DOS profiles...")
	profiles, err := f5.DOSProfiles()
	if err != nil {
		log.Printf("Error getting DOS profiles: %v", err)
		return
	}
	fmt.Printf("Found %d DOS profiles\n", len(profiles.DOSProfiles))

	// READ: Get specific DOS profile
	fmt.Println("Getting specific DOS profile...")
	profile, err := f5.GetDOSProfile("test-dos-profile")
	if err != nil {
		log.Printf("Error getting DOS profile: %v", err)
		return
	}
	if profile != nil {
		fmt.Printf("Retrieved DOS profile: %s with sensitivity: %s\n", profile.Name, profile.ThresholdSensitivity)
	}

	// UPDATE: Modify DOS profile
	fmt.Println("Updating DOS profile...")
	dosProfile.ThresholdSensitivity = "high"
	err = f5.ModifyDOSProfile("test-dos-profile", dosProfile)
	if err != nil {
		log.Printf("Error updating DOS profile: %v", err)
		return
	}
	fmt.Println("DOS profile updated successfully")

	// Verify update
	updatedProfile, err := f5.GetDOSProfile("test-dos-profile")
	if err != nil {
		log.Printf("Error getting updated DOS profile: %v", err)
		return
	}
	if updatedProfile != nil {
		fmt.Printf("Updated DOS profile sensitivity: %s\n", updatedProfile.ThresholdSensitivity)
	}

	// DELETE: Remove DOS profile
	fmt.Println("Deleting DOS profile...")
	err = f5.DeleteDOSProfile("test-dos-profile")
	if err != nil {
		log.Printf("Error deleting DOS profile: %v", err)
		return
	}
	fmt.Println("DOS profile deleted successfully")

	// Verify deletion
	_, err = f5.GetDOSProfile("test-dos-profile")
	if err != nil && err.Error() == "01020036:3: The requested DOS profile (/Common/test-dos-profile) was not found." {
		fmt.Println("DOS profile deletion confirmed")
		return
	}
}

// testFirewallPolicy tests CRUD operations for Firewall policies
func testFirewallPolicy(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Firewall Policy CRUD Operations ===")

	// Test policy configuration
	firewallPolicy := &bigip.FirewallPolicy{
		Name: "test-firewall-policy",
	}

	// CREATE: Add Firewall policy
	fmt.Println("Creating Firewall policy...")
	err := f5.AddFirewallPolicy(firewallPolicy)
	if err != nil {
		log.Printf("Error creating Firewall policy: %v", err)
		return
	}
	fmt.Println("Firewall policy created successfully")

	// READ: Get all Firewall policies
	fmt.Println("Getting all Firewall policies...")
	policies, err := f5.FirewallPolicies()
	if err != nil {
		log.Printf("Error getting Firewall policies: %v", err)
		return
	}
	fmt.Printf("Found %d Firewall policies\n", len(policies.FirewallPolicies))

	// READ: Get specific Firewall policy
	fmt.Println("Getting specific Firewall policy...")
	policy, err := f5.GetFirewallPolicy("test-firewall-policy")
	if err != nil {
		log.Printf("Error getting Firewall policy: %v", err)
		return
	}
	if policy != nil {
		fmt.Printf("Retrieved Firewall policy: %s\n", policy.Name)
	}

	//// UPDATE: Modify Firewall policy
	//fmt.Println("Updating Firewall policy...")
	//// Note: For firewall policies, we mainly update metadata fields
	//err = f5.ModifyFirewallPolicy("test-firewall-policy", firewallPolicy)
	//if err != nil {
	//	log.Printf("Error updating Firewall policy: %v", err)
	//	return
	//}
	//fmt.Println("Firewall policy updated successfully")

	// DELETE: Remove Firewall policy
	fmt.Println("Deleting Firewall policy...")
	err = f5.DeleteFirewallPolicy("test-firewall-policy")
	if err != nil {
		log.Printf("Error deleting Firewall policy: %v", err)
		return
	}
	fmt.Println("Firewall policy deleted successfully")

	// Verify deletion
	_, err = f5.GetFirewallPolicy("test-firewall-policy")
	if err != nil {
		if err.Error() == "01020036:3: The requested firewall policy (/Common/test-firewall-policy) was not found." {
			fmt.Println("Firewall policy deletion confirmed")
		}
		return
	}
}

// testIpIntelligence tests CRUD operations for IP Intelligence policies
func testIpIntelligence(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing IP Intelligence Policy CRUD Operations ===")

	// Test policy configuration
	ipIntelPolicy := &bigip.IPIntelligencePolicy{
		Name:                            "test-ip-intelligence-policy",
		DefaultAction:                   "drop",
		DefaultLogBlacklistHitOnly:      "no",
		DefaultLogBlacklistWhitelistHit: "no",
	}

	// CREATE: Add IP Intelligence policy
	fmt.Println("Creating IP Intelligence policy...")
	err := f5.AddIPIntelligencePolicy(ipIntelPolicy)
	if err != nil {
		log.Printf("Error creating IP Intelligence policy: %v", err)
		return
	}
	fmt.Println("IP Intelligence policy created successfully")

	// READ: Get all IP Intelligence policies
	fmt.Println("Getting all IP Intelligence policies...")
	policies, err := f5.IPIntelligencePolicies()
	if err != nil {
		log.Printf("Error getting IP Intelligence policies: %v", err)
		return
	}
	fmt.Printf("Found %d IP Intelligence policies\n", len(policies.IPIntelligencePolicies))

	// READ: Get specific IP Intelligence policy
	fmt.Println("Getting specific IP Intelligence policy...")
	policy, err := f5.GetIPIntelligencePolicy("test-ip-intelligence-policy")
	if err != nil {
		log.Printf("Error getting IP Intelligence policy: %v", err)
		return
	}
	if policy != nil {
		fmt.Printf("Retrieved IP Intelligence policy: %s with default action: %s\n", policy.Name, policy.DefaultAction)
	}

	// UPDATE: Modify IP Intelligence policy
	fmt.Println("Updating IP Intelligence policy...")
	ipIntelPolicy.DefaultAction = "accept"
	ipIntelPolicy.DefaultLogBlacklistHitOnly = "yes"
	err = f5.ModifyIPIntelligencePolicy("test-ip-intelligence-policy", ipIntelPolicy)
	if err != nil {
		log.Printf("Error updating IP Intelligence policy: %v", err)
		return
	}
	fmt.Println("IP Intelligence policy updated successfully")

	// Verify update
	updatedPolicy, err := f5.GetIPIntelligencePolicy("test-ip-intelligence-policy")
	if err != nil {
		log.Printf("Error getting updated IP Intelligence policy: %v", err)
		return
	}
	if updatedPolicy.DefaultLogBlacklistHitOnly == "yes" && ipIntelPolicy.DefaultAction == "accept" {
		fmt.Printf("Updated IP Intelligence policy default action: %s\n", updatedPolicy.DefaultAction)
	}

	// DELETE: Remove IP Intelligence policy
	fmt.Println("Deleting IP Intelligence policy...")
	err = f5.DeleteIPIntelligencePolicy("test-ip-intelligence-policy")
	if err != nil {
		log.Printf("Error deleting IP Intelligence policy: %v", err)
		return
	}
	fmt.Println("IP Intelligence policy deleted successfully")

	// Verify deletion
	_, err = f5.GetIPIntelligencePolicy("test-ip-intelligence-policy")
	if err != nil {
		if err.Error() == "01020036:3: The requested Firewall dynamic black/white lists policy (/Common/test-ip-intelligence-policy) was not found." {
			fmt.Println("IP Intelligence policy deletion confirmed")
		}
	}
}

// testLogProfile tests CRUD operations for Security Log profiles
func testLogProfile(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Security Log Profile CRUD Operations ===")

	// Test profile configuration - using updated structure with correct JSON tags
	logProfile := &bigip.SecurityLogProfile{
		Name:        "test-security-log-profile",
		Description: "Test security log profile for demonstration",
		IPIntelligence: struct {
			AggregateRate        int    `json:"aggregate-rate,omitempty"`
			LogGeo               string `json:"log-geo,omitempty"`
			LogRtbh              string `json:"log-rtbh,omitempty"`
			LogScrubber          string `json:"log-scrubber,omitempty"`
			LogShun              string `json:"log-shun,omitempty"`
			LogTranslationFields string `json:"log-translation-fields,omitempty"`
		}{
			AggregateRate:        100,
			LogGeo:               "disabled",
			LogRtbh:              "disabled",
			LogScrubber:          "disabled",
			LogShun:              "disabled",
			LogTranslationFields: "enabled",
		},
	}

	// CREATE: Add Security Log profile
	fmt.Println("Creating Security Log profile...")
	err := f5.AddSecurityLogProfile(logProfile)
	if err != nil {
		log.Printf("Error creating Security Log profile: %v", err)
		return
	}
	fmt.Println("Security Log profile created successfully")

	// READ: Get all Security Log profiles
	fmt.Println("Getting all Security Log profiles...")
	profiles, err := f5.SecurityLogProfiles()
	if err != nil {
		log.Printf("Error getting Security Log profiles: %v", err)
		return
	}
	fmt.Printf("Found %d Security Log profiles\n", len(profiles.SecurityLogProfiles))

	// READ: Get specific Security Log profile
	fmt.Println("Getting specific Security Log profile...")
	profile, err := f5.GetSecurityLogProfile("test-security-log-profile")
	if err != nil {
		log.Printf("Error getting Security Log profile: %v", err)
		return
	}
	if profile != nil {
		fmt.Printf("Retrieved Security Log profile: %s with description: %s\n", profile.Name, profile.Description)
	}

	// UPDATE: Modify Security Log profile
	fmt.Println("Updating Security Log profile...")
	logProfile.Description = "Updated test security log profile"
	logProfile.IPIntelligence.AggregateRate = 200
	logProfile.IPIntelligence.LogGeo = "disabled"
	err = f5.ModifySecurityLogProfile("test-security-log-profile", logProfile)
	if err != nil {
		log.Printf("Error updating Security Log profile: %v", err)
		return
	}
	fmt.Println("Security Log profile updated successfully")

	// Verify update
	updatedProfile, err := f5.GetSecurityLogProfile("test-security-log-profile")
	if err != nil {
		log.Printf("Error getting updated Security Log profile: %v", err)
		return
	}
	if updatedProfile.IPIntelligence.AggregateRate == 200 {
		fmt.Printf("Updated aggregate rate: %d\n", updatedProfile.IPIntelligence.AggregateRate)
	}

	// DELETE: Remove Security Log profile
	fmt.Println("Deleting Security Log profile...")
	err = f5.DeleteSecurityLogProfile("test-security-log-profile")
	if err != nil {
		log.Printf("Error deleting Security Log profile: %v", err)
		return
	}
	fmt.Println("Security Log profile deleted successfully")

	// Verify deletion
	_, err = f5.GetSecurityLogProfile("test-security-log-profile")
	if err != nil {
		if err.Error() == "01020036:3: The requested Security Log profile (/Common/test-security-log-profile) was not found." {
			fmt.Println("Security Log profile deletion confirmed")

		}
	}
}
