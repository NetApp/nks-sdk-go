package main

import (
	"fmt"
	"log"
	"time"

	nks "github.com/NetApp/nks-sdk-go/nks"
)

func main() {
	// Set up HTTP client with environment variables for API token and URL
	client, err := nks.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	orgID, err := nks.GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		log.Fatal(err)
	}

	// Get cluster ID from user
	var clusterID int
	fmt.Printf("Enter cluster ID: ")
	fmt.Scanf("%d", &clusterID)

	// Get event type to watch
	var eventType string
	fmt.Printf("Enter event type: ")
	fmt.Scanf("%s", &eventType)

	// Get event state to watch for
	var eventState string
	fmt.Printf("Enter event state: ")
	fmt.Scanf("%s", &eventState)

	// Get timeout from user
	var timeout int
	fmt.Printf("Enter timeout in seconds: ")
	fmt.Scanf("%d", &timeout)

	for i := 1; i < timeout; i++ {
		bls, err := client.GetBuildLogs(orgID, clusterID)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(bls); i++ {
			fmt.Printf("BuildLog Entry(%d): %s=%s\n", bls[i].ID, bls[i].EventType, bls[i].EventState)
			if bls[i].EventType == eventType {
				fmt.Printf("...eventType %s is currently at %s state\n", bls[i].EventType, bls[i].EventState)
				if bls[i].EventState == eventState {
					return
				}
			}
		}
		if len(bls) == 0 {
			fmt.Printf("Sorry, no build logs found\n")
			return
		}
		time.Sleep(time.Second)
	}
	fmt.Printf("Timeout (%d seconds) reached before build reached state success state\n",
		timeout)
}
