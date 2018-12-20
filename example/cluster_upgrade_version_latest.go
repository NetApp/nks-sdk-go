package main

import (
	"fmt"
	"log"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
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

	// Fetch list of configured clusters
	clusters, err := client.GetClusters(orgID)
	if err != nil {
		log.Fatal(err.Error())
	}

	// List clusters
	for i := 0; i < len(clusters); i++ {
		fmt.Printf("Clusters(%d): %v\n", clusters[i].ID, clusters[i].Name)
	}
	if len(clusters) == 0 {
		fmt.Println("Sorry, no clusters defined yet")
		return
	}

	// Get cluster ID to upgrade from user
	var clusterID int
	fmt.Printf("Enter cluster ID to upgrade: ")
	fmt.Scanf("%d", &clusterID)

	// Find cluster with ID from list, get index of it
	cIndex := -1
	for i := 0; i < len(clusters); i++ {
		if clusters[i].ID == clusterID {
			cIndex = i
			break
		}
	}
	if cIndex == -1 {
		log.Fatalf("Invalid cluster ID entered: %i\n", clusterID)
	}
	// Show versions available to upgrade to
	fmt.Printf("Cluster's current kubernetes version is: %s\n", clusters[cIndex].KubernetesVersion)
	for i := 0; i < len(clusters[cIndex].KubernetesMigrationVersions); i++ {
		fmt.Printf("Version available: %v\n", clusters[cIndex].KubernetesMigrationVersions[i])
	}
	if len(clusters[cIndex].KubernetesMigrationVersions) == 0 {
		fmt.Println("Sorry, no upgrade versions available")
		return
	}
	// Get version to upgrade to from user
	fmt.Println("Upgrading cluster to latest kubernetes version...")

	if err := client.UpgradeClusterToLatestVersion(clusters[cIndex]); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster should upgrade shortly\n")
}
