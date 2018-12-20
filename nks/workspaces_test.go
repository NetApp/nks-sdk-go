package nks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetWorkspacesMock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/1/workspaces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"pk": 12345, "name": "Default", "slug": "default", "org": 1}]`)
	})

	workspaces, err := client.GetWorkspaces(1)
	if err != nil {
		t.Errorf("GetWorkspaces returned error: %v", err)
	}

	expected := []Workspace{{ID: 12345, Name: "Default", Slug: "default", Org: 1}}
	if !reflect.DeepEqual(workspaces, expected) {
		t.Errorf("GetWorkspaces\n got=%#v\nwant=%#v", workspaces, expected)
	}
}

func TestGetWorkspaceMock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/1/workspaces/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"pk":12345}`)
	})

	workspace, err := client.GetWorkspace(1, 12345)
	if err != nil {
		t.Errorf("CetWorkspace returned error: %v", err)
	}

	expected := &Workspace{ID: 12345}
	if !reflect.DeepEqual(workspace, expected) {
		t.Errorf("GetWorkspace\n got=%#v\nwant=%#v", workspace, expected)
	}
}

func TestCreateWorkspaceMock(t *testing.T) {
	setup()
	defer teardown()

	createRequest := Workspace{
		ID:   12345,
		Org:  1,
		Name: "default",
	}

	mux.HandleFunc("/orgs/1/workspaces", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		expected := map[string]interface{}{
			"name":            "default",
			"created":         "0001-01-01T00:00:00Z",
			"clusters":        nil,
			"user_solutions":  nil,
			"team_workspaces": nil,
			"federations":     nil,
			"pk":              float64(12345),
			"org":             float64(1),
			"is_default":      false,
			"slug":            "",
		}

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Errorf("CreateWorkspace\n got=%#v\nwant=%#v", v, expected)
		}

		fmt.Fprintf(w, `{"name": "default", "team_workspaces": []}`)
	})

	workspace, err := client.CreateWorkspace(1, createRequest)
	if err != nil {
		t.Errorf("CreateWorkspace returned error: %v", err)
	}

	if id := workspace.ID; id != 0 {
		t.Errorf("expected id '%d', received '%d'", 1, id)
	}
}

func TestDeleteWorkspaceMock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/orgs/1/workspaces/12345", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		testMethod(t, r, http.MethodDelete)
	})

	err := client.DeleteWorkspace(1, 12345)
	if err != nil {
		t.Errorf("DeleteWorkspace returned error: %v", err)
	}
}

func TestGetWorkspaces(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	workspaces, err := c.GetWorkspaces(orgID)
	if err != nil {
		t.Error(err)
	}
	if len(workspaces) == 0 {
		fmt.Println("No workspaces found, but no error")
	}
}

func TestGetWorkspace(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	workspaces, err := c.GetWorkspaces(orgID)
	if err != nil {
		t.Error(err)
	}
	if len(workspaces) > 0 {
		workspaceID := workspaces[0].ID
		workspace, err := c.GetWorkspace(orgID, workspaceID)
		if err != nil {
			t.Error(err)
		}
		if workspace == nil {
			t.Errorf("Could not fetch key for workspace: %d", workspaceID)
		}
	}
}
