package account

import (
	"net/http"

	"github.com/henrieto/account/handlers"
	"github.com/henrieto/jax"
)

var Routes = []jax.Route{
	{
		Path:    "/signup",
		Handler: handlers.Signup,
		Method:  http.MethodGet,
		Name:    "signup",
	},
}
