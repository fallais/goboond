package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// SearchSynthesisReporting.
func (endpoint *Endpoint) SearchSynthesisReporting(ctx context.Context, r SearchSynthesisReportingRequest) (*SearchSynthesisReportingResponse, error) {
	// Options
	options := []Option{}
	if r.Period != nil {
		options = append(options, WithParam("period", string(*r.Period)))
	}
	if r.StartDate != nil {
		options = append(options, WithParam("startDate", r.StartDate.Format(DateFormat)))
	}
	if r.EndDate != nil {
		options = append(options, WithParam("endDate", r.EndDate.Format(DateFormat)))
	}
	if r.PeriodDynamic != nil {
		options = append(options, WithParam("periodDynamic", *r.PeriodDynamic))
	}
	if r.PeriodDynamicParameters != nil {
		options = append(options, WithParam("periodDynamicParameters", *r.PeriodDynamicParameters))
	}
	if r.Projects != nil {
		options = append(options, WithParam("projects", formatIntArray(r.Projects)))
	}
	if r.ReportingCategory != nil {
		options = append(options, WithParam("reportingCategory", *r.ReportingCategory))
	}
	if r.ReportingType != nil {
		options = append(options, WithParam("reportingType", *r.ReportingType))
	}
	if r.UseCache != nil {
		options = append(options, WithParam("useCache", *r.UseCache))
	}
	if r.CurrentView != nil {
		options = append(options, WithParam("currentView", *r.CurrentView))
	}
	if r.PerimeterManagers != nil {
		options = append(options, WithParam("perimeterManagers", formatIntArray(r.PerimeterManagers)))
	}
	if r.PerimeterPoles != nil {
		options = append(options, WithParam("perimeterPoles", formatIntArray(r.PerimeterPoles)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/reporting-synthesis", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *SearchSynthesisReportingResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
