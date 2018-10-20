package middleware

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/viriyahendarta/butler-core/config"
	"github.com/viriyahendarta/butler-core/infra/contextx"
	"github.com/viriyahendarta/butler-core/infra/errorx"
	serviceresource "github.com/viriyahendarta/butler-core/resource/service"
)

const (
	AuthHTTPHeader  = "Authorization"
	AuthIDClaimName = "auth_id"
)

type AuthenticationMiddleware interface {
	Middleware
}

type authenticationMiddleware struct {
	serviceResource *serviceresource.Resource
}

var mAuth AuthenticationMiddleware
var once sync.Once

func GetAuthMiddleware(resource *serviceresource.Resource) AuthenticationMiddleware {
	once.Do(func() {
		mAuth = &authenticationMiddleware{
			serviceResource: resource,
		}
	})
	return mAuth
}

func (am *authenticationMiddleware) writeError(ctx context.Context, w http.ResponseWriter, message string, err error) {
	ex := errorx.New(ctx, errorx.CodeUnauthorized, message, err)
	am.serviceResource.RenderJSON(ctx, w, nil, http.StatusUnauthorized, ex)
}

func (am *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get(AuthHTTPHeader)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Get().AuthSecretKey), nil
		})
		if err != nil || !token.Valid {
			am.writeError(r.Context(), w, "Auth token is invalid", err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); !ok {
			am.writeError(r.Context(), w, "Auth token is invalid", errorx.New(r.Context(), http.StatusUnauthorized, "Failed to get claims", nil))
		} else {
			if cAuthID, ok := claims[AuthIDClaimName]; !ok {
				am.writeError(r.Context(), w, "Auth token is invalid", errorx.New(r.Context(), http.StatusUnauthorized, "Auth ID claim is invalid", nil))
			} else if authID, ok := cAuthID.(string); !ok {
				am.writeError(r.Context(), w, "Auth token is invalid", errorx.New(r.Context(), http.StatusUnauthorized, "Malformed auth ID format", nil))
			} else {
				newRequest := r.WithContext(contextx.AppendAuthID(r.Context(), authID))
				next.ServeHTTP(w, newRequest)
			}
		}
	})
}
