package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"log"
)

const (
	keysetName = "Test 1&1 Keyset"
	provider   = "oneandone"
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

	// Gather access token for 1&1
	var oneandoneToken string
	fmt.Printf("Enter 1&1 Token: ")
	fmt.Scanf("%s", &oneandoneToken)

	newKey := spio.Key{
		Type:  "token",
		Value: oneandoneToken,
	}
	newKeyset := spio.Keyset{
		Name:       keysetName,
		Category:   "provider",
		Entity:     provider,
		Workspaces: []int{},
		Keys:       []spio.Key{newKey},
	}

	keyset, err := client.CreateKeyset(orgID, newKeyset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateKeyset created,")
	spio.PrettyPrint(keyset)
}
