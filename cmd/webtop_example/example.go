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
		Address:           os.Getenv("BIG_IP_HOST"),
		Username:          os.Getenv("BIG_IP_USER"),
		Password:          os.Getenv("BIG_IP_PASSWORD"),
		CertVerifyDisable: true,
	}

	f5 := bigip.NewSession(&config)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	conf := bigip.WebtopConfig{
		Description:        "A webtop example",
		Type:               bigip.WebtopTypePortal,
		LinkType:           bigip.LinkTypeUri,
		CustomizationType:  bigip.CustomizationTypeModern,
		CustomizationGroup: "/Common/webtop_customization",
		LocationSpecific:   false,
		MinimizeToTray:     true,
		ShowSearch:         true,
		WarningOnClose:     false,
		UrlEntryField:      true,
		ResourceSearch:     false,
		InitialState:       bigip.InitialStateCollapsed,
	}
	webtop := bigip.Webtop{
		Name:         "ExampleName",
		Partition:    "Common",
		TMPartition:  "Common",
		WebtopConfig: conf,
	}

	err := f5.CreateWebtop(ctx, webtop)
	if err != nil {
		log.Fatalf("Failed to create webtop: %v", err)
	}
	webtopGet, err := f5.GetWebtop(ctx, webtop.Name)
	if err != nil {
		log.Fatalf("Failed to get webtop: %v", err)
	}
	fmt.Printf("%+v\n", webtopGet)
	err = f5.DeleteWebtop(ctx, webtop.Name)
	if err != nil {
		log.Fatalf("Failed to delete webtop: %v", err)
	}
	log.Println("Webtop created")
}
