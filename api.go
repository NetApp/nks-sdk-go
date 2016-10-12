package stackpointio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

// APIClient references an api token and an http endpoint
type APIClient struct {
	token      string
	endpoint   string
	httpClient *http.Client
}

// NewClient returns a new api client
func NewClient(token, endpoint string, client ...*http.Client) *APIClient {
	c := &APIClient{
		token:    token,
		endpoint: endpoint,
	}
	if len(client) != 0 {
		c.httpClient = client[0]
	} else {
		c.httpClient = http.DefaultClient
	}
	return c
}

func (client *APIClient) runRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+client.token)
	req.Header.Set("User-Agent", "Stackpoint Go SDK")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)

	if err == nil && resp.StatusCode >= 400 {
		err = fmt.Errorf("Status code %d", resp.StatusCode)
	}

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (client *APIClient) get(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", client.endpoint+path, nil)
	if err != nil {
		return nil, err
	}
	content, err := client.runRequest(req)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (client *APIClient) delete(path string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", client.endpoint+path, nil)
	if err != nil {
		return nil, err
	}
	content, err := client.runRequest(req)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (client *APIClient) post(path string, dataObject interface{}) ([]byte, error) {
	data, err := json.Marshal(dataObject)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", client.endpoint+path, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	content, err := client.runRequest(req)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// GetOrganizations retrieves data organizations that the client can access
func (client *APIClient) GetOrganizations() ([]Organization, error) {
	content, err := client.get("/orgs")
	if err != nil {
		return nil, err
	}
	var organizations []Organization
	err = json.Unmarshal(content, &organizations)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

// GetOrganization retrieves data for a single organization
func (client *APIClient) GetOrganization(organizationID int) (Organization, error) {
	path := fmt.Sprintf("/orgs/%d", organizationID)
	content, err := client.get(path)
	if err != nil {
		return Organization{}, err
	}
	var organization Organization
	err = json.Unmarshal(content, &organization)
	if err != nil {
		return Organization{}, err
	}
	return organization, nil
}

<<<<<<< HEAD
// GetUser gets the StackPointCloud user
=======
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
func (client *APIClient) GetUser() (User, error) {
	content, err := client.get("/rest-auth/user/")
	if err != nil {
		return User{}, err
	}
<<<<<<< HEAD
	glog.V(8).Info(string(content))
=======
	log.Debug(string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	var user User
	err = json.Unmarshal(content, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

<<<<<<< HEAD
// GetUserProfile gets details of StackPointCloud user profile
=======
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
func (client *APIClient) GetUserProfile(username string) (UserProfile, error) {
	path := fmt.Sprintf("/userprofile/%s", username)
	content, err := client.get(path)
	if err != nil {
		return UserProfile{}, err
	}
<<<<<<< HEAD
	glog.V(8).Info(string(content))
=======
	log.Debug(string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	var profile UserProfile
	err = json.Unmarshal(content, &profile)
	if err != nil {
		return UserProfile{}, err
	}
	return profile, nil
}

<<<<<<< HEAD
// GetClusters gets all clusters associated with an organization
=======
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
func (client *APIClient) GetClusters(organizationID int) ([]Cluster, error) {
	path := fmt.Sprintf("/orgs/%d/clusters", organizationID)
	content, err := client.get(path)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	glog.V(8).Info(string(content))
=======
	log.Debug(string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	var clusters []Cluster
	err = json.Unmarshal(content, &clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

<<<<<<< HEAD
// GetCluster gets a single cluster by primary ID and organization
=======
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
func (client *APIClient) GetCluster(organizationID, clusterID int) (Cluster, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d", organizationID, clusterID)
	content, err := client.get(path)
	if err != nil {
		return Cluster{}, err
	}
<<<<<<< HEAD
	glog.V(8).Info(string(content))
=======
	log.Debug(string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	var cluster Cluster
	err = json.Unmarshal(content, &cluster)
	if err != nil {
		return Cluster{}, err
	}
	return cluster, nil
}

<<<<<<< HEAD
// GetNodes gets the nodes associated with a cluster and organization
=======
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
func (client *APIClient) GetNodes(organizationID, clusterID int) ([]Node, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d/nodes", organizationID, clusterID)
	content, err := client.get(path)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	glog.V(8).Info(string(content))
=======
	log.Debug(string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	var nodes []Node
	err = json.Unmarshal(content, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// GetNode retrieves data for a single node
func (client *APIClient) GetNode(organizationID, clusterID, nodeID int) (Node, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d/nodes/%d", organizationID, clusterID, nodeID)
	content, err := client.get(path)
	if err != nil {
		return Node{}, err
	}
<<<<<<< HEAD
	glog.V(8).Info(string(content))
=======
	log.Debug(string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	var node Node
	err = json.Unmarshal(content, &node)
	if err != nil {
		return Node{}, err
	}
	return node, nil
}

<<<<<<< HEAD
// DeleteNode makes an API call to begin deleting a node, and returns the contents of the web response
=======
// DeleteNode makes an API call to delete the node
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
func (client *APIClient) DeleteNode(organizationID, clusterID, nodeID int) ([]byte, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d/nodes/%d", organizationID, clusterID, nodeID)
	content, err := client.delete(path)
	return content, err
	// if err != nil {
	// 	return Node{}, err
	// }
<<<<<<< HEAD
	// glog.V(8).Info(string(content))
=======
	// log.Debug(string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	// var node Node
	// err = json.Unmarshal(content, &node)
	// if err != nil {
	// 	return Node{}, err
	// }
	// return node, nil
}

// AddNodes sends a request to add nodes to a cluster, returns immediately
func (client *APIClient) AddNodes(organizationID, clusterID int, nodeAdd NodeAdd) (NodeAdd, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d/add_node", organizationID, clusterID)
	content, err := client.post(path, nodeAdd)
	if err != nil {
		return NodeAdd{}, err
	}
<<<<<<< HEAD
	glog.V(8).Info("add node response: " + string(content))
=======
	log.Debug("add node response: " + string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	var response NodeAdd
	err = json.Unmarshal(content, &response)
	if err != nil {
		return NodeAdd{}, err
	}
	return response, nil
}

<<<<<<< HEAD
// GetVolumes gets the Persistent Volumes attached to a cluster
=======
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
func (client *APIClient) GetVolumes(organizationID, clusterID int) ([]PersistentVolume, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d/volumes", organizationID, clusterID)
	content, err := client.get(path)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	glog.V(8).Info(string(content))
=======
	log.Debug(string(content))
>>>>>>> 1ca9779... Support for adding and deleting nodes, interaction with cluster-autoscaler
	var volumes []PersistentVolume
	err = json.Unmarshal(content, &volumes)
	if err != nil {
		return nil, err
	}
	return volumes, nil
}
