package stackpointio

import "fmt"
import "time"

// Team struct
type Team struct {
	ID          int          `json:"pk"`
	Name        string       `json:"string"`
	Slug        string       `json:"string"`
	Org         int          `json:"org"`
	IsOrgWide   bool         `json:"is_org_wide"`
	Created     time.Time    `json:"created"`
	Memberships []Membership `json:"memberships"`
}

// Membership struct
type Membership struct {
	ID        int       `json:"pk"`
	User      User      `json:"user"`
	Team      int       `json:"team"`
	Created   time.Time `json:"created"`
}

// User struct
type User struct {
	ID         int       `json:"pk"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	FullName   string    `json:"full_name"`
	DateJoined time.Time `json:"date_joined"`
}

// GetTeams gets list of workspaces for Org ID
func (c *APIClient) GetTeams(orgID int) (team []Team, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/teams", orgID),
		ResponseObj:  &team,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetTeam returns team for Org ID and Team ID
func (c *APIClient) GetTeam(orgID, teamID int) (t *Team, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/teams/%d", orgID, teamID),
		ResponseObj:  &t,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// CreateTeam creates a team
func (c *APIClient) CreateTeam(orgID int, team Team) (t *Team, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/teams", orgID),
		ResponseObj:  &t,
		PostObj:      team,
		WantedStatus: 201,
	}
	err = c.runRequest(req)
	return
}

// DeleteTeam deletes a team
func (c *APIClient) DeleteTeam(orgID, teamID int) (err error) {
	req := &APIReq{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/orgs/%d/teams/%d", orgID, teamID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}
