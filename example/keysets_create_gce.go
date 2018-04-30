package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"io/ioutil"
	"log"
)

const (
	keysetName = "Test GCE Keyset"
	provider   = "gce"
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

	// Gather user file where JSON credential file is located
	var jsonFile string
	fmt.Printf("Enter filename for JSON credentials: ")
	fmt.Scanf("%s", &jsonFile)

	// Read file contents, if they exist
	credentialJSON, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	jsonKey := spio.Key{Type: "other",
		Value: string(credentialJSON[:])}
	newKeyset := spio.Keyset{Name: keysetName,
		Category:   "provider",
		Entity:     provider,
		Workspaces: []int{},
		Keys:       []spio.Key{jsonKey}}

        keyset, err := client.CreateKeyset(orgID, newKeyset)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("CreateKeyset created,")
        spio.PrettyPrint(keyset)
}
