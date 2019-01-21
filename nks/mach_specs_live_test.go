package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiveMachSpecBasic(t *testing.T) {
	testProvider(t, "aws")
	testProvider(t, "eks")
	testProvider(t, "gce")
	testProvider(t, "gke")
	testProvider(t, "aks")
	testProvider(t, "azure")
}

func testProvider(t *testing.T, provider string) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	endpoint, err := GetValueFromEnv("NKS_BASE_API_URL")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetInstanceSpecs(provider, endpoint)

	assert.NotEqual(t, len(list), 0, "Provider must have machine specification")
}
