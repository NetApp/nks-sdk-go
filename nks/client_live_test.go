package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveBasicClient(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetOrganizations()
	if err != nil {
		t.Error(err)
	}

	assert.NotEqual(t, len(list), 0, "Result can not be empty")
}
