package nks

import (
	"fmt"
	"strings"
	"time"
)

// IstioMeshRequest object used to create an istio mesh
type IstioMeshRequest struct {
	Name      string          `json:"name"`
	MeshType  string          `json:"mesh_type"`
	Members   []MemberRequest `json:"members"`
	Workspace int             `json:"workspace"`
}

// MemberRequest object used to by IstioMeshRequest to create an istio mesh
type MemberRequest struct {
	Role    string `json:"role,omitempty"`
	Cluster int    `json:"cluster,omitempty"`
}

// IstioMesh struct
type IstioMesh struct {
	ID        int         `json:"pk"`
	Name      string      `json:"name"`
	MeshType  string      `json:"mesh_type"`
	Org       int         `json:"org"`
	Workspace Workspace   `json:"workspace"`
	Members   []Member    `json:"members"`
	State     string      `json:"state,omitempty"`
	Config    interface{} `json:"config,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

// Member struct
type Member struct {
	ID      int         `json:"pk,omitempty"`
	Mesh    int         `json:"mesh,omitempty"`
	Gateway string      `json:"gateway,omitempty"`
	Role    string      `json:"role,omitempty"`
	Cluster Cluster     `json:"cluster,omitempty"`
	State   string      `json:"state,omitempty"`
	Config  interface{} `json:"config,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Created time.Time   `json:"created,omitempty"`
	Updated time.Time   `json:"updated,omitempty"`
}

// GetIstioMeshes gets list of meshes for Org ID and Workspace ID
func (c *APIClient) GetIstioMeshes(orgID int, workspaceID int) (m []IstioMesh, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/workspaces/%d/istio-meshes", orgID, workspaceID),
		ResponseObj:  &m,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetIstioMesh returns a mesh for Org ID, Workspace ID, and meshID
func (c *APIClient) GetIstioMesh(orgID int, workspaceID int, meshID int) (m *IstioMesh, err error) {
	items, err := c.GetIstioMeshes(orgID, workspaceID)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.ID == meshID {
			return &item, nil
		}
	}

	err = fmt.Errorf("Incorrect status code returned: 404, Status: 404 Not Found,")
	return nil, err
}

// CreateIstioMesh creates an Istio mesh
func (c *APIClient) CreateIstioMesh(orgID int, workspaceID int, mesh IstioMeshRequest) (m *IstioMesh, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/workspaces/%d/istio-meshes", orgID, workspaceID),
		ResponseObj:  &m,
		PostObj:      mesh,
		WantedStatus: 201,
	}
	err = c.runRequest(req)
	return
}

// DeleteIstioMesh deletes an Istio mesh
func (c *APIClient) DeleteIstioMesh(orgID int, workspaceID int, meshID int) (err error) {
	req := &APIReq{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/orgs/%d/istio-meshes/%d", orgID, meshID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}

// WaitIstioMeshDeleted waits until isto mesh disappears
func (c *APIClient) WaitIstioMeshDeleted(orgID, workspaceID, meshID, timeout int) error {
	for i := 1; i < timeout; i++ {
		_, err := c.GetIstioMesh(orgID, workspaceID, meshID)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("timeout (%d seconds) reached before isto mesh deleted", timeout)
}

// WaitIstioMeshCreated waits until isto mesh state is active
func (c *APIClient) WaitIstioMeshCreated(orgID, workspaceID, meshID, timeout int) error {
	for i := 1; i < timeout; i++ {
		mesh, err := c.GetIstioMesh(orgID, workspaceID, meshID)
		if err == nil && mesh.State == "active" {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("timeout (%d seconds) reached before isto mesh created", timeout)
}
