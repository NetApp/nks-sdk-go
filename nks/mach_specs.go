package nks

import (
	"fmt"
	"sort"
	"strings"
)

// ProviderSpecs instance structs
type ProviderSpecs struct {
	Name    string      `json:"name"`
	Filters interface{} `json:"filters"`
	Config  interface{} `json:"config"`
}

// Instance name and specs
type Instance struct {
	Name  string
	Specs MachineSpecs
}

// MachineSpecs machines specs details
type MachineSpecs struct {
	Memory int
	CPU    int
	Name   string
}

// GetInstanceSpecs returns list of machine types for cloud provider type
func (c *APIClient) GetInstanceSpecs(prov, endpoint string) ([]Instance, error) {
	ps := []ProviderSpecs{}
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("%s/meta/provider-instances?provider=%s", strings.TrimRight(endpoint, "/"), prov),
		ResponseObj:  &ps,
		WantedStatus: 200,
	}
	err := c.runRequest(req)
	if err != nil {
		return nil, err
	}
	// Returns list of objects, not array, so we need to read it in manually and decode JSON
	var instances []Instance
	machines := ps[0].Config.(map[string]interface{}) // at instance_name: { cpu,memory,name }
	for k, v := range machines {
		specs := v.(map[string]interface{}) // at cpu: x, memory: y, name: z
		mach := new(MachineSpecs)
		for k2, v2 := range specs {
			switch k2 {
			case "cpu":
				mach.CPU = int(v2.(float64))
			case "name":
				mach.Name = v2.(string)
			case "memory":
				mach.Memory = int(v2.(float64))
			}
		}
		instance := Instance{Name: k, Specs: *mach}
		instances = append(instances, instance)
	}
	sort.Slice(instances, func(i, j int) bool { return instances[i].Specs.Memory < instances[j].Specs.Memory })
	return instances, nil
}

// GetFormattedInstanceList takes a list of Instance objects and makes a formatted list of strings for the user
func GetFormattedInstanceList(instances []Instance) []string {
	var r []string
	for _, opt := range instances {
		r = append(r, fmt.Sprintf("%s\t(%s -- cpu: %d, memory: %d)",
			opt.Name, opt.Specs.Name, opt.Specs.CPU, opt.Specs.Memory))
	}
	return r
}

// InstanceInList returns true if instance is in list of Instances
func InstanceInList(instances []Instance, i string) bool {
	for _, opt := range instances {
		if opt.Name == i {
			return true
		}
	}
	return false
}
