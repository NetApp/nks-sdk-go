package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"io/ioutil"
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

	// Get solution JSON file from user
	var jsonFile string
	fmt.Printf("Enter JSON filename for solution to add: ")
	fmt.Scanf("%s", &jsonFile)

	// Read file contents, if they exist
	solutionJSON, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	// Add new solution using JSON config
	solution, err := client.AddSolutionFromJSON(orgID, clusterID, solutionJSON)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution creation sent (ID: %d), building...\n", solution.ID)
}
