package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
	"github.com/davidyap2002/user-go/tools"
)

func (r *userOpsResolver) Create(ctx context.Context, obj *model.UserOps, input model.NewUser) (*model.User, error) {
	return service.UserCreate(ctx, input)
}

func (r *userOpsResolver) Update(ctx context.Context, obj *model.UserOps, input model.UpdateUser) (*model.User, error) {
	return service.UserUpdate(ctx, input)
}

func (r *userOpsResolver) SoftDelete(ctx context.Context, obj *model.UserOps, id int) (string, error) {
	return service.UserSoftDelete(ctx, id)
}

func (r *userOpsResolver) RestoreDelete(ctx context.Context, obj *model.UserOps, id int) (string, error) {
	return service.UserRestoreDelete(ctx, id)
}

func (r *userOpsResolver) Delete(ctx context.Context, obj *model.UserOps, id int) (string, error) {
	return service.UserHardDelete(ctx, id)
}

func (r *userPaginationResolver) TotalData(ctx context.Context, obj *model.UserPagination) (int, error) {
	limit, page, asc, sort, status := tools.PaginationVariableGenerator(obj.Limit, obj.Page, obj.Asc, obj.SortBy, obj.Filter)
	return service.UserTotalDataPagination(ctx, limit, page, asc, sort, status, obj.Scopes)
}

func (r *userPaginationResolver) Nodes(ctx context.Context, obj *model.UserPagination) ([]*model.User, error) {
	limit, page, asc, sort, status := tools.PaginationVariableGenerator(obj.Limit, obj.Page, obj.Asc, obj.SortBy, obj.Filter)
	return service.UserPagination(ctx, limit, page, asc, sort, status, obj.Scopes)
}

// UserOps returns generated.UserOpsResolver implementation.
func (r *Resolver) UserOps() generated.UserOpsResolver { return &userOpsResolver{r} }

// UserPagination returns generated.UserPaginationResolver implementation.
func (r *Resolver) UserPagination() generated.UserPaginationResolver {
	return &userPaginationResolver{r}
}

type userOpsResolver struct{ *Resolver }
type userPaginationResolver struct{ *Resolver }
