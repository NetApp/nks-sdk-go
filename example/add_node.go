package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/StackPointCloud/stackpoint-sdk-go"
)

func main() {

	token := os.Getenv("CLUSTER_API_TOKEN")
	endpoint := os.Getenv("SPC_BASE_API_URL")
	client := stackpointio.NewClient(token, endpoint)
	cluster := stackpointio.CreateClusterClient(1, 543, client)
	// cluster := stackpointio.CreateClusterClient(6, 779, client)

	//	manager := stackpointio.CreateNodeManager("t2.medium", cluster)
	manager := stackpointio.CreateNodeManager("2GB", cluster)
	manager.Update()

	log.Info("Cluster running node count: ", manager.Size())

	count, err := manager.IncreaseSize(1)
	if err != nil {
		log.Error(err)
	}
	log.Info("Cluster running node count: ", count)

	// deleted, err := manager.DeleteNodes([]string{"spcs45yyk6-worker-4"})
	// if err != nil {
	// 	log.Error(err)
	// }
	// log.Info("Cluster deleted count: ", deleted)

}
