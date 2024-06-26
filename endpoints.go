package goboond

import (
	"context"

	"github.com/fallais/goboond/responses/actions"
	"github.com/fallais/goboond/responses/candidates"
	reportingsynthesis "github.com/fallais/goboond/responses/reporting_synthesis"
	"github.com/fallais/goboond/responses/resources"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Endpoint is an API endpoint.
type Endpoint struct {
	client *Client
}

//------------------------------------------------------------------------------
// Interfaces
//------------------------------------------------------------------------------

// Actions endpoint.
type Actions interface {
	ListActions(context.Context, string, string, int, int) (*actions.ListActionsResponse, error)
}

// Candidates endpoint.
type Candidates interface {
	ListCandidates(context.Context, string, int, int) (*candidates.ListCandidatesResponse, error)
	ListCandidatesActions(context.Context, string, int, int) (*candidates.ListActionsResponse, error)
	GetInformation(context.Context, string, int, int) (*candidates.GetInformationResponse, error)
}

// Projects endpoint.
type Projects interface {
	//ListProjects(context.Context, string, string, string, int, int) (*ProjectsPaginatedResponse, error)
}

// Resources endpoint.
type Resources interface {
	ListResources(context.Context, string, string) (*resources.ListResourcesResponse, error)
	GetResource(context.Context, string) (*resources.GetResourceResponse, error)
}

// Reporting synthesis
type ReportingSynthesis interface {
	SearchSynthesisReporting(context.Context, string, string) (*reportingsynthesis.SearchSynthesisReportingResponse, error)
}
