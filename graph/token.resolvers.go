package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/graph/model"
	"github.com/davidyap2002/user-go/service"
)

func (r *tokenOpsResolver) Login(ctx context.Context, obj *model.TokenOps, email string, password string) (*model.TokenData, error) {
	return service.UserTokenCreateByEmail(ctx, email, password)
}

// TokenOps returns generated.TokenOpsResolver implementation.
func (r *Resolver) TokenOps() generated.TokenOpsResolver { return &tokenOpsResolver{r} }

type tokenOpsResolver struct{ *Resolver }
