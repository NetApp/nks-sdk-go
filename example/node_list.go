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
	// Get cluster ID from user to list nodes from
	var clusterID int
	fmt.Printf("Enter cluster ID to list nodes from: ")
	fmt.Scanf("%d", &clusterID)

	// Get list of nodes configured
	nodes, err := client.GetNodes(orgID, clusterID)
	if err != nil {
		log.Fatal(err.Error())
	}

	// List nodes
	for i := 0; i < len(nodes); i++ {
		fmt.Printf("Node(%d): %s node is %s", nodes[i].ID, nodes[i].Role, nodes[i].State)
		if nodes[i].Role == "worker" {
			fmt.Printf(", in NodePool(%d) %s", nodes[i].NodePoolID, nodes[i].NodePoolName)
		}
		fmt.Println()
	}
	if len(nodes) == 0 {
		fmt.Printf("Sorry, no nodes found\n")
		return
	}
	// Get node ID from user to inspect
	var nodeID int
	fmt.Printf("Enter node ID to inspect: ")
	fmt.Scanf("%d", &nodeID)

	node, err := client.GetNode(orgID, clusterID, nodeID)
	if err != nil {
		log.Fatal(err)
	}
	spio.PrettyPrint(node)
}
