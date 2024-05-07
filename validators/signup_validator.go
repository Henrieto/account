package validators

import (
	"github.com/go-playground/validator"
	"github.com/henrieto/account/auth"
	"github.com/henrieto/account/models"
)

type SignupData struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Gender          string
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	Terms           string `json:"terms" validate:"required"`
}

type SignupValidationErrorData struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Terms           string `json:"terms"`
}

func (vaError SignupValidationErrorData) Error() string {
	return "validation error"
}

func (valError *SignupValidationErrorData) AddError(field string, errMsg string) {
	switch field {
	case "FirstName":
		valError.FirstName = errMsg
	case "LastName":
		valError.LastName = errMsg
	case "Email":
		valError.Email = errMsg
	case "Password":
		valError.Password = errMsg
	case "ConfirmPassword":
		valError.ConfirmPassword = errMsg
	case "Terms":
		valError.Terms = errMsg
	}
}

func NewSignupValidationErrorData(err error) *SignupValidationErrorData {
	valError := new(SignupValidationErrorData)
	for _, err := range err.(validator.ValidationErrors) {
		switch err.StructField() {
		case "FirstName":
			valError.FirstName = "your first name is required"
		case "LastName":
			valError.LastName = "your last name is required"
		case "Email":
			valError.Email = "your email is not a valid email"
		case "Password":
			valError.Password = "password is not strong enough"
		case "ConfirmPassword":
			valError.ConfirmPassword = "you need to confirm your password"
		case "Terms":
			valError.Terms = "you need to accept our terms and conditions"
		}
	}
	return valError
}

func NewSignupData() *SignupData {
	// initialized an empty signup data object
	return new(SignupData)
}

func (data *SignupData) Error(err error) *SignupValidationErrorData {
	switch valError := err.(type) {
	case validator.ValidationErrors:
		return NewSignupValidationErrorData(valError)
	case *SignupValidationErrorData:
		return valError
	default:
		return new(SignupValidationErrorData)
	}
}

func (data *SignupData) Valid() (*models.User, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return nil, err
	}
	valError := data.Error(nil)
	if data.Password != data.ConfirmPassword {
		valError.AddError("ConfirmPassword", "passwords doesn't match")
		return nil, valError
	}

	if data.Terms != "true" {
		valError.AddError("Terms", "you need to accept our terms and conditions")
		return nil, valError
	}

	err = IsStrongPassword(data.Password)
	if err != nil {
		valError.AddError("Password", err.Error())
		return nil, valError
	}

	user := models.NewUser()
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Email = data.Email
	user.PasswordHash = data.Password

	password_hash, err := auth.HashPassword(data.Password)
	if err != nil {
		panic(err)
	}
	user.PasswordHash = string(password_hash)

	return user, nil
}
