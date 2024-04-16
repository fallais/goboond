package goboond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CandidatesMeta struct {
	Version  string `json:"version"`
	IsLogged bool   `json:"isLogged"`
	Language string `json:"language"`
	Totals   struct {
		Rows int `json:"rows"`
	} `json:"totals"`
	Solr bool `json:"solr"`
}

type CandidatesAttributes struct {
	CreationDate               time.Time                 `json:"creationDate"`
	UpdateDate                 time.Time                 `json:"updateDate"`
	Civility                   int                       `json:"civility"`
	Thumbnail                  string                    `json:"thumbnail"`
	FirstName                  string                    `json:"firstName"`
	LastName                   string                    `json:"lastName"`
	TypeOf                     int                       `json:"typeOf"`
	State                      int                       `json:"state"`
	IsVisible                  bool                      `json:"isVisible"`
	Skills                     string                    `json:"skills"`
	MobilityAreas              []string                  `json:"mobilityAreas"`
	Title                      string                    `json:"title"`
	Availability               string                    `json:"availability"`
	Email1                     string                    `json:"email1"`
	Email2                     string                    `json:"email2"`
	Email3                     string                    `json:"email3"`
	Phone1                     string                    `json:"phone1"`
	Phone2                     string                    `json:"phone2"`
	Town                       string                    `json:"town"`
	Country                    string                    `json:"country"`
	Source                     CandidatesSource          `json:"source"`
	NumberOfResumes            int                       `json:"numberOfResumes"`
	NumberOfActivePositionings int                       `json:"numberOfActivePositionings"`
	SocialNetworks             []CandidatesSocialNetwork `json:"socialNetworks"`
	Diplomas                   []string                  `json:"diplomas"`
	ActivityAreas              []string                  `json:"activityAreas"`
	GlobalEvaluation           string                    `json:"globalEvaluation"`
	Languages                  []CandidatesLanguage      `json:"languages"`
	ExpertiseAreas             []string                  `json:"expertiseAreas"`
	Experience                 int                       `json:"experience"`
	References                 []CandidatesReference     `json:"references"`
	Evaluations                []CandidatesEvaluation    `json:"evaluations"`
	Tools                      []CandidatesTool          `json:"tools"`
	CanShowTechnicalData       bool                      `json:"canShowTechnicalData"`
	CanShowActions             bool                      `json:"canShowActions"`
}

type CandidatesData struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    CandidatesAttributes   `json:"attributes"`
	Relationships CandidatesRelationship `json:"relationships"`
}

type CandidatesSource struct {
	TypeOf int    `json:"typeOf"`
	Detail string `json:"detail"`
}

type CandidatesSocialNetwork struct {
	Network string `json:"network"`
	URL     string `json:"url"`
}

type CandidatesLanguage struct {
	Language string `json:"language"`
	Level    string `json:"level"`
}

type CandidatesReference struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CandidatesEvaluation struct {
	ID        string               `json:"id"`
	Notations []CandidatesNotation `json:"notations"`
	Date      string               `json:"date"`
	Comments  string               `json:"comments"`
	Manager   CandidatesManager    `json:"manager"`
}

type CandidatesNotation struct {
	Criteria   int    `json:"criteria"`
	Evaluation string `json:"evaluation"`
}

type CandidatesManager struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CandidatesTool struct {
	Tool  string `json:"tool"`
	Level int    `json:"level"`
}

type CandidatesRelationship struct {
	MainManager    CandidatesRelationshipData `json:"mainManager"`
	HRManager      CandidatesRelationshipData `json:"hrManager"`
	Agency         CandidatesRelationshipData `json:"agency"`
	Pole           CandidatesRelationshipData `json:"pole"`
	PreviousAction CandidatesRelationshipData `json:"previousAction"`
	NextAction     CandidatesRelationshipData `json:"nextAction"`
	LastAction     CandidatesRelationshipData `json:"lastAction"`
}

type CandidatesRelationshipData struct {
	Data Data `json:"data"`
}

type CandidatesResponse struct {
	Meta     Meta             `json:"meta"`
	Data     []CandidatesData `json:"data"`
	Included []CandidatesData `json:"included"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListCandidates returns the list of the candidates.
func (endpoint *Endpoint) ListCandidates(ctx context.Context, candidateStates string) (*CandidatesResponse, error) {
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
	var response *CandidatesResponse

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}
