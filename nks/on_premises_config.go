package nks

import (
	"fmt"
	"sort"
)

// OnPremisesConfig is a configuration used during on-premises region registration
type OnPremisesConfig struct {
	Name string      `json:"name"`
	Info Info        `json:"filters"`
	Data interface{} `json:"config"`
}

// Info about the config file
type Info struct {
	Version string `json:"version,omitempty"`
	Patch   int    `json:"patch,omitempty"`
}

// GetOnPremisesConfigs returns list of on-premises configs
func (c *APIClient) GetOnPremisesConfigs() (opcList []OnPremisesConfig, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         "/meta/on-premises-config",
		ResponseObj:  &opcList,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetOnPremisesConfig returns a on-premises config of a specific version
func (c *APIClient) GetOnPremisesConfig(version string) (*OnPremisesConfig, bool, error) {
	opcList := []OnPremisesConfig{}
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/meta/on-premises-config?version=%s", version),
		ResponseObj:  &opcList,
		WantedStatus: 200,
	}
	err := c.runRequest(req)
	if err != nil {
		return nil, false, err
	}

	if len(opcList) == 0 {
		return nil, false, nil
	}

	sort.Slice(opcList, func(i, j int) bool {
		return opcList[i].Info.Patch > opcList[j].Info.Patch
	})

	return &opcList[0], true, nil
}
