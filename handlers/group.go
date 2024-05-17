package handlers

import (
	"net/http"
	"strconv"

	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/utils"
	db_utils "github.com/henrieto/account/utils/db"
	"github.com/henrieto/account/validators"
)

// CreateGroup an endpoint for creating user group
//
//	@Summary      create a user group
//	@Description  create group
//	@Tags         create user group
//	@Accept       json
//	@Produce      json
//
// @Param        group_data   body      swagger.GroupData  true  "Group Name"
// @Success		200		{object}	swagger.GroupSuccess			"ok"
// @Failure		404		{object}	swagger.GroupFailed	     "failed response"
// @Router       /account/group/create [post]
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewGroupData()
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
		response.Send(w, http.StatusBadRequest)
		return
	}
	//  check if the data is valid
	group, err := data_validator.Valid()
	// if the data is not valid , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// add the group data to the database
	group, err = repository.Group.Create(Background, group)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not add group data to the database"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = " group created successfully"
	//  set the data for the response
	response.Data = group
	// send the response
	response.Send(w, http.StatusOK)
}

// UpdateGroup an endpoint for updating user group
//
//	@Summary      update user group
//	@Description  update group
//	@Tags         update user group
//	@Accept       json
//	@Produce      json
//
// @Param        name   body      swagger.GroupData true  "Group Name"
// @Success		200		{object}	swagger.GroupSuccess			"ok"
// @Failure		404		{object}	swagger.GroupFailed	     "failed response"
// @Router       /account/group/update [put]
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	// retrieve the group id from the url
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
	// initialize a data validator
	data_validator := validators.NewGroupData()
	// bind the request data to the validator
	err = utils.BindJson(r, data_validator)
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
	group, err := data_validator.Valid()
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
	// add the group data to the database
	group, err = repository.Group.Update(Background, group, int32(id))
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could update group data "
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
	response.Data = group
	// send the response
	response.Send(w, http.StatusOK)
}

// ListGroups an endpoint for creating user group
//
//	@Summary      list user groups
//	@Description  list groups
//	@Tags         list user groups
//	@Produce      json
//
// @Success		200		{object}	swagger.GroupArraySuccess		"ok"
// @Failure		404		{object}	swagger.GroupFailed	     "failed response"
// @Router       /account/group/list [get]
func ListGroups(w http.ResponseWriter, r *http.Request) {
	// fetch the user groups
	groups, err := repository.Group.List(Background, "name")
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could fetch the user groups "
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
	response.Data = groups
	// send the response
	response.Send(w, http.StatusOK)
}

// PaginateGroups an endpoint for creating user group
//
//	@Summary      list user groups
//	@Description  list groups
//	@Tags         list user groups
//	@Produce      json
//
// @Param			offset	path		int				true	"pagination parameter"
// @Param			limit	path		int				true	"pagination parameter"
// @Success		200		{object}	swagger.GroupArraySuccess		"ok"
// @Failure		404		{object}	swagger.GroupFailed	     "failed response"
// @Router       /account/group/paginate?offset={offset}&limit={limit} [get]
func PaginateGroups(w http.ResponseWriter, r *http.Request) {
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
	// fetch the user groups
	groups, err := repository.Group.Paginate(Background, "name", uint(paginator.Page), uint(paginator.PageSize))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could fetch the user groups "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// get permissions count
	group_count, err := repository.Group.Count(Background)
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
	paginator.Total = int(group_count)
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = " group updated successfully"
	//  set the data for the response
	response.Data = utils.PaginatorResponse{
		Data:      groups,
		Paginator: paginator,
	}
	// send the response
	response.Send(w, http.StatusOK)
}

// GetGroup an endpoint for retrieving a user group
//
//	@Summary      get a user group with id
//	@Description  retrieve a user group using the group id
//	@Tags         get user group
//	@Produce      json
//
// @Param			name	path		string			true	"Some ID"
// @Success		200		{object}	swagger.GroupSuccess		"ok"
// @Failure		404		{object}	swagger.GroupFailed	     "failed response"
// @Router       /account/group/get/{name} [get]
func GetGroup(w http.ResponseWriter, r *http.Request) {
	// retrieve the group name from the url
	name := r.PathValue("name")
	// fetch the user group , using the name
	group, err := repository.Group.Get(Background, name)
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not fetch user group with the id provided "
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
	response.Msg = " group retrieved successfully"
	//  set the data for the response
	response.Data = group
	// send the response
	response.Send(w, http.StatusOK)
}

// DeleteGroup an endpoint for retrieving a user group
//
//	@Summary      delete a user group using the group  id
//	@Description  delete a user group using the group  id
//	@Tags         delete user group
//	@Produce      json
//
// @Param			id	path		string			true	"Some ID"
// @Success		200		{object}	swagger.GroupDeleteSuccess		"ok"
// @Failure		404		{object}	swagger.GroupFailed	     "failed response"
// @Router       /account/group/delete/{id} [delete]
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	// retrieve the group id from the url
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
	// delete the group data from the database
	err = repository.Group.Delete(Background, int32(id))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not delete group with the id provided "
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
	response.Msg = " group deleted successfully"
	// send the response
	response.Send(w, http.StatusOK)
}

// AddUserToGroup an endpoint for updating user group
//
//	@Summary      add a user to a group
//	@Description  add a user to a group
//	@Tags         add user to a group
//	@Accept       json
//	@Produce      json
//
// @Param        group_id   body      int32 true  "user id"
// @Success		200		{object}	swagger.GroupSuccess			"ok"
// @Failure		404		{object}	swagger.GroupFailed	     "failed response"
// @Router       /account/group/user/add [post]
func AddUserToGroup(w http.ResponseWriter, r *http.Request) {
	// initalize a data validator
	data_validator := validators.NewUserGroupData()
	// bind request data to the validator
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
	// check if the data is valid
	user, group, err := data_validator.Valid()
	// if the data is  not valid
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
	// add user to group
	err = repository.Group.AddToGroup(Background, group, user)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not add user to group"
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
	response.Msg = " successfully added user to group "
	// send the response
	response.Send(w, http.StatusOK)
}
