package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"log"
)

const (
	provider       = "aws"
	clusterName    = "Test AWS Cluster Go SDK"
	awsRegion      = "us-west-2"
	awsZone        = "us-west-2b"
	awsNetworkID   = "__new__"   // replace with your AWS VPC ID if you have one
	awsNetworkCidr = "172.23.0.0/16"
	awsSubnetID    = "__new__"   // replace with your AWS subnet ID if you have one
	awsSubnetCidr  = "172.23.1.0/24"
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

	awsKeysetID, err := spio.GetIDFromEnv("SPC_AWS_KEYSET")
	if err != nil {
		log.Fatal(err.Error())
	}

        // Get list of instance types for provider
        mOptions, err := client.GetInstanceSpecs(provider)
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
		ProviderKey:        awsKeysetID,
		MasterCount:        1,
		MasterSize:         nodeSize,
		WorkerCount:        2,
		WorkerSize:         nodeSize,
		Region:             awsRegion,
		Zone:               awsZone,
		ProviderNetworkID:  awsNetworkID,
		ProviderNetworkCdr: awsNetworkCidr,
		ProviderSubnetID:   awsSubnetID,
		ProviderSubnetCidr: awsSubnetCidr,
		KubernetesVersion:  "v1.8.7",
		RbacEnabled:        true,
		DashboardEnabled:   true,
		EtcdType:           "self_hosted",
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
