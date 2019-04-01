package nks

import "fmt"

// RunProviderProxyRequest runs a request on the provider API proxy using a keyset for the provider
func (c *APIClient) RunProviderProxyRequest(provider string, keysetID int, cmd string, obj interface{}) (err error) {
	req := &APIReq{
		Method:       "GET",
		Path:         fmt.Sprintf("/provider/%s/keyset/%d/proxy?%s", provider, keysetID, cmd),
		ResponseObj:  &obj,
		WantedStatus: 200,
	}

	err = c.runRequest(req)
	return
}
