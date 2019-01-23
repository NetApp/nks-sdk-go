package nks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testIstioAwsCluster = Cluster{
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

var testIstioMeshClusterIDs = make([]int, 0)
var testIstioMeshWorkspace, meshID int

func TestLiveIstioMeshBasic(t *testing.T) {

	testIstioAwsCluster.Solutions = append(testIstioAwsCluster.Solutions, Solution{
		Solution: "istio",
		State:    "draft",
	})

	cluster1ID := 0
	cluster2ID := 0

	t.Run("create clusters", func(t *testing.T) {
		t.Run("Cluster 1", func(t *testing.T) {
			cluster1ID = testIstioMeshCreateCluster(t)
		})
		t.Run("Cluster 2", func(t *testing.T) {
			cluster2ID = testIstioMeshCreateCluster(t)
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
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetWorkspaces(orgID)

	for _, workspace := range list {
		if workspace.IsDefault {
			return workspace.ID
		}
	}

	t.Error(errors.New("Could not find default workspace"))

	return 0
}

func testIstioMeshCreateCluster(t *testing.T) int {
	t.Parallel()
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

	testIstioAwsCluster.ProviderKey = awsKeysetID
	testIstioAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testIstioAwsCluster)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	return cluster.ID
}

func testIstioMeshCreateIstioMesh(t *testing.T, workspaceID, cluster1ID, cluster2ID int) int {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

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

	mesh, err := c.CreateIstioMesh(orgID, workspaceID, newMesh)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitIstioMeshCreated(orgID, workspaceID, mesh.ID, timeout)
	if err != nil {
		t.Error(err)
	}

	return mesh.ID
}

func testIstioMeshList(t *testing.T, worspaceID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetIstioMeshes(orgID, worspaceID)
	if err != nil {
		t.Error(err)
	}

	assert.NotEqual(t, len(list), 0, "At least one istio mesh must exist")
}

func testIstioMeshGet(t *testing.T, worspaceID, meshId int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	mesh, err := c.GetIstioMesh(orgID, worspaceID, meshId)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, mesh, "Istio mesh must exist")
}

func testIstioMeshDeleteIstioMesh(t *testing.T, workspaceID, istioMeshID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteIstioMesh(orgID, workspaceID, istioMeshID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitIstioMeshDeleted(orgID, workspaceID, istioMeshID, timeout)
	if err != nil {
		t.Error(err)
	}
}
func testIstioMeshDeleteCluster(t *testing.T, clusterID int) {
	t.Parallel()

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
