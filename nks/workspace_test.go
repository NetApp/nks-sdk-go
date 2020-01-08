package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testWorkspaceLiveID int
var testWorkspace = Workspace{
	Name:           "Test Go SDK" + GetTicks(),
	Slug:           "test_go_sdk_" + GetTicks(),
	TeamWorkspaces: []TeamWorkspace{},
}

func TestLiveBasicWorkspace(t *testing.T) {
	testLiveWorkspaceCreate(t)
	testLiveWorkspaceList(t)
	testLiveWorkspaceGet(t)
	testLiveWorkspaceDelete(t)
}

func testLiveWorkspaceCreate(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	testWorkspace.Org = orgID

	workspace, err := client.CreateWorkspace(orgID, testWorkspace)
	if err != nil {
		t.Error(err)
	}

	testWorkspaceLiveID = workspace.ID

	assert.Contains(t, testWorkspace.Name, workspace.Name, "Name should be equal")
}

func testLiveWorkspaceList(t *testing.T) {

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := client.GetWorkspaces(orgID)
	if err != nil {
		t.Error(err)
	}

	var workspace Workspace
	for _, item := range list {
		if item.ID == testWorkspaceLiveID {
			workspace = item
		}
	}

	assert.NotNil(t, workspace)
	assert.Contains(t, testWorkspace.Name, workspace.Name, "Name should be equal")
}

func testLiveWorkspaceGet(t *testing.T) {

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	workspace, err := client.GetWorkspace(orgID, testWorkspaceLiveID)
	if err != nil {
		t.Error(err)
	}

	assert.Contains(t, testWorkspace.Name, workspace.Name, "Name should be equal")
}

func testLiveWorkspaceDelete(t *testing.T) {

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = client.DeleteWorkspace(orgID, testWorkspaceLiveID)
	if err != nil {
		t.Error(err)
	}
}
