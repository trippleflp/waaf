package auth

import (
	"context"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/libs/token"
	"net/http"
	"strings"
)

type contextKey string

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userIDKey contextKey = "userId"

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				next.ServeHTTP(w, r)
				return
			}
			bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
			err := token.Validate(bearerToken)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			userId, err := token.ParseToken(bearerToken)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, userId)
			r = r.WithContext(ctx)

			// executing next
			next.ServeHTTP(w, r)

			// after executing next
		})
	}
}

// UserId finds the user from the context. REQUIRES Middleware to have run.
func UserId(ctx context.Context) *string {
	userId, _ := ctx.Value(userIDKey).(*string)
	return userId
}
