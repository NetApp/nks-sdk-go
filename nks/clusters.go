package nks

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	ClusterRunningStateString  = "running"
	ClusterBuildLogEventType   = "provider_build"
	ClusterBuildLogEventFailed = "failure"
)

// Cluster describes a Kubernetes cluster in the NetApp Kubernetes Service system
type Cluster struct {
	ID                          int                `json:"pk"`
	Name                        string             `json:"name"`
	OrganizationKey             int                `json:"org"`
	InstanceID                  string             `json:"instance_id"`
	Provider                    string             `json:"provider"`
	ProviderKey                 int                `json:"provider_keyset"`
	ProviderKeyName             string             `json:"provider_keyset_name"`
	ProviderResourceGp          string             `json:"provider_resource_group,omitempty"`
	ProviderNetworkID           string             `json:"provider_network_id,omitempty"`
	ProviderNetworkCdr          string             `json:"provider_network_cidr,omitempty"`
	ProviderSubnetID            string             `json:"provider_subnet_id,omitempty"`
	ProviderSubnetCidr          string             `json:"provider_subnet_cidr,omitempty"`
	ProviderBalancerID          string             `json:"provider_balancer_id,omitempty"`
	Region                      string             `json:"region"`
	Zone                        string             `json:"zone,omitempty"`
	State                       string             `json:"state,omitempty"`
	IsFailed                    bool               `json:"is_failed"`
	ProjectID                   string             `json:"project_id,omitempty"`
	Workspace                   Workspace          `json:"workspace"`
	Owner                       int                `json:"owner"`
	Notified                    bool               `json:"notified,omitempty"`
	KubernetesVersion           string             `json:"k8s_version"`
	Created                     time.Time          `json:"created"`
	Updated                     time.Time          `json:"updated,omitempty"`
	DashboardEnabled            bool               `json:"k8s_dashboard_enabled"`
	DashboardInstalled          bool               `json:"k8s_dashboard_installed"`
	KubeconfigPath              string             `json:"kubeconfig_path"`
	RbacEnabled                 bool               `json:"k8s_rbac_enabled"`
	MasterCount                 int                `json:"master_count"`
	WorkerCount                 int                `json:"worker_count"`
	MasterSize                  string             `json:"master_size"`
	WorkerSize                  string             `json:"worker_size"`
	NodeCount                   int                `json:"node_count"`
	MaxNodeCount                int                `json:"max_node_count"`
	MinNodeCount                int                `json:"min_node_count"`
	EtcdType                    string             `json:"etcd_type"`
	Platform                    string             `json:"platform"`
	Image                       string             `json:"image"`
	Channel                     string             `json:"channel"`
	SSHKeySet                   int                `json:"user_ssh_keyset"`
	Solutions                   []Solution         `json:"solutions"`
	NetworkComponents           []NetworkComponent `json:"network_components"`
	KubernetesMigrationVersions []string           `json:"version_migrations"`
}

// GetClusters gets all clusters associated with an organization
func (c *APIClient) GetClusters(orgID int) (cls []Cluster, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters", orgID),
		ResponseObj:  &cls,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetCluster gets a single cluster by primary ID and organization
func (c *APIClient) GetCluster(orgID, clusterID int) (cl *Cluster, err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d", orgID, clusterID),
		ResponseObj:  &cl,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// GetKubeConfig returns kubeconfig string.
func (c *APIClient) GetKubeConfig(orgID, clusterID int) (kubeconfig string, err error) {
	req := &APIReq{
		Method:         "GET",
		Path:           fmt.Sprintf("/orgs/%d/clusters/%d/kubeconfig", orgID, clusterID),
		ResponseObj:    kubeconfig,
		WantedStatus:   200,
		DontUnmarsahal: true,
	}
	err = c.runRequest(req)
	kubeconfig = req.ResponseString
	return
}

// CreateCluster requests cluster creation
func (c *APIClient) CreateCluster(orgID int, cluster Cluster) (cl *Cluster, err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/clusters", orgID),
		ResponseObj:  &cl,
		PostObj:      cluster,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	return
}

// DeleteCluster deletes cluster
func (c *APIClient) DeleteCluster(orgID, clusterID int) (err error) {
	req := &APIReq{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d", orgID, clusterID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}

// ForceDeleteCluster forcing deletion of a cluster
func (c *APIClient) ForceDeleteCluster(orgID, clusterID int) (err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/force_delete", orgID, clusterID),
		WantedStatus: 204,
	}
	err = c.runRequest(req)
	return
}

// UpgradeClusterToVersion upgrades cluster to supplied k8s version
func (c *APIClient) UpgradeClusterToVersion(cl Cluster, version string) (err error) {
	req := &APIReq{
		Method:       "POST",
		Path:         fmt.Sprintf("/orgs/%d/clusters/%d/migrate_version", cl.OrganizationKey, cl.ID),
		Payload:      strings.NewReader(`{"version":"` + version + `"}`),
		WantedStatus: 202,
	}
	err = c.runRequest(req)
	return
}

// UpgradeClusterToLatestVersion upgrades cluster to latest k8s version available
func (c *APIClient) UpgradeClusterToLatestVersion(cl Cluster) error {
	if len(cl.KubernetesMigrationVersions) < 1 {
		return fmt.Errorf("No migration versions listed for UpgradeClusterToLatestVersion\n")
	}
	majorV, minorV, patchV, err := convertVersionToInts(cl.KubernetesVersion)
	if err != nil {
		return err
	}
	// Cycle through new versions, find latest version
	for _, mv := range cl.KubernetesMigrationVersions {
		major, minor, patch, err := convertVersionToInts(mv)
		if err != nil {
			return err
		}
		if major > majorV {
			majorV, minorV, patchV = major, minor, patch
		} else if major == majorV {
			if minor > minorV {
				majorV, minorV, patchV = major, minor, patch
			} else if minor == minorV {
				if patch > patchV {
					majorV, minorV, patchV = major, minor, patch
				}
			}
		}
	}
	version := "v" + strconv.Itoa(majorV) + "." + strconv.Itoa(minorV) + "." + strconv.Itoa(patchV)
	return c.UpgradeClusterToVersion(cl, version)
}

// convertVersionToInts converts a kubernetes version major, minor, patch to ints
func convertVersionToInts(v string) (major, minor, patch int, err error) {
	versionPieces := strings.Split(v, ".")
	if len(versionPieces) != 3 {
		err = fmt.Errorf("Invalid version string fed to convertVersionToInts")
		return
	}
	major, err = strconv.Atoi(versionPieces[0][1:])
	if err != nil {
		return
	}
	minor, err = strconv.Atoi(versionPieces[1])
	if err != nil {
		return
	}
	patch, err = strconv.Atoi(versionPieces[2])
	if err != nil {
		return
	}
	return
}

// WaitClusterRunning waits until cluster reaches the running state (configured as const above)
func (c *APIClient) WaitClusterRunning(orgID, clusterID int, isProvisioning bool, timeout int) error {
	for i := 1; i < timeout; i++ {
		cl, err := c.GetCluster(orgID, clusterID)
		if err != nil {
			return err
		}
		// Check if state is running, return if it is
		if cl.State == ClusterRunningStateString {
			return nil
		}
		if isProvisioning {
			// Pull build logs, check if provider build failed
			bls, err := c.GetBuildLogs(orgID, clusterID)
			if err == nil {
				bl := c.GetBuildLogEventState(bls, ClusterBuildLogEventType)
				if bl != nil && bl.EventState == ClusterBuildLogEventFailed {
					return fmt.Errorf("Cluster build failed, build log message for event %s was: %s\n",
						ClusterBuildLogEventType, bl.Message)
				}
			}
			if cl.IsFailed {
				return fmt.Errorf("Cluster build failed, is_failed: %t\n", cl.IsFailed)
			}
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before cluster reached state (%s)\n",
		timeout, ClusterRunningStateString)
}

// WaitClusterDeleted waits until cluster disappears
func (c *APIClient) WaitClusterDeleted(orgID, clusterID, timeout int) error {
	for i := 1; i < timeout; i++ {
		_, err := c.GetCluster(orgID, clusterID)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout (%d seconds) reached before cluster deleted\n", timeout)
}
