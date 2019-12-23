package mock

import (
	"os"
	"testing"

	"github.com/NetApp/nks-sdk-go/nks"
	"gopkg.in/h2non/gock.v1"
)

var mockServer = "http://foo.bar"
var MockClient *nks.APIClient

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	defer gock.Off()

	os.Exit(code)
}

func setup() {
	MockClient = nks.NewClient("MOCK_TOKEN", mockServer)
	setupMockServer()
}

func setupMockServer() {

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(200).
		JSON(mockCluster)

	gock.New("http://foo.bar").
		Get("/orgs").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(200).
		JSON(mockOrg)
}

var mockCluster = `[{
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
	"state": "draft",
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

var mockOrg = `  [{
    "pk": 1,
    "name": "Holy Math",
    "slug": "holy-math",
    "logo": null,
    "enable_experimental_features": false,
    "created": "2019-12-05T20:01:06.784326Z",
    "updated": "2019-12-05T20:01:06.784346Z"
  }]`
