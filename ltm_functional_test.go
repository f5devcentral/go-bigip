package bigip

import (
	"os"
	"testing"
)

// Functional test for IRules() using a real or integration environment
func TestFunctionalIRules(t *testing.T) {
	address := os.Getenv("BIGIP_HOST")
	username := os.Getenv("BIGIP_USER")
	password := os.Getenv("BIGIP_PASSWORD")
	t.Logf("BIGIP_HOST=%s, BIGIP_USER=%s", address, username) // Debug log for env vars
	if address == "" || username == "" || password == "" {
		t.Skip("BIGIP environment variables not set")
	}
	cfg := &Config{
		Address:           address,
		Username:          username,
		Password:          password,
		CertVerifyDisable: true, // Disable certificate validation for testing
	}
	t.Logf("Config: %+v", cfg) // Debug log for config
	client, err := NewTokenSession(cfg)
	if err != nil {
		t.Fatalf("NewTokenSession error: %v", err)
	}

	rules, err := client.IRules()
	if err != nil {
		t.Fatalf("IRules() error: %v", err)
	}
	if rules == nil || len(rules.IRules) == 0 {
		t.Errorf("No iRules returned, got: %+v", rules)
	} else {
		for _, rule := range rules.IRules {
			t.Logf("iRule: %+v", rule)
		}
	}
}

// Functional test to get a specific iRule by name
// Functional test to test the IRule() method for fetching a specific iRule by name
func TestGetSpecificIRule(t *testing.T) {
	address := os.Getenv("BIGIP_HOST")
	username := os.Getenv("BIGIP_USER")
	password := os.Getenv("BIGIP_PASSWORD")

	iruleName := "default_f5_healt"
	t.Logf("BIGIP_HOST=%s, BIGIP_USER=%s, BIGIP_IRULE_NAME=%s", address, username, iruleName)
	if address == "" || username == "" || password == "" || iruleName == "" {
		t.Skip("BIGIP environment variables not set or iRule name missing")
	}
	cfg := &Config{
		Address:           address,
		Username:          username,
		Password:          password,
		CertVerifyDisable: true,
	}
	client, err := NewTokenSession(cfg)
	if err != nil {
		t.Fatalf("NewTokenSession error: %v", err)
	}

	irule, err := client.IRule(iruleName)
	if err != nil {
		t.Fatalf("IRule error: %v", err)
	}
	if irule == nil {
		t.Errorf("iRule '%s' not found", iruleName)
	} else {
		t.Logf("Found iRule '%s': %+v", iruleName, irule)
	}
}

// Add more functional tests for other LTM features as needed
