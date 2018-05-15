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
	// Get cluster ID from user to add solution to
	var clusterID int
	fmt.Printf("Enter cluster ID to add solution to: ")
	fmt.Scanf("%d", &clusterID)

	// Get solution selection from user
	var solutionName string
	fmt.Printf("Enter solution to add: ")
	fmt.Scanf("%s", &solutionName)

	// Set up new master node
	newSolution := spio.Solution{
		Solution: solutionName,
		State:    "draft",
	}
	// Add new solution
	solution, err := client.AddSolution(orgID, clusterID, newSolution)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution creation sent (ID: %d), building...\n", solution.ID)
}
