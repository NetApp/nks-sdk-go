package stackpointio

import (
	"fmt"
	"time"
)

const SolutionInstalledStateString = "installed"

// Solution struct to hold software packages running on a cluster
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

// WaitSolutionInstalled waits until solution is installed
func (c *APIClient) WaitSolutionInstalled(orgID, clusterID, solutionID, timeout int) error {
	for i := 1; i < timeout; i++ {
		state, err := c.GetSolutionState(orgID, clusterID, solutionID)
		if err != nil {
			return err
		}
		if state == SolutionInstalledStateString {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before solution reached state (%s)\n",
		timeout, SolutionInstalledStateString)
}
