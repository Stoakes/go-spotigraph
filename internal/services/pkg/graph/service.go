package graph

import (
	"context"
	"net/http"

	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type service struct {
	users    repositories.User
	squads   repositories.Squad
	chapters repositories.Chapter
	guilds   repositories.Guild
	tribes   repositories.Tribe
}

// New returns a service instance
func New(users repositories.User, squads repositories.Squad, chapters repositories.Chapter, guilds repositories.Guild, tribes repositories.Tribe) services.Graph {
	return &service{
		users:    users,
		squads:   squads,
		chapters: chapters,
		guilds:   guilds,
		tribes:   tribes,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Expand(ctx context.Context, req *spotigraph.NodeInfoReq) (*spotigraph.GraphRes, error) {
	res := &spotigraph.GraphRes{}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	// Resolve node first
	node, err := s.resolveNode(ctx, req.Urn)
	if err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusNotFound,
			Message: "Unable to resolve requested node",
		}
		return res, err
	}

	// Expand edges
	res.Graph, err = s.expandNode(ctx, node)
	if err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Unable to retrieve node's edges",
		}
		return res, err
	}

	// Return result
	return res, nil
}

// -----------------------------------------------------------------------------

func (s *service) resolveNode(ctx context.Context, urn string) (*spotigraph.Graph_Node, error) {
	return nil, nil
}

func (s *service) expandNode(ctx context.Context, n *spotigraph.Graph_Node) (*spotigraph.Graph, error) {
	return nil, nil
}
