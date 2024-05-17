package validators

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/henrieto/account/models/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// swagger:GroupData
type GroupData struct {
	// the name for the group
	// required: true
	// min length: 3
	Name string `json:"name" validate:"required"`
}

func NewGroupData() *GroupData {
	return new(GroupData)
}

func (data *GroupData) Valid() (*db.Group, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return nil, err
	}
	return &db.Group{
		Name:      data.Name,
		CreatedAt: pgtype.Timestamptz{Time: time.Now()},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now()},
	}, nil
}

type UserGroupData struct {
	User  int32 `json:"user,omitempty" validate:"required"`
	Group int32 `json:"group,omitempty" validate:"required"`
}

func NewUserGroupData() *UserGroupData {
	return new(UserGroupData)
}

func (data *UserGroupData) Valid() (int32, int32, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return 0, 0, err
	}
	return data.User, data.Group, nil
}
