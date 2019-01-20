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

func TestLiveTeamBasic(t *testing.T) {
	testLiveTeamCreate(t)
	testLiveTeamList(t)
	testLiveTeamGet(t)
	testLiveTeamDelete(t)
}

func testLiveTeamCreate(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	Team, err := c.CreateTeam(orgID, testTeam)
	if err != nil {
		t.Error(err)
	}

	testTeamLiveID = Team.ID

	assert.Equal(t, testTeam.Name, Team.Name, "Name should be equal")
}

func testLiveTeamList(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetTeams(orgID)
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
	assert.Equal(t, testTeam.Name, Team.Name, "Name should be equal")
}

func testLiveTeamGet(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	Team, err := c.GetTeam(orgID, testTeamLiveID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, testTeam.Name, Team.Name, "Name should be equal")
}

func testLiveTeamDelete(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteTeam(orgID, testTeamLiveID)
	if err != nil {
		t.Error(err)
	}
}
