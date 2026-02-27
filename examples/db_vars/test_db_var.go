package main

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
)

/////////////////////////////////////////////////////////////////////////////////
//////////////              Database Variable Testing              //////////////
/////////////////////////////////////////////////////////////////////////////////

func testDBVariableRead(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing System DB Variable Read Operation ===")

	// Test DB variable configuration
	dbVar, err := f5.GetDBVariable("ui.advisory.enabled")
	if err != nil {
		log.Printf("Error getting Database Variable: %v", err)
		return
	}
	if dbVar != nil {
		fmt.Printf("Retrieved Database Variable:\n")
		fmt.Printf("  Name: %s\n", dbVar.Name)
		fmt.Printf("  Value: %s\n", dbVar.Value)
	}
}

func testDBVariableCreation(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing System DB Variable Creation Operation ===")

	// Test DB variable configuration
	dbVar := &bigip.DBVariable{
		Name:  "ui.advisory.enabled",
		Value: "true",
	}

	// CREATE: Set DB Variable value
	fmt.Println("Setting System DB Variable...")
	err := f5.CreateDBVariable(dbVar)
	if err != nil {
		log.Printf("Error creating DB Variable: %v", err)
		return
	}
	fmt.Printf("DB Variable ui.advisory.enabled changed successfully.\n")

}

func testDBVariableUpdate(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing System DB Variable Modify Operation ===")

	// Test DB variable configuration
	dbVar := &bigip.DBVariable{
		Name:  "ui.advisory.color",
		Value: "orange",
	}

	// CREATE: Set DB Variable value
	fmt.Println("Modifying System DB Variable...")
	err := f5.ModifyDBVariable(dbVar)
	if err != nil {
		log.Printf("Error modifying DB Variable: %v", err)
		return
	}
	fmt.Printf("DB Variable ui.advisory.color changed successfully.\n")

}

func testDBVariableDelete(f5 *bigip.BigIP) {
	fmt.Println("\n=== Testing DB Variable Removal/Reset-To-Defaults Operation ===")
	// Test DB variable configuration
	dbVar := &bigip.DBVariable{
		Name: "ui.advisory.enabled",
	}
	// DELETE: Reset DB Variable to default
	fmt.Println("Resetting DB Variable to default...")
	err := f5.DeleteDBVariable(dbVar.Name)
	if err != nil {
		log.Printf("Error resetting DB Variable to defaults: %v", err)
		return
	}
	fmt.Printf("DB Variable %s was set back to default value '%v' successfully.\n", dbVar.Name, bigip.DefaultDBValues[dbVar.Name])

}
