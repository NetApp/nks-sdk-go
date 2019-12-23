package mock

import (
	"fmt"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestGetOrganizations(t *testing.T) {
	defer gock.Off()

	orgs, err := MockClient.GetOrganizations()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(orgs) == 0 {
		fmt.Println("No orgs found, but no error")
		t.Fail()
	}
}
