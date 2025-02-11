package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SearchPositioningsRequest struct {
	Colomns   []string   `json:"columns"`
	Encoding  *string    `json:"encoding"`
	Period    *Period    `json:"period"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`

	SearchOptions
}

type SearchPositioningsResponse struct {
	Meta struct {
		Totals struct {
			Rows int `json:"rows"`
		} `json:"totals"`
		Version           string `json:"version"`
		AndroidMinVersion string `json:"androidMinVersion"`
		IosMinVersion     string `json:"iosMinVersion"`
		IsLogged          bool   `json:"isLogged"`
		Language          string `json:"language"`
		Timestamp         int64  `json:"timestamp"`
		Login             string `json:"login"`
		Customer          string `json:"customer"`
	} `json:"meta"`
	Data []PositioningData `json:"data"`
}

type PositioningData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		CreationDate string `json:"creationDate"`
		UpdateDate   string `json:"updateDate"`
		State        int    `json:"state"`
		StateReason  struct {
			TypeOf int    `json:"typeOf"`
			Detail string `json:"detail"`
		} `json:"stateReason"`
		StartDate                  string `json:"startDate"`
		EndDate                    string `json:"endDate"`
		InformationComments        string `json:"informationComments"`
		CanReadPositioning         bool   `json:"canReadPositioning"`
		CanShowDependsOn           bool   `json:"canShowDependsOn"`
		CanShowOpportunity         bool   `json:"canShowOpportunity"`
		CanShowCreationDate        bool   `json:"canShowCreationDate"`
		CanShowUpdateDate          bool   `json:"canShowUpdateDate"`
		CanShowState               bool   `json:"canShowState"`
		CanShowStartDate           bool   `json:"canShowStartDate"`
		CanShowEndDate             bool   `json:"canShowEndDate"`
		CanShowInformationComments bool   `json:"canShowInformationComments"`
	} `json:"attributes"`
	Relationships struct {
		CreatedBy struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"createdBy"`
		Opportunity struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"opportunity"`
		DependsOn struct {
			Data struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"dependsOn"`
	} `json:"relationships"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListResources returns the list of the resources.
func (endpoint *Endpoint) SearchPositionings(ctx context.Context, r SearchPositioningsRequest) (*SearchPositioningsResponse, error) {
	// Options
	options := []Option{}
	if r.Colomns != nil {
		options = append(options, WithParam("columns", formatStringArray(r.Colomns)))
	}
	if r.Encoding != nil {
		options = append(options, WithParam("encoding", *r.Encoding))
	}
	if r.Period != nil {
		options = append(options, WithParam("period", string(*r.Period)))
	}
	if r.StartDate != nil {
		options = append(options, WithParam("startDate", r.StartDate.Format(DateFormat)))
	}
	if r.EndDate != nil {
		options = append(options, WithParam("endDate", r.EndDate.Format(DateFormat)))
	}
	if r.MaxResults != 0 {
		options = append(options, WithParam("maxResults", fmt.Sprintf("%d", r.MaxResults)))
	}
	if r.Page != 0 {
		options = append(options, WithParam("page", fmt.Sprintf("%d", r.Page)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/positionings", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *SearchPositioningsResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
