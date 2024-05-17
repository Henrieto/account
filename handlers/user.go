package handlers

import (
	"net/http"
	"strconv"

	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/utils"
	db_utils "github.com/henrieto/account/utils/db"
	"github.com/henrieto/account/validators"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewUserData()
	// bind the request data to the data validator
	err := utils.BindJson(r, data_validator)
	// if an error occured , return a failed response
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
	user, err := data_validator.Valid()
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
	// add the user data to the database
	user, err = repository.User.Create(Background, user)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not add user data to the database"
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
	response.Msg = " user created successfully"
	//  set the data for the response
	response.Data = user
	// send the response
	response.Send(w, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewUserData()
	// bind the request data to the data validator
	err := utils.BindJson(r, data_validator)
	// if an error occured , return a failed response
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
	user, err := data_validator.Valid()
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
	// update the user data
	user, err = repository.User.Update(Background, user)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not update user data "
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
	response.Msg = " user data updated successfully"
	//  set the data for the response
	response.Data = user
	// send the response
	response.Send(w, http.StatusOK)
}

func PaginateUsers(w http.ResponseWriter, r *http.Request) {
	// initialize a paginator object
	paginator, err := db_utils.NewPaginator(r)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " limit and offset are not valid"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// fetch users
	users, err := repository.User.Paginate(Background, "username", uint(paginator.Page), uint(paginator.PageSize))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could fetch the users "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// get permissions count
	users_count, err := repository.User.Count(Background)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " database failure"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// set the paginator total
	paginator.Total = int(users_count)
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = " group updated successfully"
	//  set the data for the response
	response.Data = utils.PaginatorResponse{
		Data:      users,
		Paginator: paginator,
	}
	// send the response
	response.Send(w, http.StatusOK)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	// retrieve users from the database
	users, err := repository.User.List(Background, "username")
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not retrieve users data"
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
	response.Msg = " group updated successfully"
	//  set the data for the response
	response.Data = users
	// send the response
	response.Send(w, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	// retrieve the permission id from the url
	// convert it to int
	id, err := strconv.Atoi(r.PathValue("id"))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " the id provided is not valid "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// fetch the user data
	user, err := repository.User.Get(Background, int32(id))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not fetch user with the id provided "
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
	response.Msg = " permission retrieved successfully"
	//  set the data for the response
	response.Data = user
	// send the response
	response.Send(w, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// retrieve the permission id from the url
	// convert it to int
	id, err := strconv.Atoi(r.PathValue("id"))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " the id provided is not valid "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// delete the user data from the database
	err = repository.User.Delete(Background, int32(id))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not delete user with the id provided "
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
	response.Msg = " user deleted successfully"
	// send the response
	response.Send(w, http.StatusOK)
}
