package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	reportingsynthesis "github.com/fallais/goboond/responses/reporting_synthesis"
)

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// SearchSynthesisReporting.
func (endpoint *Endpoint) SearchSynthesisReporting(ctx context.Context) (*reportingsynthesis.SearchSynthesisReportingResponse, error) {
	// Options
	options := []Option{}
	/* 	if resourceStates != "" {
	   		options = append(options, WithParam("resourceStates", resourceStates))
	   	}
	   	if resourceTypes != "" {
	   		options = append(options, WithParam("resourceTypes", resourceTypes))
	   	} */

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/reporting-synthesis", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyBytes))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *reportingsynthesis.SearchSynthesisReportingResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
