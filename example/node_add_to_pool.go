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
	fmt.Printf("Enter nodepool ID to add node to: ")
	fmt.Scanf("%d", &nodepoolID)

	// Get number of nodes to add from user
	var nodeCount int
	fmt.Printf("Enter number of nodes to add: ")
	fmt.Scanf("%d", &nodeCount)

	// Use NodeAddToPool struct for this endpoint
	newNode := nks.NodeAddToPool{
		Count:      nodeCount,
		Role:       "worker",
		NodePoolID: nodepoolID,
	}

	nodes, err := client.AddNodesToNodePool(orgID, clusterID, nodepoolID, newNode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Node creation sent,")
	for _, n := range nodes {
		fmt.Printf("ID: %d, InstanceID: %s", n.ID, n.InstanceID)
	}
	fmt.Println("...building...")
}
