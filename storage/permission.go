package storage

import (
	"context"
	"time"

	"github.com/henrieto/account/models/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type Permission struct {
	Querier db.Querier
}

func NewPermissionStorage(querier db.Querier) *Permission {
	return &Permission{Querier: querier}
}

// (context , permission data) (db permission data , error)
func (strg *Permission) Create(ctx context.Context, permission *db.Permission) (*db.Permission, error) {
	params := db.CreatePermissionParams{
		Model:     permission.Model,
		Name:      permission.Name,
		Codename:  permission.Codename,
		CreatedAt: permission.CreatedAt,
		UpdatedAt: permission.UpdatedAt,
	}
	db_permission, err := strg.Querier.CreatePermission(ctx, params)
	if err != nil {
		return nil, err
	}
	return &db_permission, nil
}

// (context) (number of permissions , error )
func (strg *Permission) Count(ctx context.Context) (int64, error) {
	count, err := strg.Querier.CountPermissions(ctx)
	if err != nil {
		return 0, err
	}
	return count, err
}

// (context , order by , offset  , limit) (db permissions , error)
func (strg *Permission) Paginate(ctx context.Context, order_by string, offset uint, limit uint) ([]db.Permission, error) {
	params := db.GetAllPermssionsParams{
		Column1: order_by,
		Offset:  int32(offset),
		Limit:   int32(limit),
	}
	permissions, err := strg.Querier.GetAllPermssions(ctx, params)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// (context , group id) (db group , error)
func (strg *Permission) Get(ctx context.Context, id int32) (*db.Permission, error) {
	permission, err := strg.Querier.GetPermission(ctx, id)
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// (context , permission data , the permission id to update) (db permission data , error)
func (strg *Permission) Update(ctx context.Context, permission *db.Permission, id int32) (*db.Permission, error) {
	params := db.UpdatePermissionParams{
		Model:     permission.Model,
		Name:      permission.Name,
		Codename:  permission.Codename,
		UpdatedAt: permission.UpdatedAt,
		ID:        id,
	}
	db_permission, err := strg.Querier.UpdatePermission(ctx, params)
	if err != nil {
		return nil, err
	}
	return &db_permission, nil
}

// (context , group id) (error)
func (strg *Permission) Delete(ctx context.Context, id int32) error {
	err := strg.Querier.DeletePermission(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// (context , group id , permission id) (error)
func (strg *Permission) AddPermissionToGroup(ctx context.Context, group_id int32, permission_id int32) (*db.GroupPermission, error) {
	params := db.AddPermissionToGroupParams{
		GroupID:      pgtype.Int4{Valid: true, Int32: group_id},
		PermissionID: pgtype.Int4{Valid: true, Int32: permission_id},
		CreatedAt:    pgtype.Timestamptz{Time: time.Now()},
		UpdatedAt:    pgtype.Timestamptz{Time: time.Now()},
	}
	group_permission, err := strg.Querier.AddPermissionToGroup(ctx, params)
	if err != nil {
		return nil, err
	}
	return &group_permission, nil
}

// (context , user id)(user permissions count , error)
func (strg *Permission) CountGroupPermissions(ctx context.Context, user_id int32) (int64, error) {
	count, err := strg.Querier.CountGroupPermissions(ctx, pgtype.Int4{Int32: user_id, Valid: true})
	if err != nil {
		return 0, err
	}
	return count, nil
}

// (context , group id , permission id) (error)
func (strg *Permission) RemovePermissionFromGroup(ctx context.Context, group_id int32, permission_id int32) error {
	params := db.DeleteGroupPermissionParams{
		GroupID:      pgtype.Int4{Valid: true, Int32: group_id},
		PermissionID: pgtype.Int4{Valid: true, Int32: permission_id},
	}
	err := strg.Querier.DeleteGroupPermission(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

// (context , group id ) (db group permissions , error)
func (strg *Permission) GetAllGroupPermissions(ctx context.Context, group_id int32) ([]db.GetAllGroupPermissionsRow, error) {
	permissions, err := strg.Querier.GetAllGroupPermissions(ctx, pgtype.Int4{Int32: group_id, Valid: true})
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// (context , group id  , offset , limit) (db group permissions , error)
func (strg *Permission) GetGroupPermissions(ctx context.Context, group_id int32, offset uint, limit uint) ([]db.GetGroupPermissionsRow, error) {
	params := db.GetGroupPermissionsParams{
		GroupID: pgtype.Int4{Int32: group_id, Valid: true},
		Offset:  int32(offset),
		Limit:   int32(limit),
	}
	permissions, err := strg.Querier.GetGroupPermissions(ctx, params)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// (context , user id , permission id) (error)
func (strg *Permission) AddPermissionToUser(ctx context.Context, user_id int32, permission_id int32) (*db.UserPermission, error) {
	params := db.AddPermissionToUserParams{
		UserID:       pgtype.Int4{Int32: user_id, Valid: true},
		PermissionID: pgtype.Int4{Int32: permission_id, Valid: true},
		CreatedAt:    pgtype.Timestamptz{Time: time.Now()},
		UpdatedAt:    pgtype.Timestamptz{Time: time.Now()},
	}
	user_permission, err := strg.Querier.AddPermissionToUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user_permission, nil
}

// (context , user id)(user permissions count , error)
func (strg *Permission) CountUserPermissions(ctx context.Context, user_id int32) (int64, error) {
	count, err := strg.Querier.CountUserPermissions(ctx, pgtype.Int4{Int32: user_id, Valid: true})
	if err != nil {
		return 0, err
	}
	return count, nil
}

// (context , user id , permission id) (error)
func (strg *Permission) RemovePermissionFromUser(ctx context.Context, user_id int32, permission_id int32) error {
	params := db.DeleteUserPermissionParams{
		UserID:       pgtype.Int4{Int32: user_id, Valid: true},
		PermissionID: pgtype.Int4{Int32: permission_id, Valid: true},
	}
	err := strg.Querier.DeleteUserPermission(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

// (context , user id , offset , limit) (db user permissions , error)
func (strg *Permission) GetUserPermissions(ctx context.Context, user_id int32, offset uint, limit uint) ([]db.GetUserPermissionsRow, error) {
	params := db.GetUserPermissionsParams{
		UserID: pgtype.Int4{Int32: user_id, Valid: true},
		Offset: int32(offset),
		Limit:  int32(limit),
	}
	user_permissions, err := strg.Querier.GetUserPermissions(ctx, params)
	if err != nil {
		return nil, err
	}
	return user_permissions, nil
}

// (context , group id ) (db user permissions , error)
func (strg *Permission) GetAllUserPermissions(ctx context.Context, user_id int32) ([]db.GetAllUserPermissionsRow, error) {
	permissions, err := strg.Querier.GetAllUserPermissions(ctx, pgtype.Int4{Int32: user_id, Valid: true})
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
