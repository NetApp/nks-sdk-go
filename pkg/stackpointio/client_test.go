package stackpointio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getSimpleGetMuxDummy(path, responseBody string) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, r.URL.Path)
	}))
	mux.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if http.MethodGet != r.Method {
			w.WriteHeader(http.StatusNotImplemented)
		}
		fmt.Fprint(w, responseBody)
	}))

	return mux
}

func TestGetOrganizations(t *testing.T) {

	mux := getSimpleGetMuxDummy("/orgs", "[{\"name\":\"Misty Fire\",\"pk\":1}]")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	orgs, err := client.GetOrganizations()
	require.Nil(t, err)

	assert.Equal(t, 1, len(orgs))
	assert.Equal(t, 1, orgs[0].PrimaryKey)
	assert.Equal(t, "Misty Fire", orgs[0].Name)
}

func TestGetOrganization(t *testing.T) {

	mux := getSimpleGetMuxDummy("/orgs/1", "{\"name\":\"Misty Fire\",\"pk\":1}")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	org, err := client.GetOrganization(1)
	require.Nil(t, err)

	assert.Equal(t, 1, org.PrimaryKey)
	assert.Equal(t, "Misty Fire", org.Name)
}

func TestGetNodes(t *testing.T) {

	orgID := 1
	clusterID := 108
	nodeSerializationString := "[{\"pk\":330,\"cluster\":108,\"pool\":null,\"pool_name\":\"\",\"instance_id\":\"spcp19ss7c-master-1\",\"role\":\"master\",\"group_name\":\"\",\"private_ip\":\"10.138.0.2\",\"public_ip\":\"35.203.163.5\",\"platform\":\"coreos\",\"image\":\"coreos-stable\",\"channel\":\"stable\",\"location\":\"us-west1-a\",\"provider_subnet_id\":\"\",\"provider_subnet_cidr\":\"\",\"size\":\"n1-standard-1\",\"state\":\"provisioned\",\"created\":\"2017-10-19T17:07:38.272197Z\",\"updated\":\"2017-10-19T17:07:38.272216Z\"},{\"pk\":331,\"cluster\":108,\"pool\":108,\"pool_name\":\"Default Worker Pool\",\"instance_id\":\"spcp19ss7c-worker-1\",\"role\":\"worker\",\"group_name\":\"autoscaling-spcp19ss7c-pool-1\",\"private_ip\":\"10.138.0.3\",\"public_ip\":\"35.203.171.134\",\"platform\":\"coreos\",\"image\":\"coreos-stable\",\"channel\":\"stable\",\"location\":\"us-west1-a\",\"provider_subnet_id\":\"\",\"provider_subnet_cidr\":\"\",\"size\":\"n1-standard-1\",\"state\":\"provisioned\",\"created\":\"2017-10-19T17:07:38.276642Z\",\"updated\":\"2017-10-19T17:07:38.276663Z\"},{\"pk\":332,\"cluster\":108,\"pool\":108,\"pool_name\":\"Default Worker Pool\",\"instance_id\":\"spcp19ss7c-worker-2\",\"role\":\"worker\",\"group_name\":\"autoscaling-spcp19ss7c-pool-1\",\"private_ip\":\"10.138.0.4\",\"public_ip\":\"35.203.128.110\",\"platform\":\"coreos\",\"image\":\"coreos-stable\",\"channel\":\"stable\",\"location\":\"us-west1-a\",\"provider_subnet_id\":\"\",\"provider_subnet_cidr\":\"\",\"size\":\"n1-standard-1\",\"state\":\"provisioned\",\"created\":\"2017-10-19T17:07:38.277660Z\",\"updated\":\"2017-10-19T17:07:38.277677Z\"}]"

	mux := getSimpleGetMuxDummy("/orgs/1/clusters/108/nodes", nodeSerializationString)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	nodes, err := client.GetNodes(orgID, clusterID)
	require.Nil(t, err)

	assert.Equal(t, 3, len(nodes))
	assert.Equal(t, "spcp19ss7c-master-1", nodes[0].InstanceID)
}

const addNodeResponse = `[
	{
	  "pk": 9226,
	  "cluster": 2851,
	  "pool": null,
	  "pool_name": "",
	  "instance_id": "spcaufj5gy-master-1",
	  "role": "master",
	  "group_name": "",
	  "private_ip": "10.136.66.195",
	  "public_ip": "192.241.159.155",
	  "platform": "coreos",
	  "image": "coreos-stable",
	  "channel": "stable",
	  "location": "nyc1",
	  "provider_subnet_id": "",
	  "provider_subnet_cidr": "",
	  "size": "2gb",
	  "state": "running",
	  "created": "2017-10-30T20:19:44.700845Z",
	  "updated": "2017-10-30T20:19:44.700862Z"
	},
	{
	  "pk": 9227,
	  "cluster": 2851,
	  "pool": 1224,
	  "pool_name": "Default Worker Pool",
	  "instance_id": "spcaufj5gy-worker-1",
	  "role": "worker",
	  "group_name": "autoscaling-spcaufj5gy-pool-1",
	  "private_ip": "10.136.67.56",
	  "public_ip": "174.138.39.50",
	  "platform": "coreos",
	  "image": "coreos-stable",
	  "channel": "stable",
	  "location": "nyc1",
	  "provider_subnet_id": "",
	  "provider_subnet_cidr": "",
	  "size": "2gb",
	  "state": "running",
	  "created": "2017-10-30T20:19:44.702390Z",
	  "updated": "2017-10-30T20:19:44.702406Z"
	},
	{
	  "pk": 9228,
	  "cluster": 2851,
	  "pool": 1224,
	  "pool_name": "Default Worker Pool",
	  "instance_id": "spcaufj5gy-worker-2",
	  "role": "worker",
	  "group_name": "autoscaling-spcaufj5gy-pool-1",
	  "private_ip": "10.136.67.92",
	  "public_ip": "165.227.94.149",
	  "platform": "coreos",
	  "image": "coreos-stable",
	  "channel": "stable",
	  "location": "nyc1",
	  "provider_subnet_id": "",
	  "provider_subnet_cidr": "",
	  "size": "2gb",
	  "state": "running",
	  "created": "2017-10-30T20:19:44.703842Z",
	  "updated": "2017-10-30T20:19:44.703858Z"
	},
	{
	  "pk": 9229,
	  "cluster": 2851,
	  "pool": 1224,
	  "pool_name": "Default Worker Pool",
	  "instance_id": "spcaufj5gy-worker-3",
	  "role": "worker",
	  "group_name": "autoscaling-spcaufj5gy-pool-1",
	  "private_ip": "",
	  "public_ip": "",
	  "platform": "coreos",
	  "image": "coreos-stable",
	  "channel": "stable",
	  "location": "nyc1:nyc1",
	  "provider_subnet_id": "",
	  "provider_subnet_cidr": "",
	  "size": "2gb",
	  "state": "draft",
	  "created": "2017-10-30T20:32:11.102766Z",
	  "updated": "2017-10-30T20:32:11.102791Z"
	},
	{
	  "pk": 9230,
	  "cluster": 2851,
	  "pool": 1224,
	  "pool_name": "Default Worker Pool",
	  "instance_id": "spcaufj5gy-worker-4",
	  "role": "worker",
	  "group_name": "autoscaling-spcaufj5gy-pool-1",
	  "private_ip": "",
	  "public_ip": "",
	  "platform": "coreos",
	  "image": "coreos-stable",
	  "channel": "stable",
	  "location": "nyc1:nyc1",
	  "provider_subnet_id": "",
	  "provider_subnet_cidr": "",
	  "size": "2gb",
	  "state": "draft",
	  "created": "2017-10-30T20:32:11.107449Z",
	  "updated": "2017-10-30T20:32:11.107466Z"
	}
  ]`

func TestAddNode(t *testing.T) {

	organizationKey := 123
	clusterKey := 2851
	nodeAdd := NodeAdd{
		Count:      2,
		Size:       "2gb",
		NodePoolID: 200,
		Role:       "worker",
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	mux.Handle(fmt.Sprintf("/orgs/%d/clusters/%d/add_node", organizationKey, clusterKey),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			postedData, err := ioutil.ReadAll(r.Body)
			require.Nil(t, err)
			assert.True(t, 0 < len(postedData), "postedData non-zero length")
			fmt.Fprint(w, addNodeResponse)
		}))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	nodes, err := client.AddNodes(organizationKey, clusterKey, nodeAdd)
	require.Nil(t, err)
	assert.True(t, 2 < len(nodes), "at least 2 nodes added")

	filtered := nodes[:0]
	for _, n := range nodes {
		if n.State == "draft" {
			filtered = append(filtered, n)
		}
	}
	assert.Equal(t, 2, len(filtered), "2 nodes added in draft state")

	assert.Equal(t, "draft", filtered[0].State, "returned Node in state \"draft\"")
	assert.Equal(t, "autoscaling-spcaufj5gy-pool-1", filtered[0].Group, "returned Node in group \"autoscaling\"")
	assert.Equal(t, nodeAdd.Role, filtered[0].Role, "returned Node in group \"autoscaling\"")

}

func TestAddNodeValidation(t *testing.T) {

	client := NewClient("", "")
	_, err := client.AddNodes(0, 0, NodeAdd{
		Count: 2,
		Size:  "2gb",
	})
	assert.IsType(t, &ValidationError{}, err)
	_, err = client.AddNodes(0, 0, NodeAdd{
		Size:       "2gb",
		NodePoolID: 100,
	})
	assert.IsType(t, &ValidationError{}, err)
	_, err = client.AddNodes(0, 0, NodeAdd{
		Count:      2,
		NodePoolID: 100,
	})
	assert.IsType(t, &ValidationError{}, err)
}

const nodePoolResponse = `[
	{
	  "pk": 108,
	  "cluster": 108,
	  "name": "Default Worker Pool",
	  "instance_id": "spcp19ss7c-pool-1",
	  "instance_size": "n1-standard-1",
	  "platform": "coreos",
	  "channel": "stable",
	  "zone": "us-west1-a",
	  "provider_subnet_id": "",
	  "provider_subnet_cidr": "",
	  "node_count": 2,
	  "role": "worker",
	  "state": "active",
	  "is_default": true,
	  "created": "2017-10-19T17:07:38.268902Z",
	  "updated": "2017-10-19T17:07:38.337931Z"
	}
  ]`

func TestGetNodePools(t *testing.T) {

	orgID := 1
	clusterID := 108

	mux := getSimpleGetMuxDummy("/orgs/1/clusters/108/nodepools", nodePoolResponse)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	pools, err := client.GetNodePools(orgID, clusterID)
	require.Nil(t, err)

	assert.Equal(t, 1, len(pools))
	assert.Equal(t, "Default Worker Pool", pools[0].Name)
	assert.True(t, pools[0].Default)
	assert.False(t, pools[0].Autoscaled)
	assert.Equal(t, "n1-standard-1", pools[0].Size)
	assert.Equal(t, "1", pools[0].CPU)
	assert.Equal(t, "3750", pools[0].Memory)
}

func TestBuildLogSerialization(t *testing.T) {
	serialized := "{\"cluster\":1,\"event_category\":\"provider\",\"event_type\":\"provider_communication\",\"event_state\":\"failure\",\"message\":\"Alert: this is a critical alert\",\"details\":\"It's only a test actually\",\"reference\":\"http://162.243.174.231:30902/dashboard/db/nodes\",\"created\":\"2017-06-29T00:52:34.766665Z\"}"
	var log BuildLogEntry
	json.Unmarshal([]byte(serialized), &log)
	assert.Equal(t, "failure", log.EventState)
	assert.Equal(t, "provider_communication", log.EventType)
}

const buildLogResponse = `[
	{
	  "cluster": 1,
	  "event_category": "",
	  "event_type": "cloud_network",
	  "event_state": "success",
	  "message": "Finished tagging resource",
	  "details": "",
	  "reference": "TagResource-1-ci0lrkvm6r",
	  "created": "2017-06-29T00:07:18.718882Z"
	},
	{
	  "cluster": 1,
	  "event_category": "",
	  "event_type": "cloud_network",
	  "event_state": "started",
	  "message": "Started tagging resource subnet-d19d44b6",
	  "details": "",
	  "reference": "TagResource-1-ci0lrkvm6r",
	  "created": "2017-06-29T00:07:18.090554Z"
	}
  ]`

func TestGetBuildLog(t *testing.T) {

	mux := getSimpleGetMuxDummy("/orgs/1/clusters/1/logs", buildLogResponse)

	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	orgID := 1
	clusterID := 1
	logs, err := client.GetLogs(orgID, clusterID)
	require.Nil(t, err)

	assert.Equal(t, 2, len(logs))
	assert.Equal(t, "success", logs[0].EventState)
	assert.Equal(t, "started", logs[1].EventState)

	assert.Equal(t, "cloud_network", logs[1].EventType)

	assert.Equal(t, "TagResource-1-ci0lrkvm6r", logs[0].Reference)
}

func TestPostLog(t *testing.T) {

	organizationKey := 1
	clusterKey := 1
	buildLog := BuildLogEntry{
		ClusterID:     500,
		EventCategory: "solution",
		EventType:     "launch_operation",
		EventState:    "started",
		Message:       "TestPostLog",
		Details:       "stackpoint-sdk-go",
		Reference:     "github.com/Stackpoint/stackpoint-sdk-go",
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	responseText := "{\"cluster\":1,\"event_category\":\"solution\",\"event_type\":\"launch_operation\",\"event_state\":\"started\",\"message\":\"TestPostLog\",\"details\":\"stackpoint-sdk-go\",\"reference\":\"github.com/Stackpoint/stackpoint-sdk-go\",\"created\":\"0001-01-01T00:00:00Z\"}"

	mux.Handle(fmt.Sprintf("/orgs/%d/clusters/%d/buildlogs", organizationKey, clusterKey),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			postedData, err := ioutil.ReadAll(r.Body)
			require.Nil(t, err)
			assert.True(t, 0 < len(postedData), "postedData non-zero length")
			fmt.Fprint(w, responseText)
		}))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	responseLog, err := client.PostBuildLog(organizationKey, clusterKey, buildLog)
	require.Nil(t, err)
	assert.Equal(t, "launch_operation", responseLog.EventType)
	assert.Equal(t, "started", responseLog.EventState)

}

func TestPostAlert(t *testing.T) {

	organizationKey := 1
	clusterKey := 1

	message := "OutOfDisk"
	details := "Disk usage on node foobar exceeds 95%"
	reference := "https://lookit.alerting.com/mycluster/789235213"

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	responseText := "{\"cluster\":0,\"event_category\":\"kubernetes\",\"event_type\":\"provider_communication\",\"event_state\":\"failure\",\"message\":\"OutOfDisk\",\"details\":\"Disk usage on node foobar exceeds 95%\",\"reference\":\"https://lookit.alerting.com/mycluster/789235213\",\"created\":\"0001-01-01T00:00:00Z\"}"
	mux.Handle(fmt.Sprintf("/orgs/%d/clusters/%d/buildlogs", organizationKey, clusterKey),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			postedData, err := ioutil.ReadAll(r.Body)
			require.Nil(t, err)
			assert.True(t, 0 < len(postedData), "postedData non-zero length")
			fmt.Fprint(w, responseText)
		}))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	responseLog, err := client.PostAlert(organizationKey, clusterKey, message, details, reference)
	require.Nil(t, err)
	assert.Equal(t, "kubernetes", responseLog.EventCategory)
	assert.Equal(t, "failure", responseLog.EventState)
	assert.Equal(t, message, responseLog.Message)
	assert.Equal(t, details, responseLog.Details)
	assert.Equal(t, reference, responseLog.Reference)

}
