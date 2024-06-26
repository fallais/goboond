package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fallais/goboond/responses/resources"
)

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListResources returns the list of the resources.
func (endpoint *Endpoint) ListResources(ctx context.Context, resourceStates, resourceTypes string) (*resources.ListResourcesResponse, error) {
	// Options
	options := []Option{}
	if resourceStates != "" {
		options = append(options, WithParam("resourceStates", resourceStates))
	}
	if resourceTypes != "" {
		options = append(options, WithParam("resourceTypes", resourceTypes))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/resources", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *resources.ListResourcesResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListResources returns the list of the resources.
func (endpoint *Endpoint) GetResource(ctx context.Context, id string) (*resources.GetResourceResponse, error) {
	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/resources/"+id)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *resources.GetResourceResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
