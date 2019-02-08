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
		log.Fatal(err)
	}

	orgID, err := nks.GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		log.Fatal(err)
	}

	// Get list of configured clusters
	clusters, err := client.GetClusters(orgID)
	if err != nil {
		log.Fatal(err)
	}

	// Print list of clusters
	for i := 0; i < len(clusters); i++ {
		fmt.Printf("Cluster(%d): %v\n", clusters[i].ID, clusters[i].Name)
	}
	if len(clusters) == 0 {
		fmt.Println("Sorry, no clusters defined yet")
		return
	}
	// Get cluster ID from user to list nodes from
	var clusterID int
	fmt.Printf("Enter cluster ID to list nodes from: ")
	fmt.Scanf("%d", &clusterID)

	// Get list of buildlog events
	bls, err := client.GetBuildLogs(orgID, clusterID)
	if err != nil {
		log.Fatal(err)
	}

	// List events
	for i := 0; i < len(bls); i++ {
		fmt.Printf("BuildLog Entry(%d): %s=%s\n", bls[i].ID, bls[i].EventType, bls[i].EventState)
	}
	if len(bls) == 0 {
		fmt.Printf("Sorry, no build logs found\n")
		return
	}
	// Get buildlog ID from user to inspect
	var buildlogID int
	fmt.Printf("Enter buildlog ID to inspect: ")
	fmt.Scanf("%d", &buildlogID)

	buildlog, err := client.GetBuildLog(bls[:], buildlogID)
	if err != nil {
		log.Fatal(err)
	}
	nks.PrettyPrint(buildlog)
}
