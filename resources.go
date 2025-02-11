package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ListResourcesRequest struct {
	ResourceStates []int
	ResourceTypes  []int

	SearchOptions
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListResources returns the list of the resources.
func (endpoint *Endpoint) ListResources(ctx context.Context, r ListResourcesRequest) (*ListResourcesResponse, error) {
	// Options
	options := []Option{}
	if r.ResourceStates != nil {
		options = append(options, WithParam("resourceStates", formatIntArray(r.ResourceStates)))
	}
	if r.ResourceTypes != nil {
		options = append(options, WithParam("resourceTypes", formatIntArray(r.ResourceTypes)))
	}
	if r.MaxResults != 0 {
		options = append(options, WithParam("maxResults", fmt.Sprintf("%d", r.MaxResults)))
	}
	if r.Page != 0 {
		options = append(options, WithParam("page", fmt.Sprintf("%d", r.Page)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/resources", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *ListResourcesResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListResources returns the list of the resources.
func (endpoint *Endpoint) GetResource(ctx context.Context, id string) (*GetResourceResponse, error) {
	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/resources/"+id, nil)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *GetResourceResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
