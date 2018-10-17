package stackpointio

import (
	"fmt"
	"strings"
	"time"
)

const (
	SolutionInstalledStateString = "installed"
	HelmTillerInstallWaitTimeout = 120
	HelmTillerSolutionName       = "helm_tiller"
)

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
	err = c.WaitHelmTillerInstalled(orgID, clusterID, HelmTillerInstallWaitTimeout)
	if err != nil {
		return
	}
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

// FindSolutionByName finds solution ID by name given
func (c *APIClient) FindSolutionByName(orgID, clusterID int, solutionName string) (int, error) {
	sols, err := c.GetSolutions(orgID, clusterID)
	if err != nil {
		return 0, err
	}
	for _, sol := range sols {
		if sol.Solution == solutionName {
			return sol.ID, nil
		}
	}
	return 0, fmt.Errorf("No solution by by the name: %s\n", solutionName)
}

// WaitHelmTillerInstalled waits until Tiller is installed, or errors if Tiller is not installed or not going to install,
// since Tiller is needed for any Helm install
func (c *APIClient) WaitHelmTillerInstalled(orgID, clusterID, timeout int) error {
	helmID, err := c.FindSolutionByName(orgID, clusterID, HelmTillerSolutionName)
	if err != nil {
		return err
	}
	for i := 1; i < timeout; i++ {
		sol, err := c.GetSolution(orgID, clusterID, helmID)
		if err != nil {
			return err
		}
		if sol.State == SolutionInstalledStateString {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before Tiller reached state (%s)\n",
		timeout, SolutionInstalledStateString)
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
