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
	const wrkspcName = "ExampleWorkspce"
	err := f5.CreateWorkspace(ctx, wrkspcName)
	if err != nil {
		panic(err)
	}
	result, err := f5.GetWorkspace(ctx, wrkspcName)
	if err != nil {
		panic(err)
	}
	log.Printf("Workspace: %v", result)
	opts := bigip.ExtensionConfig{
		WorkspaceName: wrkspcName,
		Name:          "exampleExt",
		Partition:     "Common",
	}
	err = f5.CreateExtension(ctx, opts)
	if err != nil {
		panic(err)
	}

	err = f5.WriteExtensionFile(ctx, opts, "{}", bigip.PackageJSON)
	if err != nil {
		panic(err)
	}

	err = f5.WriteExtensionFile(ctx, opts, "const a = 12;", bigip.IndexJS)
	if err != nil {
		panic(err)
	}

	content, err := f5.ReadExtensionFile(ctx, opts, bigip.IndexJS)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Content: %+v\n", content)
}
