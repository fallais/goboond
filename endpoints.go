package goboond

import (
	"context"
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
	//ListCandidates(context.Context, string, string, string, int, int) (*ProjectsPaginatedResponse, error)
}

// Projects endpoint.
type Projects interface {
	//ListProjects(context.Context, string, string, string, int, int) (*ProjectsPaginatedResponse, error)
}

// Resources endpoint.
type Resources interface {
	ListResources(context.Context, string) (*ResourcesPaginatedResponse, error)
	GetResource(context.Context, string) (*ResourceResponse, error)
}
