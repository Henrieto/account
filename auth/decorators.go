package auth

import (
	"net/http"
	"strings"

	jwt_auth "github.com/henrieto/account/auth/jwt"
	"github.com/henrieto/account/config"
	"github.com/henrieto/jax"
)

func Protected(handler http.HandlerFunc) http.HandlerFunc {
	// set the jwt auth secret
	secret := config.SECRET
	// return a http handler func
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
		// get the jwt token
		token := BearerToken[1]
		// check if the token is valid
		claims := jwt_auth.NewClaims()
		// validate the jwt token
		err := claims.ValidateJwtToken(secret, token)
		// if an error occured , return
		if err != nil {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		// get the jwt claims
		err = claims.GetJwtTokenClaims(secret, token)
		// if an error occured , return
		if err != nil {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		// get user
		user, err := GetUser(claims.Object)
		// if an error occured , return
		if err != nil {
			jax.Json(w, "user is not authenticated", http.StatusUnauthorized)
			return
		}
		// add the user to the request context
		r = AddUserToRequest(r, user)
		// call the handler
		handler(w, r)
	}
}
