package service

import (
	"context"
	"net/http"
	"strings"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			header := strings.Trim(r.Header.Get("Authorization"), " ")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			auth := strings.Split(header, " ")

			if strings.Trim(auth[0], " ") != "Bearer" {
				http.Error(w, "Auth Type Not Included", http.StatusBadRequest)
				return
			}

			if len(auth) != 2 {
				http.Error(w, "Invalid Token", http.StatusBadRequest)
				return
			}

			tokenBeforeClaims, err := UserTokenValidate(strings.Trim(auth[1], " "))

			if err != nil {
				http.Error(w, "Invalid Token", http.StatusForbidden)
				return
			}

			claims, ok := tokenBeforeClaims.Claims.(*UserClaims)
			if !ok && !tokenBeforeClaims.Valid {
				http.Error(w, "Invalid Token", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), userCtxKey, claims)

			// and call the next with our new context
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *UserClaims {
	raw, _ := ctx.Value(userCtxKey).(*UserClaims)
	return raw
}
