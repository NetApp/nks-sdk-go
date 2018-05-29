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

	// Get list of configured clusters
	clusters, err := client.GetClusters(orgID)
	if err != nil {
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
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

	buildlog, err := client.GetBuildLog(orgID, clusterID, buildlogID)
	if err != nil {
		log.Fatal(err)
	}
	spio.PrettyPrint(buildlog)
}
