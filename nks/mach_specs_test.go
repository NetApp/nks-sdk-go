package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLiveBasicMachSpec(t *testing.T) {
	testProvider(t, "aws")
	testProvider(t, "gce")
	testProvider(t, "azure")
}

func testProvider(t *testing.T, provider string) {
	endpoint, err := GetValueFromEnv("NKS_BASE_API_URL")
	require.NoError(t, err)

	list, err := client.GetInstanceSpecs(provider, endpoint)

	assert.NotEqual(t, len(list), 0, "Provider must have machine specification")
}
