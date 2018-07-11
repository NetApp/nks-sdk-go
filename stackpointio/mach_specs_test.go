package stackpointio

import (
	"fmt"
	"testing"
)

func TestGetInstanceSpecs(t *testing.T) {
	fmt.Println("GetInstanceSpecs testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	insts, err := c.GetInstanceSpecs("do", "")
	if err != nil {
		t.Error(err)
	}
	if len(insts) == 0 {
		t.Errorf("No instances returned")
	}
	fl := GetFormattedInstanceList(insts)
	if fl[0] == "" {
		t.Errorf("No instances to format")
	}
	if !InstanceInList(insts, insts[0].Name) {
		t.Errorf("Invalid data in instance list")
	}
}
