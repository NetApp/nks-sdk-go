package nks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var solutionName = "haproxy"

var testSolutionAwsCluster = Cluster{
	Name:                  "Test AWS Cluster Go SDK " + GetTicks(),
	Provider:              "aws",
	MasterCount:           1,
	MasterSize:            "t2.medium",
	WorkerCount:           2,
	WorkerSize:            "t2.medium",
	Region:                "eu-west-1",
	Zone:                  "eu-west-1b",
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

func TestLiveBasicSolution(t *testing.T) {
	clusterID := testSolutionCreateCluster(t)
	solutionID := testSolutionAdd(t, clusterID)
	testSolutionList(t, clusterID)
	testSolutionGet(t, clusterID)
	testSolutionDelete(t, clusterID, solutionID)
	testSolutionDeleteCluster(t, clusterID)
}

func testSolutionCreateCluster(t *testing.T) int {
	c, err := NewTestClientFromEnv()
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

	testSolutionAwsCluster.ProviderKey = awsKeysetID
	testSolutionAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testSolutionAwsCluster)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	return cluster.ID
}

func testSolutionAdd(t *testing.T, clusterID int) int {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	newSolution := Solution{
		Solution: solutionName,
		State:    "draft",
	}

	solution, err := c.AddSolution(orgID, clusterID, newSolution)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitSolutionInstalled(orgID, clusterID, solution.ID, timeout)
	if err != nil {
		t.Error(err)
	}

	return solution.ID
}

func testSolutionList(t *testing.T, clusterID int) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetSolutions(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(list)

	assert.Equal(t, len(list), 2, "Two solutins have to be installed")
	assert.Equal(t, list[0].Solution, solutionName, solutionName+" solution has to be installed")
}

func testSolutionGet(t *testing.T, clusterID int) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	solutionID, err := c.FindSolutionByName(orgID, clusterID, solutionName)
	if err != nil {
		t.Error(err)
	}

	solution, err := c.GetSolution(orgID, clusterID, solutionID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, solution.Solution, solutionName, solutionName+"solution has to be installed")
	assert.Equal(t, solution.State, SolutionInstalledStateString, solutionName+"solution has to be installed")
}

func testSolutionDelete(t *testing.T, clusterID int, solutionID int) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteSolution(orgID, clusterID, solutionID)
	if err != nil {
		t.Error(err)
	}
	if testEnv != "mock" {
		err = c.WaitSolutionDeleted(orgID, clusterID, solutionID, timeout)
		if err != nil {
			t.Error(err)
		}
	}
}

func testSolutionDeleteCluster(t *testing.T, clusterID int) {
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
