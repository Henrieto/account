package jwt_auth

type Error string

func (err Error) Error() string {
	return string(err)
}

func NewError(msg string) Error {
	return Error(msg)
}

var (
	ErrJwtTokenNotValid Error = "jwt token not valid"
)
