package main

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/StackPointCloud/stackpoint-sdk-go"
)

func main() {

	token := os.Getenv("CLUSTER_API_TOKEN")
	endpoint := os.Getenv("SPC_BASE_API_URL")
	client := stackpointio.NewClient(token, endpoint)

	orgs, err := client.GetOrganizations()
	if err != nil {
		log.Fatal(err)
	}

	for _, org := range orgs {
		fmt.Printf("%s (%d)\n", org.Name, org.PrimaryKey)
	}

	text, _ := json.MarshalIndent(orgs, "", "   ")
	fmt.Println(string(text))

	user, err := client.GetUser()
	if err != nil {
		log.Warn(err)
	}
	text, _ = json.MarshalIndent(user, "", "   ")
	fmt.Println(string(text))

	userProfile, err := client.GetUserProfile(user.Username)
	if err != nil {
		log.Warn(err)
	}
	text, _ = json.MarshalIndent(userProfile, "", "   ")
	fmt.Println(string(text))

	organizationKey := userProfile.OrgMemberships[0].PrimaryKey
	clusters, err := client.GetClusters(organizationKey)
	if err != nil {
		log.Warn(err)
	}
	text, _ = json.MarshalIndent(clusters, "", "   ")
	fmt.Println(string(text))

	clusterKey := (clusters)[0].PrimaryKey
	cluster, err := client.GetCluster(organizationKey, clusterKey)
	if err != nil {
		log.Warn(err)
	}
	text, _ = json.MarshalIndent(cluster, "", "   ")
	fmt.Println(string(text))

	nodes, err := client.GetNodes(organizationKey, clusterKey)
	if err != nil {
		log.Warn(err)
	}
	text, _ = json.MarshalIndent(nodes, "", "   ")
	fmt.Println(string(text))

	// nodeAdd := stackpointio.NodeAdd{Count: 1, Size: "t2.medium"}
	// content, _ := client.AddNodes(organizationKey, clusterKey, nodeAdd)
	// fmt.Println(string(content))
}
