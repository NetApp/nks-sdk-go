package stackpointio

import (
	"fmt"
	"strings"
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
func (c *APIClient) GetSolutions(orgID, clusterID int) (sols []Solution, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/solutions", orgID, clusterID),
		ResponseObj:  &sols,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetSolution retrieves data for a single solution
func (c *APIClient) GetSolution(orgID, clusterID, solutionID int) (sol *Solution, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/solutions/%d", orgID, clusterID, solutionID),
		ResponseObj:  &sol,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// DeleteSolution makes an API call to begin deleting a solution
func (c *APIClient) DeleteSolution(orgID, clusterID, solutionID int) (err error) {
	req := &APIReq{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/solutions/%d", orgID, clusterID, solutionID),
		WantedStatus: 202,
	}
	err = c.runRequest(req)
	return
}

// AddSolution sends a request to add a solution to a cluster, returns list of solutions added
func (c *APIClient) AddSolution(orgID, clusterID int, newSolution Solution) (sol *Solution, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/solutions", orgID, clusterID),
		ResponseObj:  &sol,
		PostObj:      newSolution,
		WantedStatus: 201,
	}
	err = c.runRequest(req)
	return
}

// AddSolutionFromJSON sends a request to add a solution to a cluster using the supplied JSON, returns list of solutions added
func (c *APIClient) AddSolutionFromJSON(orgID, clusterID int, solJSON string) (sol *Solution, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/solutions", orgID, clusterID),
		ResponseObj:  &sol,
		Payload:      strings.NewReader(solJSON),
		WantedStatus: 201,
	}
	err = c.runRequest(req)
	return
}

// WaitSolutionInstalled waits until solution is installed
func (c *APIClient) WaitSolutionInstalled(orgID, clusterID, solutionID, timeout int) error {
	for i := 1; i < timeout; i++ {
		sol, err := c.GetSolution(orgID, clusterID, solutionID)
		if err != nil {
			return err
		}
		if sol.State == SolutionInstalledStateString {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before solution reached state (%s)\n",
		timeout, SolutionInstalledStateString)
}

// WaitSolutionDeleted waits until solution disappears
func (c *APIClient) WaitSolutionDeleted(orgID, clusterID, solutionID, timeout int) error {
	for i := 1; i < timeout; i++ {
		_, err := c.GetSolution(orgID, clusterID, solutionID)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before solution deleted\n", timeout)
}
