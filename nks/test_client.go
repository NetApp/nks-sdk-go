package nks

import (
	"errors"
	"net/http"
	"os"

	"gopkg.in/h2non/gock.v1"
)

var mockServer = "http://foo.bar"

// NewTestClientFromEnv creates either a mock or live test client based on env variable NKS_TEST_ENV
func NewTestClientFromEnv() (*APIClient, error) {
	testEnv := os.Getenv("NKS_TEST_ENV")
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

	//nodes
	gock.New("http://foo.bar").
		Post("/orgs/1/clusters/1/add_node").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusCreated).
		JSON(mockNodes)

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/nodes/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockNode)

	gock.New("http://foo.bar").
		Delete("/orgs/1/clusters/1/nodes/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusNoContent)

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/nodes").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockNodes)

	//nodepools
	gock.New("http://foo.bar").
		Post("/orgs/1/clusters/1/nodepools/1/add").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusCreated).
		JSON(mockNodes)

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/nodepools/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockNodePool)

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/nodepools").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockNodePools)

	gock.New("http://foo.bar").
		Post("/orgs/1/clusters/1/nodepools").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusAccepted).
		JSON(mockNodePool)

	gock.New("http://foo.bar").
		Delete("/orgs/1/clusters/1/nodepools").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusNoContent)

	//mach specs
	gock.New("http://foo.bar").
		Get("/meta/provider-instances").
		MatchParam("provider", "aws").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockMachSpecs)

	gock.New("http://foo.bar").
		Get("/meta/provider-instances").
		MatchParam("provider", "gce").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockMachSpecs)

	gock.New("http://foo.bar").
		Get("/meta/provider-instances").
		MatchParam("provider", "azure").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockMachSpecs)

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
		Reply(http.StatusAccepted)

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

	//teams
	gock.New("http://foo.bar").
		Get("/orgs/1/teams/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockTeam)

	gock.New("http://foo.bar").
		Delete("/orgs/1/teams/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusNoContent)

	gock.New("http://foo.bar").
		Get("/orgs/1/teams").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockTeams)

	gock.New("http://foo.bar").
		Post("/orgs/1/teams").
		MatchHeader("Authorization", "MOCK_TOKEN").
		MatchHeader("Content-Type", "application/json").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusCreated).
		JSON(mockTeam)

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
    "name": "haproxy",
    "instance_id": "solul8hytt",
    "cluster": 30635,
    "solution": "haproxy",
    "installer": "ansible_helm",
    "keyset": null,
    "keyset_name": "",
    "version": "latest",
    "version_migrations": [],
    "state": "installed",    
    "max_nodes": null,
    "initial": false,
    "config": {},
    "extra_data": {},
    "created": "2020-01-07T15:13:12.434024Z",
    "updated": "2020-01-07T15:13:12.434049Z",
    "is_deleteable": false
  },
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
  }
]`

var mockSolution = `{
  "pk": 1,
  "name": "haproxy",
  "instance_id": "solul8hytt",
  "cluster": 30635,
  "solution": "haproxy",
  "installer": "ansible_helm",
  "keyset": null,
  "keyset_name": "",
  "version": "latest",
  "version_migrations": [],
  "state": "installed",
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

var mockMachSpecs = `[
{
	"name":"provider-instances",
	"filters":{
		"gpuEnabled":true,
		"provider":"aws"
	},
	"config":{
		"m3.large":{
			"memory":7500,
			"cpu":2,
			"value":"m3.large",
			"name":"m3.large"
		},
		"m4.2xlarge":{
			"memory":32000,
			"cpu":8,
			"value":"m4.2xlarge",
			"name":"m4.2xlarge"
		},
		"m5.xlarge":{
			"memory":16000,
			"cpu":4,
			"value":"m5.xlarge",
			"name":"m5.xlarge"
		},
		"p2.16xlarge":{
			"gpu":16,
			"memory":732000,
			"cpu":8,
			"value":"p2.16xlarge",
			"name":"p2.16xlarge"
		},
		"t2.large":{
			"memory":8000,
			"cpu":2,
			"value":"t2.large",
			"name":"t2.large"
		},
		"m3.2xlarge":{
			"memory":30000,
			"cpu":8,
			"value":"m3.2xlarge",
			"name":"m3.2xlarge"
		},
		"m5.24xlarge":{
			"memory":384000,
			"cpu":96,
			"value":"m5.24xlarge",
			"name":"m5.24xlarge"
		},
		"r4.large":{
			"memory":15250,
			"cpu":2,
			"value":"r4.large",
			"name":"r4.large"
		},
		"r4.4xlarge":{
			"memory":122000,
			"cpu":16,
			"value":"r4.4xlarge",
			"name":"r4.4xlarge"
		},
		"r4.2xlarge":{
			"memory":61000,
			"cpu":8,
			"value":"r4.2xlarge",
			"name":"r4.2xlarge"
		},
		"c4.8xlarge":{
			"memory":60000,
			"cpu":36,
			"value":"c4.8xlarge",
			"name":"c4.8xlarge"
		},
		"t2.2xlarge":{
			"memory":32000,
			"cpu":8,
			"value":"t2.2xlarge",
			"name":"t2.2xlarge"
		},
		"m4.16xlarge":{
			"memory":256000,
			"cpu":64,
			"value":"m4.16xlarge",
			"name":"m4.16xlarge"
		},
		"p2.8xlarge":{
			"gpu":8,
			"memory":488000,
			"cpu":4,
			"value":"p2.8xlarge",
			"name":"p2.8xlarge"
		},
		"p2.xlarge":{
			"gpu":1,
			"memory":61000,
			"cpu":2,
			"value":"p2.xlarge",
			"name":"p2.xlarge"
		},
		"p3.8xlarge":{
			"gpu":4,
			"memory":244000,
			"cpu":32,
			"value":"p3.8xlarge",
			"name":"p3.8xlarge"
		},
		"c4.xlarge":{
			"memory":7500,
			"cpu":4,
			"value":"c4.xlarge",
			"name":"c4.xlarge"
		},
		"m5.12xlarge":{
			"memory":192000,
			"cpu":48,
			"value":"m5.12xlarge",
			"name":"m5.12xlarge"
		},
		"m4.large":{
			"memory":8000,
			"cpu":2,
			"value":"m4.large",
			"name":"m4.large"
		},
		"c5.4xlarge":{
			"memory":32000,
			"cpu":16,
			"value":"c5.4xlarge",
			"name":"c5.4xlarge"
		},
		"t2.medium":{
			"memory":4000,
			"cpu":2,
			"value":"t2.medium",
			"name":"t2.medium"
		},
		"c5.large":{
			"memory":4000,
			"cpu":2,
			"value":"c5.large",
			"name":"c5.large"
		},
		"m5.2xlarge":{
			"memory":32000,
			"cpu":8,
			"value":"m5.2xlarge",
			"name":"m5.2xlarge"
		},
		"t2.xlarge":{
			"memory":16000,
			"cpu":4,
			"value":"t2.xlarge",
			"name":"t2.xlarge"
		},
		"m5.4xlarge":{
			"memory":64000,
			"cpu":16,
			"value":"m5.4xlarge",
			"name":"m5.4xlarge"
		},
		"c4.4xlarge":{
			"memory":30000,
			"cpu":16,
			"value":"c4.4xlarge",
			"name":"c4.4xlarge"
		},
		"m5.large":{
			"memory":8000,
			"cpu":2,
			"value":"m5.large",
			"name":"m5.large"
		},
		"g3.4xlarge":{
			"gpu":1,
			"memory":122000,
			"cpu":16,
			"value":"g3.4xlarge",
			"name":"g3.4xlarge"
		},
		"c5.2xlarge":{
			"memory":16000,
			"cpu":8,
			"value":"c5.2xlarge",
			"name":"c5.2xlarge"
		},
		"m4.xlarge":{
			"memory":16000,
			"cpu":4,
			"value":"m4.xlarge",
			"name":"m4.xlarge"
		},
		"c4.2xlarge":{
			"memory":15000,
			"cpu":8,
			"value":"c4.2xlarge",
			"name":"c4.2xlarge"
		},
		"c5.xlarge":{
			"memory":8000,
			"cpu":4,
			"value":"c5.xlarge",
			"name":"c5.xlarge"
		},
		"r4.16xlarge":{
			"memory":488000,
			"cpu":64,
			"value":"r4.16xlarge",
			"name":"r4.16xlarge"
		},
		"r4.xlarge":{
			"memory":30500,
			"cpu":4,
			"value":"r4.xlarge",
			"name":"r4.xlarge"
		},
		"m4.10xlarge":{
			"memory":160000,
			"cpu":40,
			"value":"m4.10xlarge",
			"name":"m4.10xlarge"
		},
		"g3.8xlarge":{
			"gpu":2,
			"memory":244000,
			"cpu":32,
			"value":"g3.8xlarge",
			"name":"g3.8xlarge"
		},
		"c4.large":{
			"memory":3750,
			"cpu":2,
			"value":"c4.large",
			"name":"c4.large"
		},
		"r4.8xlarge":{
			"memory":244000,
			"cpu":32,
			"value":"r4.8xlarge",
			"name":"r4.8xlarge"
		},
		"c5.9xlarge":{
			"memory":72000,
			"cpu":36,
			"value":"c5.9xlarge",
			"name":"c5.9xlarge"
		},
		"g3.16xlarge":{
			"gpu":4,
			"memory":488000,
			"cpu":64,
			"value":"g3.16xlarge",
			"name":"g3.16xlarge"
		},
		"p3.16xlarge":{
			"gpu":8,
			"memory":488000,
			"cpu":64,
			"value":"p3.16xlarge",
			"name":"p3.16xlarge"
		},
		"m4.4xlarge":{
			"memory":64000,
			"cpu":16,
			"value":"m4.4xlarge",
			"name":"m4.4xlarge"
		},
		"p3.2xlarge":{
			"gpu":1,
			"memory":61000,
			"cpu":8,
			"value":"p3.2xlarge",
			"name":"p3.2xlarge"
		},
		"m3.xlarge":{
			"memory":15000,
			"cpu":4,
			"value":"m3.xlarge",
			"name":"m3.xlarge"
		}
	}
}
]`

var mockNodes = `[
    {
        "pk":1,
        "cluster":30667,
        "pool":null,
        "pool_name":"",
        "instance_id":"netp2h4pg8-master-1",
        "provider_node_id":"i-038f00f3b164f86bb",
        "role":"master",
        "group_name":"",
        "private_ip":"172.23.4.102",
        "public_ip":"54.218.62.37",
        "platform":"ubuntu",
        "image":"ami-0a7d051a1c4b54f65",
        "channel":"18.04-lts",
        "etcd_state":"kube_etcd",
        "root_disk_size":50,
        "gpu_instance_size":"",
        "gpu_core_count":null,
        "location":"us-west-2:us-west-2a",
        "zone":"us-west-2a",
        "provider_subnet_id":"subnet-065b60e89fb5acf33",
        "provider_subnet_cidr":"172.23.4.0/24",
        "size":"m5.large",
        "state":"running",
        "is_failed":false,
        "created":"2020-01-08T11:39:53.872530Z",
        "updated":"2020-01-08T11:39:53.872541Z"
    },
    {
        "pk":1,
        "cluster":30667,
        "pool":27288,
        "pool_name":"Default Worker Pool",
        "instance_id":"netp2h4pg8-worker-1",
        "provider_node_id":"i-0dd18342301f04607",
        "role":"worker",
        "group_name":"autoscaling-netp2h4pg8-pool-1",
        "private_ip":"172.23.4.104",
        "public_ip":"18.237.248.178",
        "platform":"ubuntu",
        "image":"ami-0a7d051a1c4b54f65",
        "channel":"18.04-lts",
        "etcd_state":"kube_etcd",
        "root_disk_size":50,
        "gpu_instance_size":"",
        "gpu_core_count":null,
        "location":"us-west-2:us-west-2a",
        "zone":"us-west-2a",
        "provider_subnet_id":"subnet-065b60e89fb5acf33",
        "provider_subnet_cidr":"172.23.4.0/24",
        "size":"m5.large",
        "state":"running",
        "is_failed":false,
        "created":"2020-01-08T11:39:53.873670Z",
        "updated":"2020-01-08T11:39:53.873680Z"
    },
    {
        "pk":1,
        "cluster":30667,
        "pool":27288,
        "pool_name":"Default Worker Pool",
        "instance_id":"netp2h4pg8-worker-2",
        "provider_node_id":"i-09afa4461e8d1f5c4",
        "role":"worker",
        "group_name":"autoscaling-netp2h4pg8-pool-1",
        "private_ip":"172.23.4.220",
        "public_ip":"18.236.132.53",
        "platform":"ubuntu",
        "image":"ami-0a7d051a1c4b54f65",
        "channel":"18.04-lts",
        "etcd_state":"kube_etcd",
        "root_disk_size":50,
        "gpu_instance_size":"",
        "gpu_core_count":null,
        "location":"us-west-2:us-west-2a",
        "zone":"us-west-2a",
        "provider_subnet_id":"subnet-065b60e89fb5acf33",
        "provider_subnet_cidr":"172.23.4.0/24",
        "size":"m5.large",
        "state":"running",
        "is_failed":false,
        "created":"2020-01-08T11:39:53.874550Z",
        "updated":"2020-01-08T11:39:53.874562Z"
    },
    {
        "pk":1,
        "cluster":30667,
        "pool":27288,
        "pool_name":"Default Worker Pool",
        "instance_id":"netp2h4pg8-worker-2",
        "provider_node_id":"i-09afa4461e8d1f5c4",
        "role":"worker",
        "group_name":"autoscaling-netp2h4pg8-pool-1",
        "private_ip":"172.23.4.220",
        "public_ip":"18.236.132.53",
        "platform":"ubuntu",
        "image":"ami-0a7d051a1c4b54f65",
        "channel":"18.04-lts",
        "etcd_state":"kube_etcd",
        "root_disk_size":50,
        "gpu_instance_size":"",
        "gpu_core_count":null,
        "location":"us-west-2:us-west-2a",
        "zone":"us-west-2a",
        "provider_subnet_id":"subnet-065b60e89fb5acf33",
        "provider_subnet_cidr":"172.23.4.0/24",
        "size":"m5.large",
        "state":"running",
        "is_failed":false,
        "created":"2020-01-08T11:39:53.874550Z",
        "updated":"2020-01-08T11:39:53.874562Z"
    }
]`

var mockNode = `{
        "pk":1,
        "cluster":30667,
        "pool":null,
        "pool_name":"",
        "instance_id":"netp2h4pg8-master-1",
        "provider_node_id":"i-038f00f3b164f86bb",
        "role":"worker",
        "group_name":"",
        "private_ip":"172.23.4.102",
        "public_ip":"54.218.62.37",
        "platform":"ubuntu",
        "image":"ami-0a7d051a1c4b54f65",
        "channel":"18.04-lts",
        "etcd_state":"kube_etcd",
        "root_disk_size":50,
        "gpu_instance_size":"",
        "gpu_core_count":null,
        "location":"us-west-2:us-west-2a",
        "zone":"us-west-2a",
        "provider_subnet_id":"subnet-065b60e89fb5acf33",
        "provider_subnet_cidr":"172.23.4.0/24",
        "size":"m5.large",
        "state":"running",
        "is_failed":false,
        "created":"2020-01-08T11:39:53.872530Z",
        "updated":"2020-01-08T11:39:53.872541Z"
    }`

var mockNodePools = `[
    {
        "pk":1,
        "cluster":30667,
        "name":"Default Worker Pool",
        "instance_id":"netp2h4pg8-pool-1",
        "instance_size":"m5.large",
        "platform":"ubuntu",
        "channel":"18.04-lts",
        "root_disk_size":50,
        "zone":"us-west-2a",
        "provider_subnet_id":"subnet-065b60e89fb5acf33",
        "provider_subnet_cidr":"172.23.4.0/24",
        "node_count":2,
        "autoscaled":false,
        "min_count":0,
        "max_count":0,
        "network_components":[

        ],
        "gpu_instance_size":"",
        "gpu_core_count":null,
        "labels":"",
        "role":"worker",
        "state":"active",
        "is_default":true,
        "config":{

        },
        "created":"2020-01-08T11:39:53.871453Z",
        "updated":"2020-01-08T11:39:53.930118Z"
    },
 {
        "pk":1,
        "cluster":30667,
        "name":"Default Worker Pool",
        "instance_id":"netp2h4pg8-pool-1",
        "instance_size":"m5.large",
        "platform":"ubuntu",
        "channel":"18.04-lts",
        "root_disk_size":50,
        "zone":"us-west-2a",
        "provider_subnet_id":"subnet-065b60e89fb5acf33",
        "provider_subnet_cidr":"172.23.4.0/24",
        "node_count":2,
        "autoscaled":false,
        "min_count":0,
        "max_count":0,
        "network_components":[

        ],
        "gpu_instance_size":"",
        "gpu_core_count":null,
        "labels":"",
        "role":"worker",
        "state":"active",
        "is_default":true,
        "config":{

        },
        "created":"2020-01-08T11:39:53.871453Z",
        "updated":"2020-01-08T11:39:53.930118Z"
    }
]`

var mockNodePool = `{
        "pk":1,
        "cluster":30667,
        "name":"Default Worker Pool",
        "instance_id":"netp2h4pg8-pool-1",
        "instance_size":"m5.large",
        "platform":"ubuntu",
        "channel":"18.04-lts",
        "root_disk_size":50,
        "zone":"us-west-2a",
        "provider_subnet_id":"subnet-065b60e89fb5acf33",
        "provider_subnet_cidr":"172.23.4.0/24",
        "node_count":2,
        "autoscaled":false,
        "min_count":0,
        "max_count":0,
        "network_components":[

        ],
        "gpu_instance_size":"",
        "gpu_core_count":null,
        "labels":"",
        "role":"worker",
        "state":"active",
        "is_default":true,
        "config":{

        },
        "created":"2020-01-08T11:39:53.871453Z",
        "updated":"2020-01-08T11:39:53.930118Z"
    }`

var mockTeams = `[
    {
        "pk":1,
        "name":"Test Go SDK1",
        "slug":"Test Go SDK1",
        "org":1,
        "is_org_wide":true,
        "created":"2018-10-09T05:43:48.656349Z",
        "memberships":[
            {
                "pk":1,
                "user":{
                    "pk":1,
                    "username":"test",
                    "email":"test@netapp.com",
                    "first_name":"test",
                    "last_name":"test",
                    "full_name":"test test",
                    "date_joined":"2018-10-26T20:02:22.808014Z"
                },
                "team":1,
                "created":"2019-12-09T22:47:51.580402Z"
            },
            {
                "pk":1,
                "user":{
                    "pk":11700,
                    "username":"test2-runar",
                    "email":"test2@netapp.com",
                    "first_name":"test2",
                    "last_name":"test2",
                    "full_name":"test2",
                    "date_joined":"2018-10-08T14:31:21.131137Z"
                },
                "team":12164,
                "created":"2019-12-09T22:47:51.546908Z"
            }
        ]
    },
    {
        "pk":1,
        "name":"Dev Team",
        "slug":"dev-team",
        "org":1,
        "is_org_wide":false,
        "created":"2019-07-11T23:51:13.452413Z",
        "memberships":[
            {
                "pk":1,
                "user":{
                    "pk":1,
                    "username":"devtest",
                    "email":"devtest@netapp.com",
                    "first_name":"devtest",
                    "last_name":"devtest",
                    "full_name":"devtest",
                    "date_joined":"2018-10-08T14:31:21.131137Z"
                },
                "team":1,
                "created":"2019-12-06T19:42:20.628920Z"
            },
            {
                "pk":1,
                "user":{
                    "pk":1,
                    "username":"devtest2",
                    "email":"devtest2@antcolony.io",
                    "first_name":"devtest2",
                    "last_name":"devtest2",
                    "full_name":"devtest2",
                    "date_joined":"2019-12-05T20:01:06.770149Z"
                },
                "team":20675,
                "created":"2019-12-05T20:23:35.029754Z"
            }
        ]
    }
]`

var mockTeam = `{
        "pk":1,
        "name":"Test Go SDK",
        "slug":"Test Go SDK1",
        "org":1,
        "is_org_wide":true,
        "created":"2018-10-09T05:43:48.656349Z",
        "memberships":[
            {
                "pk":1,
                "user":{
                    "pk":1,
                    "username":"test",
                    "email":"test@netapp.com",
                    "first_name":"test",
                    "last_name":"test",
                    "full_name":"test test",
                    "date_joined":"2018-10-26T20:02:22.808014Z"
                },
                "team":1,
                "created":"2019-12-09T22:47:51.580402Z"
            },
            {
                "pk":1,
                "user":{
                    "pk":11700,
                    "username":"test2-runar",
                    "email":"test2@netapp.com",
                    "first_name":"test2",
                    "last_name":"test2",
                    "full_name":"test2",
                    "date_joined":"2018-10-08T14:31:21.131137Z"
                },
                "team":12164,
                "created":"2019-12-09T22:47:51.546908Z"
            }
        ]
    }`
