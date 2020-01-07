package nks

import (
	"errors"
	"net/http"
	"os"

	"gopkg.in/h2non/gock.v1"
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
	//solutions
	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/solutions/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockSolution)

	gock.New("http://foo.bar").
		Delete("/orgs/1/clusters/1/solutions/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusNoContent)

	gock.New("http://foo.bar").
		Post("/orgs/1/clusters/1/solutions").
		MatchType("json").
		MatchHeader("Authorization", "MOCK_TOKEN").
		MatchHeader("Content-Type", "application/json").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusCreated).
		JSON(mockSolution)

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/solutions").
		MatchType("json").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockSolutions)

	//istio mesh
	gock.New("http://foo.bar").
		Get("/orgs/1/workspaces/1/istio-meshes").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockIstioMeshes)

	gock.New("http://foo.bar").
		Post("/orgs/1/workspaces/1/istio-meshes").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(201).
		JSON(mockIstioMesh)

	gock.New("http://foo.bar").
		Delete("/orgs/1/istio-meshes/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusNoContent)

	//build logs
	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/logs").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockBuildLogs)

	//cluster endpoints
	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockCluster)

	gock.New("http://foo.bar").
		Delete("/orgs/1/clusters/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusNoContent)

	gock.New("http://foo.bar").
		Post("/orgs/1/clusters").
		MatchHeader("Authorization", "MOCK_TOKEN").
		MatchHeader("Content-Type", "application/json").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockCluster)

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockClusters)

	// workspace
	gock.New("http://foo.bar").
		Get("orgs/1/workspaces").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockWorkspaces)

	//organization endpoints
	gock.New("http://foo.bar").
		Get("/orgs/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockOrg)

	gock.New("http://foo.bar").
		Get("/orgs").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockOrgs)

}

var mockClusters = `[{
    "pk": 1,
	"instance_id": "spcbsk7u4u",
	"name": "My Cluster",
	"org": 4,
	"provider": "do",
	"workspace": {
		"pk": 1,
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
			"pk": 1,
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
		"pk": 1,
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
			"pk": 1,
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

var mockBuildLogs = `[
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "solution",
    "event_type": "solution_install",
    "event_state": "started",
    "message": "Installing helm tiller",
    "details": "",
    "reference": "helm_tiller",
    "created": "2020-01-07T14:47:46.771662Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "kubernetes",
    "event_type": "k8s_master_config",
    "event_state": "success",
    "message": "Successfully configured Kubernetes master",
    "details": "",
    "reference": "",
    "created": "2020-01-07T14:45:55.869876Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "kubernetes",
    "event_type": "k8s_master_config",
    "event_state": "started",
    "message": "Configuring Kubernetes master",
    "details": "",
    "reference": "",
    "created": "2020-01-07T14:44:45.229735Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "kubernetes",
    "event_type": "k8s_node_config",
    "event_state": "success",
    "message": "Successfully configured Kubernetes nodes",
    "details": "",
    "reference": "",
    "created": "2020-01-07T14:44:43.744242Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "kubernetes",
    "event_type": "k8s_node_config",
    "event_state": "started",
    "message": "Configuring Kubernetes nodes",
    "details": "",
    "reference": "",
    "created": "2020-01-07T14:41:17.835235Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "provider",
    "event_type": "provider_build",
    "event_state": "success",
    "message": "Successfully provisioned servers",
    "details": "",
    "reference": "aws",
    "created": "2020-01-07T14:41:16.258392Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "provider",
    "event_type": "provider_build",
    "event_state": "started",
    "message": "Provisioning servers",
    "details": "",
    "reference": "aws",
    "created": "2020-01-07T14:39:35.567654Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "provider",
    "event_type": "cloud_authorization",
    "event_state": "success",
    "message": "Finished checking AWS authorizations",
    "details": "",
    "reference": "VerifyAuthorization-337944-po8msl4bul",
    "created": "2020-01-07T14:38:41.339529Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "provider",
    "event_type": "cloud_authorization",
    "event_state": "started",
    "message": "Started checking AWS authorizations",
    "details": "",
    "reference": "VerifyAuthorization-337944-po8msl4bul",
    "created": "2020-01-07T14:38:40.036049Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "provider",
    "event_type": "cloud_network",
    "event_state": "success",
    "message": "Finished tagging resource",
    "details": "",
    "reference": "TagResource-337943-5itmpq9zzo",
    "created": "2020-01-07T14:38:39.964638Z"
  },
  {
    "pk": 1,
    "cluster": 30635,
    "event_category": "provider",
    "event_type": "cloud_network",
    "event_state": "started",
    "message": "Started tagging resource 'subnet-065b60e89fb5acf33'",
    "details": "",
    "reference": "TagResource-337943-5itmpq9zzo",
    "created": "2020-01-07T14:38:39.594358Z"
  }
]`

var mockWorkspaces = `[
  {
    "pk": 1,
    "name": "Demo Space",
    "slug": "demo-space-2",
    "org": 11949,
    "is_default": false,
    "is_pinned": false,
    "created": "2019-10-21T22:03:39.767164Z",
    "clusters": [],
    "federations": [],
    "user_solutions": [],
    "team_workspaces": [
      {
        "pk": 1,
        "team": {
          "pk": 1,
          "name": "Everyone",
          "slug": "everyone",
          "org": 11949,
          "is_org_wide": true,
          "created": "2018-10-09T05:43:48.656349Z"
        },
        "workspace": 26010,
        "created": "2019-10-21T22:03:51.049410Z"
      }
    ]
  },
  {
    "pk": 1,
    "name": "Dev Space",
    "slug": "dev-space-2",
    "org": 11949,
    "is_default": false,
    "is_pinned": false,
    "created": "2019-07-11T23:51:23.087470Z",
    "clusters": [],
    "federations": [],
    "user_solutions": [],
    "team_workspaces": [
      {
        "pk": 1,
        "team": {
          "pk": 1,
          "name": "Dev Team",
          "slug": "dev-team",
          "org": 11949,
          "is_org_wide": false,
          "created": "2019-07-11T23:51:13.452413Z"
        },
        "workspace": 20801,
        "created": "2019-07-11T23:51:23.089504Z"
      }
    ]
  },
  {
    "pk": 1,
    "name": "Default",
    "slug": "default-10167",
    "org": 11949,
    "is_default": true,
    "is_pinned": false,
    "created": "2018-10-09T05:43:48.650889Z",
    "clusters": [
      {
        "pk": 1,
        "name": "SDK GO TEST",
        "provider": "aws",
        "state": "running"
      }
    ],
    "federations": [],
    "user_solutions": [],
    "team_workspaces": [
      {
        "pk": 1,
        "team": {
          "pk": 1,
          "name": "Everyone",
          "slug": "everyone",
          "org": 11949,
          "is_org_wide": true,
          "created": "2018-10-09T05:43:48.656349Z"
        },
        "workspace": 12265,
        "created": "2018-10-09T05:43:48.671135Z"
      }
    ]
  }
]`

var mockSolutions = `[
  {
    "pk": 1,
    "name": "Helm Tiller",
    "instance_id": "sole28lbi1",
    "cluster": 30635,
    "solution": "helm_tiller",
    "installer": "ansible_custom",
    "keyset": null,
    "keyset_name": "",
    "version": "latest",
    "version_migrations": [],
    "state": "installed",
    "url": "",
    "username": "",
    "password": "",
    "max_nodes": null,
    "git_repo": "",
    "git_path": "",
    "initial": true,
    "config": {},
    "extra_data": {},
    "created": "2020-01-07T14:38:39.347700Z",
    "updated": "2020-01-07T14:38:39.347717Z",
    "is_deleteable": false
  },
  {
    "pk": 1,
    "name": "Gitlab",
    "instance_id": "solul8hytt",
    "cluster": 30635,
    "solution": "gitlab",
    "installer": "ansible_helm",
    "keyset": null,
    "keyset_name": "",
    "version": "latest",
    "version_migrations": [],
    "state": "installed",
    "url": "https://gitlab.gl.nettbievan.stackpoint.io",
    "username": "xx",
    "password": "xx",
    "max_nodes": null,
    "git_repo": "",
    "git_path": "",
    "initial": false,
    "config": {},
    "extra_data": {},
    "created": "2020-01-07T15:13:12.434024Z",
    "updated": "2020-01-07T15:13:12.434049Z",
    "is_deleteable": false
  }
]`

var mockSolution = `{
  "pk": 1,
  "name": "Gitlab",
  "instance_id": "solul8hytt",
  "cluster": 30635,
  "solution": "gitlab",
  "installer": "ansible_helm",
  "keyset": null,
  "keyset_name": "",
  "version": "latest",
  "version_migrations": [],
  "state": "installed",
  "url": "",
  "username": "",
  "password": "",
  "max_nodes": null,
  "git_repo": "",
  "git_path": "",
  "initial": false,
  "config": {},
  "extra_data": {},
  "created": "2020-01-07T15:13:12.434024Z",
  "updated": "2020-01-07T15:13:12.434049Z",
  "is_deleteable": false
}`

var mockIstioMeshes = `[
  {
    "pk": 1,
    "name": "New Mesh",
    "state": "active",
    "mesh_type": "cross_cluster",
    "workspace": {
      "pk": 1,
      "name": "Default",
      "slug": "default",
      "org": 4,
      "is_default": true,
      "created": "2016-12-03T04:42:29.612800Z"
    },
    "members": [
      {
        "cluster": {
          "pk": 1
        },
        "role": "host"
      },
      {
        "cluster": {
          "pk": 1
        },
        "role": "guest"
      }
    ]
  }
]`

var mockIstioMesh = `{
  "pk": 1,
  "name": "New Mesh",
  "state": "active",
  "mesh_type": "cross_cluster",
  "workspace": {
    "pk": 1,
    "name": "Default",
    "slug": "default",
    "org": 4,
    "is_default": true,
    "created": "2016-12-03T04:42:29.612800Z"
  },
  "members": [
    {
      "cluster": {
        "pk": 1
      },
      "role": "host"
    },
    {
      "cluster": {
        "pk": 1
      },
      "role": "guest"
    }
  ]
}`
