package nks

import (
	"fmt"
)

// Organization is the top level of the hierarchy
type Organization struct {
	ID   int    `json:"pk"`
	Name string `json:"name"`
}

// GetOrganizations gets the organizations the API token is associated with
func (c *APIClient) GetOrganizations() (orgs []Organization, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         "/orgs",
		ResponseObj:  &orgs,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetOrganization gets the organization by the supplied org ID
func (c *APIClient) GetOrganization(orgID int) (org *Organization, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d", orgID),
		ResponseObj:  &org,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}
