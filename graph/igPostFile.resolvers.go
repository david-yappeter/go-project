package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
)

func (r *igPostFileResolver) ViewLink(ctx context.Context, obj *model.IgPostFile) (string, error) {
	return service.GdriveViewLink(ctx, obj.FileID)
}

func (r *igPostFileResolver) DownloadLink(ctx context.Context, obj *model.IgPostFile) (string, error) {
	return service.GdriveDownloadLink(ctx, obj.FileID)
}

// IgPostFile returns generated.IgPostFileResolver implementation.
func (r *Resolver) IgPostFile() generated.IgPostFileResolver { return &igPostFileResolver{r} }

type igPostFileResolver struct{ *Resolver }
