package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"log"
)

const nodepoolName = "Test Nodepool"

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
	providers := make(map[int]string)
	for i := 0; i < len(clusters); i++ {
		fmt.Printf("Cluster(%d): %v\n", clusters[i].ID, clusters[i].Name)
		providers[clusters[i].ID] = clusters[i].Provider
	}
	if len(clusters) == 0 {
		fmt.Println("Sorry, no clusters defined yet")
		return
	}
	// Get cluster ID from user to add node to
	var clusterID int
	fmt.Printf("Enter cluster ID to add node to: ")
	fmt.Scanf("%d", &clusterID)

	// Get number of worker nodes to add to nodepool
	var nodeCount int
	fmt.Printf("Enter number of worker nodes to add into pool: ")
	fmt.Scanf("%v", &nodeCount)

	if nodeCount < 1 {
		log.Fatal("You must add at least one node to the new nodepool")
	}
	// Get list of instance types for provider
	mOptions, err := client.GetInstanceSpecs(providers[clusterID], "")
	if err != nil {
		log.Fatal(err.Error())
	}

	// List instance types
	fmt.Printf("Node size options for provider %s:\n", providers[clusterID])
	for _, opt := range spio.GetFormattedInstanceList(mOptions) {
		fmt.Println(opt)
	}
	// Get node size selection from user
	var nodeSize string
	fmt.Printf("Enter node size: ")
	fmt.Scanf("%s", &nodeSize)

	newNodepool := spio.NodePool{
		Name:      nodepoolName,
		NodeCount: nodeCount,
		Size:      nodeSize,
		Platform:  "coreos",
	}

	// Create new nodepool
	pool, err := client.CreateNodePool(orgID, clusterID, newNodepool)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("NodePool creation sent (ID: %d), building...\n", pool.ID)
}
