package main

import (
	"fmt"
	"log"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
)

const (
	keysetName = "Test 1&1 Keyset"
	provider   = "oneandone"
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

	// Gather access token for 1&1
	var oneandoneToken string
	fmt.Printf("Enter 1&1 Token: ")
	fmt.Scanf("%s", &oneandoneToken)

	newKey := nks.Key{
		Type:  "token",
		Value: oneandoneToken,
	}
	newKeyset := nks.Keyset{
		Name:       keysetName,
		Category:   "provider",
		Entity:     provider,
		Workspaces: []int{},
		Keys:       []nks.Key{newKey},
	}

	keyset, err := client.CreateKeyset(orgID, newKeyset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateKeyset created,")
	nks.PrettyPrint(keyset)
}
