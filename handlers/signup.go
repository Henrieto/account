package handlers

import (
	"net/http"

	"github.com/henrieto/jax"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	jax.Json(w, " signup route", 200)
}
