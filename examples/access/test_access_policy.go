package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/f5devcentral/go-bigip"
)

func testAccessPolicy(f5 *bigip.BigIP) {

	fmt.Println("=== F5 BIG-IP Access Policy CRUD Operations Test ===")

	// Test 1: List all existing Access Policies
	fmt.Println("\n1. Listing all existing Access Policies...")
	listAccessPolicies(f5)

	// Test 2: Create a new Access Policy
	fmt.Println("\n2. Creating a new Access Policy...")
	policyName := "test-access-policy"
	createAccessPolicy(f5, policyName)

	// Test 3: Get the specific Access Policy we just created
	fmt.Println("\n3. Retrieving the created Access Policy...")
	getAccessPolicy(f5, policyName)

	// Test 4: Modify the Access Policy
	fmt.Println("\n4. Modifying the Access Policy...")
	modifyAccessPolicy(f5, policyName)

	// Test 5: Get the modified Access Policy to verify changes
	fmt.Println("\n5. Retrieving the modified Access Policy...")
	getAccessPolicy(f5, policyName)

	// Test 6: Clean up - Delete the created Access Policy
	fmt.Println("\n6. Cleaning up - Deleting created Access Policy...")
	deleteAccessPolicy(f5, policyName)

	fmt.Println("\n=== Access Policy CRUD Operations Test Complete ===")
}

// listAccessPolicies demonstrates how to retrieve all access policies
func listAccessPolicies(f5 *bigip.BigIP) {
	// Create a context with timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Use a channel to handle the result
	type result struct {
		policies *bigip.AccessPolicies
		err      error
	}

	resultChan := make(chan result, 1)

	go func() {
		policies, err := f5.AccessPolicies()
		resultChan <- result{policies: policies, err: err}
	}()

	select {
	case res := <-resultChan:
		if res.err != nil {
			log.Printf("Error retrieving access policies: %v", res.err)
			return
		}

		if len(res.policies.Items) == 0 {
			fmt.Println("No access policies found.")
			return
		}

		fmt.Printf("Found %d access policy(ies):\n", len(res.policies.Items))
		for i, policy := range res.policies.Items {
			fmt.Printf("  %d. Name: %s, Partition: %s, Type: %s\n",
				i+1, policy.Name, policy.Partition, policy.Type)
			if policy.DefaultEnding != "" {
				fmt.Printf("     Default Ending: %s\n", policy.DefaultEnding)
			}
			fmt.Printf("     Full Path: %s\n", policy.FullPath)
		}
	case <-ctx.Done():
		fmt.Printf("Timeout: Failed to retrieve access policies within 30 seconds\n")
	}
}

// createAccessPolicy demonstrates how to create a basic access policy
func createAccessPolicy(f5 *bigip.BigIP, name string) {
	policy := &bigip.AccessPolicy{
		Name:              name,
		Partition:         "Common",
		Type:              "per-rq-policy",
		DefaultEnding:     "allow", // Use simple allow instead of non-existent item
		MaxMacroLoopCount: 1,
		// StartItem:         "fallback", // Add proper start item
		PerReqPolicyProperties: []bigip.PerReqPolicyProperty{
			{
				Name:             "PerReqPolicyProperty",
				Partition:        "Common",
				IncompleteAction: "allow",
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resultChan := make(chan error, 1)

	go func() {
		err := f5.CreateAccessPolicy(policy)
		resultChan <- err
	}()

	select {
	case err := <-resultChan:
		if err != nil {
			log.Printf("Error creating access policy '%s': %v", name, err)
			return
		}
		fmt.Printf("Successfully created access policy: %s\n", name)

		// Wait a moment for the policy to be fully created
		time.Sleep(2 * time.Second)
	case <-ctx.Done():
		fmt.Printf("Timeout: Failed to create access policy within 30 seconds\n")
	}
}

// getAccessPolicy demonstrates how to retrieve a specific access policy
func getAccessPolicy(f5 *bigip.BigIP, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	type result struct {
		policy *bigip.AccessPolicy
		err    error
	}

	resultChan := make(chan result, 1)

	go func() {
		policy, err := f5.GetAccessPolicy(name)
		resultChan <- result{policy: policy, err: err}
	}()

	select {
	case res := <-resultChan:
		if res.err != nil {
			log.Printf("Error retrieving access policy '%s': %v", name, res.err)
			return
		}

		if res.policy == nil {
			fmt.Printf("Access policy '%s' not found.\n", name)
			return
		}

		fmt.Printf("Access Policy Details:\n")
		fmt.Printf("  Name: %s\n", res.policy.Name)
		fmt.Printf("  Partition: %s\n", res.policy.Partition)
		fmt.Printf("  Full Path: %s\n", res.policy.FullPath)
		fmt.Printf("  Type: %s\n", res.policy.Type)
		fmt.Printf("  Default Ending: %s\n", res.policy.DefaultEnding)
		fmt.Printf("  Max Macro Loop Count: %d\n", res.policy.MaxMacroLoopCount)
		fmt.Printf("  Oneshot Macro: %s\n", res.policy.OneshotMacro)
		fmt.Printf("  Start Item: %s\n", res.policy.StartItem)
		if len(res.policy.PerReqPolicyProperties) > 0 {
			fmt.Printf("  Per-Request Policy Properties:\n")
			for i, prop := range res.policy.PerReqPolicyProperties {
				fmt.Printf("    %d. Name: %s, Partition: %s, Incomplete Action: %s\n",
					i+1, prop.Name, prop.Partition, prop.IncompleteAction)
			}
		}
		if res.policy.Generation > 0 {
			fmt.Printf("  Generation: %d\n", res.policy.Generation)
		}
	case <-ctx.Done():
		fmt.Printf("Timeout: Failed to retrieve access policy within 30 seconds\n")
	}
}

// modifyAccessPolicy demonstrates how to modify an existing access policy
func modifyAccessPolicy(f5 *bigip.BigIP, name string) {
	currentPolicy, err := f5.GetAccessPolicy(name)
	if err != nil {
		log.Printf("Error retrieving access policy '%s' for modification: %v", name, err)
		return
	}
	// Create a modified version
	currentPolicy.PerReqPolicyProperties[0].IncompleteAction = "deny"
	modifiedPolicy := &bigip.AccessPolicy{
		Name:              currentPolicy.Name,
		Partition:         currentPolicy.Partition,
		Type:              "per-rq-policy",
		DefaultEnding:     "/Common/test-access-policy_end_reject",
		MaxMacroLoopCount: 1,
		// Preserve existing PerReqPolicyProperties - don't modify them to avoid deletion error
		PerReqPolicyProperties: currentPolicy.PerReqPolicyProperties,
	}

	err = f5.ModifyAccessPolicy(name, modifiedPolicy)
	if err != nil {
		log.Printf("Error modifying access policy '%s': %v", name, err)
		return
	}
}

// deleteAccessPolicy demonstrates how to delete an access policy
func deleteAccessPolicy(f5 *bigip.BigIP, name string) {
	err := f5.DeleteAccessPolicy(name)
	if err != nil {
		log.Printf("Error deleting access policy '%s': %v", name, err)
		return
	}
}
