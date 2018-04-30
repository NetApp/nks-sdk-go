package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"log"
)

const (
	provider    = "gke"
	clusterName = "Test GKE Cluster Go SDK"
	region      = "us-west1-a"
)

func main() {
	// Set up HTTP client with environment variables for API token and URL
	client, err := spio.NewClientFromEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

        orgid, err := spio.GetIDFromEnv("SPC_ORG_ID")
        if err != nil {
                log.Fatal(err.Error())
        }

	sshKeysetid, err := spio.GetIDFromEnv("SPC_SSH_KEYSET")
	if err != nil {
		log.Fatal(err.Error())
	}

	gkeKeysetid, err := spio.GetIDFromEnv("SPC_GKE_KEYSET")
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
		Provider:          provider,
		ProviderKey:       gkeKeysetid,
		MasterCount:       1,
		MasterSize:        nodeSize,
		WorkerCount:       2,
		WorkerSize:        nodeSize,
		Region:            region,
		KubernetesVersion: "latest",
		RbacEnabled:       true,
		DashboardEnabled:  true,
		EtcdType:          "self_hosted",
		Platform:          "gci",
		Channel:           "stable",
		SSHKeySet:         sshKeysetid,
		Solutions:         []spio.Solution{newSolution}}

	cluster, err := client.CreateCluster(orgid, newCluster)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster created (ID: %d) (instance name: %s), building...\n", cluster.ID, cluster.InstanceID)
}
