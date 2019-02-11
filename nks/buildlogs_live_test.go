package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testBuildLogAwsCluster = Cluster{
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

func TestLiveBasicBuildLogs(t *testing.T) {
	clusterID := testBuildLogsCreateCluster(t)
	testBuildLogsGet(t, clusterID)
	testBuildLogsDeleteCluster(t, clusterID)
}

func testBuildLogsGet(t *testing.T, clusterID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	logs, err := c.GetBuildLogs(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}

	assert.NotEqual(t, 0, len(logs), "Logs should be present when a cluster is created")
}

func testBuildLogsCreateCluster(t *testing.T) int {
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

	testBuildLogAwsCluster.ProviderKey = awsKeysetID
	testBuildLogAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testBuildLogAwsCluster)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	return cluster.ID
}

func testBuildLogsDeleteCluster(t *testing.T, clusterID int) {
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
