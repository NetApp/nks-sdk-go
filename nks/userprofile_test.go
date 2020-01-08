package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveBasicUserprofile(t *testing.T) {
	up, err := client.GetUserProfile()
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, up, "User profile must exist")
}
