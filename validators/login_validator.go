package validators

import "github.com/go-playground/validator"

type LoginData struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

func NewLoginData() *LoginData {
	return new(LoginData)
}

func (data *LoginData) Valid() (string, string, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return "", "", nil
	}
	return data.Email, data.Password, nil
}
