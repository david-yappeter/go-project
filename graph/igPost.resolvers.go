package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidyap2002/user-go/dataloader"
	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
)

func (r *igPostResolver) Files(ctx context.Context, obj *model.IgPost) ([]*model.IgPostFile, error) {
	return dataloader.For(ctx).IgPostFileByPostIds.Load(obj.ID)
}

func (r *igPostResolver) User(ctx context.Context, obj *model.IgPost) (*model.User, error) {
	return service.UserGetByID(ctx, obj.UserID, nil)
}

func (r *igPostOpsResolver) Create(ctx context.Context, obj *model.IgPostOps, input model.NewIgPost) (*model.IgPost, error) {
	return service.IgPostCreate(ctx, input)
}

func (r *igPostOpsResolver) Update(ctx context.Context, obj *model.IgPostOps, input model.UpdateIgPost) (*model.IgPost, error) {
	return service.IgPostUpdate(ctx, input)
}

func (r *igPostOpsResolver) Archive(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	return service.IgPostArchive(ctx, id)
}

func (r *igPostOpsResolver) Unarchive(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	return service.IgPostUnarchive(ctx, id)
}

func (r *igPostOpsResolver) SoftDelete(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	return service.IgPostSoftDelete(ctx, id)
}

func (r *igPostOpsResolver) RestoreDelete(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	return service.IgPostRestoreDelete(ctx, id)
}

func (r *igPostOpsResolver) HardDelete(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	return service.IgPostDelete(ctx, id)
}

func (r *igPostPaginationResolver) TotalData(ctx context.Context, obj *model.IgPostPagination) (int, error) {
	return service.IgPostTotalDataPagination(ctx, obj.Limit, obj.Page, obj.Asc, obj.SortBy, obj.Scopes)
}

func (r *igPostPaginationResolver) Nodes(ctx context.Context, obj *model.IgPostPagination) ([]*model.IgPost, error) {
	return service.IgPostPagination(ctx, obj.Limit, obj.Page, obj.Asc, obj.SortBy, obj.Scopes)
}

// IgPost returns generated.IgPostResolver implementation.
func (r *Resolver) IgPost() generated.IgPostResolver { return &igPostResolver{r} }

// IgPostOps returns generated.IgPostOpsResolver implementation.
func (r *Resolver) IgPostOps() generated.IgPostOpsResolver { return &igPostOpsResolver{r} }

// IgPostPagination returns generated.IgPostPaginationResolver implementation.
func (r *Resolver) IgPostPagination() generated.IgPostPaginationResolver {
	return &igPostPaginationResolver{r}
}

type igPostResolver struct{ *Resolver }
type igPostOpsResolver struct{ *Resolver }
type igPostPaginationResolver struct{ *Resolver }
