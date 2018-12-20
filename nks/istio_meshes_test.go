package nks

import (
	"fmt"
	"testing"
)

func TestGetIstioMeshes(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	workspaceID, err := GetIDFromEnv("NKS_WORKSPACE_ID")
	if err != nil {
		t.Error(err)
	}
	meshes, err := c.GetIstioMeshes(orgID, workspaceID)
	if err != nil {
		t.Error(err)
	}
	if len(meshes) == 0 {
		fmt.Println("No Istio meshes found, but no error")
	}
}

func TestGetIstioMesh(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	workspaceID, err := GetIDFromEnv("NKS_WORKSPACE_ID")
	if err != nil {
		t.Error(err)
	}
	meshes, err := c.GetIstioMeshes(orgID, workspaceID)
	if err != nil {
		t.Error(err)
	}
	if len(meshes) > 0 {
		meshID := meshes[0].ID
		mesh, err := c.GetIstioMesh(orgID, workspaceID, meshID)
		if err != nil {
			t.Error(err)
		}
		if mesh == nil {
			t.Errorf("Could not fetch key for mesh: %d", meshID)
		}
	}
}
