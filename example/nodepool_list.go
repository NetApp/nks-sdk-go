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

	// Print list of clusters, saving map of providers for later use
	for i := 0; i < len(clusters); i++ {
		fmt.Printf("Cluster(%d): %v\n", clusters[i].ID, clusters[i].Name)
	}
	if len(clusters) == 0 {
		fmt.Println("Sorry, no clusters defined yet")
		return
	}
	// Get cluster ID from user to list nodepools from
	var clusterID int
	fmt.Printf("Enter cluster ID to list nodepools from: ")
	fmt.Scanf("%d", &clusterID)

	// Get list of nodepools to select from
	nps, err := client.GetNodePools(orgID, clusterID)
	if err != nil {
		log.Fatal(err.Error())
	}

	// List nodepools
	for i := 0; i < len(nps); i++ {
		fmt.Printf("Nodepool(%d): %v (node count: %d)\n", nps[i].ID, nps[i].Name, nps[i].NodeCount)
	}
	if len(nps) == 0 {
		fmt.Println("Sorry, no nodepools found")
		return
	}
	// Get nodepool ID from user
	var nodepoolID int
	fmt.Printf("Enter nodepool ID to inspect: ")
	fmt.Scanf("%d", &nodepoolID)

	nodepool, err := client.GetNodePool(orgID, clusterID, nodepoolID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s\nInstanceID: %s\n", nodepool.Name, nodepool.InstanceID)
	spio.PrettyPrint(nodepool)
}
