package bigip

import (
	"os"
	"testing"
)

func TestGetManagedDevicesFunc(t *testing.T) {
	host := os.Getenv("BIGIQ_HOST")
	user := os.Getenv("BIGIQ_USER")
	pass := os.Getenv("BIGIQ_PASSWORD")
	if host == "" || user == "" || pass == "" {
		t.Skip("Environment variables BIGIQ_HOST, BIGIQ_USER, BIGIQ_PASSWORD must be set")
	}
	b, err := NewTokenSession(&Config{
		Address:           host,
		Username:          user,
		Password:          pass,
		CertVerifyDisable: true, // Disable cert verification for testing
	})
	if err != nil {
		t.Fatalf("NewTokenSession failed: %v", err)
	}
	devices, err := b.GetManagedDevices()
	if err != nil {
		t.Fatalf("GetManagedDevices failed: %v", err)
	}
	if devices == nil || len(devices.DevicesInfo) == 0 {
		t.Errorf("No managed devices found")
	}
	for _, d := range devices.DevicesInfo {
		t.Logf("Managed Device: Address=%s, Hostname=%s, UUID=%s", d.Address, d.Hostname, d.UUID)
	}
}
