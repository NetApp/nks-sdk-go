package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testNodePoolAwsCluster = Cluster{
	Name:               "Test AWS Cluster Go SDK " + GetTicks(),
	Provider:           "aws",
	MasterCount:        1,
	MasterSize:         "t2.medium",
	WorkerCount:        2,
	WorkerSize:         "t2.medium",
	Region:             "us-east-1",
	Zone:               "us-east-1a",
	ProviderNetworkID:  "__new__",
	ProviderNetworkCdr: "172.23.0.0/16",
	ProviderSubnetID:   "__new__",
	ProviderSubnetCidr: "172.23.1.0/24",
	KubernetesVersion:  "v1.13.1",
	RbacEnabled:        true,
	DashboardEnabled:   true,
	EtcdType:           "classic",
	Platform:           "coreos",
	Channel:            "stable",
	NetworkComponents:  []NetworkComponent{},
	Solutions:          []Solution{Solution{Solution: "helm_tiller"}},
}

func TestLiveNodePoolBasic(t *testing.T) {
	clusterID, nodePoolID := testNodePoolClusterCreate(t)
	nodeID := testNodePoolCreate(t, clusterID, nodePoolID)
	testNodePoolList(t, clusterID)
	testNodePoolGet(t, clusterID, nodeID)
	testNodePoolDelete(t, clusterID, nodeID)
	testNodePoolClusterDelete(t, clusterID)
}

func testNodePoolClusterCreate(t *testing.T) (int, int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testNodePoolAwsCluster.ProviderKey = awsKeysetID
	testNodePoolAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testNodePoolAwsCluster)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	pools, err := c.GetNodePools(orgID, cluster.ID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(pools), "Cluster must have a nood pool")
	return cluster.ID, pools[0].ID
}

func testNodePoolCreate(t *testing.T, clusterID, nodePoolID int) int {
	c, err := NewClientFromEnv()
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
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetNodes(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 4, len(list), "There should be 4 nodes")
}

func testNodePoolGet(t *testing.T, clusterID, nodeID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	node, err := c.GetNode(orgID, clusterID, nodeID)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, node)
	assert.Equal(t, node.ID, nodeID, "Worker node must exist")
	assert.Equal(t, "worker", node.Role, "Node must be of type worker")
}

func testNodePoolDelete(t *testing.T, clusterID, nodeID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteNode(orgID, clusterID, nodeID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitNodeDeleted(orgID, clusterID, nodeID, timeout)
	if err != nil {
		t.Error(err)
	}
}

func testNodePoolClusterDelete(t *testing.T, clusterID int) {
	c, err := NewClientFromEnv()
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

	err = c.WaitClusterDeleted(orgID, clusterID, timeout)
	if err != nil {
		t.Error(err)
	}
}
