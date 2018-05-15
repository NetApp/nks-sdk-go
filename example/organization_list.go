package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"log"
)

func main() {
	// Set up HTTP client with environment variables for API token and URL
	client, err := spio.NewClientFromEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get list of configured organizations
	orgs, err := client.GetOrganizations()
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(orgs) == 0 {
		fmt.Println("Sorry, no organizations defined yet")
		return
	}

	// Print list of organizations
	for i := 0; i < len(orgs); i++ {
		fmt.Printf("Org(%d): %v\n", orgs[i].ID, orgs[i].Name)
	}

	// Get organization ID from user to inspect
	var orgID int
	fmt.Printf("Enter org ID to inspect: ")
	fmt.Scanf("%d", &orgID)

	org, err := client.GetOrganization(orgID)
	if err != nil {
		log.Fatal(err)
	}
	spio.PrettyPrint(org)
}
