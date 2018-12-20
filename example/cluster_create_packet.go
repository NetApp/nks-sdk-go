package main

import (
	"fmt"
	"log"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
)

const (
	provider    = "packet"
	clusterName = "Test Packet Cluster Go SDK"
	projectID   = "93125c2a-8b78-4d4f-a3c4-7367d6b7cca8"
	region      = "sjc1"
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

	sshKeysetID, err := nks.GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		log.Fatal(err.Error())
	}

	pktKeysetID, err := nks.GetIDFromEnv("NKS_PKT_KEYSET")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get list of instance types for provider
	mOptions, err := client.GetInstanceSpecs(provider, "")
	if err != nil {
		log.Fatal(err.Error())
	}

	// List instance types
	fmt.Printf("Node size options for provider %s:\n", provider)
	for _, opt := range nks.GetFormattedInstanceList(mOptions) {
		fmt.Println(opt)
	}

	// Get node size selection from user
	var nodeSize string
	fmt.Printf("Enter node size: ")
	fmt.Scanf("%s", &nodeSize)

	// Validate machine type selection
	if !nks.InstanceInList(mOptions, nodeSize) {
		log.Fatalf("Invalid option: %s\n", nodeSize)
	}

	newSolution := nks.Solution{Solution: "helm_tiller"}
	newCluster := nks.Cluster{Name: clusterName,
		Provider:          provider,
		ProviderKey:       pktKeysetID,
		ProjectID:         projectID,
		MasterCount:       1,
		MasterSize:        nodeSize,
		WorkerCount:       2,
		WorkerSize:        nodeSize,
		Region:            region,
		KubernetesVersion: "v1.8.7",
		RbacEnabled:       true,
		DashboardEnabled:  true,
		EtcdType:          "classic",
		Platform:          "coreos",
		Channel:           "stable",
		SSHKeySet:         sshKeysetID,
		Solutions:         []nks.Solution{newSolution}}
	cluster, err := client.CreateCluster(orgID, newCluster)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster created (ID: %d) (instance name: %s), building...\n", cluster.ID, cluster.InstanceID)
}
