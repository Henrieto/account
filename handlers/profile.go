package handlers

import (
	"context"
	"net/http"

	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/utils"
	"github.com/henrieto/account/validators"
	"github.com/henrieto/jax"
)

var (
	Background = context.Background()
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewSignupData()
	// bind the request data to the validator object
	err := utils.BindJson(r, data_validator)
	// if there was an error , return a failed response
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "data is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	// check if the user is valid
	user, err := data_validator.Valid()
	// if the user is not valid return a failed response
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "data is invalid",
			"data":   err.(*validators.SignupValidationErrorData),
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	// check if the user email already exists in the database
	email_exists := repository.UserRepository.EmailExists(Background, user.Email)
	if email_exists {
		_err := data_validator.Error(nil)
		_err.AddError("Email", "email is already taken")
		response := map[string]any{
			"status": "failed",
			"msg":    "email is already taken",
			"data":   _err,
		}
		jax.Json(w, response, http.StatusResetContent)
		return
	}
	//create the user
	user, err = repository.UserRepository.Create(Background, user)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "could not create user",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusOK)
		return
	}
	// return back the user data
	response := map[string]any{
		"status": "success",
		"msg":    "user created successfully",
		"data":   user,
	}
	jax.Json(w, response, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {}

func Profile(w http.ResponseWriter, r *http.Request) {}

func VerifyIdentity(w http.ResponseWriter, r *http.Request) {}
