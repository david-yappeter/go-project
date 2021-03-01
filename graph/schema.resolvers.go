package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
)

func (r *mutationResolver) File(ctx context.Context) (*model.FileUploadOps, error) {
	return &model.FileUploadOps{}, nil
}

func (r *mutationResolver) Token(ctx context.Context) (*model.TokenOps, error) {
	return &model.TokenOps{}, nil
}

func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

func (r *mutationResolver) IgPost(ctx context.Context) (*model.IgPostOps, error) {
	return &model.IgPostOps{}, nil
}

func (r *queryResolver) User(ctx context.Context, id int, scopes *bool) (*model.User, error) {
	return service.UserGetByID(ctx, id, scopes)
}

func (r *queryResolver) Users(ctx context.Context, limit *int, page *int, asc *bool, sortBy *string, scopes *bool) (*model.UserPagination, error) {
	return &model.UserPagination{Limit: limit, Page: page, Asc: asc, SortBy: sortBy, Scopes: scopes}, nil
}

func (r *queryResolver) File(ctx context.Context, id int, scopes *bool) (*model.FileUpload, error) {
	return service.FileUploadGetByID(ctx, id, scopes)
}

func (r *queryResolver) Files(ctx context.Context, userID *int, limit *int, page *int, asc *bool, sortBy *string, scopes *bool) (*model.FileUploadPagination, error) {
	return &model.FileUploadPagination{Limit: limit, Page: page, Asc: asc, SortBy: sortBy, Scopes: scopes, UserID: userID}, nil
}

func (r *queryResolver) GithubRepositories(ctx context.Context, username string) ([]*model.UserGithubRepository, error) {
	return service.GithubGetUserRepositories(ctx, username)
}

func (r *queryResolver) IgPost(ctx context.Context, id int) (*model.IgPost, error) {
	return service.IgPostGetByID(ctx, id)
}

func (r *queryResolver) IgPosts(ctx context.Context, limit *int, page *int, asc *bool, sortBy *string, scopes *bool) (*model.IgPostPagination, error) {
	return &model.IgPostPagination{Limit: limit, Page: page, Asc: asc, SortBy: sortBy, Scopes: scopes}, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return service.UserGetByToken(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
