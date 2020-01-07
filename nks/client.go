package nks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

var Debug = "false"

const ClientUserAgentString = "NetApp Kubernetes Service Go SDK v2.0.10"
const defaultNKSApiURL = "https://api.nks.netapp.io"

// APIClient references an api token and an http endpoint
type APIClient struct {
	Token      string
	Endpoint   string
	HttpClient *http.Client
}

// APIReq struct holds data for runRequest method to operate http request on
type APIReq struct {
	Method         string
	Path           string
	PostObj        interface{}
	Payload        io.Reader
	ResponseObj    interface{}
	WantedStatus   int
	ResponseString string
	DontUnmarsahal bool
}

// NewClient returns a new api client
func NewClient(token, endpoint string) *APIClient {
	Debug = os.Getenv("DEBUG")
	c := &APIClient{
		Token:      token,
		Endpoint:   strings.TrimRight(endpoint, "/"),
		HttpClient: http.DefaultClient,
	}
	return c
}

// NewClientFromEnv creates a new client from environment variables
func NewClientFromEnv() (*APIClient, error) {
	token := os.Getenv("NKS_API_TOKEN")
	if token == "" {
		return nil, errors.New("Missing token env in NKS_API_TOKEN")
	}
	endpoint := os.Getenv("NKS_API_URL")
	if endpoint == "" {
		endpoint = defaultNKSApiURL
	}
	return NewClient(token, endpoint), nil
}

// runRequest performs HTTP request, takes APIReq object
func (c *APIClient) runRequest(req *APIReq) error {
	// If method is POST and postObjNeedsEncoding, encode data object and set up payload
	if req.Method == "POST" && req.Payload == nil {
		data, err := json.Marshal(req.PostObj)
		if err != nil {
			return err
		}
		req.Payload = bytes.NewBuffer(data)
	}

	// If path is not fully qualified URL, then prepend with endpoint URL
	if req.Path[0:4] != "http" {
		req.Path = c.Endpoint + req.Path
	}

	// Set up new HTTP request
	httpReq, err := http.NewRequest(req.Method, req.Path, req.Payload)
	debug(httputil.DumpRequestOut(httpReq, true))
	if err != nil {
		return err
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.Token)
	httpReq.Header.Set("User-Agent", ClientUserAgentString)
	httpReq.Header.Set("Content-Type", "application/json")

	// Run HTTP request, catching response
	resp, err := c.HttpClient.Do(httpReq)
	debug(httputil.DumpResponse(resp, true))
	if err != nil {
		return err
	}

	// Check Status Code versus what the caller wanted, error if not correct
	if req.WantedStatus != resp.StatusCode {
		body, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("Incorrect status code returned: %d, Status: %s\n%s", resp.StatusCode, resp.Status, string(body))
		return err
	}

	// If DELETE operation, return
	if req.Method == "DELETE" || req.ResponseObj == nil {
		return nil
	}

	// Store response from remote server, if not a delete operation
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	req.ResponseString = string(body)

	if req.DontUnmarsahal {
		return err
	}

	// Unmarshal response into ResponseObj struct, return ResponseObj and error, if there is one
	return json.Unmarshal(body, req.ResponseObj)
}

func debug(data []byte, err error) {
	if Debug == "true" || Debug == "1" {
		if err == nil {
			fmt.Printf("%s\n\n", data)
		} else {
			log.Fatalf("%s\n\n", err)
		}
	}
}
