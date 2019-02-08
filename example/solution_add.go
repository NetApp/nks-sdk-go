package main

import (
	"fmt"
	"log"

	nks "github.com/NetApp/nks-sdk-go/nks"
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

	// Set up new solution
	newSolution := nks.Solution{
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
