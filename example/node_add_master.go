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

	// Get list of machine types for provider
	mOptions, err := client.GetMachSpecs(providers[clusterID])
	if err != nil {
		log.Fatal(err.Error())
	}

	// List machine types
	fmt.Printf("Node size options for provider %s:\n", providers[clusterID])
	for _, opt := range mOptions {
		fmt.Println(opt)
	}
	// Get node size selection from user
	var nodeSize string
	fmt.Printf("Enter node size: ")
	fmt.Scanf("%s", &nodeSize)

	// Validate machine type selection
	if !spio.StringInSlice(nodeSize, mOptions) {
		log.Fatalf("Invalid option: %s\n", nodeSize)
	}

	// Set up new master node
	newNode := spio.NodeAdd{
		Count: 1,
		Role:  "master",
		Size:  nodeSize}

	// Add new node
	nodes, err := client.AddNode(orgID, clusterID, newNode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Master node creation sent (ID: %d), building...\n", nodes[0].ID)
}
