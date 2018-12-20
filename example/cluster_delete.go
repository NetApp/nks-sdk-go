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

	// Get cluster ID to delete from user
	var clusterID int
	fmt.Printf("Enter cluster ID to delete: ")
	fmt.Scanf("%d", &clusterID)

	if err := client.DeleteCluster(orgID, clusterID); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Cluster should delete shortly\n")
}
