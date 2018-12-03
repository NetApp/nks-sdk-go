package stackpointio

import "fmt"
import "time"

// IstioMesh struct
type IstioMesh struct {
	ID        int         `json:"pk"`
	Name      string      `json:"name"`
	MeshType  string      `json:"mesh_type"`
	Org       int         `json:"org"`
	Workspace Workspace   `json:"workspace"`
	Members   []Member    `json:"members"`
	State     string      `json:"state"`
	Config    interface{} `json:"config,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

// Member struct
type Member struct {
	ID      int         `json:"pk"`
	Mesh    int         `json:"mesh"`
	Gateway string      `json:"gateway"`
	Role    string      `json:"role"`
	Cluster Cluster     `json:"cluster"`
	State   string      `json:"state"`
	Config  interface{} `json:"config,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Created time.Time   `json:"created"`
	Updated time.Time   `json:"updated"`
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
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/workspaces/%d/istio-meshes/%d", orgID, workspaceID, meshID),
		ResponseObj:  &m,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// CreateIstioMesh creates an Istio mesh
func (c *APIClient) CreateIstioMesh(orgID int, workspaceID int, mesh IstioMesh) (m *IstioMesh, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/workspaces/%d/istio-mesh", orgID, workspaceID),
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
		Path:         fmt.Sprintf("/orgs/%d/workspaces/%d/istio-mesh/%d", orgID, workspaceID, meshID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}
