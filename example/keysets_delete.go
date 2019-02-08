package main

import (
	"fmt"
	"log"

	nks "github.com/NetApp/nks-sdk-go/nks"
)

func main() {
	// Set up HTTP client with environment variables for API token and URL
	client, err := nks.NewClientFromEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	orgID, err := nks.GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Gather list of keysets
	keysets, err := client.GetKeysets(orgID)
	if err != nil {
		log.Fatal(err.Error())
	}

	// List keysets configured
	for i := 0; i < len(keysets); i++ {
		fmt.Printf("Keysets(%d): %v\n", keysets[i].ID, keysets[i].Name)
	}
	// Get keyset ID to inspect from user
	var keysetID int
	fmt.Printf("Enter keyset ID to delete: ")
	fmt.Scanf("%d", &keysetID)

	// Do keyset ID deletion
	if err = client.DeleteKeyset(orgID, keysetID); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Keyset should delete shortly\n")
}
