package nks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLiveOrganization(t *testing.T) {
	t.Run("get clusters", func(t *testing.T) {
		t.Run("list", testGetOrganizations)
		t.Run("get", testGetOrganization)
	})

}

func testGetOrganizations(t *testing.T) {
	fmt.Println("GetOrganizations testing")
	orgs, err := client.GetOrganizations()
	require.NoError(t, err)
	if len(orgs) == 0 {
		fmt.Println("No orgs found, but no error")
	}
}

func testGetOrganization(t *testing.T) {
	fmt.Println("GetOrganization testing")
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)
	org, err := client.GetOrganization(orgID)
	require.NoError(t, err)
	if org == nil {
		fmt.Println("No org found, but no error")
	}
}
