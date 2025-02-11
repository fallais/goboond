package goboond

import "time"

type SearchSynthesisReportingRequest struct {
	// StartDate required, string matching ^[0-9]{4}-[0-9]{2}-[0-9]{2}$
	StartDate               *time.Time `json:"startDate,omitempty"`
	EndDate                 *time.Time `json:"endDate,omitempty"`
	Period                  *Period    `json:"period,omitempty"`
	PeriodDynamic           *string    `json:"periodDynamic,omitempty"`
	PeriodDynamicParameters *string    `json:"periodDynamicParameters,omitempty"`
	Projects                []int      `json:"projects,omitempty"`
	ReportingCategory       *string    `json:"reportingCategory,omitempty"`
	ReportingType           *string    `json:"reportingType,omitempty"`
	UseCache                *string    `json:"useCache,omitempty"`
	CurrentView             *string    `json:"currentView,omitempty"`
	PerimeterManagers       []int      `json:"perimeterManagers,omitempty"`
	PerimeterPoles          []int      `json:"perimeterPoles,omitempty"`
}

type SearchSynthesisReportingResponse struct {
	Meta struct {
		Cache struct {
			State      bool   `json:"state"`
			DateUpdate string `json:"dateUpdate"`
		} `json:"cache"`
		Dates struct {
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
		} `json:"dates"`
		Version           string `json:"version"`
		AndroidMinVersion string `json:"androidMinVersion"`
		IosMinVersion     string `json:"iosMinVersion"`
		IsLogged          bool   `json:"isLogged"`
		Language          string `json:"language"`
		Timestamp         int64  `json:"timestamp"`
		Login             string `json:"login"`
		Customer          string `json:"customer"`
	} `json:"meta"`
	Data []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			EndDate   string `json:"endDate"`
			Scorecard struct {
				Category  string `json:"category"`
				Reference string `json:"reference"`
				TypeOf    string `json:"typeOf"`
			} `json:"scorecard"`
			StartDate string `json:"startDate"`
			Target    any    `json:"target"`
			Value     string `json:"value"`
		} `json:"attributes"`
	} `json:"data"`
}
