package storage

import (
	"context"
	"time"

	"github.com/henrieto/account/models/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type Group struct {
	Querier db.Querier
}

func NewGroupStorage(querier db.Querier) *Group {
	return &Group{querier}
}

// (context , group data) (db group data , error)
func (strg *Group) Create(ctx context.Context, group *db.Group) (*db.Group, error) {
	// create the db user ceate parameters
	params := db.CreateUserGroupParams{
		Name:      group.Name,
		CreatedAt: pgtype.Timestamptz{Time: time.Now()},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now()},
	}
	// create the group
	user_group, err := strg.Querier.CreateUserGroup(ctx, params)
	// if an error occured , return an error
	if err != nil {
		return nil, err
	}
	return &user_group, nil
}

// (context) (number of groups , error )
func (strg *Group) Count(ctx context.Context) (int64, error) {
	count, err := strg.Querier.CountGroups(ctx)
	if err != nil {
		return 0, err
	}
	return count, err
}

// (context , order by)(db user groups , error)
func (strg *Group) List(ctx context.Context, order_by string) ([]db.Group, error) {
	// retrieve the user groups
	groups, err := strg.Querier.GetAllGroups(ctx, order_by)
	// if an error occured , return a failed response
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// (context , order by , offset  , limit) (db groups , error)
func (strg *Group) Paginate(ctx context.Context, order_by string, offset uint, limit uint) ([]db.Group, error) {
	params := db.PaginateGroupsParams{
		Column1: order_by,
		Offset:  int32(offset),
		Limit:   int32(limit),
	}
	groups, err := strg.Querier.PaginateGroups(ctx, params)
	if err != nil {
		return nil, err
	}
	return groups, err
}

// (context , group id) (db group , error)
func (strg *Group) Get(ctx context.Context, name string) (*db.Group, error) {
	group, err := strg.Querier.GetGroup(ctx, name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// (context , group data , group id to update) (db group data , error)
func (strg *Group) Update(ctx context.Context, group *db.Group, id int32) (*db.Group, error) {
	params := db.UpdateGroupParams{
		Name:      group.Name,
		UpdatedAt: group.UpdatedAt,
		ID:        id,
	}
	user_group, err := strg.Querier.UpdateGroup(ctx, params)
	if err != nil {
		return nil, err
	}
	return &user_group, nil
}

// (context , group id) (error)
func (strg *Group) Delete(ctx context.Context, id int32) error {
	err := strg.Querier.DeleteGroup(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// (context , group id , user id ) (error)
func (strg *Group) AddToGroup(ctx context.Context, group_id int32, user_id int32) error {
	params := db.AddUserToGroupParams{
		GroupID:   pgtype.Int4{Int32: group_id},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now()},
		ID:        user_id,
	}
	err := strg.Querier.AddUserToGroup(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
