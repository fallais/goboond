package goboond

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const BaseURL = "https://ui.boondmanager.com/api"
const JWTAppHeader = "X-Jwt-App-Boondmanager"
const JWTClientHeader = "X-Jwt-Client-Boondmanager"
const ModeNormal = "normal"
const ModeGod = "god"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Client is a client for QRadar REST API.
type Client struct {
	client *http.Client

	Version string

	// User token
	UserToken string

	// Client authentication
	ClientToken string
	ClientKey   string

	// App authentication
	AppToken string
	AppKey   string

	// Basic authentication
	Username string
	Password string

	// Authentication method
	AuthenticationMethod AuthenticationMethod

	// Endpoints
	Projects           Projects
	Resources          Resources
	Candidates         Candidates
	Actions            Actions
	ReportingSynthesis ReportingSynthesis
	Positionings       Positionings
	Contacts           Contacts
	Companies          Companies

	debug bool
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewClientWithClientToken returns a new Boondmanager client with client token authentication.
func NewClientWithClientToken(httpClient *http.Client, debug bool, userToken, clientToken, clientKey string) *Client {
	return newClient(httpClient, debug, AuthenticationMethodClient, userToken, clientToken, clientKey, "", "", "", "")
}

// NewClientWithAppToken returns a new Boondmanager client with app token authentication.
func NewClientWithAppToken(httpClient *http.Client, debug bool, userToken, appToken, appKey string) *Client {
	return newClient(httpClient, debug, AuthenticationMethodApp, userToken, "", "", appToken, appKey, "", "")
}

// NewClientWithBasicAuth returns a new Boondmanager client with basic authentication.
func NewClientWithBasicAuth(httpClient *http.Client, debug bool, userToken, username, password string) *Client {
	return newClient(httpClient, debug, AuthenticationMethodBasic, "", "", "", "", "", username, password)
}

func newClient(httpClient *http.Client, debug bool, authMethod AuthenticationMethod, userToken, clientToken, clientKey, appToken, appKey, username, password string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Create the client
	c := &Client{
		client:               httpClient,
		AuthenticationMethod: authMethod,
		UserToken:            userToken,
		ClientToken:          clientToken,
		ClientKey:            clientKey,
		AppToken:             appToken,
		AppKey:               appKey,
		Username:             username,
		Password:             password,
		debug:                debug,
	}

	// Add the endpoints
	c.Actions = &Endpoint{client: c}
	c.Candidates = &Endpoint{client: c}
	c.Projects = &Endpoint{client: c}
	c.Resources = &Endpoint{client: c}
	c.ReportingSynthesis = &Endpoint{client: c}
	c.Positionings = &Endpoint{client: c}
	c.Contacts = &Endpoint{client: c}
	c.Companies = &Endpoint{client: c}

	return c
}

func (c *Client) do(ctx context.Context, method, endpoint string, body interface{}, opts ...Option) (*http.Response, error) {
	// Options
	var apiOptions options

	// Add options
	for _, op := range opts {
		err := op(&apiOptions)
		if err != nil {
			return nil, err
		}
	}

	// Raw URL
	rawURL := fmt.Sprintf("%s%s", BaseURL, endpoint)

	// Build query
	queryURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	// Assign query parameters
	if apiOptions.Params != nil {
		queryURL.RawQuery = apiOptions.Params.Encode()
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Initialize request
	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), buf)
	if err != nil {
		return nil, err
	}

	// Generate JWT
	jwt, err := c.generateJWT()
	if err != nil {
		return nil, err
	}

	// Default headers
	headers := http.Header{}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Add debug header if needed
	if c.debug {
		headers.Add("x-Debug-Boondmanager", "true")
	}

	// Add JWT header
	switch c.AuthenticationMethod {
	case AuthenticationMethodClient:
		headers.Add(JWTClientHeader, jwt)
	case AuthenticationMethodApp:
		headers.Add(JWTAppHeader, jwt)
	case AuthenticationMethodBasic:
		req.SetBasicAuth(c.Username, c.Password)
	}

	// Assign new headers
	req.Header = headers

	// Dump request if debug
	if c.debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		fmt.Println(string(dump))
	}

	// Do the query
	resp, err := c.client.Do(req)
	if err != nil {
		// Dump response if debug
		if c.debug {
			dumpResponse(resp)
		}

		return nil, fmt.Errorf("error while doing the request: %s", err)
	}

	return resp, err
}

func (c *Client) generateJWT() (string, error) {
	switch c.AuthenticationMethod {
	case AuthenticationMethodClient:
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userToken":   c.UserToken,
			"clientToken": c.ClientToken,
			"time":        time.Now().Unix(),
			"mode":        ModeGod,
		})
		return token.SignedString([]byte(c.ClientKey))
	case AuthenticationMethodApp:
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userToken": c.UserToken,
			"appToken":  c.AppToken,
			"time":      time.Now().Unix(),
			"mode":      ModeGod,
		})
		return token.SignedString([]byte(c.AppKey))
	}

	return "", nil
}

func dumpResponse(resp *http.Response) {
	httpDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("error while dumping the response", err)
		return
	}

	fmt.Println(string(httpDump))
}
