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

	// Get nodepool ID from user
	var nodepoolID int
	fmt.Printf("Enter nodepool ID: ")
	fmt.Scanf("%d", &nodepoolID)

	// Get timeout from user
	var timeout int
	fmt.Printf("Enter timeout in seconds: ")
	fmt.Scanf("%d", &timeout)

	// Use WaitNodePoolProvisioned function, but run it in separate routine, waitDoneCh set to true when finished
	waitDoneCh := make(chan bool)
	waitSuccessCh := make(chan bool)
	waitResultCh := make(chan string)
	go func() {
		err = client.WaitNodePoolProvisioned(orgID, clusterID, nodepoolID, timeout)
		waitDoneCh <- true
		if err != nil {
			waitSuccessCh <- false
			waitResultCh <- err.Error()
		}
		waitSuccessCh <- true
		waitResultCh <- "NodePool provisioned successfully"
	}()

	// Enter loop for main routine, checking cluster status
	for i := 1; ; i++ {
		np, err := client.GetNodePool(orgID, clusterID, nodepoolID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("\033[200D")
		fmt.Print("\033[0K")
		fmt.Printf("(Check: %d) NodePool at ID %d is: %v", i, nodepoolID, np.State)

		// Check channel to see if routine is done waiting
		if true == <-waitDoneCh {
			if true == <-waitSuccessCh {
				fmt.Printf("\nSUCCESS: %s\n", <-waitResultCh)
			} else {
				log.Fatalf("\n %s\n", <-waitResultCh)
			}
			break
		}
		time.Sleep(time.Second)
	}
}
