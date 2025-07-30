// Package main provides the entry point for the test client application.
package main

import (
	"fmt"
	"os"

	"car-price-prediction/cmd/test_client/testclient"
)

func main() {
	// Check if the user provided a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/test_client/main_runner.go [minimal|full]")
		return
	}

	// Get the client type from the command-line argument
	clientType := os.Args[1]

	// Run the appropriate client
	switch clientType {
	case "minimal":
		testclient.MinimalClient()
	case "full":
		testclient.FullClient()
	default:
		fmt.Println("Invalid client type. Use 'minimal' or 'full'.")
	}
}