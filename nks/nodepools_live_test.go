package nks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var awsZone = "eu-west-2a"
var testNodePoolAwsCluster = Cluster{
	Name:                  "Test AWS Cluster Go SDK " + GetTicks(),
	Provider:              "aws",
	MasterCount:           1,
	MasterSize:            "t2.medium",
	WorkerCount:           2,
	WorkerSize:            "t2.medium",
	Region:                "eu-west-2",
	Zone:                  awsZone,
	ProviderNetworkID:     "__new__",
	ProviderNetworkCdr:    "172.23.0.0/16",
	ProviderSubnetID:      "__new__",
	ProviderSubnetCidr:    "172.23.1.0/24",
	KubernetesVersion:     "v1.15.5",
	KubernetesPodCidr:     "10.2.0.0",
	KubernetesServiceCidr: "10.3.0.0",
	RbacEnabled:           true,
	DashboardEnabled:      true,
	EtcdType:              "classic",
	Platform:              "coreos",
	Channel:               "stable",
	NetworkComponents:     []NetworkComponent{},
	Solutions:             []Solution{Solution{Solution: "helm_tiller"}},
}

func TestLiveBasicNodePool(t *testing.T) {
	clusterID, nodePoolID := testNodePoolClusterCreate(t)
	testAddNodeToNodePool(t, clusterID, nodePoolID)
	testNodePoolList(t, clusterID)
	testNodePoolGet(t, clusterID, nodePoolID)
	testNodePoolDelete(t, clusterID, nodePoolID)
	testNodePoolClusterDelete(t, clusterID)
}

func testNodePoolClusterCreate(t *testing.T) (int, int) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Fatal(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Fatal(err)
	}

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	if err != nil {
		t.Fatal(err)
	}

	testNodePoolAwsCluster.ProviderKey = awsKeysetID
	testNodePoolAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testNodePoolAwsCluster)
	fmt.Println(cluster.ID, err)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitClusterRunning(orgID, cluster.ID, true, timeout)


	newNodePool := NodePool{
		Name:               "test sdk np",
		Platform:           "coreos",
		NodeCount:          1,
		Size:               "t2.medium",
		Zone:               awsZone,
		ProviderSubnetID:   "__new__",
		ProviderSubnetCidr: "172.23.4.0/24",
	}

	nodePool, err := c.CreateNodePool(orgID, cluster.ID, newNodePool)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitNodePoolProvisioned(orgID, cluster.ID, nodePool.ID, timeout)

	pools, err := c.GetNodePools(orgID, cluster.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(pools), "Cluster must have a node pool")
	return cluster.ID, nodePool.ID
}

func testAddNodeToNodePool(t *testing.T, clusterID, nodePoolID int) int {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	nodeAdd := NodeAddToPool{
		Count: 1,
		Role:  "worker",
	}

	nodes, err := c.AddNodesToNodePool(orgID, clusterID, nodePoolID, nodeAdd)
	if err != nil {
		t.Error(err)
	}

	node := nodes[0]

	err = c.WaitNodeProvisioned(orgID, clusterID, node.ID, timeout)
	if err != nil {
		t.Error(err)
	}

	return node.ID
}

func testNodePoolList(t *testing.T, clusterID int) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetNodePools(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, len(list), "There should be 2 nodepool")

}

func testNodePoolGet(t *testing.T, clusterID int, nodePoolID int) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	nodePool, err := c.GetNodePool(orgID, clusterID, nodePoolID)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, nodePool)
	assert.Equal(t, nodePool.ID, nodePoolID, "different node pool ID")
}

func testNodePoolDelete(t *testing.T, clusterID, nodepoolID int) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	//delete all nodes before deleting node pool
	nodesInNodePool, err := c.GetNodesInPool(orgID, clusterID, nodepoolID)

	for _, node := range nodesInNodePool {
		err = c.DeleteNode(orgID, clusterID, node.ID)

	}
	for _, node := range nodesInNodePool {
		if testEnv != "mock" {
			err = c.WaitNodeDeleted(orgID, clusterID, node.ID, timeout)
		}
	}

	err = c.DeleteNodePool(orgID, clusterID, nodepoolID)
	if err != nil {
		t.Error(err)
	}

}

func testNodePoolClusterDelete(t *testing.T, clusterID int) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteCluster(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}

	if testEnv != "mock" {
		err = c.WaitClusterDeleted(orgID, clusterID, timeout)
		if err != nil {
			t.Error(err)
		}
	}
}
