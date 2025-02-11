package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fallais/goboond/responses/candidates"
)

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListCandidates returns the list of the candidates.
func (endpoint *Endpoint) ListCandidates(ctx context.Context, candidateStates string, maxResults, page int) (*candidates.ListCandidatesResponse, error) {
	// Options
	options := []Option{}
	if candidateStates != "" {
		options = append(options, WithParam("candidateStates", candidateStates))
	}
	if maxResults != 0 {
		options = append(options, WithParam("maxResults", strconv.Itoa(maxResults)))
	}
	if page != 0 {
		options = append(options, WithParam("page", strconv.Itoa(page)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/candidates", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *candidates.ListCandidatesResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListCandidatesActions returns the list of the actions for a candidate.
func (endpoint *Endpoint) ListCandidatesActions(ctx context.Context, candidateID string, maxResults, page int) (*candidates.ListActionsResponse, error) {
	// Options
	options := []Option{}
	if maxResults != 0 {
		options = append(options, WithParam("maxResults", strconv.Itoa(maxResults)))
	}
	if page != 0 {
		options = append(options, WithParam("page", strconv.Itoa(page)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/candidates/"+candidateID+"/actions", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *candidates.ListActionsResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetInformation returns the information of the candidate.
func (endpoint *Endpoint) GetInformation(ctx context.Context, candidateID string, maxResults, page int) (*candidates.GetInformationResponse, error) {
	// Options
	options := []Option{}
	if maxResults != 0 {
		options = append(options, WithParam("maxResults", strconv.Itoa(maxResults)))
	}
	if page != 0 {
		options = append(options, WithParam("page", strconv.Itoa(page)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/candidates/"+candidateID+"/information", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *candidates.GetInformationResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
