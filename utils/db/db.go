package db_utils

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/henrieto/account/config"
	"github.com/henrieto/account/models"
	"github.com/henrieto/account/models/database/db"
)

// a paginator object
type Paginator struct {
	Next     int
	Previous int
	PageSize int
	Page     int
	Total    int
}

// constructor funtion to initialize a new paginator
func NewPaginator(r *http.Request, page_size ...uint) (*Paginator, error) {
	// get the offset and limit
	offset, limit, err := GetOffsetAndLimit(r, page_size...)
	// if an error occured return the error
	if err != nil {
		return nil, err
	}
	//  set the next and prevoius
	var next, previous int
	next = int(offset + limit)
	previous = int(offset - limit)
	// set the previous
	if uint(previous) == 0 {
		previous = 1
	}
	return &Paginator{
		Next:     int(next),
		Previous: previous,
		PageSize: int(limit),
		Page:     int(offset),
	}, nil
}

// function to retrive the offset and limit from the url query
func GetOffsetAndLimit(r *http.Request, page_size ...uint) (int32, int32, error) {
	// check if the url query contains the offset and limit
	// if it doesn't , set the default values
	offset, limit := default_offset_and_limit(r.URL.Query(), page_size...)
	// if the offset or limit exists ie (offset = 0 or limit = 0)
	// set there values

	if offset == 0 {
		// get the offset as string
		offset_string := r.URL.Query().Get("offset")
		// convert the string offset to int
		offset_int, err := strconv.Atoi(offset_string)
		// if an error occured  , return the error
		if err != nil {
			return 0, 0, err
		}
		// set the int offset to int32 offset
		offset = int32(offset_int)
	}
	if limit == 0 {
		// get the limit as string
		limit_string := r.URL.Query().Get("limit")
		// convert the string limit to int
		limit_int, err := strconv.Atoi(limit_string)
		// if an error occured  , return the error
		if err != nil {
			return 0, 0, err
		}
		// set the int limit to int32 limit
		limit = int32(limit_int)
	}
	// return the offset and limit with the error as nil
	return offset, limit, nil
}

func default_offset_and_limit(query url.Values, page_size ...uint) (int32, int32) {
	// initialize the offset and limit variables
	var offset, limit int32
	// check if the offset exists in the url query
	has_offset := query.Has("offset")
	// check if the limit exists in the url query
	has_limit := query.Has("limit")
	// if the offset doesn't exists , set the default offset
	if !has_offset {
		offset = 1
	}
	// if the limit doesn't exists , set the default limit
	if !has_limit {
		if len(page_size) != 0 {
			limit = int32(page_size[0])
		} else {
			limit = int32(config.PaginatorPageSize)
		}
	}
	return offset, limit
}

// a model to db user translator
func DbUserToModelUser(_user db.User) models.User {
	return models.User{
		ID:           _user.ID,
		Username:     _user.Username,
		FirstName:    _user.FirstName,
		LastName:     _user.LastName,
		Email:        _user.Email,
		Gender:       _user.Gender,
		Birthday:     _user.Birthday,
		PasswordHash: _user.PasswordHash,
		Verified:     _user.Verified,
		AuthID:       _user.AuthID,
		Staff:        _user.Staff,
		Superuser:    _user.Superuser,
		CreatedAt:    _user.CreatedAt,
		UpdatedAt:    _user.UpdatedAt,
		GroupID:      _user.GroupID,
	}
}

// GetUserPermissionsRow to permission
// GetAllUserPermissionsRow to permission
// GetGroupPermissionsRow to permission
// GetAllGroupPermissionsRow to permission

func TranslatePermission(perm any) []*db.Permission {
	switch permissions := perm.(type) {
	case []db.GetUserPermissionsRow:
		return GetUserPermissionRowTranslator(permissions)
	case []db.GetAllUserPermissionsRow:
		return GetAllUserPermissionRowTranslator(permissions)
	case []db.GetGroupPermissionsRow:
		return GetGroupPermissionRowTranslator(permissions)
	case []db.GetAllGroupPermissionsRow:
		return GetAllGroupPermissionRowTranslator(permissions)
	default:
		return []*db.Permission{}
	}
}

// GetUserPermissionsRow to permission
func GetUserPermissionRowTranslator(permissions []db.GetUserPermissionsRow) []*db.Permission {
	permList := []*db.Permission{}
	for _, pem := range permissions {
		new_perm := db.Permission{
			ID:        pem.PermissionID.Int32,
			Model:     pem.Model,
			Name:      pem.Name,
			Codename:  pem.Codename,
			CreatedAt: pem.CreatedAt,
			UpdatedAt: pem.UpdatedAt,
		}
		permList = append(permList, &new_perm)
	}
	return permList
}

// GetAllUserPermissionsRow to permission
func GetAllUserPermissionRowTranslator(permissions []db.GetAllUserPermissionsRow) []*db.Permission {
	permList := []*db.Permission{}
	for _, pem := range permissions {
		new_perm := db.Permission{
			ID:        pem.PermissionID.Int32,
			Model:     pem.Model,
			Name:      pem.Name,
			Codename:  pem.Codename,
			CreatedAt: pem.CreatedAt,
			UpdatedAt: pem.UpdatedAt,
		}
		permList = append(permList, &new_perm)
	}
	return permList
}

// GetGroupPermissionsRow to permission
func GetGroupPermissionRowTranslator(permissions []db.GetGroupPermissionsRow) []*db.Permission {
	permList := []*db.Permission{}
	for _, pem := range permissions {
		new_perm := db.Permission{
			ID:        pem.PermissionID.Int32,
			Model:     pem.Model,
			Name:      pem.Name,
			Codename:  pem.Codename,
			CreatedAt: pem.CreatedAt,
			UpdatedAt: pem.UpdatedAt,
		}
		permList = append(permList, &new_perm)
	}
	return permList
}

// GetAllGroupPermissionsRow to permission
func GetAllGroupPermissionRowTranslator(permissions []db.GetAllGroupPermissionsRow) []*db.Permission {
	permList := []*db.Permission{}
	for _, pem := range permissions {
		new_perm := db.Permission{
			ID:        pem.PermissionID.Int32,
			Model:     pem.Model,
			Name:      pem.Name,
			Codename:  pem.Codename,
			CreatedAt: pem.CreatedAt,
			UpdatedAt: pem.UpdatedAt,
		}
		permList = append(permList, &new_perm)
	}
	return permList
}
