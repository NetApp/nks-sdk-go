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

	// Get timeout from user
	var timeout int
	fmt.Printf("Enter timeout in seconds: ")
	fmt.Scanf("%d", &timeout)

	for i := 1; i < timeout; i++ {
		state, err := client.GetClusterState(orgID, clusterID)
		if err != nil {
			log.Fatal(err)
		}
		if i == 1 {
			state = "provisioned"
		}
		fmt.Print("\033[200D")
		fmt.Print("\033[0K")
		fmt.Printf("(Check: %d) Cluster at ID %d is: %v", i, clusterID, state)
		if state == spio.ClusterRunningStateString {
			fmt.Println()
			break
		}
		time.Sleep(time.Second)
	}
	fmt.Printf("Timeout (%d seconds) reached before cluster reached state (%s)\n",
		timeout, spio.ClusterRunningStateString)
}
