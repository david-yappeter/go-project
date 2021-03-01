package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
)

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

// IgPost returns generated.IgPostResolver implementation.
func (r *Resolver) IgPost() generated.IgPostResolver { return &igPostResolver{r} }

// IgPostOps returns generated.IgPostOpsResolver implementation.
func (r *Resolver) IgPostOps() generated.IgPostOpsResolver { return &igPostOpsResolver{r} }

type igPostResolver struct{ *Resolver }
type igPostOpsResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *igPostOpsResolver) Delete(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	panic("a")
}
