package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ListContactsRequest struct {
	StartDate  *time.Time `json:"startDate,omitempty"`
	EndDate    *time.Time `json:"endDate,omitempty"`
	MaxResults int        `json:"maxResults,omitempty"`
	Page       int        `json:"page,omitempty"`
}

type ListContactsResponse struct {
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
			CreationDate    string   `json:"creationDate,omitempty"`
			Civility        int      `json:"civility,omitempty"`
			Thumbnail       string   `json:"thumbnail,omitempty"`
			FirstName       string   `json:"firstName,omitempty"`
			LastName        string   `json:"lastName,omitempty"`
			State           int      `json:"state,omitempty"`
			Function        string   `json:"function,omitempty"`
			Department      string   `json:"department,omitempty"`
			Email1          string   `json:"email1,omitempty"`
			Email2          string   `json:"email2,omitempty"`
			Email3          string   `json:"email3,omitempty"`
			Phone1          string   `json:"phone1,omitempty"`
			Phone2          string   `json:"phone2,omitempty"`
			Town            string   `json:"town,omitempty"`
			Country         string   `json:"country,omitempty"`
			UpdateDate      string   `json:"updateDate,omitempty"`
			CanReadContact  bool     `json:"canReadContact,omitempty"`
			CanWriteContact bool     `json:"canWriteContact,omitempty"`
			CanShowAction   bool     `json:"canShowAction,omitempty"`
			TypesOf         []string `json:"typesOf,omitempty"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			MainManager struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"mainManager,omitempty"`
			Company struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"company,omitempty"`
			LastAction struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"lastAction,omitempty"`
			Agency struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"agency,omitempty"`
			Pole struct {
				Data struct {
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
		ID         string                 `json:"id"`
		Type       string                 `json:"type"`
		Attributes map[string]interface{} `json:"attributes,omitempty"`
	} `json:"included,omitempty"`
}

func (endpoint *Endpoint) ListContacts(ctx context.Context, r ListContactsRequest) (*ListContactsResponse, error) {
	// Options
	options := []Option{}
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
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/contacts", nil, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Prepare the response
	var response *ListContactsResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
