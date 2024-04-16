package goboond

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const DefaultBaseURL = "https://ui.boondmanager.com"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Client is a client for QRadar REST API.
type Client struct {
	client *http.Client

	BaseURL     string
	UserToken   string
	ClientToken string
	ClientKey   string
	Version     string

	// Endpoints
	Projects   Projects
	Resources  Resources
	Candidates Candidates
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewClient returns a new QRadar API client.
func NewClient(httpClient *http.Client, baseURL, userToken, clientToken, clientKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Create the client
	c := &Client{
		client:      httpClient,
		BaseURL:     DefaultBaseURL,
		UserToken:   userToken,
		ClientToken: clientToken,
		ClientKey:   clientKey,
	}

	if baseURL != "" {
		c.BaseURL = baseURL
	}

	// Add the endpoints
	c.Candidates = &Endpoint{client: c}
	c.Projects = &Endpoint{client: c}
	c.Resources = &Endpoint{client: c}

	return c
}

func (c *Client) do(ctx context.Context, method, endpoint string, opts ...Option) (*http.Response, error) {
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
	rawURL := fmt.Sprintf("%s/api%s", c.BaseURL, endpoint)

	// Build query
	queryURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	// Assign query parameters
	if apiOptions.Params != nil {
		queryURL.RawQuery = apiOptions.Params.Encode()
	}

	// Initialize request
	req, err := http.NewRequestWithContext(ctx, method, queryURL.String(), nil)
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
	headers.Add("Accept", "application/json")
	headers.Add("X-Jwt-Client-Boondmanager", jwt)

	// Assign new headers
	req.Header = headers

	// Do the query
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while doing the request: %s", err)
	}

	return resp, err
}

func (c *Client) generateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userToken":   c.UserToken,
		"clientToken": c.ClientToken,
		"time":        time.Now().Unix(),
		"mode":        "god",
	})

	return token.SignedString([]byte(c.ClientKey))
}
