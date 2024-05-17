package validators

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/henrieto/account/models"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserData struct {
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Gender    string    `json:"gender" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Verified  bool      `json:"verified" validate:"required"`
	Birthday  time.Time `json:"birthday" validate:"required"`
	Staff     bool      `json:"staff" validate:"required"`
	Superuser bool      `json:"superuser" validate:"required"`
	AuthID    string    `json:"auth_id" validate:"required"`
	GroupID   int32     `json:"group_id" validate:"required"`
}

func NewUserData() *UserData {
	return new(UserData)
}

func (data *UserData) Valid() (*models.User, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Username:     data.Username,
		Email:        data.Email,
		Phone:        data.Phone,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Gender:       data.Gender,
		PasswordHash: data.Password,
		Verified:     pgtype.Bool{Bool: data.Verified},
		Birthday:     pgtype.Timestamptz{Time: data.Birthday},
		Staff:        pgtype.Bool{Bool: data.Staff},
		Superuser:    pgtype.Bool{Bool: data.Superuser},
		GroupID:      pgtype.Int4{Int32: data.GroupID},
	}, nil
}
