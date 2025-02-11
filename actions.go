package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type ListActionsRequest struct {
	ActionTypes             []int      `json:"actionTypes,omitempty"`
	Period                  *Period    `json:"period,omitempty"`
	PeriodDynamic           *string    `json:"periodDynamic,omitempty"`
	PeriodDynamicParameters *string    `json:"periodDynamicParameters,omitempty"`
	PerimeterManagers       []int      `json:"perimeterManagers,omitempty"`
	PerimeterDynamic        *string    `json:"perimeterDynamic,omitempty"`
	ReturnRelatedActions    *bool      `json:"returnRelatedActions,omitempty"`
	StartDate               *time.Time `json:"startDate,omitempty"`
	EndDate                 *time.Time `json:"endDate,omitempty"`
	MaxResults              int        `json:"maxResults,omitempty"`
	Page                    int        `json:"page,omitempty"`
}

// ListActions returns the list of the actions for a candidate.
func (endpoint *Endpoint) ListActions(ctx context.Context, r ListActionsRequest) (*ListActionsResponse, error) {
	// Options
	options := []Option{}
	if r.ActionTypes != nil {
		options = append(options, WithParam("actionTypes", formatIntArray(r.ActionTypes)))
	}
	if r.Period != nil {
		options = append(options, WithParam("period", string(*r.Period)))
	}
	if r.PeriodDynamic != nil {
		options = append(options, WithParam("periodDynamic", *r.PeriodDynamic))
	}
	if r.PerimeterManagers != nil {
		options = append(options, WithParam("perimeterManagers", formatIntArray(r.PerimeterManagers)))
	}
	if r.PerimeterDynamic != nil {
		options = append(options, WithParam("perimeterDynamic", *r.PerimeterDynamic))
	}
	if r.ReturnRelatedActions != nil {
		options = append(options, WithParam("returnRelatedActions", strconv.FormatBool(*r.ReturnRelatedActions)))
	}
	if r.MaxResults != 0 {
		options = append(options, WithParam("maxResults", strconv.Itoa(r.MaxResults)))
	}
	if r.Page != 0 {
		options = append(options, WithParam("page", strconv.Itoa(r.Page)))
	}
	if r.StartDate != nil {
		options = append(options, WithParam("startDate", r.StartDate.Format(DateFormat)))
	}
	if r.EndDate != nil {
		options = append(options, WithParam("endDate", r.EndDate.Format(DateFormat)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/actions", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *ListActionsResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
