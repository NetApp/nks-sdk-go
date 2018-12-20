package nks

import (
	"fmt"
	"time"
)

const BuildLogEventStateSuccess = "success"

// BuildLog struct to hold buildlog entry
type BuildLog struct {
	ID            int       `json:"pk"`
	ClusterID     int       `json:"cluster"`
	EventCategory string    `json:"event_category"`
	EventType     string    `json:"event_type"`
	EventState    string    `json:"event_state"`
	Message       string    `json:"message"`
	Details       string    `json:"details"`
	Reference     string    `json:"reference"`
	Created       time.Time `json:"created"`
}

// GetBuildLogs gets the list of buildlog entries associated with a cluster and organization
func (c *APIClient) GetBuildLogs(orgID, clusterID int) (bls []BuildLog, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/logs", orgID, clusterID),
		ResponseObj:  &bls,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetBuildLog retrieves buildlog entry for buildlog ID, or error if not found
func (c *APIClient) GetBuildLog(bls []BuildLog, buildlogID int) (*BuildLog, error) {
	for i, _ := range bls {
		if bls[i].ID == buildlogID {
			return &bls[i], nil
		}
	}
	return nil, fmt.Errorf("No build log found by the ID: %d\n", buildlogID)
}

// GetBuildLogEventState takes a list of buildlogs, returns most recent build log
// entry that matches eventType string, or returns nil if not found
func (c *APIClient) GetBuildLogEventState(bls []BuildLog, eventType string) *BuildLog {
	for i, _ := range bls {
		if bls[i].EventType == eventType {
			return &bls[i]
		}
	}
	return nil
}

// WaitEventSuccess waits until event type has success state
func (c *APIClient) WaitBuildLogEventSuccess(orgID, clusterID, timeout int, eventType string) error {
	for i := 1; i < timeout; i++ {
		bls, err := c.GetBuildLogs(orgID, clusterID)
		if err != nil {
			return err
		}
		bl := c.GetBuildLogEventState(bls, eventType)
		if bl != nil && bl.EventState == BuildLogEventStateSuccess {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before eventType (%s) reached state (%s)\n",
		timeout, eventType, BuildLogEventStateSuccess)
}
