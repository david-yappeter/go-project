package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
)

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

func (r *mutationResolver) Token(ctx context.Context) (*model.TokenOps, error) {
	return &model.TokenOps{}, nil
}

func (r *queryResolver) User(ctx context.Context, id int, scopes *bool) (*model.User, error) {
	return service.UserGetByID(ctx, id, scopes)
}

func (r *queryResolver) Users(ctx context.Context, limit *int, page *int, asc *bool, sortBy *string, filter []*int, scopes *bool) (*model.UserPagination, error) {
	return &model.UserPagination{Limit: limit, Page: page, Asc: asc, SortBy: sortBy, Filter: filter, Scopes: scopes}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
