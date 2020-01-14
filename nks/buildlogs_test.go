package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testBuildLogAwsCluster = Cluster{
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
	KubernetesVersion:     "v1.13.1",
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

func TestLiveBasicBuildLogs(t *testing.T) {
	clusterID := testBuildLogsCreateCluster(t)
	testBuildLogsGet(t, clusterID)
	testBuildLogsDeleteCluster(t, clusterID)
}

func testBuildLogsGet(t *testing.T, clusterID int) {

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	logs, err := client.GetBuildLogs(orgID, clusterID)
	require.NoError(t, err)

	assert.NotEqual(t, 0, len(logs), "Logs should be present when a cluster is created")
}

func testBuildLogsCreateCluster(t *testing.T) int {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	require.NoError(t, err)

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	require.NoError(t, err)

	testBuildLogAwsCluster.ProviderKey = awsKeysetID
	testBuildLogAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := client.CreateCluster(orgID, testBuildLogAwsCluster)
	require.NoError(t, err)

	err = client.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	return cluster.ID
}

func testBuildLogsDeleteCluster(t *testing.T, clusterID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	err = client.DeleteCluster(orgID, clusterID)
	require.NoError(t, err)

	if testEnv != "mock" {
		err = client.WaitClusterDeleted(orgID, clusterID, timeout)
		require.NoError(t, err)
	}
}
