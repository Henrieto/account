package handlers

import (
	"net/http"

	"github.com/henrieto/account/auth"
	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/utils"
	"github.com/henrieto/account/validators"
	"github.com/henrieto/jax"
	"github.com/jackc/pgx/v5/pgtype"
)

func ForgortPassword(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator for the request data
	data_validator := validators.NewForgotPasswordDataValidator()
	// bind the request data to the validator
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
	// check if the data is valid
	email, err := data_validator.Valid()
	// if there was an error return a failed response
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "data is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	user, err := repository.UserRepository.Filter(Background, "email", email)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "email doesn't exists",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	resetToken, validityId, err := auth.GetPasswordResetToken(email)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "password reset failed",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	user.AuthID = pgtype.Text{String: validityId}
	_, err = repository.UserRepository.Update(Background, user)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "password reset failed",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}

	response := map[string]any{
		"status": "success",
		"msg":    "move to change password",
		"data": map[string]any{
			"token": resetToken,
		},
	}
	jax.Json(w, response, http.StatusOK)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	// rereive the password reset token
	resetToken := r.PathValue("token")
	// if the token is empty , return a failed respone
	if resetToken == "" {
		response := map[string]any{
			"status": "failed",
			"msg":    "token can't empty",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	//  deconstruct the token into a map object
	data, err := auth.PasswordResetToken(resetToken)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "token is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	// get user email from token
	_email, ok := data["email"]
	if !ok {
		response := map[string]any{
			"status": "failed",
			"msg":    "email is empty",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	var email string
	switch __email := _email.(type) {
	case string:
		email = __email
	default:
		response := map[string]any{
			"status": "failed",
			"msg":    "email is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	// get validity string
	_validity, ok := data["validity"]
	if !ok {
		response := map[string]any{
			"status": "failed",
			"msg":    "validity is empty",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	// retrieve the user with the email
	user, err := repository.UserRepository.Filter(Background, "email", email)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "password change failed (email is invlaid)",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}

	var validity string
	switch __validity := _validity.(type) {
	case string:
		validity = __validity
	default:
		response := map[string]any{
			"status": "failed",
			"msg":    "validity string is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	valid := auth.ComparePassword([]byte(user.AuthID.String), []byte(validity))
	if !valid {
		response := map[string]any{
			"status": "failed",
			"msg":    "validity string is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	// initialize a data validator
	data_validator := validators.NewPasswordChangeData()
	// bind the request data to the validator
	err = utils.BindJson(r, data_validator)
	// if an error occured return a failed response
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "data is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	// check if the data is valid
	password, err := data_validator.Valid()
	// if an error occured return a failed response
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "data is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	hashed_password, err := auth.HashPassword(password)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "password change failed",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	user.PasswordHash = string(hashed_password)
	_, err = repository.UserRepository.Update(Background, user)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "password change failed",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	response := map[string]any{
		"status": "success",
		"msg":    "password changed successfully ",
		"data":   nil,
	}
	jax.Json(w, response, http.StatusOK)
}
