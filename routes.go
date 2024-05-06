package account

import (
	"net/http"

	"github.com/henrieto/jax"
	"github.com/henrieto/plugins/handlers"
)

var Routes = []jax.Route{
	{
		Path:    "/signup",
		Handler: handlers.Signup,
		Method:  http.MethodGet,
		Name:    "signup",
	},
}
