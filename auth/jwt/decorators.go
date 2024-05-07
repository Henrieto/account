package jwt_auth

import (
	"net/http"
	"strings"

	"github.com/henrieto/jax"
)

func Protected(secret string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the authorization header
		AuthHeader := r.Header.Get("Authorization")
		// split the Authorization header , to get the token
		BearerToken := strings.Split(AuthHeader, "")
		// get the token from the authorization header
		if len(BearerToken) != 2 {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		token := BearerToken[1]
		// check if the token is valid
		claims := NewClaims()
		err := claims.ValidateJwtToken(secret, token)
		if err != nil {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		// get the jwt claims
		err = claims.GetJwtTokenClaims(secret, token)
		if err != nil {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		// get user
		user, err := GetUser(claims.Object)
		if err != nil {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		r = AddUserToRequest(r, user)
		handler(w, r)
	}
}
