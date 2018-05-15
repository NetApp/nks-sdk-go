package stackpointio

import (
	"fmt"
	"time"
)

// Solution is a application or process running with or on a kubernetes cluster
type Solution struct {
	ID         int       `json:"pk"`
	Name       string    `json:"name"`
	InstanceID string    `json:"instance_id"`
	Solution   string    `json:"solution"`
	Installer  string    `json:"installer"`
	State      string    `json:"state,omitempty"`
	URL        string    `json:"url,omitempty"`
	Deleteable bool      `json:"is_deleteable"`
	Keyset     int       `json:"keyset,omitempty"` // only for turbonomic and sysdig
	KeysetName string    `json:"keyset_name,omitempty"`
	MaxNodes   int       `json:"max_nodes,omitempty"` // only for autoscaler
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated,omitempty"`
}

// GetSolutions gets the solutions associated with a cluster and organization
func (c *APIClient) GetSolutions(orgID, clusterID int) ([]Solution, error) {
	r := []Solution{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d/solutions", orgID, clusterID), nil, &r, 200)
	return r, err
}

// GetSolution retrieves data for a single solution
func (c *APIClient) GetSolution(orgID, clusterID, solutionID int) (*Solution, error) {
	r := &Solution{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d/solutions/%d",
		orgID, clusterID, solutionID), nil, r, 200)
	return r, err
}

// DeleteSolution makes an API call to begin deleting a solution
func (c *APIClient) DeleteSolution(orgID, clusterID, solutionID int) error {
	return c.runRequest("DELETE", fmt.Sprintf("/orgs/%d/clusters/%d/solutions/%d",
		orgID, clusterID, solutionID), nil, nil, 204)
}

// AddSolution sends a request to add a solution to a cluster, returns list of solutions added
func (c *APIClient) AddSolution(orgID, clusterID int, newSolution Solution) (*Solution, error) {
	r := &Solution{}
	err := c.runRequest("POST", fmt.Sprintf("/orgs/%d/clusters/%d/solutions", orgID, clusterID), newSolution, r, 201)
	return r, err
}

// GetSolutionState returns state of solution
func (c *APIClient) GetSolutionState(orgID, clusterID, solutionID int) (string, error) {
	r := &Solution{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d/solutions/%d", orgID, clusterID, solutionID), nil, r, 200)
	if err != nil {
		return "", err
	}
	return r.State, nil
}
