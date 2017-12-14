package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/StackPointCloud/stackpoint-sdk-go/pkg/stackpointio"
	"github.com/golang/glog"
	"github.com/urfave/cli"
)

// StackpointClient is an client for a particular cluster and organization
type StackpointClient struct {
	id     int
	org    int
	client *stackpointio.APIClient
}

// NewClusterClient creates a ClusterClient from environment
// variables CLUSTER_API_TOKEN, SPC_BASE_API_URL, ORGANIZATION_ID, CLUSTER_ID
func NewClusterClient(c *cli.Context) (*StackpointClient, error) {

	token := c.GlobalString("token")
	endpoint := c.GlobalString("url")
	orgPk := c.GlobalInt("org")
	clusterPk := c.GlobalInt("cluster")

	if token == "" {
		return nil, fmt.Errorf("Environment variable CLUSTER_API_TOKEN not defined")
	}
	if endpoint == "" {
		return nil, fmt.Errorf("Environment variable SPC_API_BASE_URL not defined")
	}

	client := stackpointio.NewClient(token, endpoint)
	glog.V(5).Infof("Using stackpoint io api server [%s]", endpoint)
	glog.V(5).Infof("Using stackpoint organization [%s], cluster [%s]", orgPk, clusterPk)

	spcClient := &StackpointClient{
		id:     clusterPk,
		org:    orgPk,
		client: client,
	}
	return spcClient, nil
}

func (cluster *StackpointClient) getOrganizations() ([]stackpointio.Organization, error) {
	return cluster.client.GetOrganizations()
}

func (cluster *StackpointClient) getOrganization() (stackpointio.Organization, error) {
	return cluster.client.GetOrganization(cluster.org)
}

func (cluster *StackpointClient) getNodes() ([]stackpointio.Node, error) {
	return cluster.client.GetNodes(cluster.org, cluster.id)
}

func (cluster *StackpointClient) getNodePools() ([]stackpointio.NodePool, error) {
	return cluster.client.GetNodePools(cluster.org, cluster.id)
}

func (cluster *StackpointClient) getNodePool(nodePoolID string) (stackpointio.NodePool, error) {
	id, err := strconv.Atoi(nodePoolID)
	if err != nil {
		return stackpointio.NodePool{}, fmt.Errorf("Invalid nodepool id \"%s\"", nodePoolID)
	}
	return cluster.client.GetNodePool(cluster.org, cluster.id, id)
}

func (cluster *StackpointClient) getUser() (stackpointio.User, error) {
	return cluster.client.GetUser()
}

func get(c *cli.Context) error {
	cluster, err := NewClusterClient(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	request := c.Args().First()

	var jsonSerialization []byte
	var jsonError error

	var response interface{}
	var apiErr error

	switch request {

	case "orgs":
		response, apiErr = cluster.getOrganizations()

	case "org":
		response, apiErr = cluster.getOrganization()

	case "user":
		response, apiErr = cluster.getUser()

	case "nodes":
		response, apiErr = cluster.getNodes()

	case "nodepools":
		response, apiErr = cluster.getNodePools()

	case "nodepool":
		response, apiErr = cluster.getNodePool(c.Args().Get(1))

	default:
		return cli.NewExitError(fmt.Sprintf("Invalid get request ? : %s", request), 1)
	}

	if apiErr != nil {
		return cli.NewExitError(apiErr, 1)
	}
	jsonSerialization, jsonError = json.Marshal(response)

	if jsonError != nil {
		return jsonError
	}
	fmt.Printf("%v\n", string(jsonSerialization))

	return nil
}

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "url",
			Value:  "http://api-dev.stackpoint.io:30800/",
			Usage:  "url of the stackpoint api server",
			EnvVar: "SPC_API_BASE_URL",
		},
		cli.StringFlag{
			Name:   "token",
			Value:  "",
			Usage:  "stackpoint api access token",
			EnvVar: "CLUSTER_API_TOKEN",
		},
		cli.IntFlag{
			Name:   "org",
			Value:  1,
			Usage:  "integer key of the organization",
			EnvVar: "ORGANIZATION_ID",
		},
		cli.IntFlag{
			Name:   "cluster",
			Value:  1,
			Usage:  "integer key of the cluster",
			EnvVar: "CLUSTER_ID",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get information on an object",
			Action:  get,
		},
	}

	app.Run(os.Args)
}
