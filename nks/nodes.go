package nks

import (
	"fmt"
	"strings"
	"time"
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
func (c *APIClient) GetNodes(orgID, clusterID int) (ns []Node, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/nodes", orgID, clusterID),
		ResponseObj:  &ns,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetNode retrieves data for a single node
func (c *APIClient) GetNode(orgID, clusterID, nodeID int) (n *Node, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/nodes/%d", orgID, clusterID, nodeID),
		ResponseObj:  &n,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// DeleteNode makes an API call to begin deleting a node
func (c *APIClient) DeleteNode(orgID, clusterID, nodeID int) (err error) {
	req := &APIReq{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/nodes/%d", orgID, clusterID, nodeID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}

// AddNodes sends a request to add master nodes to a cluster, returns list of Node objects created
func (c *APIClient) AddNode(orgID, clusterID int, newNode NodeAdd) (ns []Node, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/add_node", orgID, clusterID),
		ResponseObj:  &ns,
		PostObj:      newNode,
		WantedStatus: 201,
	}
	err = c.runRequest(req)
	return
}

// WaitNodeProvisioned waits until node reaches the running state (configured as const above)
func (c *APIClient) WaitNodeProvisioned(orgID, clusterID, nodeID, timeout int) error {
	for i := 1; i < timeout; i++ {
		node, err := c.GetNode(orgID, clusterID, nodeID)
		if err != nil {
			return err
		}
		if node.State == NodeRunningStateString {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before node reached state (%s)\n",
		timeout, NodeRunningStateString)
}

// WaitNodeDeleted waits until node disappears
func (c *APIClient) WaitNodeDeleted(orgID, clusterID, nodeID, timeout int) error {
	for i := 1; i < timeout; i++ {
		_, err := c.GetNode(orgID, clusterID, nodeID)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before node deleted\n", timeout)
}
