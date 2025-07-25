package main

import (
	"fmt"
	"log"
	"time"

	"github.com/f5devcentral/go-bigip"
)

func main() {
	// Connect to the BIG-IP system.
	// Replace with your actual BIG-IP credentials and hostname
	config := &bigip.Config{
		Address:           "https://192.168.1.1",
		Username:          "admin",
		Password:          "secret",
		CertVerifyDisable: true, // Disable certificate verification for testing purposes
	}
	f5 := bigip.NewSession(config)

	fmt.Println("=== F5 BIG-IP Access Profile CRUD Operations Example ===")

	// Example 1: List all existing Access Profiles
	fmt.Println("1. Listing all existing Access Profiles...")
	listAccessProfiles(f5)

	// Example 2: Create a new Access Profile
	fmt.Println("\n2. Creating a new Access Profile...")
	profileName := "example-access-profile"
	createAccessProfile(f5, profileName)

	// Example 3: Get the specific Access Profile we just created
	fmt.Println("\n3. Retrieving the created Access Profile...")
	getAccessProfile(f5, profileName)

	// Example 4: Modify the Access Profile
	fmt.Println("\n4. Modifying the Access Profile...")
	modifyAccessProfile(f5, profileName)

	// Example 5: Get the modified Access Profile to verify changes
	fmt.Println("\n5. Retrieving the modified Access Profile...")
	getAccessProfile(f5, profileName)

	// Example 6: Create additional profiles with different configurations
	fmt.Println("\n6. Creating additional Access Profiles with different configurations...")
	createAdvancedAccessProfiles(f5)

	// Example 7: List all Access Profiles again to see all created profiles
	fmt.Println("\n7. Listing all Access Profiles after creation...")
	listAccessProfiles(f5)

	// Example 8: Clean up - Delete the created Access Profiles
	fmt.Println("\n8. Cleaning up - Deleting created Access Profiles...")
	deleteAccessProfile(f5, profileName)
	deleteAccessProfile(f5, "secure-access-profile")
	deleteAccessProfile(f5, "vpn-access-profile")
	deleteAccessProfile(f5, "portal-access-profile")

	fmt.Println("\n=== Access Profile CRUD Operations Example Complete ===")
}

// listAccessProfiles demonstrates how to retrieve all access profiles
func listAccessProfiles(f5 *bigip.BigIP) {
	profiles, err := f5.AccessProfiles()
	if err != nil {
		log.Printf("Error retrieving access profiles: %v", err)
		return
	}

	if len(profiles.AccessProfiles) == 0 {
		fmt.Println("No access profiles found.")
		return
	}

	fmt.Printf("Found %d access profile(s):\n", len(profiles.AccessProfiles))
	for i, profile := range profiles.AccessProfiles {
		fmt.Printf("  %d. Name: %s, Partition: %s, Type: %s\n",
			i+1, profile.Name, profile.Partition, profile.Type)
		if profile.Description != "" {
			fmt.Printf("     Description: %s\n", profile.Description)
		}
		fmt.Printf("     Full Path: %s\n", profile.FullPath)
	}
}

// createAccessProfile demonstrates how to create a basic access profile
func createAccessProfile(f5 *bigip.BigIP, name string) {
	profile := &bigip.AccessProfile{
		Name:         name,
		Partition:    "Common",
		Description:  "Example access profile created via go-bigip",
		Type:         "system-authentication",
		AccessPolicy: "/Common/access_policy", // Required: Reference to access policy
		AcceptLanguages: []string{
			"en",
			"es",
		},
		DefaultLanguage:       "en",
		InactivityTimeout:     900,    // 15 minutes
		AccessPolicyTimeout:   300,    // 5 minutes
		MaxSessionTimeout:     604800, // 7 days
		MaxConcurrentSessions: 100,
		MaxConcurrentUsers:    50,
		LogSettings: []string{
			"/Common/default-log-setting",
		},
		Services: []string{
			"http",
			"https",
		},
	}

	err := f5.CreateAccessProfile(profile)
	if err != nil {
		log.Printf("Error creating access profile '%s': %v", name, err)
		return
	}

	fmt.Printf("Successfully created access profile: %s\n", name)
}

// getAccessProfile demonstrates how to retrieve a specific access profile
func getAccessProfile(f5 *bigip.BigIP, name string) {
	profile, err := f5.GetAccessProfile(name)
	if err != nil {
		log.Printf("Error retrieving access profile '%s': %v", name, err)
		return
	}

	if profile == nil {
		fmt.Printf("Access profile '%s' not found.\n", name)
		return
	}

	fmt.Printf("Access Profile Details:\n")
	fmt.Printf("  Name: %s\n", profile.Name)
	fmt.Printf("  Partition: %s\n", profile.Partition)
	fmt.Printf("  Full Path: %s\n", profile.FullPath)
	fmt.Printf("  Description: %s\n", profile.Description)
	fmt.Printf("  Type: %s\n", profile.Type)
	fmt.Printf("  Inactivity Timeout: %d seconds\n", profile.InactivityTimeout)
	fmt.Printf("  Access Policy Timeout: %d seconds\n", profile.AccessPolicyTimeout)
	fmt.Printf("  Max Session Timeout: %d seconds\n", profile.MaxSessionTimeout)
	fmt.Printf("  Max Concurrent Sessions: %d\n", profile.MaxConcurrentSessions)
	fmt.Printf("  Max Concurrent Users: %d\n", profile.MaxConcurrentUsers)
	fmt.Printf("  User Identity Method: %s\n", profile.UserIdentityMethod)
	fmt.Printf("  Secure Cookie: %s\n", profile.SecureCookie)
	fmt.Printf("  HTTP Only Cookie: %s\n", profile.HTTPOnlyCookie)
	fmt.Printf("  Persistent Cookie: %s\n", profile.PersistentCookie)
	fmt.Printf("  Accept Languages: %v\n", profile.AcceptLanguages)
	fmt.Printf("  Services: %v\n", profile.Services)
	if profile.Generation > 0 {
		fmt.Printf("  Generation: %d\n", profile.Generation)
	}
}

// modifyAccessProfile demonstrates how to modify an existing access profile
func modifyAccessProfile(f5 *bigip.BigIP, name string) {
	// First, get the current profile to modify
	currentProfile, err := f5.GetAccessProfile(name)
	if err != nil {
		log.Printf("Error retrieving access profile for modification: %v", err)
		return
	}

	if currentProfile == nil {
		fmt.Printf("Access profile '%s' not found for modification.\n", name)
		return
	}

	// Create a modified version
	modifiedProfile := &bigip.AccessProfile{
		Name:         currentProfile.Name,
		Partition:    currentProfile.Partition,
		Description:  "Modified example access profile - updated via go-bigip",
		Type:         "all",
		AccessPolicy: "/Common/access_policy", // Required: Reference to access policy
		AcceptLanguages: []string{
			"en",
			"es",
			"fr", // Added French
		},
		DefaultLanguage:          "en",
		InactivityTimeout:        1800, // Changed to 30 minutes
		AccessPolicyTimeout:      600,  // Changed to 10 minutes
		MaxSessionTimeout:        604800,
		MaxConcurrentSessions:    200, // Increased limit
		MaxConcurrentUsers:       100, // Increased limit
		UserIdentityMethod:       "http",
		SecureCookie:             "true",
		HTTPOnlyCookie:           "true",
		PersistentCookie:         "false",
		SamesiteCookie:           "true",
		SamesiteCookieAttrValue:  "lax", // Changed from strict to lax
		RestrictToSingleClientIP: "false",
		UseHTTP503OnError:        "false",
		WebtopRedirectOnRootURI:  "true",
		LogSettings: []string{
			"/Common/default-log-setting",
			"/Common/access-log-setting", // Added additional log setting
		},
		Services: []string{
			"http",
			"https",
			"ftp", // Added FTP service
		},
	}

	err = f5.ModifyAccessProfile(name, modifiedProfile)
	if err != nil {
		log.Printf("Error modifying access profile '%s': %v", name, err)
		return
	}

	fmt.Printf("Successfully modified access profile: %s\n", name)
}

// createAdvancedAccessProfiles demonstrates creating profiles with different configurations
func createAdvancedAccessProfiles(f5 *bigip.BigIP) {
	// Create a secure access profile
	secureProfile := &bigip.AccessProfile{
		Name:                     "secure-access-profile",
		Partition:                "Common",
		Description:              "High-security access profile for sensitive applications",
		Type:                     "all",
		AccessPolicy:             "/Common/access_policy", // Required: Reference to access policy
		AcceptLanguages:          []string{"en"},
		DefaultLanguage:          "en",
		InactivityTimeout:        300,   // 5 minutes for high security
		AccessPolicyTimeout:      180,   // 3 minutes
		MaxSessionTimeout:        28800, // 8 hours max
		MaxConcurrentSessions:    10,    // Limited concurrent sessions
		MaxConcurrentUsers:       5,     // Limited concurrent users
		UserIdentityMethod:       "http",
		SecureCookie:             "true",
		HTTPOnlyCookie:           "true",
		PersistentCookie:         "false",
		SamesiteCookie:           "true",
		SamesiteCookieAttrValue:  "strict",
		RestrictToSingleClientIP: "true", // Enhanced security
		UseHTTP503OnError:        "true",
		WebtopRedirectOnRootURI:  "true",
		Services:                 []string{"https"}, // HTTPS only
	}

	err := f5.CreateAccessProfile(secureProfile)
	if err != nil {
		log.Printf("Error creating secure access profile: %v", err)
	} else {
		fmt.Println("Created secure access profile")
	}

	// Wait a moment between creations
	time.Sleep(1 * time.Second)

	// Create a VPN access profile
	vpnProfile := &bigip.AccessProfile{
		Name:                     "vpn-access-profile",
		Partition:                "Common",
		Description:              "VPN access profile for remote workers",
		Type:                     "all",
		AccessPolicy:             "/Common/access_policy", // Required: Reference to access policy
		AcceptLanguages:          []string{"en", "es", "fr", "de"},
		DefaultLanguage:          "en",
		InactivityTimeout:        3600,  // 1 hour
		AccessPolicyTimeout:      900,   // 15 minutes
		MaxSessionTimeout:        86400, // 24 hours
		MaxConcurrentSessions:    500,   // Higher limit for VPN
		MaxConcurrentUsers:       250,   // Higher limit for VPN
		UserIdentityMethod:       "http",
		SecureCookie:             "true",
		HTTPOnlyCookie:           "true",
		PersistentCookie:         "true", // Allow persistent for VPN
		SamesiteCookie:           "true",
		SamesiteCookieAttrValue:  "lax",
		RestrictToSingleClientIP: "false",
		UseHTTP503OnError:        "false",
		WebtopRedirectOnRootURI:  "true",
		Services:                 []string{"http", "https", "ftp", "ssh"},
	}

	err = f5.CreateAccessProfile(vpnProfile)
	if err != nil {
		log.Printf("Error creating VPN access profile: %v", err)
	} else {
		fmt.Println("Created VPN access profile")
	}

	// Wait a moment between creations
	time.Sleep(1 * time.Second)

	// Create a portal access profile
	portalProfile := &bigip.AccessProfile{
		Name:                     "portal-access-profile",
		Partition:                "Common",
		Description:              "Web portal access profile for general users",
		Type:                     "all",
		AccessPolicy:             "/Common/access_policy", // Required: Reference to access policy
		AcceptLanguages:          []string{"en", "es"},
		DefaultLanguage:          "en",
		InactivityTimeout:        1800,  // 30 minutes
		AccessPolicyTimeout:      600,   // 10 minutes
		MaxSessionTimeout:        43200, // 12 hours
		MaxConcurrentSessions:    1000,  // High limit for portal
		MaxConcurrentUsers:       500,   // High limit for portal
		UserIdentityMethod:       "http",
		SecureCookie:             "true",
		HTTPOnlyCookie:           "false", // Less restrictive for portal
		PersistentCookie:         "true",
		SamesiteCookie:           "false", // Less restrictive
		SamesiteCookieAttrValue:  "none",
		RestrictToSingleClientIP: "false",
		UseHTTP503OnError:        "false",
		WebtopRedirectOnRootURI:  "true",
		Services:                 []string{"http", "https"},
	}

	err = f5.CreateAccessProfile(portalProfile)
	if err != nil {
		log.Printf("Error creating portal access profile: %v", err)
	} else {
		fmt.Println("Created portal access profile")
	}
}

// deleteAccessProfile demonstrates how to delete an access profile
func deleteAccessProfile(f5 *bigip.BigIP, name string) {
	// First check if the profile exists
	profile, err := f5.GetAccessProfile(name)
	if err != nil {
		log.Printf("Error checking access profile '%s' before deletion: %v", name, err)
		return
	}

	if profile == nil {
		fmt.Printf("Access profile '%s' not found, skipping deletion.\n", name)
		return
	}

	// Delete the profile
	err = f5.DeleteAccessProfile(name)
	if err != nil {
		log.Printf("Error deleting access profile '%s': %v", name, err)
		return
	}

	fmt.Printf("Successfully deleted access profile: %s\n", name)

	// Verify deletion
	deletedProfile, err := f5.GetAccessProfile(name)
	if err != nil {
		log.Printf("Error verifying deletion of access profile '%s': %v", name, err)
		return
	}

	if deletedProfile == nil {
		fmt.Printf("Verified: Access profile '%s' has been deleted.\n", name)
	} else {
		fmt.Printf("Warning: Access profile '%s' still exists after deletion attempt.\n", name)
	}
}
