package nks

import (
	"fmt"
	"testing"
)

func TestGetUserProfile(t *testing.T) {
	fmt.Println("GetUserProfile testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	up, err := c.GetUserProfile()
	if err != nil {
		t.Error(err)
	}
	if len(up) == 0 {
		fmt.Println("No userprofile found, but no error")
	}
}

func TestGetUserProfileDefaultOrg(t *testing.T) {
	fmt.Println("GetUserProfileDefaultOrg testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	up, err := c.GetUserProfile()
	if err != nil {
		t.Error(err)
	}
	if len(up) == 0 {
		fmt.Println("No userprofile found, but no error")
		return
	}
	orgID, err := c.GetUserProfileDefaultOrg(&up[0])
	if err != nil {
		t.Error(err)
	}
	if orgID == 0 {
		t.Error("No default org found")
	}
}

func TestGetUserProfileKeysetID(t *testing.T) {
	fmt.Println("GetUserProfileKeysetID testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	up, err := c.GetUserProfile()
	if err != nil {
		t.Error(err)
	}
	if len(up) == 0 {
		fmt.Println("No userprofile found, but no error")
		return
	}
	orgID, err := c.GetUserProfileDefaultOrg(&up[0])
	if err != nil {
		t.Error(err)
	}
	if orgID == 0 {
		t.Error("No default org found")
	}
	var ksids []int
	for _, prov := range []string{"aws", "azure", "do", "gce", "gke", "oneandone", "packet", "user_ssh"} {
		ksid, _ := c.GetUserProfileKeysetID(&up[0], prov)
		ksids = append(ksids, ksid)
	}
	if len(ksids) == 0 {
		fmt.Println("No keysets found, but no error")
		return
	}
}
