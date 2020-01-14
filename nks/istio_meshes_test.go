package nks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testIstioAwsCluster = Cluster{
	Name:                  "Test AWS Cluster Go SDK " + GetTicks(),
	Provider:              "aws",
	MasterCount:           1,
	MasterSize:            "t2.medium",
	WorkerCount:           2,
	WorkerSize:            "t2.medium",
	Region:                "us-west-2",
	Zone:                  "us-west-2a",
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
	Platform:              "ubuntu",
	Channel:               "18.04-lts",
	MasterRootDiskSize:    50,
	WorkerRootDiskSize:    50,
	NetworkComponents:     []NetworkComponent{},
	Solutions:             []Solution{Solution{Solution: "helm_tiller"}},
}

var testIstioMeshClusterIDs = make([]int, 0)
var testIstioMeshWorkspace, meshID int

func TestLiveBasicIstioMesh(t *testing.T) {
	cluster1ID := 0
	cluster2ID := 0

	t.Run("create clusters", func(t *testing.T) {
		t.Run("Cluster 1", func(t *testing.T) {
			cluster1ID = testIstioMeshCreateCluster(t, "1")
		})
		t.Run("Cluster 2", func(t *testing.T) {
			cluster2ID = testIstioMeshCreateCluster(t, "2")
		})
	})

	workspaceID := testIstioMeshGetDefaultWorkspace(t)
	meshID := testIstioMeshCreateIstioMesh(t, workspaceID, cluster1ID, cluster2ID)

	testIstioMeshList(t, workspaceID)
	testIstioMeshGet(t, workspaceID, meshID)

	testIstioMeshDeleteIstioMesh(t, workspaceID, meshID)

	t.Run("delete clusters", func(t *testing.T) {
		t.Run("Cluster 1", func(t *testing.T) {
			testIstioMeshDeleteCluster(t, cluster1ID)
		})
		t.Run("Cluster 2", func(t *testing.T) {
			testIstioMeshDeleteCluster(t, cluster2ID)
		})
	})
}

func testIstioMeshGetDefaultWorkspace(t *testing.T) int {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	list, err := client.GetWorkspaces(orgID)

	for _, workspace := range list {
		if workspace.IsDefault {
			return workspace.ID
		}
	}

	t.Fatal(errors.New("Could not find default workspace"))

	return 0
}

func testIstioMeshCreateCluster(t *testing.T, index string) int {
	t.Parallel()
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	require.NoError(t, err)

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	require.NoError(t, err)

	testIstioAwsCluster.ProviderKey = awsKeysetID
	testIstioAwsCluster.SSHKeySet = sshKeysetID
	testIstioAwsCluster.Name = testIstioAwsCluster.Name + index

	cluster, err := client.CreateCluster(orgID, testIstioAwsCluster)
	require.NoError(t, err)

	err = client.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	newSolution := Solution{
		Solution: "istio",
		State:    "draft",
	}

	solution, err := client.AddSolution(orgID, cluster.ID, newSolution)
	require.NoError(t, err)

	err = client.WaitSolutionInstalled(orgID, cluster.ID, solution.ID, timeout)
	require.NoError(t, err)
	return cluster.ID
}

func testIstioMeshCreateIstioMesh(t *testing.T, workspaceID, cluster1ID, cluster2ID int) int {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	newMesh := IstioMeshRequest{
		Name:      "Test AWS Istio Mesh Go SDK " + GetTicks(),
		MeshType:  "cross_cluster",
		Workspace: workspaceID,
		Members: []MemberRequest{
			MemberRequest{
				Cluster: cluster1ID,
				Role:    "host",
			},
			MemberRequest{
				Cluster: cluster2ID,
				Role:    "guest",
			},
		},
	}

	mesh, err := client.CreateIstioMesh(orgID, workspaceID, newMesh)
	require.NoError(t, err)

	err = client.WaitIstioMeshCreated(orgID, workspaceID, mesh.ID, timeout)
	require.NoError(t, err)

	return mesh.ID
}

func testIstioMeshList(t *testing.T, worspaceID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	list, err := client.GetIstioMeshes(orgID, worspaceID)
	require.NoError(t, err)

	assert.NotEqual(t, len(list), 0, "At least one istio mesh must exist")
}

func testIstioMeshGet(t *testing.T, worspaceID, meshID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	mesh, err := client.GetIstioMesh(orgID, worspaceID, meshID)
	require.NoError(t, err)

	assert.NotNil(t, mesh, "Istio mesh must exist")
}

func testIstioMeshDeleteIstioMesh(t *testing.T, workspaceID, istioMeshID int) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	err = client.DeleteIstioMesh(orgID, workspaceID, istioMeshID)
	require.NoError(t, err)

	if testEnv != "mock" {
		err = client.WaitIstioMeshDeleted(orgID, workspaceID, istioMeshID, timeout)
		require.NoError(t, err)
	}
}
func testIstioMeshDeleteCluster(t *testing.T, clusterID int) {
	t.Parallel()

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	err = client.DeleteCluster(orgID, clusterID)
	require.NoError(t, err)

	if testEnv != "mock" {
		err = client.WaitClusterDeleted(orgID, clusterID, timeout)
		require.NoError(t, err)
	}
}
