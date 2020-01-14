package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLiveBasicClient(t *testing.T) {
	list, err := client.GetOrganizations()
	require.NoError(t, err)

	assert.NotEqual(t, len(list), 0, "Result can not be empty")
}
