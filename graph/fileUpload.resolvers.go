package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
)

func (r *fileUploadResolver) ViewLink(ctx context.Context, obj *model.FileUpload) (string, error) {
	return service.GdriveViewLink(ctx, obj.FileID)
}

func (r *fileUploadResolver) DownloadLink(ctx context.Context, obj *model.FileUpload) (string, error) {
	return service.GdriveDownloadLink(ctx, obj.FileID)
}

func (r *fileUploadOpsResolver) UploadSingle(ctx context.Context, obj *model.FileUploadOps, file graphql.Upload) (*model.FileUpload, error) {
	return service.UploadFile(ctx, file)
}

func (r *fileUploadOpsResolver) UploadBatch(ctx context.Context, obj *model.FileUploadOps, files []*graphql.Upload) ([]*model.FileUpload, error) {
	return service.UploadFileBatch(ctx, files)
}

func (r *fileUploadPaginationResolver) TotalData(ctx context.Context, obj *model.FileUploadPagination) (int, error) {
	return service.FileUploadTotalDataPagination(ctx, obj.Limit, obj.Page, obj.Asc, obj.SortBy, obj.Scopes, obj.UserID)
}

func (r *fileUploadPaginationResolver) Nodes(ctx context.Context, obj *model.FileUploadPagination) ([]*model.FileUpload, error) {
	return service.FileUploadPagination(ctx, obj.Limit, obj.Page, obj.Asc, obj.SortBy, obj.Scopes, obj.UserID)
}

// FileUpload returns generated.FileUploadResolver implementation.
func (r *Resolver) FileUpload() generated.FileUploadResolver { return &fileUploadResolver{r} }

// FileUploadOps returns generated.FileUploadOpsResolver implementation.
func (r *Resolver) FileUploadOps() generated.FileUploadOpsResolver { return &fileUploadOpsResolver{r} }

// FileUploadPagination returns generated.FileUploadPaginationResolver implementation.
func (r *Resolver) FileUploadPagination() generated.FileUploadPaginationResolver {
	return &fileUploadPaginationResolver{r}
}

type fileUploadResolver struct{ *Resolver }
type fileUploadOpsResolver struct{ *Resolver }
type fileUploadPaginationResolver struct{ *Resolver }
