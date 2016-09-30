package stackpointio

import (
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

	mux.Handle(fmt.Sprintf("/orgs/%d/clusters/%d/add_node", organizationKey, clusterKey),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			postedData, err := ioutil.ReadAll(r.Body)
			require.Nil(t, err)
			assert.True(t, 0 < len(postedData), "postedData non-zero length")
			fmt.Fprint(w, string(postedData))
		}))

	ts := httptest.NewServer(mux)
	defer ts.Close()

	token := "not used"
	client := NewClient(token, ts.URL)

	_, err := client.AddNodes(organizationKey, clusterKey, nodeAdd)
	require.Nil(t, err)

}
