package handlers

import (
	"net/http"

	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/utils"
	"github.com/henrieto/account/validators"
	"github.com/jackc/pgx/v5/pgtype"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	// retrieve the user auth token
	auth_token := r.PathValue("token")
	// initialize a data validator
	data_validator := validators.NewPasswordChangeData()
	// bind the request data to the validator
	err := utils.BindJson(r, data_validator)
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	//  check if the data is valid
	password, err := data_validator.Valid()
	// if the data is not valid , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// retrieve the user using the auth token
	user, err := repository.User.GetByAuthId(Background, auth_token)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " password reset failed "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// update user password
	user.PasswordHash = password
	// reset user auth id
	user.AuthID = pgtype.Text{String: ""}
	// update the user database data
	_, err = repository.User.Update(Background, user)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " password reset failed  (colud not update user)"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = " password changed successfully"
	// send the response
	response.Send(w, http.StatusOK)
}
