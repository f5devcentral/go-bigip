package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////       Device Self Testing                        /////////////////
/////////////////////////////////////////////////////////////////////////////////

func testDeviceSelfRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Device Self Read Operation ===")

	fmt.Println("Getting Device Self...")
	deviceSelf, err := f5.GetDeviceSelf("ltm02.f5lab.com")
	if err != nil {
		log.Printf("Error getting Device Self: %v", err)
		return
	}

	if deviceSelf != nil {
		fmt.Printf("Retrieved Device Self:\n")
		fmt.Printf("  Name: %s\n", deviceSelf.Name)
		fmt.Printf("  Mirror IP: %s\n", deviceSelf.MirrorIp)
		fmt.Printf("  Mirror Secondary IP: %s\n", deviceSelf.MirrorSecondaryIp)
		fmt.Printf("  Configsync IP: %s\n", deviceSelf.ConfigsyncIp)
		fmt.Printf("  Unicast Addresses:\n")
		for i, addr := range deviceSelf.UnicastAddress {
			fmt.Printf("    [%d] IP: %s, Port: %d, EffectiveIP: %s, EffectivePort: %d\n",
				i, addr.IP, addr.Port, addr.EffectiveIP, addr.EffectivePort)
		}
		fmt.Println()
	}
}

func testDeviceSelfCreate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Device Self Create Operation ===")

	deviceSelf := &bigip.DeviceSelf{
		Name:              "ltm02.f5lab.com",
		MirrorIp:          "any6",
		MirrorSecondaryIp: "any6",
		ConfigsyncIp:      "10.1.131.38",
		UnicastAddress: []bigip.UnicastAddress{
			{
				IP:            "10.1.131.38",
				Port:          1026,
				EffectiveIP:   "10.1.131.38",
				EffectivePort: 1026,
			},
			{
				IP:            "management-ip",
				Port:          1026,
				EffectiveIP:   "management-ip",
				EffectivePort: 1026,
			},
		},
	}

	fmt.Println("Creating Device Self...")
	err := f5.CreateDeviceSelf(deviceSelf)
	if err != nil {
		log.Printf("Error creating Device Self: %v", err)
		return
	}
	fmt.Printf("Device Self created successfully.\n")
}

func testDeviceSelfUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Device Self Update Operation ===")

	deviceSelf := &bigip.DeviceSelf{
		Name:              "ltm02.f5lab.com",
		MirrorIp:          "any6",
		MirrorSecondaryIp: "any6",
		ConfigsyncIp:      "10.1.131.38",
		UnicastAddress: []bigip.UnicastAddress{
			{
				IP:            "10.1.131.38",
				Port:          1026,
				EffectiveIP:   "10.1.131.38",
				EffectivePort: 1026,
			},
		},
	}

	fmt.Println("Updating Device Self...")
	err := f5.ModifyDeviceSelf(deviceSelf)
	if err != nil {
		log.Printf("Error updating Device Self: %v", err)
		return
	}
	fmt.Printf("Device Self updated successfully.\n")
}

func testDeviceSelfDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing Device Self Delete Operation ===")

	fmt.Println("Deleting (resetting) Device Self...")
	err := f5.DeleteDeviceSelf("ltm02.f5lab.com")
	if err != nil {
		log.Printf("Error deleting Device Self: %v", err)
		return
	}
	fmt.Printf("Device Self reset to defaults successfully.\n")

}
