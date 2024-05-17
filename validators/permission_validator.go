package validators

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/henrieto/account/models/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

type PermissionData struct {
	Model    string `json:"model" validate:"required,email"`
	Name     string `json:"name" validate:"required,email"`
	Codename string `json:"codename" validate:"required,email"`
}

func NewPermissionData() *PermissionData {
	return new(PermissionData)
}

func (data *PermissionData) Valid() (*db.Permission, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return nil, err
	}
	return &db.Permission{
		Model:     data.Model,
		Name:      data.Name,
		Codename:  data.Codename,
		CreatedAt: pgtype.Timestamptz{Time: time.Now()},
		UpdatedAt: pgtype.Timestamptz{Time: time.Now()},
	}, nil
}

type GroupPermissionData struct {
	Permission int32 `json:"permission,omitempty" validate:"required,email"`
	Group      int32 `json:"group,omitempty" validate:"required,email"`
}

func NewGroupPermissionData() *GroupPermissionData {
	return new(GroupPermissionData)
}

func (data *GroupPermissionData) Valid() (int32, int32, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return 0, 0, err
	}
	return data.Permission, data.Group, nil
}

type UserPermissionData struct {
	Permission int32 `json:"permission,omitempty" validate:"required,email"`
	User       int32 `json:"user,omitempty" validate:"required,email"`
}

func NewUserPermissionData() *UserPermissionData {
	return new(UserPermissionData)
}

func (data *UserPermissionData) Valid() (int32, int32, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return 0, 0, err
	}
	return data.Permission, data.User, nil
}
