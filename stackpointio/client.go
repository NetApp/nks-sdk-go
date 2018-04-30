package stackpointio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const clientUserAgentString = "Stackpoint Go SDK v0.9"

// APIClient references an api token and an http endpoint
type APIClient struct {
	token      string
	endpoint   string
	httpClient *http.Client
}

// NewClient returns a new api client
func NewClient(token, endpoint string) *APIClient {
	c := &APIClient{
		token:      token,
		endpoint:   strings.TrimRight(endpoint, "/"),
		httpClient: http.DefaultClient,
	}
	return c
}

// NewClientFromEnv creates a new client from environment variables
func NewClientFromEnv() (*APIClient, error) {
	token := os.Getenv("SPC_API_TOKEN")
	if token == "" {
		return nil, errors.New("Missing token env in SPC_API_TOKEN")
	}
	endpoint := os.Getenv("SPC_BASE_API_URL")
	if endpoint == "" {
		return nil, errors.New("Missing endpoint env in SPC_BASE_API_URL")
	}
	return NewClient(token, endpoint), nil
}

// runRequest performs HTTP request
func (c *APIClient) runRequest(method, path string, postObj interface{}, response interface{}, wantedStatus int) error {
	var payload io.Reader

	// If method is POST, encode data object and set up payload
	if method == "POST" {
		data, err := json.Marshal(postObj)
		if err != nil {
			return err
		}
		payload = bytes.NewBuffer(data)
	}

	// If path is not fully qualified URL, then prepend with endpoint URL
	if path[0:4] != "http" {
		path = c.endpoint + path
	}

	// Set up new HTTP request
	req, err := http.NewRequest(method, path, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", clientUserAgentString)
	req.Header.Set("Content-Type", "application/json")

	// Run HTTP request, catching response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	// Check Status Code versus what the caller wanted, error if not correct
	if wantedStatus != resp.StatusCode {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("Incorrect status code returned: %d, Status: %s\n%s", 
				resp.StatusCode, resp.Status, string(body)))
	}

	// If method is DELETE, don't care about response
	if method == "DELETE" {
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, response)
}
