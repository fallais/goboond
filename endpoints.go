package goboond

import (
	"context"

	"github.com/fallais/goboond/responses/candidates"
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

// Candidates endpoint.
type Candidates interface {
	ListCandidates(context.Context, string) (*candidates.ListCandidatesResponse, error)
}

// Projects endpoint.
type Projects interface {
	//ListProjects(context.Context, string, string, string, int, int) (*ProjectsPaginatedResponse, error)
}

// Resources endpoint.
type Resources interface {
	ListResources(context.Context, string) (*resources.ListResourcesResponse, error)
	GetResource(context.Context, string) (*resources.GetResourceResponse, error)
}
