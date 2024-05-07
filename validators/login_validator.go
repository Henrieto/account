package validators

type LoginData struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewLoginData() *LoginData {
	return new(LoginData)
}

func (data *LoginData) Valid() (string, string, error) {
	return data.Email, data.Password, nil
}
