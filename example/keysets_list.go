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

        orgID, err := spio.GetIDFromEnv("SPC_ORG_ID")
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
	fmt.Printf("Enter keyset ID to inspect: ")
	fmt.Scanf("%d", &keysetID)

	keyset, err := client.GetKeyset(orgID, keysetID)
	if err != nil {
		log.Fatal(err.Error())
	}
	spio.PrettyPrint(keyset)
}
