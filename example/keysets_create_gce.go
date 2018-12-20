package main

import (
	"fmt"
	"io/ioutil"
	"log"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
)

const (
	keysetName = "Test GCE Keyset"
	provider   = "gce"
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

	// Gather user file where JSON credential file is located
	var jsonFile string
	fmt.Printf("Enter filename for JSON credentials: ")
	fmt.Scanf("%s", &jsonFile)

	// Read file contents, if they exist
	credentialJSON, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	jsonKey := nks.Key{Type: "other",
		Value: string(credentialJSON[:])}
	newKeyset := nks.Keyset{Name: keysetName,
		Category:   "provider",
		Entity:     provider,
		Workspaces: []int{},
		Keys:       []nks.Key{jsonKey}}

	keyset, err := client.CreateKeyset(orgID, newKeyset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateKeyset created,")
	nks.PrettyPrint(keyset)
}
