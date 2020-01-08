package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	if err != nil {
		t.Error(err)
	}

	Team, err := client.CreateTeam(orgID, testTeam)
	if err != nil {
		t.Error(err)
	}

	testTeamLiveID = Team.ID

	assert.Contains(t, testTeam.Name, Team.Name, "Name should be equal")
}

func testLiveTeamList(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := client.GetTeams(orgID)
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}

	Team, err := client.GetTeam(orgID, testTeamLiveID)
	if err != nil {
		t.Error(err)
	}

	assert.Contains(t, testTeam.Name, Team.Name, "Name should be equal")
}

func testLiveTeamDelete(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = client.DeleteTeam(orgID, testTeamLiveID)
	if err != nil {
		t.Error(err)
	}
}
