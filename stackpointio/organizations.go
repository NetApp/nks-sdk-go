package stackpointio

import (
	"fmt"
)

// Organization is the top level of the hierarchy
type Organization struct {
	Name string `json:"name"`
	ID   int    `json:"pk"`
}

// GetOrganizations retrieves data organizations that the client can access
func (c *APIClient) GetOrganizations() ([]Organization, error) {
	r := []Organization{}
	err := c.runRequest("GET", "/orgs", nil, &r, 200)
	return r, err
}

// GetOrganization retrieves data for a single organization
func (c *APIClient) GetOrganization(orgID int) (*Organization, error) {
	r := &Organization{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d", orgID), nil, r, 200)
	return r, err
}
