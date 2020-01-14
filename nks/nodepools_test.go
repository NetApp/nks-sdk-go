package nks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	require.NoError(t, err)

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	require.NoError(t, err)

	testNodePoolAwsCluster.ProviderKey = awsKeysetID
	testNodePoolAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := client.CreateCluster(orgID, testNodePoolAwsCluster)
	fmt.Println(cluster.ID, err)
	require.NoError(t, err)

	err = client.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	newNodePool := NodePool{
		Name:               "test sdk np",
		Platform:           "coreos",
		NodeCount:          1,
		Size:               "t2.medium",
		Zone:               awsZone,
		ProviderSubnetID:   "__new__",
		ProviderSubnetCidr: "172.23.4.0/24",
	}

	nodePool, err := client.CreateNodePool(orgID, cluster.ID, newNodePool)
	require.NoError(t, err)

	err = client.WaitNodePoolProvisioned(orgID, cluster.ID, nodePool.ID, timeout)

	pools, err := client.GetNodePools(orgID, cluster.ID)
	require.NoError(t, err)

	assert.Equal(t, 2, len(pools), "Cluster must have a node pool")
	return cluster.ID, nodePool.ID
}

func testAddNodeToNodePool(t *testing.T, clusterID, nodePoolID int) int {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	nodeAdd := NodeAddToPool{
		Count: 1,
		Role:  "worker",
	}

	nodes, err := client.AddNodesToNodePool(orgID, clusterID, nodePoolID, nodeAdd)
	require.NoError(t, err)

	node := nodes[0]

	err = client.WaitNodeProvisioned(orgID, clusterID, node.ID, timeout)
	require.NoError(t, err)

	return node.ID
}

func testNodePoolList(t *testing.T, clusterID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	list, err := client.GetNodePools(orgID, clusterID)
	require.NoError(t, err)

	assert.Equal(t, 2, len(list), "There should be 2 nodepool")

}

func testNodePoolGet(t *testing.T, clusterID int, nodePoolID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	nodePool, err := client.GetNodePool(orgID, clusterID, nodePoolID)
	require.NoError(t, err)

	assert.NotNil(t, nodePool)
	assert.Equal(t, nodePool.ID, nodePoolID, "different node pool ID")
}

func testNodePoolDelete(t *testing.T, clusterID, nodepoolID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	//delete all nodes before deleting node pool
	nodesInNodePool, err := client.GetNodesInPool(orgID, clusterID, nodepoolID)

	for _, node := range nodesInNodePool {
		err = client.DeleteNode(orgID, clusterID, node.ID)

	}
	for _, node := range nodesInNodePool {
		if testEnv != "mock" {
			err = client.WaitNodeDeleted(orgID, clusterID, node.ID, timeout)
		}
	}

	err = client.DeleteNodePool(orgID, clusterID, nodepoolID)
	require.NoError(t, err)
}

func testNodePoolClusterDelete(t *testing.T, clusterID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	err = client.DeleteCluster(orgID, clusterID)
	require.NoError(t, err)

	if testEnv != "mock" {
		err = client.WaitClusterDeleted(orgID, clusterID, timeout)
		require.NoError(t, err)
	}
}
