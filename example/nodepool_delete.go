package main

import (
	"fmt"
	"log"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
)

const nodepoolName = "Test Nodepool"

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
	// Get nodepool ID from user to delete
	var nodepoolID int
	fmt.Printf("Enter nodepool ID to delete: ")
	fmt.Scanf("%d", &nodepoolID)

	// Create new nodepool
	err = client.DeleteNodePool(orgID, clusterID, nodepoolID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nodepool should delete shortly")
}
