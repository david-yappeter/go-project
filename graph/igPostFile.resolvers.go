package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
)

func (r *igPostFileResolver) ViewLink(ctx context.Context, obj *model.IgPostFile) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *igPostFileResolver) DownloadLink(ctx context.Context, obj *model.IgPostFile) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// IgPostFile returns generated.IgPostFileResolver implementation.
func (r *Resolver) IgPostFile() generated.IgPostFileResolver { return &igPostFileResolver{r} }

type igPostFileResolver struct{ *Resolver }
