package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLiveBasicUserprofile(t *testing.T) {
	up, err := client.GetUserProfile()
	require.NoError(t, err)

	assert.NotEmpty(t, up, "User profile must exist")
}
