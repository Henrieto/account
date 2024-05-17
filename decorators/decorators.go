package decorators

import (
	"net/http"

	"github.com/henrieto/account/auth"
	"github.com/henrieto/jax"
)

func HasPermission(handler http.HandlerFunc, permissions ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.GetUserFromRequest(r)
		if err != nil {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		if user.HasPermissions(permissions...) {
			handler(w, r)
		}
	}
}
