package stackpointio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// Client references an api token and an http endpoint
type Client struct {
	token      string
	endpoint   string
	httpClient *http.Client
}

// NewClient returns a new api client
func NewClient(token, endpoint string, client ...*http.Client) *Client {
	c := &Client{
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

func (client *Client) runRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+client.token)
	req.Header.Set("User-Agent", "Stackpoint Go SDK")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)

	if resp.StatusCode >= 400 {
		err = fmt.Errorf("Status code %d", resp.StatusCode)
	}

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (client *Client) get(path string) ([]byte, error) {
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

func (client *Client) post(path string, dataObject interface{}) ([]byte, error) {
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

func (client *Client) GetOrganizations() ([]Organization, error) {
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

func (client *Client) GetOrganization(organizationID int) (Organization, error) {
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

func (client *Client) GetUser() (User, error) {
	content, err := client.get("/rest-auth/user/")
	if err != nil {
		return User{}, err
	}
	log.Info(string(content))
	var user User
	err = json.Unmarshal(content, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (client *Client) GetUserProfile(username string) (UserProfile, error) {
	path := fmt.Sprintf("/userprofile/%s", username)
	content, err := client.get(path)
	if err != nil {
		return UserProfile{}, err
	}
	log.Info(string(content))
	var profile UserProfile
	err = json.Unmarshal(content, &profile)
	if err != nil {
		return UserProfile{}, err
	}
	return profile, nil
}

func (client *Client) GetClusters(organizationID int) ([]Cluster, error) {
	path := fmt.Sprintf("/orgs/%d/clusters", organizationID)
	content, err := client.get(path)
	if err != nil {
		return nil, err
	}
	log.Info(string(content))
	var clusters []Cluster
	err = json.Unmarshal(content, &clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

func (client *Client) GetCluster(organizationID, clusterID int) (Cluster, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d", organizationID, clusterID)
	content, err := client.get(path)
	if err != nil {
		return Cluster{}, err
	}
	log.Info(string(content))
	var cluster Cluster
	err = json.Unmarshal(content, &cluster)
	if err != nil {
		return Cluster{}, err
	}
	return cluster, nil
}

func (client *Client) GetNodes(organizationID, clusterID int) ([]Node, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d/nodes", organizationID, clusterID)
	content, err := client.get(path)
	if err != nil {
		return nil, err
	}
	log.Info(string(content))
	var nodes []Node
	err = json.Unmarshal(content, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (client *Client) AddNodes(organizationID, clusterID int, nodeAdd NodeAdd) ([]byte, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d/add_node", organizationID, clusterID)
	content, err := client.post(path, nodeAdd)
	if err != nil {
		return nil, err
	}
	log.Info("add node response: " + string(content))
	return content, nil
}

func (client *Client) GetVolumes(organizationID, clusterID int) ([]PersistentVolume, error) {
	path := fmt.Sprintf("/orgs/%d/clusters/%d/volumes", organizationID, clusterID)
	content, err := client.get(path)
	if err != nil {
		return nil, err
	}
	log.Info(string(content))
	var volumes []PersistentVolume
	err = json.Unmarshal(content, &volumes)
	if err != nil {
		return nil, err
	}
	return volumes, nil
}
