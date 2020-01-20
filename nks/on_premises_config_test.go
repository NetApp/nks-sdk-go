package nks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const onPremisesConfigVersion = "v1.2"

func TestLiveOnPremisesConfig(t *testing.T) {
	testGetOnPremisesConfigs(t)
	testGetOnPremisesConfig(t)
}

func testGetOnPremisesConfigs(t *testing.T) {
	fmt.Println("GetOnPremisesConfigs testing")
	configs, err := client.GetOnPremisesConfigs()
	require.NoError(t, err)
	if len(configs) == 0 {
		fmt.Println("No configs found, but no error")
	}
}

func testGetOnPremisesConfig(t *testing.T) {
	fmt.Println("GetOnPremisesConfig testing")
	config, _, err := client.GetOnPremisesConfig(onPremisesConfigVersion)
	require.NoError(t, err)
	if config == nil {
		fmt.Println("No config found, but no error")
	}
}

