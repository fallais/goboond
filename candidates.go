package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ListCandidatesResponse struct {
	Meta CandidatesMeta `json:"meta"`
	Data []Data         `json:"data"`
}
type CandidatesMetaTotals struct {
	Rows int `json:"rows"`
}
type CandidatesMeta struct {
	Totals            CandidatesMetaTotals `json:"totals"`
	Solr              bool                 `json:"solr"`
	Version           string               `json:"version"`
	AndroidMinVersion string               `json:"androidMinVersion"`
	IosMinVersion     string               `json:"iosMinVersion"`
	IsLogged          bool                 `json:"isLogged"`
	Language          string               `json:"language"`
	Timestamp         int64                `json:"timestamp"`
	Customer          string               `json:"customer"`
}
type CandidatesSource struct {
	TypeOf int    `json:"typeOf"`
	Detail string `json:"detail"`
}
type CandidatesSocialNetworks struct {
	Network string `json:"network"`
	URL     string `json:"url"`
}
type CandidatesAttributes struct {
	CreationDate               string                     `json:"creationDate"`
	UpdateDate                 string                     `json:"updateDate"`
	Civility                   int                        `json:"civility"`
	FirstName                  string                     `json:"firstName"`
	LastName                   string                     `json:"lastName"`
	State                      int                        `json:"state"`
	TypeOf                     int                        `json:"typeOf"`
	IsVisible                  bool                       `json:"isVisible"`
	Thumbnail                  string                     `json:"thumbnail"`
	Availability               int                        `json:"availability"`
	Skills                     string                     `json:"skills"`
	Diplomas                   []any                      `json:"diplomas"`
	MobilityAreas              []string                   `json:"mobilityAreas"`
	ActivityAreas              []any                      `json:"activityAreas"`
	GlobalEvaluation           string                     `json:"globalEvaluation"`
	Languages                  []any                      `json:"languages"`
	ExpertiseAreas             []any                      `json:"expertiseAreas"`
	Experience                 int                        `json:"experience"`
	References                 []any                      `json:"references"`
	Evaluations                []any                      `json:"evaluations"`
	Tools                      []any                      `json:"tools"`
	Title                      string                     `json:"title"`
	Email1                     string                     `json:"email1"`
	Email2                     string                     `json:"email2"`
	Email3                     string                     `json:"email3"`
	Phone1                     string                     `json:"phone1"`
	Phone2                     string                     `json:"phone2"`
	Town                       string                     `json:"town"`
	Country                    string                     `json:"country"`
	Source                     CandidatesSource           `json:"source"`
	NumberOfResumes            int                        `json:"numberOfResumes"`
	NumberOfActivePositionings int                        `json:"numberOfActivePositionings"`
	SocialNetworks             []CandidatesSocialNetworks `json:"socialNetworks"`
	CanShowTechnicalData       bool                       `json:"canShowTechnicalData"`
	CanShowActions             bool                       `json:"canShowActions"`
}

type CandidatesMainManager struct {
	Data CandidatesData `json:"data"`
}
type CandidatesAgency struct {
	Data CandidatesData `json:"data"`
}
type CandidatesPole struct {
	Data any `json:"data"`
}
type CandidatesRelationships struct {
	MainManager MainManager `json:"mainManager"`
	Agency      Agency      `json:"agency"`
	Pole        Pole        `json:"pole"`
}
type CandidatesData struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListCandidates returns the list of the candidates.
func (endpoint *Endpoint) ListCandidates(ctx context.Context, candidateStates string) (*ListCandidatesResponse, error) {
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
	var response *ListCandidatesResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
