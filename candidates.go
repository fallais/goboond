package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fallais/goboond/responses/candidates"
)

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListCandidates returns the list of the candidates.
func (endpoint *Endpoint) ListCandidates(ctx context.Context, candidateStates string) (*candidates.ListCandidatesResponse, error) {
	// Options
	options := []Option{}
	if candidateStates != "" {
		options = append(options, WithParam("candidateStates", candidateStates))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/candidates", options...)
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
