package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
)

func (r *igPostResolver) User(ctx context.Context, obj *model.IgPost) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *igPostOpsResolver) Create(ctx context.Context, obj *model.IgPostOps, input model.NewIgPost) (*model.IgPost, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *igPostOpsResolver) Update(ctx context.Context, obj *model.IgPostOps, input model.UpdateIgPost) (*model.IgPost, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *igPostOpsResolver) Archive(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *igPostOpsResolver) Unarchive(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *igPostOpsResolver) Delete(ctx context.Context, obj *model.IgPostOps, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// IgPost returns generated.IgPostResolver implementation.
func (r *Resolver) IgPost() generated.IgPostResolver { return &igPostResolver{r} }

// IgPostOps returns generated.IgPostOpsResolver implementation.
func (r *Resolver) IgPostOps() generated.IgPostOpsResolver { return &igPostOpsResolver{r} }

type igPostResolver struct{ *Resolver }
type igPostOpsResolver struct{ *Resolver }
