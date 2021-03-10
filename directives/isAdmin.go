package directives

import (
	"context"

	"github.com/davidyap2002/user-go/service"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

//IsAdmin IsAdmin Directive
func IsAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := service.ForContext(ctx)

	if user == nil {
		return nil, &gqlerror.Error{
			Message: "Access denied",
			Extensions: map[string]interface{}{
				"code": "UNAUTHENTICATE",
			},
		}
	}

	getUser, err := service.UserGetByID(ctx, user.ID, nil)

	if err == gorm.ErrRecordNotFound {
		return nil, &gqlerror.Error{
			Message: "Invalid Token",
			Extensions: map[string]interface{}{
				"code": "UNAUTHENTICATE",
			},
		}
	}

	if getUser.IsAdmin == 0 {
		return nil, &gqlerror.Error{
			Message: "Access denied: Not Admin",
			Extensions: map[string]interface{}{
				"code": "UNAUTHENTICATE",
			},
		}
	}

	return next(ctx)
}
