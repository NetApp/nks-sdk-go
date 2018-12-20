package main

import (
	"fmt"
	"log"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
)

func main() {
	// Set up HTTP client with environment variables for API token and URL
	client, err := nks.NewClientFromEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get userprofile
	up, err := client.GetUserProfile()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get default org ID
	orgID, err := client.GetUserProfileDefaultOrg(&up[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Org ID found is: %d\n", orgID)

	// Loop through supported providers, see if a key can be fetched for any of them
	for _, prov := range []string{"aws", "azure", "do", "gce", "gke", "oneandone", "packet", "user_ssh"} {
		ksid, _ := client.GetUserProfileKeysetID(&up[0], prov)
		fmt.Printf("Keyset ID found for %s is: %d\n", prov, ksid)
	}
}
