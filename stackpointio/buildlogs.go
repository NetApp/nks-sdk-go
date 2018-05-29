package stackpointio

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
	err = c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d/logs", orgID, clusterID), nil, &bls, 200)
	return
}

// GetBuildLog retrieves buildlog entry for buildlog ID
func (c *APIClient) GetBuildLog(orgID, clusterID, buildlogID int) (bl *BuildLog, err error) {
	err = c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d/logs/%d",
		orgID, clusterID, buildlogID), nil, bl, 200)
	return
}

// GetBuildLogEventState takes list of buildlogs, returns state of entry for eventType string
func (c *APIClient) GetBuildLogEventState(bls []BuildLog, eventType string) string {
	for _, bl := range bls {
		if bl.EventType == eventType {
			return bl.EventState
		}
	}
	return ""
}

// WaitEventSuccess waits until event type has success state
func (c *APIClient) WaitBuildLogEventSuccess(orgID, clusterID, timeout int, eventType string) error {
	for i := 1; i < timeout; i++ {
		bls, err := c.GetBuildLogs(orgID, clusterID)
		if err != nil {
			return err
		}
		if BuildLogEventStateSuccess == c.GetBuildLogEventState(bls, eventType) {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before eventType (%s) reached state (%s)\n",
		timeout, eventType, BuildLogEventStateSuccess)
}
