package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ListCompaniesRequest struct {
	MaxResults int `url:"max_results,omitempty"`
	Page       int `url:"page,omitempty"`
}

type ListCompaniesResponse struct {
	Meta struct {
		Version  string `json:"version"`
		IsLogged bool   `json:"isLogged"`
		Language string `json:"language"`
		Totals   struct {
			Rows int `json:"rows"`
		} `json:"totals,omitempty"`
		Solr bool `json:"solr,omitempty"`
	} `json:"meta"`
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name                      string `json:"name"`
			ExpertiseArea             string `json:"expertiseArea,omitempty"`
			State                     int    `json:"state,omitempty"`
			InformationComments       string `json:"informationComments,omitempty"`
			Thumbnail                 string `json:"thumbnail,omitempty"`
			Website                   string `json:"website,omitempty"`
			Phone1                    string `json:"phone1,omitempty"`
			Town                      string `json:"town,omitempty"`
			Country                   string `json:"country,omitempty"`
			CreationDate              string `json:"creationDate,omitempty"`
			NumberOfActiveOpportunity int    `json:"numberbOfActiveOpportunity,omitempty"`
			UpdateDate                string `json:"updateDate,omitempty"`
		} `json:"attributes"`
		Relationships struct {
			MainManager struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"mainManager,omitempty"`
			Agency struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"agency,omitempty"`
			Pole struct {
				Data *struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"pole,omitempty"`
			PreviousAction struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"previousAction,omitempty"`
			NextAction struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"nextAction,omitempty"`
		} `json:"relationships,omitempty"`
	} `json:"data"`
	Included []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			FirstName string `json:"firstName,omitempty"`
			LastName  string `json:"lastName,omitempty"`
			Name      string `json:"name,omitempty"`
			StartDate string `json:"startDate,omitempty"`
			TypeOf    int    `json:"typeOf,omitempty"`
			Text      string `json:"text,omitempty"`
		} `json:"attributes,omitempty"`
	} `json:"included,omitempty"`
}

func (endpoint *Endpoint) ListCompanies(ctx context.Context, r ListCompaniesRequest) (*ListCompaniesResponse, error) {
	// Options
	options := []Option{}
	if r.MaxResults != 0 {
		options = append(options, WithParam("maxResults", fmt.Sprintf("%d", r.MaxResults)))
	}
	if r.Page != 0 {
		options = append(options, WithParam("page", fmt.Sprintf("%d", r.Page)))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/companies", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *ListCompaniesResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
