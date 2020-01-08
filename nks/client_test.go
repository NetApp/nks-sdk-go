package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveBasicClient(t *testing.T) {
	list, err := client.GetOrganizations()
	if err != nil {
		t.Error(err)
	}

	assert.NotEqual(t, len(list), 0, "Result can not be empty")
}
