package handlers

import (
	"net/http"
	"strconv"

	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/utils"
	db_utils "github.com/henrieto/account/utils/db"
	"github.com/henrieto/account/validators"
)

// CreatePermission an endpoint for creating user permissions
//
//	@Summary      create a user permission
//	@Description  create permission
//	@Tags         create user permission
//	@Accept       json
//	@Produce      json
//
// @Param        permission_data   body      swagger.PermissionData  true  "Permission Data"
// @Success		200		{object}	swagger.PermissionSuccess			"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/create [post]
func CreatePermission(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewPermissionData()
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
	permission, err := data_validator.Valid()
	// if the data is invalid , return a failed response
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
	// add the permission data to the database
	permission, err = repository.Permission.Create(Background, permission)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not add permission data to the database"
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
	response.Msg = " permission created successfully"
	//  set the data for the response
	response.Data = permission
	// send the response
	response.Send(w, http.StatusOK)
}

// UpdatePermission an endpoint for updating user permissions
//
//	@Summary      update user permission
//	@Description  update permission
//	@Tags         update user permission
//	@Accept       json
//	@Produce      json
//
// // @Param			id	path		int32			true	"permission ID"
// @Param        permission_data   body      swagger.PermissionData  true  "Permission Data"
// @Success		200		{object}	swagger.PermissionSuccess			"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/update/{id} [put]
func UpdatePermission(w http.ResponseWriter, r *http.Request) {
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
	// initialize a data validator
	data_validator := validators.NewPermissionData()
	// bind request data to the validator
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
	// check if the data is valid
	permission, err := data_validator.Valid()
	// if the data is invalid , return a failed response
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
	// add the permission data to the database
	permission, err = repository.Permission.Update(Background, permission, int32(id))
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not update permission data "
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
	response.Msg = " permission updated successfully"
	//  set the data for the response
	response.Data = permission
	// send the response
	response.Send(w, http.StatusOK)
}

// ListPermission an endpoint for retrieving user permissions
//
//	@Summary      list user permission
//	@Description  list permission
//	@Tags         list user permission
//	@Accept       json
//	@Produce      json
//
// @Success		200		{object}	swagger.PermissionArraySuccess		"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/list [post]
func ListPermission(w http.ResponseWriter, r *http.Request) {
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
	// retrieve the permissions
	permissions, err := repository.Permission.Paginate(Background, "name", uint(paginator.Page), uint(paginator.PageSize))
	// if an error occured  , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not retrieve permissions"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// get permissions count
	permissions_count, err := repository.Permission.Count(Background)
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
	paginator.Total = int(permissions_count)

	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = " permission retrieved successfully"
	//  set the data for the response
	response.Data = utils.PaginatorResponse{
		Data:      permissions,
		Paginator: paginator,
	}
	// send the response
	response.Send(w, http.StatusOK)
}

// GetPermission an endpoint for retrieving a particular permission
//
//	@Summary      get permission with id
//	@Description  retrieve a permission using the permission id
//	@Tags         get permission
//	@Produce      json
//
// @Param			id	path		int32			true	"permission ID"
// @Success		200		{object}	swagger.PermissionSuccess		"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/get/{id} [get]
func GetPermission(w http.ResponseWriter, r *http.Request) {
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
	// fetch the permission
	permsission, err := repository.Permission.Get(Background, int32(id))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not fetch permission with the id provided "
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
	response.Data = permsission
	// send the response
	response.Send(w, http.StatusOK)
}

// DeletePermission an endpoint for deleting a particular permission
//
//	@Summary      delete permission with id
//	@Description  delete a permission using the permission id
//	@Tags         delete permission
//	@Produce      json
//
// @Param			id	path		int32			true	"permission ID"
// @Success		200		{object}	swagger.PermissionSuccess		"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/delete/{id} [get]
func DeletePermission(w http.ResponseWriter, r *http.Request) {
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
	// delete the permission data from the database
	err = repository.Permission.Delete(Background, int32(id))
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not delete permission with the id provided "
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
	response.Msg = " permission deleted successfully"
	// send the response
	response.Send(w, http.StatusOK)
}

// AddPermissionsToGroup an endpoint for adding permission to a user group
//
//	@Summary      add permission to a user group
//	@Description  add permission to a user group
//	@Tags         add permission to a user group
//	@Produce      json
//
// @Param        add_permission_to_group_data   body      swagger.PermissionsGroupData  true  "Add Permission To Group Data"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/group/add [post]
func AddPermissionsToGroup(w http.ResponseWriter, r *http.Request) {
	// initalize a data validator
	data_validator := validators.NewGroupPermissionData()
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
	permission, group, err := data_validator.Valid()
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
	// add permission to group
	_, err = repository.Permission.AddPermissionToGroup(Background, group, permission)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not add permission to group"
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
	response.Msg = " successfully added permission to group "
	// send the response
	response.Send(w, http.StatusOK)
}

// RemoveGroupPermisions an endpoint for deleting a particular group permission
//
//	@Summary      delete permission with id
//	@Description  delete a permission using the permission id
//	@Tags         delete permission
//	@Produce      json
//
// @Param        remove_permission_from_group_data   body      swagger.PermissionsGroupData  true  "Remove Permission from Group Data"
// @Success		200		{object}	swagger.GetPermissionDeleteSuccess		"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/group/remove [post]
func RemoveGroupPermisions(w http.ResponseWriter, r *http.Request) {
	// initalize a data validator
	data_validator := validators.NewGroupPermissionData()
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
	permission, group, err := data_validator.Valid()
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
	// remove permission from group
	err = repository.Permission.RemovePermissionFromGroup(Background, group, permission)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not remove permission from group"
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
	response.Msg = " successfully removed permission from group "
	// send the response
	response.Send(w, http.StatusOK)
}

// GetGroupPermisions an endpoint for retrieving a particular group permission
//
//	@Summary      retrieving group permissions using the group  id
//	@Description  retrieving group permissions using the group  id
//	@Tags         retrieve group permissions
//	@Produce      json
//
// // @Param			id	path		int32			true	"Group ID"
// @Success		200		{object}	swagger.GetGroupPermissionSuccess		"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/group/{id} [get]
func GetGroupPermissions(w http.ResponseWriter, r *http.Request) {
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
	// retrieve group permissions
	permissions, err := repository.Permission.GetGroupPermissions(Background, int32(id), uint(paginator.Page), uint(paginator.PageSize))

	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not retrieve group permissions  "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// get group permissions count
	group_permissions_count, err := repository.Permission.CountGroupPermissions(Background, int32(id))
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
	paginator.Total = int(group_permissions_count)
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = " group permissions retrieved successfully"
	//  set the data for the response
	response.Data = utils.PaginatorResponse{
		Data:      permissions,
		Paginator: paginator,
	}
	// send the response
	response.Send(w, http.StatusOK)
}

// AddPermissionsToUser an endpoint for adding permission to a user
//
//	@Summary      add permission to a user
//	@Description  add permission to a user
//	@Tags         add permission to user
//	@Produce      json
//
// @Param        add_permission_to_user_data   body      swagger.PermissionsUserData  true  "Add Permission To User Data"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/user/add [post]
func AddPermissionsToUser(w http.ResponseWriter, r *http.Request) {
	// initalize a data validator
	data_validator := validators.NewUserPermissionData()
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
	permission, user, err := data_validator.Valid()
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
	// add permission to user
	_, err = repository.Permission.AddPermissionToUser(Background, user, permission)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not add permission to user"
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
	response.Msg = " successfully added permission to user "
	// send the response
	response.Send(w, http.StatusOK)
}

// GetUserPermisions an endpoint for retrieving a particular user permissions
//
//	@Summary      retrieving user permissions using the user  id
//	@Description  retrieving user permissions using the user  id
//	@Tags         retrieve user permissions
//	@Produce      json
//
// // @Param			id	path		int32			true	"User ID"
// @Success		200		{object}	swagger.GetUserPermissionSuccess		"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/user/{id} [get]
func GetUserPermissions(w http.ResponseWriter, r *http.Request) {
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
	// retrieve the user id from the url
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
	// retrieve user permissions
	permissions, err := repository.Permission.GetUserPermissions(Background, int32(id), uint(paginator.Page), uint(paginator.PageSize))

	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not retrieve user permissions  "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// get user permissions count
	user_permissions_count, err := repository.Permission.CountUserPermissions(Background, int32(id))
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
	paginator.Total = int(user_permissions_count)
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = " user permissions retrieved successfully"
	//  set the data for the response
	response.Data = utils.PaginatorResponse{
		Data:      permissions,
		Paginator: paginator,
	}
	// send the response
	response.Send(w, http.StatusOK)
}

// RemoveUserPermisions an endpoint for deleting a particular user permission
//
//	@Summary      remove a particular user permission
//	@Description  delete user permission using the user id and the permission id
//	@Tags         delete user permission
//	@Produce      json
//
// @Param        remove_permission_from_group_data   body      swagger.PermissionsUserData  true  "Remove Permission from Group Data"
// @Success		200		{object}	swagger.GetPermissionDeleteSuccess		"ok"
// @Failure		404		{object}	swagger.PermissionFailed	     "failed response"
// @Router       /account/permission/user/remove [post]
func RemoveUserPermisions(w http.ResponseWriter, r *http.Request) {
	// initalize a data validator
	data_validator := validators.NewUserPermissionData()
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
	permission, user, err := data_validator.Valid()
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
	// remove permission from user
	err = repository.Permission.RemovePermissionFromUser(Background, user, permission)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not remove permission from user"
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
	response.Msg = " successfully removed permission from user "
	// send the response
	response.Send(w, http.StatusOK)
}
