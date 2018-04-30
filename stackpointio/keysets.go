package stackpointio

import "fmt"

// Key struct
type Key struct {
	ID          int    `json:"pk"`
	Type        string `json:"key_type"`
	Value       string `json:"key"`
	Fingerprint string `json:"fingerprint"`
	User        int    `json:"user"`
}

// Keyset struct
type Keyset struct {
	Name       string `json:"name"`
	ID         int    `json:"pk"`
	Category   string `json:"category"`
	Entity     string `json:"entity"`
	Org        int    `json:"org"`
	Workspaces []int  `json:"workspaces"`
	User       int    `json:"user"`
	IsDefault  bool   `json:"is_default"`
	Keys       []Key  `json:"keys"`
	Created    string `json:"created"`
}

// GetKeysets gets list of keysets for Org ID
func (c *APIClient) GetKeysets(orgID int) ([]Keyset, error) {
	r := []Keyset{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/keysets", orgID), nil, &r, 200)
	return r, err
}

// GetKeyset returns keyset for Org ID and Keyset ID
func (c *APIClient) GetKeyset(orgID, keysetID int) (*Keyset, error) {
	r := &Keyset{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/keysets/%d", orgID, keysetID), nil, r, 200)
	return r, err
}

// CreateKeyset creates keyset
func (c *APIClient) CreateKeyset(orgID int, keyset Keyset) (*Keyset, error) {
	r := &Keyset{}
	err := c.runRequest("POST", fmt.Sprintf("/orgs/%d/keysets", orgID), keyset, r, 200)
	return r, err
}

// DeleteKeyset deletes keyset
func (c *APIClient) DeleteKeyset(orgID, keysetID int) error {
	return c.runRequest("DELETE", fmt.Sprintf("/orgs/%d/keysets/%d", orgID, keysetID), nil, nil, 204)
}
