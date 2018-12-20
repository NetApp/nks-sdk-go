package nks

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
func (c *APIClient) GetKeysets(orgID int) (kss []Keyset, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/keysets", orgID),
		ResponseObj:  &kss,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetKeyset returns keyset for Org ID and Keyset ID
func (c *APIClient) GetKeyset(orgID, keysetID int) (ks *Keyset, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/keysets/%d", orgID, keysetID),
		ResponseObj:  &ks,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// CreateKeyset creates keyset
func (c *APIClient) CreateKeyset(orgID int, keyset Keyset) (ks *Keyset, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/keysets", orgID),
		ResponseObj:  &ks,
		PostObj:      keyset,
		WantedStatus: 201,
	}
	err = c.runRequest(req)
	return
}

// DeleteKeyset deletes keyset
func (c *APIClient) DeleteKeyset(orgID, keysetID int) (err error) {
	req := &APIReq{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/orgs/%d/keysets/%d", orgID, keysetID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}
