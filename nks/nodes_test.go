package nks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testNodeAwsCluster = Cluster{
	Name:                  "Test AWS Cluster Go SDK " + GetTicks(),
	Provider:              "aws",
	MasterCount:           1,
	MasterSize:            "t2.medium",
	WorkerCount:           2,
	WorkerSize:            "t2.medium",
	Region:                "eu-west-3",
	Zone:                  "eu-west-3a",
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

func TestLiveBasicNode(t *testing.T) {
	clusterID := testNodeClusterCreate(t)
	nodeID := testNodeCreate(t, clusterID)
	testNodeList(t, clusterID)
	testNodeGet(t, clusterID, nodeID)
	testNodeDelete(t, clusterID, nodeID)
	testNodeClusterDelete(t, clusterID)
}

func testNodeClusterCreate(t *testing.T) int {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	require.NoError(t, err)

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	require.NoError(t, err)

	testNodeAwsCluster.ProviderKey = awsKeysetID
	testNodeAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := client.CreateCluster(orgID, testNodeAwsCluster)
	fmt.Println(cluster.ID)
	require.NoError(t, err)

	err = client.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	return cluster.ID
}

func testNodeCreate(t *testing.T, clusterID int) int {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	nodeAdd := NodeAdd{
		Count:              1,
		Size:               "t2.medium",
		Role:               "master",
		Zone:               "eu-west-3a",
		ProviderSubnetID:   "__new__",
		ProviderSubnetCidr: "172.23.1.0/24",
		RootDiskSize:       50,
	}

	nodes, err := client.AddNode(orgID, clusterID, nodeAdd)
	require.NoError(t, err)

	node := nodes[0]

	err = client.WaitNodeProvisioned(orgID, clusterID, node.ID, timeout)
	require.NoError(t, err)

	return node.ID
}

func testNodeList(t *testing.T, clusterID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	list, err := client.GetNodes(orgID, clusterID)
	require.NoError(t, err)

	assert.Equal(t, len(list), 4, "There should be 4 nodes")
}

func testNodeGet(t *testing.T, clusterID, nodeID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	node, err := client.GetNode(orgID, clusterID, nodeID)
	require.NoError(t, err)

	assert.NotNil(t, node)
	assert.Equal(t, node.ID, nodeID, "Master node must exist")
}

func testNodeDelete(t *testing.T, clusterID, nodeID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	err = client.DeleteNode(orgID, clusterID, nodeID)
	require.NoError(t, err)

	if testEnv != "mock" {
		err = client.WaitNodeDeleted(orgID, clusterID, nodeID, timeout)
		require.NoError(t, err)
	}
}

func testNodeClusterDelete(t *testing.T, clusterID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	err = client.DeleteCluster(orgID, clusterID)
	require.NoError(t, err)
	if testEnv != "mock" {
		err = client.WaitClusterDeleted(orgID, clusterID, timeout)
		require.NoError(t, err)
	}
}
