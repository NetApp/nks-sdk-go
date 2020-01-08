package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveBasicMachSpec(t *testing.T) {
	testProvider(t, "aws")
	testProvider(t, "gce")
	testProvider(t, "azure")
}

func testProvider(t *testing.T, provider string) {
	endpoint, err := GetValueFromEnv("NKS_BASE_API_URL")
	if err != nil {
		t.Error(err)
	}

	list, err := client.GetInstanceSpecs(provider, endpoint)

	assert.NotEqual(t, len(list), 0, "Provider must have machine specification")
}
