package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fallais/goboond/responses/actions"
)

// ListActions returns the list of the actions for a candidate.
func (endpoint *Endpoint) ListActions(ctx context.Context, actionTypes, period, periodDynamic string, maxResults, page int) (*actions.ListActionsResponse, error) {
	// Options
	options := []Option{}
	if actionTypes != "" {
		options = append(options, WithParam("actionTypes", actionTypes))
	}
	if period != "" {
		options = append(options, WithParam("period", period))
	}
	if periodDynamic != "" {
		options = append(options, WithParam("periodDynamic", periodDynamic))
	}
	if maxResults != 0 {
		options = append(options, WithParam("maxResults", strconv.Itoa(maxResults)))
	}
	if page != 0 {
		options = append(options, WithParam("page", strconv.Itoa(page)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/actions", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *actions.ListActionsResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
