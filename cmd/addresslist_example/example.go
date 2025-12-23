package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/f5devcentral/go-bigip"
)

func main() {
	// Connect to the BIG-IP system.
	config := bigip.Config{
		Address:  os.Getenv("BIG_IP_HOST"),
		Username: os.Getenv("BIG_IP_USER"),
		Password: os.Getenv("BIG_IP_PASSWORD"),
	}

	f5 := bigip.NewSession(&config)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create an Address List
	addressList := &bigip.AddressList{
		Name: "addresslist-example",
		Partition: "Common",
		Addresses: []bigip.AddressListAddress{
			{ Name: "1.2.3.4" },
		},
	}

	err := f5.AddAddressList(ctx, addressList)

	if err != nil {
    log.Fatalf("Failed to create address list: %v", err)
	}

	// Get the address list
	addressListGet, err := f5.GetAddressList(ctx, addressList.Name)
	if err != nil {
		log.Fatalf("Failed to get address list: %v", err)
	}
	fmt.Printf("%+v\n", addressListGet)

	// Create a traffic matching criteria
	trafficMatchingCriteria := &bigip.TrafficMatchingCriteria{
		Name: "tmc-example",
		Partition: "Common",
		DestinationAddressInline: "192.168.1.150",
		DestinationPortInline: "443",
		SourceAddressList: fmt.Sprintf("/%s/%s", addressList.Partition, addressList.Name),
		SourcePortInline: 0,
	}

	err = f5.AddTrafficMatchingCriteria(ctx, trafficMatchingCriteria)

	if err != nil {
		log.Fatalf("Failed to create traffic matching criteria: %v", err)
	}

	// Get the traffic matching criteria
	trafficMatchingCriteriaGet, err := f5.GetTrafficMatchingCriteria(ctx, trafficMatchingCriteria.Name)
	if err != nil {
		log.Fatalf("Failed to get traffic matching criteria: %v", err)
	}
	fmt.Printf("%+v\n", trafficMatchingCriteriaGet)

	// Create a virtual server using the address list and traffic matching criteria
	virtualServer := &bigip.VirtualServer{
		Name: "virtualserver-example",
		Partition: "Common",
		Destination: ":0",
		Source: "0.0.0.0/0",
		IPProtocol: "tcp",
		Mask: "255.255.255.255",
		TrafficMatchingCriteria: fmt.Sprintf("/%s/%s", trafficMatchingCriteria.Partition, trafficMatchingCriteria.Name),
	}

	err = f5.AddVirtualServer(virtualServer)
	if err != nil {
		log.Fatalf("Failed to create virtual server: %v", err)
	}

	// Get the virtual server
	virtualServerGet, err := f5.GetVirtualServer(virtualServer.Name)
	if err != nil {
		log.Fatalf("Failed to get virtual server: %v", err)
	}
	fmt.Printf("%+v\n", virtualServerGet)

	fmt.Println("Virtual server with address list created")

	err = f5.DeleteVirtualServer(virtualServer.Name)
	if err != nil {
		log.Fatalf("Failed to delete virtual server: %v", err)
	}

	err = f5.DeleteTrafficMatchingCriteria(ctx, trafficMatchingCriteria.Name)
	if err != nil {
		log.Fatalf("Failed to delete traffic matching criteria: %v", err)
	}

	err = f5.DeleteAddressList(ctx, addressList.Name)
	if err != nil {
		log.Fatalf("Failed to delete address list: %v", err)
	}

	fmt.Println("virutal server, traffic matching criteria, and address list deleted.")
}

