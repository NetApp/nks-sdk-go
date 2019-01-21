package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveUserprofileBasic(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	up, err := c.GetUserProfile()
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, up, "User profile must exist")
}
