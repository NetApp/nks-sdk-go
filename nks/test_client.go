package nks

import (
	"errors"
	"gopkg.in/h2non/gock.v1"
	"os"
)

var mockServer = "http://foo.bar"

// NewTestClientFromEnv creates a new client from environment variables checks if mock or live tests needed
func NewTestClientFromEnv() (*APIClient, error) {
	testEnv := os.Getenv("TEST_ENV")
	if testEnv == "live" {
		token := os.Getenv("NKS_API_TOKEN")
		if token == "" {
			return nil, errors.New("Missing token env in NKS_API_TOKEN")
		}
		endpoint := os.Getenv("NKS_API_URL")
		if endpoint == "" {
			endpoint = defaultNKSApiURL
		}
		return NewClient(token, endpoint), nil
	} else {
		mockClient := NewClient("MOCK_TOKEN", mockServer)
		setupMockServer()
		return mockClient, nil

	}
}

func setupMockServer() {

	//cluster endpoints
	gock.New("http://foo.bar").
		Post("/orgs/1/clusters").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(200).
		JSON(mockCluster)

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(200).
		JSON(mockCluster)

	gock.New("http://foo.bar").
		Delete("/orgs/1/clusters/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(204).
		JSON(mockCluster)

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(200).
		JSON(mockClusters)

	//organization endpoints
	gock.New("http://foo.bar").
		Get("/orgs/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(200).
		JSON(mockOrg)

	gock.New("http://foo.bar").
		Get("/orgs").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(200).
		JSON(mockOrgs)

}

var mockClusters = `[{
	"pk": 1,
	"instance_id": "spcbsk7u4u",
	"name": "My Cluster",
	"org": 4,
	"provider": "do",
	"workspace": {
		"pk": 4,
		"name": "Default",
		"slug": "default",
		"org": 4,
		"is_default": true,
		"created": "2016-12-03T04:42:29.612800Z"
	},
	"k8s_version": "v1.8.3",
	"node_count": 3,
	"etcd_type": "self_hosted",
	"platform": "coreos",
	"channel": "stable",
	"region": "nyc1",
	"zone": "",
	"state": "running",
	"solutions": [
		{
			"pk": 4409,
			"name": "Helm Tiller",
			"solution": "helm_tiller",
			"state": "draft"
		}
	],
	"is_failed": false,
	"federation_role": null,
	"version_migrations": [],
	"k8s_rbac_enabled": true
}]`

var mockCluster = `{
	"pk": 1,
	"instance_id": "spcbsk7u4u",
	"name": "My Cluster",
	"org": 4,
	"provider": "do",
	"workspace": {
		"pk": 4,
		"name": "Default",
		"slug": "default",
		"org": 4,
		"is_default": true,
		"created": "2016-12-03T04:42:29.612800Z"
	},
	"k8s_version": "v1.8.3",
	"node_count": 3,
	"etcd_type": "self_hosted",
	"platform": "coreos",
	"channel": "stable",
	"region": "nyc1",
	"zone": "",
	"state": "running",
	"solutions": [
		{
			"pk": 4409,
			"name": "Helm Tiller",
			"solution": "helm_tiller",
			"state": "draft"
		}
	],
	"is_failed": false,
	"federation_role": null,
	"version_migrations": [],
	"k8s_rbac_enabled": true
}`

var mockOrgs = `[{
    "pk": 1,
    "name": "Holy Math",
    "slug": "holy-math",
    "logo": null,
    "enable_experimental_features": false,
    "created": "2019-12-05T20:01:06.784326Z",
    "updated": "2019-12-05T20:01:06.784346Z"
  }]`

var mockOrg = `{
    "pk": 1,
    "name": "Holy Math",
    "slug": "holy-math",
    "logo": null,
    "enable_experimental_features": false,
    "created": "2019-12-05T20:01:06.784326Z",
    "updated": "2019-12-05T20:01:06.784346Z"
  }`
