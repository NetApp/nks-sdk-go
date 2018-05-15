package stackpointio

import (
	"fmt"
	"testing"
)

func TestNewClientFromEnv(t *testing.T) {
	fmt.Println("NewClientFromEnv testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	if c == nil {
		t.Error(err)
	}
}

func TestrunRequest(t *testing.T) {
	fmt.Println("runRequest testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	r := []Organization{}
	err = c.runRequest("GET", "/orgs", nil, &r, 200)
	if err != nil {
		t.Error(err)
	}
	if c == nil {
		t.Error(err)
	}
}
