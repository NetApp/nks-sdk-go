package stackpointio

import (
	"fmt"
	"time"
)

const clusterRunningStateString = "running"

// Cluster describes a Kubernetes cluster in the StackPointCloud system
type Cluster struct {
	ID                 int        `json:"pk"`
	Name               string     `json:"name"`
	OrganizationKey    int        `json:"org"`
	InstanceID         string     `json:"instance_id"`
	Provider           string     `json:"provider"`
	ProviderKey        int        `json:"provider_keyset"`
	ProviderKeyName    string     `json:"provider_keyset_name"`
	ProviderResourceGp string     `json:"provider_resource_group,omitempty"`
	ProviderNetworkID  string     `json:"provider_network_id,omitempty"`
	ProviderNetworkCdr string     `json:"provider_network_cidr,omitempty"`
	ProviderSubnetID   string     `json:"provider_subnet_id,omitempty"`
	ProviderSubnetCidr string     `json:"provider_subnet_cidr,omitempty"`
	ProviderBalancerID string     `json:"provider_balancer_id,omitempty"`
	Region             string     `json:"region"`
	Zone               string     `json:"zone,omitempty"`
	State              string     `json:"state,omitempty"`
	ProjectID          string     `json:"project_id,omitempty"`
	Owner              int        `json:"owner"`
	Notified           bool       `json:"notified,omitempty"`
	KubernetesVersion  string     `json:"k8s_version"`
	Created            time.Time  `json:"created"`
	Updated            time.Time  `json:"updated,omitempty"`
	DashboardEnabled   bool       `json:"k8s_dashboard_enabled"`
	DashboardInstalled bool       `json:"k8s_dashboard_installed"`
	KubeconfigPath     string     `json:"kubeconfig_path"`
	RbacEnabled        bool       `json:"k8s_rbac_enabled"`
	MasterCount        int        `json:"master_count"`
	WorkerCount        int        `json:"worker_count"`
	MasterSize         string     `json:"master_size"`
	WorkerSize         string     `json:"worker_size"`
	NodeCount          int        `json:"node_count"`
	EtcdType           string     `json:"etcd_type"`
	Platform           string     `json:"platform"`
	Image              string     `json:"image"`
	Channel            string     `json:"channel"`
	SSHKeySet          int        `json:"user_ssh_keyset"`
	Solutions          []Solution `json:"solutions"`
}

// GetClusters gets all clusters associated with an organization
func (c *APIClient) GetClusters(orgID int) ([]Cluster, error) {
	r := []Cluster{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters", orgID), nil, &r, 200)
	return r, err
}

// GetCluster gets a single cluster by primary ID and organization
func (c *APIClient) GetCluster(orgID, clusterID int) (*Cluster, error) {
	r := &Cluster{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d", orgID, clusterID), nil, r, 200)
	return r, err
}

// CreateCluster requests cluster creation
func (c *APIClient) CreateCluster(orgID int, cluster Cluster) (*Cluster, error) {
	r := &Cluster{}
	err := c.runRequest("POST", fmt.Sprintf("/orgs/%d/clusters", orgID), cluster, r, 200)
	return r, err
}

// DeleteCluster deletes cluster
func (c *APIClient) DeleteCluster(orgID, clusterID int) error {
	return c.runRequest("DELETE", fmt.Sprintf("/orgs/%d/clusters/%d", orgID, clusterID), nil, nil, 204)
}

// GetClusterState returns state of cluster
func (c *APIClient) GetClusterState(orgID, clusterID int) (string, error) {
	r := &Cluster{}
	err := c.runRequest("GET", fmt.Sprintf("/orgs/%d/clusters/%d", orgID, clusterID), nil, r, 200)
	if err != nil {
		return "", err
	}
	return r.State, nil
}

// WaitClusterProvisioned waits until cluster reaches the running state (configured as const above)
func (c *APIClient) WaitClusterProvisioned(orgID, clusterID int) error {
	for i := 1; ; i++ {
		state, err := c.GetClusterState(orgID, clusterID)
		if err != nil {
			return err
		}
		if state == clusterRunningStateString {
			return nil
		}
		time.Sleep(time.Second)
	}
}
