package nks

import (
	"fmt"
	"testing"
)

func TestGetOrganizations(t *testing.T) {
	fmt.Println("GetOrganizations testing")
	c, err := NewClientFromEnv()
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

func TestGetOrganization(t *testing.T) {
	fmt.Println("GetOrganization testing")
	c, err := NewClientFromEnv()
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
