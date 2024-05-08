package validators

import "github.com/go-playground/validator"

type ForgotPasswordData struct {
	Email string `json:"email,omitempty" validate:"required"`
}

func NewForgotPasswordDataValidator() *ForgotPasswordData {
	return new(ForgotPasswordData)
}

func (data *ForgotPasswordData) Valid() (string, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return "", nil
	}
	return data.Email, nil
}
