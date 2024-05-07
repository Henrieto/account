package middlewares

import (
	"net/http"

	jwt_auth "github.com/henrieto/account/auth/jwt"
	"github.com/henrieto/jax"
)

func HasPermission(handler http.HandlerFunc, permissions ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := jwt_auth.GetUserFromRequest(r)
		if err != nil {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		if user.HasPermissions(permissions...) {
			handler(w, r)
		}
	}
}
