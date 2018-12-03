package stackpointio

import (
	"fmt"
	"testing"
)

func TestGetTeams(t *testing.T) {
	fmt.Println("GetInstanceSpecs testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("SPC_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	teams, err := c.GetTeams(orgID)
	if err != nil {
		t.Error(err)
	}
	if len(teams) == 0 {
		fmt.Println("No teams found, but no error")
	}
}

func TestGetTeam(t *testing.T) {
	fmt.Println("GetInstanceSpecs testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("SPC_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	teams, err := c.GetTeams(orgID)
	if err != nil {
		t.Error(err)
	}
	if len(teams) > 0 {
		teamID := teams[0].ID
		team, err := c.GetTeam(orgID, teamID)
		if err != nil {
			t.Error(err)
		}
		if team == nil {
			t.Errorf("Could not fetch key for team: %d", teamID)
		}
	}
}
