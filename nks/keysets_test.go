package nks

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testKeysetLiveID int
var testKeyset = Keyset{
	Name:       "Test Go SDK " + GetTicks(),
	Category:   "user_ssh",
	Workspaces: []int{},
	IsDefault:  false,
	Keys:       []Key{},
}

func TestLiveBasicKeyset(t *testing.T) {
	testLiveKeysetCreate(t)
	testLiveKeysetList(t)
	testLiveKeysetGet(t)
	testLiveKeysetDelete(t)
}

func testLiveKeysetCreate(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Fatal(err)
	}

	idRsaPubPath, err := GetValueFromEnv("NKS_ID_RSA_PUB_PATH")
	if err != nil {
		t.Fatal(err)
	}

	idRsaPubPath, err = GetAbsPath(idRsaPubPath)
	if err != nil {
		t.Fatal(err)
	}
	content := []byte{}

	if testEnv != "mock" {
		content, err = ioutil.ReadFile(idRsaPubPath)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		content = []byte{}
	}

	testKeyset.Keys = append(testKeyset.Keys, Key{
		Type:  "pub",
		Value: string(content),
	})

	Keyset, err := client.CreateKeyset(orgID, testKeyset)
	if err != nil {
		t.Error(err)
	}

	testKeysetLiveID = Keyset.ID

	assert.Contains(t, testKeyset.Name, Keyset.Name, "Name should be equal")
	assert.NotNil(t, len(testKeyset.Keys), 1, "One key should be present")
	assert.Equal(t, testKeyset.Keys[0].Type, "pub", "A key should be pub")
}

func testLiveKeysetList(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := client.GetKeysets(orgID)
	if err != nil {
		t.Error(err)
	}

	var Keyset Keyset
	for _, item := range list {
		if item.ID == testKeysetLiveID {
			Keyset = item
		}
	}

	assert.NotNil(t, Keyset)
	assert.Contains(t, testKeyset.Name, Keyset.Name, "Name should be equal")
}

func testLiveKeysetGet(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	Keyset, err := client.GetKeyset(orgID, testKeysetLiveID)
	if err != nil {
		t.Error(err)
	}

	assert.Contains(t, testKeyset.Name, Keyset.Name, "Name should be equal")
}

func testLiveKeysetDelete(t *testing.T) {
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = client.DeleteKeyset(orgID, testKeysetLiveID)
	if err != nil {
		t.Error(err)
	}
}
