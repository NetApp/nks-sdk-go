package main

import (
	"fmt"
	spio "github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"log"
	"time"
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

	// Get cluster ID from user
	var clusterID int
	fmt.Printf("Enter cluster ID: ")
	fmt.Scanf("%d", &clusterID)

	// Get nodepool ID from user
	var nodepoolID int
	fmt.Printf("Enter nodepool ID: ")
	fmt.Scanf("%d", &nodepoolID)

	// Get timeout from user
	var timeout int
	fmt.Printf("Enter timeout in seconds: ")
	fmt.Scanf("%d", &timeout)

	for i := 1; i < timeout; i++ {
		state, err := client.GetNodePoolState(orgID, clusterID, nodepoolID)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Print("\033[200D")
		fmt.Print("\033[0K")
		fmt.Printf("(Try: %d) Nodepool at ID %d is active: %v", i, nodepoolID, state)
		if state == spio.NodePoolRunningStateString {
			fmt.Println()
			log.Fatal(err.Error())
		}
		time.Sleep(time.Second)
	}
	fmt.Printf("Timeout (%d seconds) reached before nodepool reached state (%s)\n",
		timeout, spio.NodePoolRunningStateString)
}
