package nks

import "fmt"

// Tunnel describes credentials for a Dispatch tunnel
type Tunnel struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

// GetTunnel gets credentials for a tunnel within an organization with tunnel ID
func (c *APIClient) GetTunnel(orgID int, tunnelID string) (t *Tunnel, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/tunnels/%s", orgID, tunnelID),
		ResponseObj:  &t,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}
