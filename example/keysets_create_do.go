package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"log"
)

const (
	keysetName = "Test DO Keyset"
	provider   = "do"
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

	// Gather access token for DO
	var doToken string
	fmt.Printf("Enter DigitalOcean Token: ")
	fmt.Scanf("%s", &doToken)

	newKey := spio.Key{Type: "token",
		Value: doToken}
	newKeyset := spio.Keyset{Name: keysetName,
		Category:   "provider",
		Entity:     provider,
		Workspaces: []int{},
		Keys:       []spio.Key{newKey}}

	keyset, err := client.CreateKeyset(orgID, newKeyset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateKeyset created,")
        spio.PrettyPrint(keyset)
}
