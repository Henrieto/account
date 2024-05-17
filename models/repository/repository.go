package repository

import (
	"context"

	"github.com/henrieto/account/models"
	"github.com/henrieto/account/models/database/db"
)

type IUser interface {
	// (context , user data) (db user data , error)
	Create(context.Context, *models.User) (*models.User, error)
	// (context) (number of users , error)
	Count(context.Context) (int64, error)
	// (context , user data) (db user data , error)
	CreateStaff(context.Context, *models.User) (*models.User, error)
	// (context , user data) (db user data , error)
	CreateSuperUser(context.Context, *models.User) (*models.User, error)
	// (context , user data) (db user data , error)
	Update(context.Context, *models.User) (*models.User, error)
	// (context , order by) (db users data , error)
	List(context.Context, string) ([]*models.User, error)
	// (context , order by  , offset  , limit) (db user data , error)
	Paginate(context.Context, string, uint, uint) ([]*models.User, error)
	// (context , user id) (db user data , error)
	Get(context.Context, int32) (*models.User, error)
	// (context , user email) (db user data , error)
	GetByEmail(context.Context, string) (*models.User, error)
	// (context , user email) (db user data , error)
	GetByPhone(context.Context, string) (*models.User, error)
	// (context , user email) (db user data , error)
	GetByAuthId(context.Context, string) (*models.User, error)
	// (context , field , value , order by , offset  , limit) (db users data , error)
	Filter(context.Context, string, any, string, uint, uint) ([]*models.User, error)
	// (context , user id) (error)
	Delete(context.Context, int32) error
	// (context , user email) (if user exists)
	EmailExists(context.Context, string) bool
	// (context , user phone number) (if user exists)
	PhoneExists(context.Context, string) bool
}

var User IUser

type IGroup interface {
	// (context , group data) (db group data , error)
	Create(context.Context, *db.Group) (*db.Group, error)
	// (context) (number of groups , error)
	Count(context.Context) (int64, error)
	// (context , order by)(db user groups , error)
	List(context.Context, string) ([]db.Group, error)
	// (context , order by , offset  , limit) (db groups , error)
	Paginate(context.Context, string, uint, uint) ([]db.Group, error)
	// (context , group id) (db group , error)
	Get(context.Context, string) (*db.Group, error)
	// (context , group data , group id to update) (db group data , error)
	Update(context.Context, *db.Group, int32) (*db.Group, error)
	// (context , group id) (error)
	Delete(context.Context, int32) error
	// (context , group id , user id ) (error)
	AddToGroup(context.Context, int32, int32) error
}

var Group IGroup

type IPermission interface {
	// (context , permission data) (db permission data , error)
	Create(context.Context, *db.Permission) (*db.Permission, error)
	// (context) (number of permissions , error)
	Count(context.Context) (int64, error)
	// (context , order by , offset  , limit) (db permissions , error)
	Paginate(context.Context, string, uint, uint) ([]db.Permission, error)
	// (context , group id) (db group , error)
	Get(context.Context, int32) (*db.Permission, error)
	// (context , permission data , the permission id to update) (db permission data , error)
	Update(context.Context, *db.Permission, int32) (*db.Permission, error)
	// (context , group id) (error)
	Delete(context.Context, int32) error
	// (context , group id , permission id) (error)
	AddPermissionToGroup(context.Context, int32, int32) (*db.GroupPermission, error)
	// (context) (group permissions count  , group id, error)
	CountGroupPermissions(context.Context, int32) (int64, error)
	// (context , group id , permission id) (error)
	RemovePermissionFromGroup(context.Context, int32, int32) error
	// (context , group id  ) (db group permissions , error)
	GetAllGroupPermissions(context.Context, int32) ([]db.GetAllGroupPermissionsRow, error)
	// (context , group id , offset , limit ) (db group permissions , error)
	GetGroupPermissions(context.Context, int32, uint, uint) ([]db.GetGroupPermissionsRow, error)
	// (context , user id , permission id) (error)
	AddPermissionToUser(context.Context, int32, int32) (*db.UserPermission, error)
	// (context) (user permission count , user id , error)
	CountUserPermissions(context.Context, int32) (int64, error)
	// (context , user id , permission id) (error)
	RemovePermissionFromUser(context.Context, int32, int32) error
	// (context , user id ) (db user permissions , error)
	GetAllUserPermissions(context.Context, int32) ([]db.GetAllUserPermissionsRow, error)
	// (context , user id , offset , limit) (db user permissions , error)
	GetUserPermissions(context.Context, int32, uint, uint) ([]db.GetUserPermissionsRow, error)
}

var Permission IPermission

type IVerifyIdentityData interface {
	// (context , identity verification data) (db identity verification data , error)
	Create(context.Context, *db.VerifyIdentityData) (*db.VerifyIdentityData, error)
	// (context) (number of identity verification datas , error)
	Count(context.Context) (int64, error)
	// (context , random string) (db identification data , error)
	Get(context.Context, string) (*db.VerifyIdentityData, error)
	// (context , verification data id) (error)
	Delete(context.Context, int32) error
	// (context , order by , offset , limit)
	List(context.Context, string, uint, uint) ([]db.VerifyIdentityData, error)
}

var VerifyIdentityData IVerifyIdentityData
