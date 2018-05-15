package stackpointio

import (
	"fmt"
	"testing"
)

func TestGetKeysets(t *testing.T) {
	fmt.Println("GetInstanceSpecs testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("SPC_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	keysets, err := c.GetKeysets(orgID)
	if err != nil {
		t.Error(err)
	}
	if len(keysets) == 0 {
		fmt.Println("No keysets found, but no error")
	}
}

func TestGetKeyset(t *testing.T) {
	fmt.Println("GetInstanceSpecs testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("SPC_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	keysets, err := c.GetKeysets(orgID)
	if err != nil {
		t.Error(err)
	}
	if len(keysets) > 0 {
		keysetID := keysets[0].ID
		keyset, err := c.GetKeyset(orgID, keysetID)
		if err != nil {
			t.Error(err)
		}
		if keyset == nil {
			t.Errorf("Could not fetch key for keyset: %d", keysetID)
		}
	}
}
