package nks

import "fmt"

// UserProfile struct
type UserProfile struct {
	ID           int             `json:"pk"`
	Username     string          `json:"username"`
	Email        string          `json:"email"`
	FirstName    string          `json:"first_name"`
	LastName     string          `json:"last_name"`
	FullName     string          `json:"full_name"`
	OrgMems      []OrgMembership `json:"org_memberships"`
	Keysets      []Keyset        `json:"keysets"`
	Subscription Subscription    `json:"subscription"`
}

type OrgMembership struct {
	ID        int          `json:"pk"`
	User      int          `json:"user"`
	Org       Organization `json:"org"`
	Role      string       `json:"role"`
	IsOwner   bool         `json:"is_owner"`
	IsManager bool         `json:"is_manager"`
	IsDefault bool         `json:"is_default"`
}

type Subscription struct {
	ID       int    `json:"pk"`
	State    string `json:"state"`
	IsActive bool   `json:"is_active"`
}

// GetUserProfile gets user profile for user (based on API token)
func (c *APIClient) GetUserProfile() (up []UserProfile, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         "/userprofile",
		ResponseObj:  &up,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetUserProfileDefaultOrg returns default organization ID for user profile
func (c *APIClient) GetUserProfileDefaultOrg(up *UserProfile) (int, error) {
	if up == nil {
		return 0, fmt.Errorf("userprofile object is nil")
	}
	for _, mem := range up.OrgMems {
		if mem.IsDefault {
			return mem.Org.ID, nil
		}
	}
	return 0, fmt.Errorf("no default org found in userprofile")
}

// GetUserProfileKeyset returns first keyset matching provider string
func (c *APIClient) GetUserProfileKeysetID(up *UserProfile, prov string) (int, error) {
	if up == nil {
		return 0, fmt.Errorf("userprofile object is nil")
	}
	for _, ks := range up.Keysets {
		if (prov == "user_ssh" && ks.Category == "user_ssh" && ks.IsDefault) || (ks.Category == "provider" && ks.Entity == prov) {
			return ks.ID, nil
		}
	}
	return 0, fmt.Errorf("no %s keyset found in userprofile", prov)
}
