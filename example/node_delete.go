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
	// Get cluster ID from user to delete node from
	var clusterID int
	fmt.Printf("Enter cluster ID to delete node from: ")
	fmt.Scanf("%d", &clusterID)

	// Get list of nodes configured
	nodes, err := client.GetNodes(orgID, clusterID)
	if err != nil {
		log.Fatal(err.Error())
	}

	// List nodes
	for i := 0; i < len(nodes); i++ {
		fmt.Printf("Nodes(%d): %s node is %s\n", nodes[i].ID, nodes[i].Role, nodes[i].State)
	}
	if len(nodes) == 0 {
		fmt.Printf("Sorry, no nodes found\n")
		return
	}
	// Get node ID to delete from user
	var nodeID int
	fmt.Printf("Enter node ID to delete: ")
	fmt.Scanf("%d", &nodeID)

	if err = client.DeleteNode(orgID, clusterID, nodeID); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Node should delete shortly")
}
