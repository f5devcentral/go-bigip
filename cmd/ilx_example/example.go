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

	// Create a new BIG-IP session with the provided configuration.
	f5 := bigip.NewSession(&config)

	// Create a context with a timeout of 1 second.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Define the workspace configuration.
	workspaceConfig := bigip.WorkspaceConfig{
		WorkspaceName: "ExampleWorkspace",
		Partition:     "Common",
	}

	// Create a new workspace.
	err := f5.CreateWorkspace(ctx, workspaceConfig.WorkspaceName)
	if err != nil {
		panic(err)
	}

	// Fetch the details of the created workspace.
	result, err := f5.GetWorkspace(ctx, workspaceConfig.WorkspaceName)
	if err != nil {
		panic(err)
	}
	log.Printf("Workspace: %v", result)

	// Define the extension configuration.
	opts := bigip.ExtensionConfig{
		WorkspaceConfig: workspaceConfig,
		ExtensionName:   "exampleExt",
	}

	// Create a new extension.
	err = f5.CreateExtension(ctx, opts)
	if err != nil {
		panic(err)
	}

	// Read the package.json file.
	const packagePath string = "cmd/ilx_example/ilx/package.json"
	packagejson, err := os.ReadFile(packagePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// Write the package.json file to the extension.
	err = f5.WriteExtensionFile(ctx, opts, string(packagejson), bigip.PackageJSON)
	if err != nil {
		panic(err)
	}

	// Read the index.js file.
	const indexjsPath string = "cmd/ilx_example/ilx/index.js"
	indexjs, err := os.ReadFile(indexjsPath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// Write the index.js file to the extension.
	err = f5.WriteExtensionFile(ctx, opts, string(indexjs), bigip.IndexJS)
	if err != nil {
		panic(err)
	}

	// Read the index.js file from the extension.
	content, err := f5.ReadExtensionFile(ctx, opts, bigip.IndexJS)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Content: %+v\n", content)

	// Write a TCL rule file to the workspace.
	err = f5.WriteRuleFile(ctx, workspaceConfig, "<sometcl>", "Example.tcl")
	if err != nil {
		panic(err)
	}

	// Read the TCL rule file from the workspace.
	out, err := f5.ReadRuleFile(ctx, workspaceConfig, "Example.tcl")
	if err != nil {
		panic(err)
	}
	log.Printf("Rule: %v", out.Content)
}
