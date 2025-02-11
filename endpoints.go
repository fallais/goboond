package goboond

import (
	"context"

	"github.com/fallais/goboond/responses/candidates"
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
	ListActions(context.Context, ListActionsRequest) (*ListActionsResponse, error)
}

// Candidates endpoint.
type Candidates interface {
	ListCandidates(context.Context, string, int, int) (*candidates.ListCandidatesResponse, error)
	ListCandidatesActions(context.Context, string, int, int) (*candidates.ListActionsResponse, error)
	GetInformation(context.Context, string, int, int) (*candidates.GetInformationResponse, error)
}

// Positionings endpoint.
type Positionings interface {
	SearchPositionings(context.Context, SearchPositioningsRequest) (*SearchPositioningsResponse, error)
}

// Projects endpoint.
type Projects interface {
	//ListProjects(context.Context, string, string, string, int, int) (*ProjectsPaginatedResponse, error)
}

// Resources endpoint.
type Resources interface {
	ListResources(context.Context, ListResourcesRequest) (*ListResourcesResponse, error)
	GetResource(context.Context, string) (*GetResourceResponse, error)
}

// Reporting synthesis
type ReportingSynthesis interface {
	SearchSynthesisReporting(context.Context, SearchSynthesisReportingRequest) (*SearchSynthesisReportingResponse, error)
}

// Contacts endpoint.
type Contacts interface {
	ListContacts(context.Context, ListContactsRequest) (*ListContactsResponse, error)
}

// Companies endpoint.
type Companies interface {
	ListCompanies(context.Context, ListCompaniesRequest) (*ListCompaniesResponse, error)
}
