// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddPermissionToGroup(ctx context.Context, arg AddPermissionToGroupParams) (GroupPermission, error)
	AddPermissionToUser(ctx context.Context, arg AddPermissionToUserParams) (UserPermission, error)
	AddUserToGroup(ctx context.Context, arg AddUserToGroupParams) error
	CountGroupPermissions(ctx context.Context, groupID pgtype.Int4) (int64, error)
	CountGroups(ctx context.Context) (int64, error)
	CountPermissions(ctx context.Context) (int64, error)
	CountUserPermissions(ctx context.Context, userID pgtype.Int4) (int64, error)
	CountUsers(ctx context.Context) (int64, error)
	CountVerifyIdentityDatas(ctx context.Context) (int64, error)
	CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error)
	CreateStaff(ctx context.Context, arg CreateStaffParams) (User, error)
	CreateSuperUser(ctx context.Context, arg CreateSuperUserParams) (User, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserGroup(ctx context.Context, arg CreateUserGroupParams) (Group, error)
	CreateVerifyIdentityData(ctx context.Context, arg CreateVerifyIdentityDataParams) (VerifyIdentityData, error)
	DeleteGroup(ctx context.Context, id int32) error
	DeleteGroupPermission(ctx context.Context, arg DeleteGroupPermissionParams) error
	DeletePermission(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	DeleteUserPermission(ctx context.Context, arg DeleteUserPermissionParams) error
	DeleteVerifyIdentiyData(ctx context.Context, id int32) error
	FilterUsers(ctx context.Context, arg FilterUsersParams) ([]User, error)
	GetAllGroupPermissions(ctx context.Context, groupID pgtype.Int4) ([]GetAllGroupPermissionsRow, error)
	GetAllGroups(ctx context.Context, dollar_1 interface{}) ([]Group, error)
	GetAllPermssions(ctx context.Context, arg GetAllPermssionsParams) ([]Permission, error)
	GetAllUserPermissions(ctx context.Context, userID pgtype.Int4) ([]GetAllUserPermissionsRow, error)
	GetAllUsers(ctx context.Context, dollar_1 interface{}) ([]User, error)
	GetGroup(ctx context.Context, name string) (Group, error)
	GetGroupPermissions(ctx context.Context, arg GetGroupPermissionsParams) ([]GetGroupPermissionsRow, error)
	GetPermission(ctx context.Context, id int32) (Permission, error)
	GetPermissionByName(ctx context.Context, name string) (Permission, error)
	GetUser(ctx context.Context, id int32) (User, error)
	GetUserByAuthId(ctx context.Context, authID pgtype.Text) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByPhone(ctx context.Context, phone pgtype.Text) (User, error)
	GetUserPermissions(ctx context.Context, arg GetUserPermissionsParams) ([]GetUserPermissionsRow, error)
	GetVerifyIdentityData(ctx context.Context, randomString string) (VerifyIdentityData, error)
	PaginateGroups(ctx context.Context, arg PaginateGroupsParams) ([]Group, error)
	PaginateUsers(ctx context.Context, arg PaginateUsersParams) ([]User, error)
	PaginateVerifyIdentityData(ctx context.Context, arg PaginateVerifyIdentityDataParams) ([]VerifyIdentityData, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (Group, error)
	UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
