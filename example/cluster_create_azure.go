package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"log"
)

const (
	provider      = "azure"
	clusterName   = "Test Azure Cluster Go SDK"
	region        = "eastus"
	resourceGroup = "__new__"       // Azure creates network subsystems inside of a resource group or `__new__`
	networkID     = "__new__"       // ID of existing Azure virtual network or `__new__`
	networkCIDR   = "10.0.0.0/16" // CIDR for a new network or CIDR of the existing network
	subnetID      = "__new__"       // CIDR for an existing subnet in specified network or `__new__`
	subnetCIDR    = "10.0.0.0/24" // CIDR for a new subnet or CIDR of the existing subnet
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

	sshKeysetID, err := spio.GetIDFromEnv("SPC_SSH_KEYSET")
	if err != nil {
		log.Fatal(err.Error())
	}

	azrKeysetID, err := spio.GetIDFromEnv("SPC_AZR_KEYSET")
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
	for _, opt := range spio.GetFormattedInstanceList(mOptions) {
		fmt.Println(opt)
	}

	// Get node size selection from user
	var nodeSize string
	fmt.Printf("Enter node size: ")
	fmt.Scanf("%s", &nodeSize)

	// Validate machine type selection
	if !spio.InstanceInList(mOptions, nodeSize) {
		log.Fatalf("Invalid option: %s\n", nodeSize)
	}

	newSolution := spio.Solution{Solution: "helm_tiller"}
	newCluster := spio.Cluster{Name: clusterName,
		Provider:           provider,
		ProviderKey:        azrKeysetID,
		MasterCount:        1,
		MasterSize:         nodeSize,
		WorkerCount:        2,
		WorkerSize:         nodeSize,
		Region:             region,
		ProviderResourceGp: resourceGroup,
		ProviderNetworkID:  networkID,
		ProviderNetworkCdr: networkCIDR,
		ProviderSubnetID:   subnetID,
		ProviderSubnetCidr: subnetCIDR,
		KubernetesVersion:  "v1.10.4",
		RbacEnabled:        true,
		DashboardEnabled:   true,
		EtcdType:           "classic",
		Platform:           "coreos",
		Channel:            "stable",
		SSHKeySet:          sshKeysetID,
		Solutions:          []spio.Solution{newSolution}}
	cluster, err := client.CreateCluster(orgID, newCluster)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster created (ID: %d) (instance name: %s), building...\n", cluster.ID, cluster.InstanceID)

}
