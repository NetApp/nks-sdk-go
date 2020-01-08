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

	testSolutionAwsCluster.ProviderKey = awsKeysetID
	testSolutionAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := client.CreateCluster(orgID, testSolutionAwsCluster)
	if err != nil {
		t.Error(err)
	}

	err = client.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	return cluster.ID
}

func testSolutionAdd(t *testing.T, clusterID int) int {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	newSolution := Solution{
		Solution: solutionName,
		State:    "draft",
	}

	solution, err := client.AddSolution(orgID, clusterID, newSolution)
	if err != nil {
		t.Error(err)
	}

	err = client.WaitSolutionInstalled(orgID, clusterID, solution.ID, timeout)
	if err != nil {
		t.Error(err)
	}

	return solution.ID
}

func testSolutionList(t *testing.T, clusterID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := client.GetSolutions(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(list)

	assert.Equal(t, len(list), 2, "Two solutins have to be installed")
	assert.Equal(t, list[0].Solution, solutionName, solutionName+" solution has to be installed")
}

func testSolutionGet(t *testing.T, clusterID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	solutionID, err := client.FindSolutionByName(orgID, clusterID, solutionName)
	if err != nil {
		t.Error(err)
	}

	solution, err := client.GetSolution(orgID, clusterID, solutionID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, solution.Solution, solutionName, solutionName+"solution has to be installed")
	assert.Equal(t, solution.State, SolutionInstalledStateString, solutionName+"solution has to be installed")
}

func testSolutionDelete(t *testing.T, clusterID int, solutionID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = client.DeleteSolution(orgID, clusterID, solutionID)
	if err != nil {
		t.Error(err)
	}
	if testEnv != "mock" {
		err = client.WaitSolutionDeleted(orgID, clusterID, solutionID, timeout)
		if err != nil {
			t.Error(err)
		}
	}
}

func testSolutionDeleteCluster(t *testing.T, clusterID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = client.DeleteCluster(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}
	if testEnv != "mock" {
		err = client.WaitClusterDeleted(orgID, clusterID, timeout)
		if err != nil {
			t.Error(err)
		}
	}
}
