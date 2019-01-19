package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testWorkspaceLiveID int
var testWorkspace = Workspace{
	Name:           "Test Go SDK" + getTicks(),
	Slug:           "test_go_sdk_" + getTicks(),
	TeamWorkspaces: []TeamWorkspace{},
}

func TestLiveWorkspaceBasic(t *testing.T) {
	testLiveWorkspaceCreate(t)
	testLiveWorkspaceList(t)
	testLiveWorkspaceGet(t)
	testLiveWorkspaceDelete(t)
}

func testLiveWorkspaceCreate(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	workspace, err := c.CreateWorkspace(orgID, testWorkspace)
	if err != nil {
		t.Error(err)
	}

	testWorkspaceLiveID = workspace.ID

	assert.Equal(t, testWorkspace.Name, workspace.Name, "Name should be equal")
	assert.Equal(t, testWorkspace.Name, workspace.Name, "Slug should be equal")
}

func testLiveWorkspaceList(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetWorkspaces(orgID)
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
	assert.Equal(t, testWorkspace.Name, workspace.Name, "Name should be equal")
	assert.Equal(t, testWorkspace.Name, workspace.Name, "Slug should be equal")
}

func testLiveWorkspaceGet(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	workspace, err := c.GetWorkspace(orgID, testWorkspaceLiveID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, testWorkspace.Name, workspace.Name, "Name should be equal")
	assert.Equal(t, testWorkspace.Name, workspace.Name, "Slug should be equal")
}

func testLiveWorkspaceDelete(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteWorkspace(orgID, testWorkspaceLiveID)
	if err != nil {
		t.Error(err)
	}
}
