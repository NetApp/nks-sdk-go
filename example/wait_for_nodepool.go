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

        for i := 1; ; i++ {
                isActive, err := client.IsNodePoolActive(orgID, clusterID, nodepoolID)
                if err != nil {
                        log.Fatal(err)
                }
                fmt.Print("\033[200D")
                fmt.Printf("(Try: %d) Nodepool at ID %d is active: %v", i, nodepoolID, isActive)
                if isActive {
			fmt.Println()
                        break
                }
                time.Sleep(time.Second)
        }
}
