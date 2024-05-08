package handlers

import (
	"context"
	"net/http"

	"github.com/henrieto/account/auth"
	jwt_auth "github.com/henrieto/account/auth/jwt"
	"github.com/henrieto/account/config"
	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/utils"
	"github.com/henrieto/account/validators"
	"github.com/henrieto/jax"
	"github.com/jackc/pgx/v5/pgtype"
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
	// check if the data is valid
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

func Login(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewLoginData()
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
	email, password, err := data_validator.Valid()
	// if the user is not valid return a failed response
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "data is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	// fetch the user from the database
	user, err := repository.UserRepository.Filter(Background, "email", email)
	// if there was an error ,  return a failed response
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "user does not exists",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusOK)
		return
	}
	// compare passwords
	authenticated := auth.ComparePassword([]byte(user.PasswordHash), []byte(password))
	// if passwords doesn't match return a failed response
	if !authenticated {
		response := map[string]any{
			"status": "failed",
			"msg":    "your credentials is invalid",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusOK)
		return
	}
	// initialize a jwt claim
	claims := jwt_auth.NewClaims()
	// add the user object in to the claims
	claims.Object = user
	// generate the jwt token
	token, err := claims.GenerateJwtToken(config.SECRET)
	// if there was an error , return a failed response
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "login failed",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusOK)
		return
	}
	// return a successful response
	response := map[string]any{
		"status": "success",
		"msg":    "you are successfully logged in",
		"data": map[string]any{
			"token": token,
			"user":  user,
		},
	}
	jax.Json(w, response, http.StatusOK)
	return
}

func Profile(w http.ResponseWriter, r *http.Request) {}

func VerifyIdentity(w http.ResponseWriter, r *http.Request) {
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

	user.Verified = pgtype.Bool{Bool: true}
	user, err = repository.UserRepository.Update(Background, user)
	if err != nil {
		response := map[string]any{
			"status": "failed",
			"msg":    "verification failed",
			"data":   nil,
		}
		jax.Json(w, response, http.StatusBadRequest)
		return
	}
	response := map[string]any{
		"status": "success",
		"msg":    "user verified successfully",
		"data":   user,
	}
	jax.Json(w, response, http.StatusOK)
	return
}
