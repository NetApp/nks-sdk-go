package nks

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

func TestRunRequest(t *testing.T) {
	fmt.Println("runRequest testing")
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	if c == nil {
		t.Error(err)
	}
	orgs := []Organization{}
	req := &APIReq{
		Method:       "GET",
		Path:         "/orgs",
		ResponseObj:  &orgs,
		WantedStatus: 200,
	}
	err = c.runRequest(req)
	if err != nil {
		t.Error(err)
	}

}
