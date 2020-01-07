package nks

import (
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestLiveOrganization(t *testing.T) {
	defer gock.Off()
	t.Run("get clusters", func(t *testing.T) {
		t.Run("list", testGetOrganizations)
		t.Run("get", testGetOrganization)
	})

}

func testGetOrganizations(t *testing.T) {
	fmt.Println("GetOrganizations testing")
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgs, err := c.GetOrganizations()
	if err != nil {
		t.Error(err)
	}
	if len(orgs) == 0 {
		fmt.Println("No orgs found, but no error")
	}
}

func testGetOrganization(t *testing.T) {
	fmt.Println("GetOrganization testing")
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	org, err := c.GetOrganization(orgID)
	if err != nil {
		t.Error(err)
	}
	if org == nil {
		fmt.Println("No org found, but no error")
	}
}
