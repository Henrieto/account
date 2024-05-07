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
	CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error)
	CreateStaff(ctx context.Context, arg CreateStaffParams) (User, error)
	CreateSuperUser(ctx context.Context, arg CreateSuperUserParams) (User, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserGroup(ctx context.Context, arg CreateUserGroupParams) ([]Group, error)
	DeleteGroup(ctx context.Context, id pgtype.UUID) error
	DeleteGroupPermission(ctx context.Context, id pgtype.UUID) error
	DeletePermission(ctx context.Context, id pgtype.UUID) error
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	DeleteUserPermission(ctx context.Context, id pgtype.UUID) error
	GetAllGroups(ctx context.Context, dollar_1 interface{}) ([]Group, error)
	GetAllPermssions(ctx context.Context, arg GetAllPermssionsParams) ([]Permission, error)
	GetAllUsers(ctx context.Context, dollar_1 interface{}) ([]User, error)
	GetGroup(ctx context.Context, name string) (Group, error)
	GetGroupPermissions(ctx context.Context, arg GetGroupPermissionsParams) ([]GetGroupPermissionsRow, error)
	GetPermissionByName(ctx context.Context, name string) (Permission, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserPermissions(ctx context.Context, arg GetUserPermissionsParams) ([]GetUserPermissionsRow, error)
	MakeUserVerified(ctx context.Context, arg MakeUserVerifiedParams) (User, error)
	PaginateGroups(ctx context.Context, arg PaginateGroupsParams) ([]Group, error)
	PaginateUsers(ctx context.Context, arg PaginateUsersParams) ([]User, error)
	SetUserAuthId(ctx context.Context, arg SetUserAuthIdParams) (User, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (Group, error)
	UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (User, error)
}

var _ Querier = (*Queries)(nil)