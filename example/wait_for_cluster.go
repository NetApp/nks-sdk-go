package main

import (
	"fmt"
	"log"
	"time"

	nks "github.com/StackPointCloud/nks-sdk-go/nks"
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

	// Get timeout from user
	var timeout int
	fmt.Printf("Enter timeout in seconds: ")
	fmt.Scanf("%d", &timeout)

	// Use WaitClusterProvisioned function, but run it in separate routine, waitSuccessCh set when finished
	waitSuccessCh := make(chan bool)
	waitResultCh := make(chan string)
	go func() {
		err = client.WaitClusterRunning(orgID, clusterID, true, timeout)
		if err != nil {
			waitSuccessCh <- false
			waitResultCh <- err.Error()
		}
		waitSuccessCh <- true
		waitResultCh <- "Cluster provisioned successfully"
	}()

	// Enter loop for main routine, checking cluster status
	for i := 1; ; i++ {
		cl, err := client.GetCluster(orgID, clusterID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("\033[200D")
		fmt.Print("\033[0K")
		fmt.Printf("(Check: %d) Cluster at ID %d is: %v", i, clusterID, cl.State)

		// Check channel to see if routine is finished
		select {
		case status := <-waitSuccessCh:
			if status {
				fmt.Printf("\nSUCCESS: %s\n", <-waitResultCh)
			} else {
				log.Fatalf("\n %s\n", <-waitResultCh)
			}
			return
		default:
			time.Sleep(time.Second)
		}
	}
}
