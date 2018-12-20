package main

import (
	"fmt"
	"log"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
)

const (
	awsZone       = "us-west-2a"
	awsSubnetID   = "__new__" // replace with your AWS subnet ID if you have one
	awsSubnetCidr = "172.23.5.0/24"
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

	// Get list of instance types for provider
	mOptions, err := client.GetInstanceSpecs(providers[clusterID], "")
	if err != nil {
		log.Fatal(err.Error())
	}

	// List instance types
	fmt.Printf("Node size options for provider %s:\n", providers[clusterID])
	for _, opt := range nks.GetFormattedInstanceList(mOptions) {
		fmt.Println(opt)
	}

	// Set up new master node
	newNode := nks.NodeAdd{
		Count:              1,
		Role:               "master",
		Zone:               awsZone,
		ProviderSubnetID:   awsSubnetID,
		ProviderSubnetCidr: awsSubnetCIDR,
		Size:               nodeSize,
	}

	// Add new node
	nodes, err := client.AddNode(orgID, clusterID, newNode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Master node creation sent (ID: %d), building...\n", nodes[0].ID)
}
