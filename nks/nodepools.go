package nks

import (
	"fmt"
	"time"
)
//NodePoolRunningStateString state of a running node pool is active
const NodePoolRunningStateString = "active"

// NodePool defines the characteristics of a grouping of nodes
type NodePool struct {
	ID                 int       `json:"pk"`
	Name               string    `json:"name"`
	ClusterID          int       `json:"cluster"`
	InstanceID         string    `json:"instance_id"`
	Size               string    `json:"instance_size"`
	CPU                string    `json:"cpu,omitempty"`
	Memory             string    `json:"memory,omitempty"`
	Labels             string    `json:"labels,omitempty"`
	Autoscaled         bool      `json:"autoscaled"`
	MinCount           int       `json:"min_count,omitempty"`
	MaxCount           int       `json:"max_count,omitempty"`
	Zone               string    `json:"zone,omitempty"`
	ProviderSubnetID   string    `json:"provider_subnet_id,omitempty"`
	ProviderSubnetCidr string    `json:"provider_subnet_cidr,omitempty"`
	NodeCount          int       `json:"node_count"`
	Platform           string    `json:"platform"`
	Channel            string    `json:"channel"`
	Role               string    `json:"role,omitempty"`
	State              string    `json:"state,omitempty"`
	Default            bool      `json:"is_default"`
	Created            time.Time `json:"created"`
	Updated            time.Time `json:"updated,omitempty"`
	Deleted            time.Time `json:"deleted,omitempty"`
}

// NodeAddToPool is used for adding worker nodes to pools (endpoint /clusters/<cluster_id>/nodepools/add)
type NodeAddToPool struct {
	Count              int    `json:"node_count"`
	Group              string `json:"group,omitempty"`
	NodePoolID         int    `json:"node_pool"`
	Role               string `json:"role,omitempty"`
	Zone               string `json:"zone,omitempty"`
	ProviderSubnetID   string `json:"provider_subnet_id,omitempty"`
	ProviderSubnetCidr string `json:"provider_subnet_cidr,omitempty"`
}

// AddNodesToNodePool sends a request to add worker nodes to a nodepool, returns list of Node objects created
func (c *APIClient) AddNodesToNodePool(orgID, clusterID, nodepoolID int, newNode NodeAddToPool) (nodes []Node, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/nodepools/%d/add", orgID, clusterID, nodepoolID),
		ResponseObj:  &nodes,
		PostObj:      newNode,
		WantedStatus: 201,
	}
	err = c.runRequest(req)
	return
}

// GetNodePools gets the NodePools for a cluster, returns list of NodePool objects
func (c *APIClient) GetNodePools(orgID, clusterID int) (nps []NodePool, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/nodepools", orgID, clusterID),
		ResponseObj:  &nps,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetNodePool gets a NodePool for a cluster, returns NodePool object
func (c *APIClient) GetNodePool(orgID, clusterID, nodepoolID int) (np *NodePool, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/nodepools/%d", orgID, clusterID, nodepoolID),
		ResponseObj:  &np,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// CreateNodePool creates a new nodepool for a cluster, returns NodePool object
func (c *APIClient) CreateNodePool(orgID, clusterID int, newPool NodePool) (np *NodePool, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/nodepools", orgID, clusterID),
		ResponseObj:  &np,
		PostObj:      newPool,
		WantedStatus: 202,
	}
	err = c.runRequest(req)
	return
}

// DeleteNodePool deletes nodepool
func (c *APIClient) DeleteNodePool(orgID, clusterID, nodepoolID int) (err error) {
	req := &APIReq{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/nodepools/%d", orgID, clusterID, nodepoolID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}

// GetNodesInPool returns a list of nodes that are in given nodepool ID
func (c *APIClient) GetNodesInPool(orgID, clusterID, nodepoolID int) (rNodes []Node, err error) {
	nodes, err := c.GetNodes(orgID, clusterID)
	if err != nil {
		return
	}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].NodePoolID == nodepoolID {
			rNodes = append(rNodes, nodes[i])
		}
	}
	return
}

// WaitNodePoolProvisioned waits until nodepool reaches the running state (configured as const above)
func (c *APIClient) WaitNodePoolProvisioned(orgID, clusterID, nodepoolID, timeout int) error {
	//Check if the pool is ready
	for i := 1; i < timeout; i++ {
		node, err := c.GetNodePool(orgID, clusterID, nodepoolID)
		if err != nil {
			return err
		}
		if node.State == NodePoolRunningStateString {
			return nil
		}
		time.Sleep(time.Second)
	}

	//Check if nodes are ready
	nodes, err := c.GetNodesInPool(orgID, clusterID, nodepoolID)
	if err != nil {
		return err
	}
	for _, node := range nodes {
		c.WaitNodeProvisioned(orgID, clusterID, node.ID, 60)
	}
	return fmt.Errorf("timeout (%d seconds) reached before nodepool reached state (%s)",
		timeout, NodePoolRunningStateString)
}
