package validators

import (
	"errors"

	"github.com/go-playground/validator"
)

type PasswordChangeData struct {
	Password        string `json:"password,omitempty" validate:"required"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"required"`
}

func NewPasswordChangeData() *PasswordChangeData {
	return new(PasswordChangeData)
}

func (data *PasswordChangeData) Valid() (string, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return "", err
	}
	if data.Password != data.ConfirmPassword {
		return "", errors.New(" passwords does not match")
	}
	return data.Password, nil
}
