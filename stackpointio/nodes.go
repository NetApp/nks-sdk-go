package stackpointio

import (
	"fmt"
	"time"
	"strings"
)

const NodeRunningStateString = "running"

// Node describes a node in a cluster.  The string field Size is provider-specific
type Node struct {
	ID                 int       `json:"pk"`
	Name               string    `json:"name,omitempty"`
	ClusterID          int       `json:"cluster"`
	InstanceID         string    `json:"instance_id"`
	NodePoolID         int       `json:"pool,omitempty"`
	NodePoolName       string    `json:"pool_name,omitempty"`
	Role               string    `json:"role"`
	Group              string    `json:"group_name,omitempty"`
	PrivateIP          string    `json:"private_ip"`
	PublicIP           string    `json:"public_ip"`
	Platform           string    `json:"platform"`
	Image              string    `json:"image"`
	Location           string    `json:"location"` // "location": "us-east-2:us-east-2a",
	ProviderSubnetID   string    `json:"provider_subnet_id,omitempty"`
	ProviderSubnetCidr string    `json:"provider_subnet_cidr,omitempty"`
	Size               string    `json:"size"`
	State              string    `json:"state,omitempty"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated,omitempty"`
}

// NodeAdd is used for adding master nodes only (endpoint /clusters/<cluster_id>/add_node)
type NodeAdd struct {
	Size               string `json:"size"`
	Count              int    `json:"node_count"`
	Group              string `json:"group,omitempty"`
	Role               string `json:"role,omitempty"`
	Zone               string `json:"zone,omitempty"`
	ProviderSubnetID   string `json:"provider_subnet_id,omitempty"`
	ProviderSubnetCidr string `json:"provider_subnet_cidr,omitempty"`
}

// GetNodes gets the nodes associated with a cluster and organization
func (c *APIClient) GetNodes(orgID, clusterID int) ([]Node, error) {
	r := []Node{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d/nodes", orgID, clusterID), nil, &r, 200)
	return r, err
}

// GetNode retrieves data for a single node
func (c *APIClient) GetNode(orgID, clusterID, nodeID int) (*Node, error) {
	r := &Node{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d/nodes/%d", orgID, clusterID, nodeID), nil, r, 200)
	return r, err
}

// DeleteNode makes an API call to begin deleting a node
func (c *APIClient) DeleteNode(orgID, clusterID, nodeID int) error {
	return c.runRequest("DELETE", fmt.Sprintf("/orgs/%d/clusters/%d/nodes/%d",
		orgID, clusterID, nodeID), nil, nil, 204)
}

// AddNodes sends a request to add master nodes to a cluster, returns list of Node objects created
func (c *APIClient) AddNode(orgID, clusterID int, newNode NodeAdd) ([]Node, error) {
	r := []Node{}
	err := c.runRequest("POST", fmt.Sprintf("/orgs/%d/clusters/%d/add_node", orgID, clusterID), newNode, &r, 201)
	return r, err
}

// GetNodeState returns state of node
func (c *APIClient) GetNodeState(orgID, clusterID, nodeID int) (string, error) {
	r := &Node{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d/nodes/%d", orgID, clusterID, nodeID), nil, r, 200)
	if err != nil {
		return "", err
	}
	return r.State, nil
}

// WaitNodeProvisioned waits until node reaches the running state (configured as const above)
func (c *APIClient) WaitNodeProvisioned(orgID, clusterID, nodeID int) error {
	for i := 1; ; i++ {
		state, err := c.GetNodeState(orgID, clusterID, nodeID)
		if err != nil {
			return err
		}
		if state == NodeRunningStateString {
			return nil
		}
		time.Sleep(time.Second)
	}
}

// WaitNodeDeleted waits until node disappears
func (c *APIClient) WaitNodeDeleted(orgID, clusterID, nodeID int) error {
        for i := 1; ; i++ {
                _, err := c.GetNodeState(orgID, clusterID, nodeID)
                if err != nil {
                        if strings.Contains(err.Error(), "404") {
                                return nil
                        }
                }
                time.Sleep(time.Second)
        }
}
