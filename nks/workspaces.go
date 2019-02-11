package nks

import "fmt"
import "time"

// Federation struct
type Federation struct {
	ID    int    `json:"pk"`
	Name  string `json:"name"`
	State string `json:"state"`
}

// TeamWorkspace struct
type TeamWorkspace struct {
	ID        int       `json:"pk"`
	Team      Team      `json:"team"`
	Workspace int       `json:"workspace"`
	Created   time.Time `json:"created"`
}

// Workspace struct
type Workspace struct {
	ID             int             `json:"pk"`
	Name           string          `json:"name"`
	Slug           string          `json:"slug"`
	Org            int             `json:"org"`
	IsDefault      bool            `json:"is_default"`
	Created        time.Time       `json:"created"`
	Clusters       []Cluster       `json:"clusters"`
	UserSolutions  []int           `json:"user_solutions"`
	TeamWorkspaces []TeamWorkspace `json:"team_workspaces"`
	Federations    []Federation    `json:"federations"`
}

// GetWorkspaces gets list of workspaces for Org ID
func (c *APIClient) GetWorkspaces(orgID int) (workspace []Workspace, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/workspaces", orgID),
		ResponseObj:  &workspace,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetWorkspace returns workspace for Org ID and Workspace ID
func (c *APIClient) GetWorkspace(orgID, workspaceID int) (ws *Workspace, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/workspaces/%d", orgID, workspaceID),
		ResponseObj:  &ws,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// CreateWorkspace creates a workspace
func (c *APIClient) CreateWorkspace(orgID int, workspace Workspace) (ws *Workspace, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/workspaces", orgID),
		ResponseObj:  &ws,
		PostObj:      workspace,
		WantedStatus: 201,
	}
	err = c.runRequest(req)
	return
}

// DeleteWorkspace deletes workspace
func (c *APIClient) DeleteWorkspace(orgID, workspaceID int) (err error) {
	req := &APIReq{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/orgs/%d/workspaces/%d", orgID, workspaceID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}
