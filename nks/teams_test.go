package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testTeamLiveID int
var testTeam = Team{
	Name:        "Test Go SDK" + GetTicks(),
	Memberships: []Membership{},
}

func TestLiveBasicTeam(t *testing.T) {
	testLiveTeamCreate(t)
	testLiveTeamList(t)
	testLiveTeamGet(t)
	testLiveTeamDelete(t)
}

func testLiveTeamCreate(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	Team, err := client.CreateTeam(orgID, testTeam)
	require.NoError(t, err)

	testTeamLiveID = Team.ID

	assert.Contains(t, testTeam.Name, Team.Name, "Name should be equal")
}

func testLiveTeamList(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	list, err := client.GetTeams(orgID)
	require.NoError(t, err)

	var Team Team
	for _, item := range list {
		if item.ID == testTeamLiveID {
			Team = item
		}
	}

	assert.NotNil(t, Team)
}

func testLiveTeamGet(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	Team, err := client.GetTeam(orgID, testTeamLiveID)
	require.NoError(t, err)

	assert.Contains(t, testTeam.Name, Team.Name, "Name should be equal")
}

func testLiveTeamDelete(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	require.NoError(t, err)

	err = client.DeleteTeam(orgID, testTeamLiveID)
	require.NoError(t, err)

}
