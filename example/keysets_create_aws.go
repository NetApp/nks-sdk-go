package main

import (
	"fmt"
	"log"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
)

const (
	keysetName = "Test AWS Keyset"
	provider   = "aws"
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

	// Gather access key for AWS
	var awsAccessKey string
	fmt.Printf("Enter AWS Access Key ID: ")
	fmt.Scanf("%s", &awsAccessKey)

	// Gather secret key for AWS
	var awsSecretKey string
	fmt.Printf("Enter AWS Secret Access Key: ")
	fmt.Scanf("%s", &awsSecretKey)

	pubKey := nks.Key{Type: "pub",
		Value: awsAccessKey}
	pvtKey := nks.Key{Type: "pvt",
		Value: awsSecretKey}
	newKeyset := nks.Keyset{Name: keysetName,
		Category:   "provider",
		Entity:     provider,
		Workspaces: []int{},
		Keys:       []nks.Key{pubKey, pvtKey}}

	keyset, err := client.CreateKeyset(orgID, newKeyset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CreateKeyset created,")
	nks.PrettyPrint(keyset)
}
