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

func TestAddNode(t *testing.T) {

	organizationKey := 123
	clusterKey := 456
	nodeAdd := NodeAdd{Count: 1, Size: "t2.medium"}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	responseText := "{\"channel\":\"stable\",\"cluster\":504,\"created\":\"2016-09-27T22:09:57.089819Z\",\"image\":\"ami-06af7f66\", \"instance_id\": \"spcvd7ah21-worker-1\", \"location\": \"us-west-2:us-west-2a\", \"pk\": 1031,\"platform\": \"coreos\",\"private_ip\": \"172.23.1.209\", \"public_ip\": \"54.70.151.25\",\"role\": \"worker\",\"group_name\":\"autoscaling\",\"size\":\"t2.medium\",\"state\":\"draft\",\"updated\":\"2016-09-27T22:09:57.089836Z\"}"

	mux.Handle(fmt.Sprintf("/orgs/%d/clusters/%d/add_node", organizationKey, clusterKey),
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

	node, err := client.AddNodes(organizationKey, clusterKey, nodeAdd)
	require.Nil(t, err)
	assert.Equal(t, "draft", node.State, "returned Node in state \"draft\"")
	assert.Equal(t, "autoscaling", node.Group, "returned Node in group \"autoscaling\"")

}

func TestBuildLogSerialization(t *testing.T) {
	serialized := "{\"cluster\":1,\"event_category\":\"provider\",\"event_type\":\"provider_communication\",\"event_state\":\"failure\",\"message\":\"Alert: this is a critical alert\",\"details\":\"It's only a test actually\",\"reference\":\"http://162.243.174.231:30902/dashboard/db/nodes\",\"created\":\"2017-06-29T00:52:34.766665Z\"}"
	var log BuildLogEntry
	json.Unmarshal([]byte(serialized), &log)
	assert.Equal(t, "failure", log.EventState)
	assert.Equal(t, "provider_communication", log.EventType)
}

func TestGetBuildLog(t *testing.T) {

	mux := getSimpleGetMuxDummy("/orgs/1/clusters/1/logs", "[{\"cluster\":1,\"event_category\":\"\",\"event_type\":\"cloud_network\",\"event_state\":\"success\",\"message\":\"Finished tagging resource\",\"details\":\"\",\"reference\":\"TagResource-1-ci0lrkvm6r\",\"created\":\"2017-06-29T00:07:18.718882Z\"},{\"cluster\":1,\"event_category\":\"\",\"event_type\":\"cloud_network\",\"event_state\":\"started\",\"message\":\"Started tagging resource subnet-d19d44b6\",\"details\":\"\",\"reference\":\"TagResource-1-ci0lrkvm6r\",\"created\":\"2017-06-29T00:07:18.090554Z\"}]")

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
