package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveBasicOrganization(t *testing.T) {
	tesOrganizationtList(t)
	testOrganizationGet(t)
}

func tesOrganizationtList(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetOrganizations()
	if err != nil {
		t.Error(err)
	}

	assert.NotEqual(t, len(list), 0, "An organization must exist")
}

func testOrganizationGet(t *testing.T) {
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

	assert.Equal(t, org.ID, orgID, "An organization must exist")
}
