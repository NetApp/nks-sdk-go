package nks

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

var (
	mux          *http.ServeMux
	client       *APIClient
	server       *httptest.Server
	mockServer   = "http://foo.bar"
	mockData     map[string]string
	mockDataPath = "mock_data/"
)

func TestMain(m *testing.M) {
	testEnv := os.Getenv("NKS_TEST_ENV")
	if testEnv == "live" {
		token := os.Getenv("NKS_API_TOKEN")
		endpoint := os.Getenv("NKS_API_URL")
		if endpoint == "" {
			endpoint = defaultNKSApiURL
		}
		client = NewClient(token, endpoint)
	} else {
		fmt.Println("Starting mock client")
		client = NewClient("MOCK_TOKEN", mockServer)
		loadMockData()
		setupMockServer()
	}

	os.Exit(m.Run())

}

func loadMockData() {
	fmt.Println("Loading mock data")
	mockData = make(map[string]string)
	root := "mock_data/"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fName := f.Name()
		dat, _ := ioutil.ReadFile(fmt.Sprintf("%s%s", root, fName))

		mockData[fName[:len(fName)-5]] = string(dat)
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
		JSON(mockData["nodes"])

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/nodes/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["node"])

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
		JSON(mockData["nodes"])

	//nodepools
	gock.New("http://foo.bar").
		Post("/orgs/1/clusters/1/nodepools/1/add").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusCreated).
		JSON(mockData["nodes"])

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/nodepools/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["nodepool"])

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/nodepools").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["nodepools"])

	gock.New("http://foo.bar").
		Post("/orgs/1/clusters/1/nodepools").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusAccepted).
		JSON(mockData["nodepool"])

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
		JSON(mockData["machinespecs"])

	gock.New("http://foo.bar").
		Get("/meta/provider-instances").
		MatchParam("provider", "gce").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["machinespecs"])

	gock.New("http://foo.bar").
		Get("/meta/provider-instances").
		MatchParam("provider", "azure").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["machinespecs"])

	//solutions
	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/solutions/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["solution"])

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
		JSON(mockData["solution"])

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1/solutions").
		MatchType("json").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["solutions"])

	//istio mesh
	gock.New("http://foo.bar").
		Get("/orgs/1/workspaces/1/istio-meshes").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["istiomeshes"])

	gock.New("http://foo.bar").
		Post("/orgs/1/workspaces/1/istio-meshes").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(201).
		JSON(mockData["istiomesh"])

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
		JSON(mockData["buildlogs"])

	//cluster endpoints
	gock.New("http://foo.bar").
		Get("/orgs/1/clusters/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["cluster"])

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
		JSON(mockData["cluster"])

	gock.New("http://foo.bar").
		Get("/orgs/1/clusters").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["clusters"])

	// workspace
	gock.New("http://foo.bar").
		Delete("orgs/1/workspaces/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusNoContent)

	gock.New("http://foo.bar").
		Get("orgs/1/workspaces/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["workspace"])

	gock.New("http://foo.bar").
		Post("orgs/1/workspaces").
		MatchHeader("Authorization", "MOCK_TOKEN").
		MatchHeader("Content-Type", "application/json").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusCreated).
		JSON(mockData["workspace"])

	//keyset endpoints
	gock.New("http://foo.bar").
		Get("/orgs/1/keysets/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["keyset"])
	gock.New("http://foo.bar").
		Delete("/orgs/1/keysets/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusNoContent)

	gock.New("http://foo.bar").
		Post("/orgs/1/keysets").
		MatchHeader("Authorization", "MOCK_TOKEN").
		MatchHeader("Content-Type", "application/json").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusCreated).
		JSON(mockData["keyset"])

	gock.New("http://foo.bar").
		Get("/orgs/1/keysets").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["keysets"])

	gock.New("http://foo.bar").
		Get("orgs/1/workspaces").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["workspaces"])

	//teams
	gock.New("http://foo.bar").
		Get("/orgs/1/teams/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["team"])

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
		JSON(mockData["teams"])

	gock.New("http://foo.bar").
		Post("/orgs/1/teams").
		MatchHeader("Authorization", "MOCK_TOKEN").
		MatchHeader("Content-Type", "application/json").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusCreated).
		JSON(mockData["team"])

	//userprofiles
	gock.New("http://foo.bar").
		Get("/userprofile").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["userprofile"])

	//organization endpoints
	gock.New("http://foo.bar").
		Get("/orgs/1").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["organization"])

	gock.New("http://foo.bar").
		Get("/orgs").
		MatchHeader("Authorization", "MOCK_TOKEN").
		HeaderPresent("User-Agent").
		HeaderPresent("Content-Type").
		Persist().
		Reply(http.StatusOK).
		JSON(mockData["organizations"])
}
