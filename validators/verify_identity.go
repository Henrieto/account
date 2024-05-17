package validators

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/henrieto/account/auth"
	"github.com/henrieto/account/models/database/db"
	"github.com/henrieto/account/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	// errors
	NotOperationType      utils.Error = "operation type not accepted"
	NotIdentificationType utils.Error = " identification type not accepted"

	// otp expiry time
	Expiry = time.Now().Add(time.Minute * 1)

	// random string length
	RandomStringLength = 32
)

const (
	// operation types
	PASSWORDRESET  = "password-reset"
	VERIFYIDENTITY = "verify-identity"

	// identification types
	PHONE = "phone"
	EMAIL = "email"
)

type VerifyIdentityData struct {
	OperationType      string
	IdentificationType string `json:"type,omitempty" validate:"required"`
	Value              string `json:"email,omitempty" validate:"required"`
}

func NewerifyIdentityData() *VerifyIdentityData {
	return new(VerifyIdentityData)
}

func (data *VerifyIdentityData) Valid() (*db.VerifyIdentityData, error) {
	// initalize the data validator
	var validate *validator.Validate = validator.New()
	// validate the verify identity data
	err := validate.Struct(data)
	// return an error , if any occured
	if err != nil {
		return nil, err
	}
	// check if the operation type is accepted
	value, ok := IsAcceptedOperationType(data.OperationType)
	// if not accepted , return an error
	if !ok {
		return nil, NotOperationType
	}
	// check if the operation type is accepted
	value, ok = IsAcceptedOperationType(data.OperationType)
	// if not accepted , return an error
	if !ok {
		return nil, NotOperationType
	}
	// generate the one time password
	otp, err := auth.GenerateToken(6)
	// if an error occured , return the error
	if err != nil {
		return nil, err
	}
	// generate the random string
	random_string := auth.GenerateRandomString(RandomStringLength)

	return &db.VerifyIdentityData{
		RandomString:        random_string,
		Otp:                 otp,
		OperationType:       data.OperationType,
		IdentificationType:  data.IdentificationType,
		IdentificationValue: value,
		Expiry:              pgtype.Timestamptz{Time: Expiry},
	}, nil
}

func IsAcceptedOperationType(oeration_type string) (string, bool) {
	switch oeration_type {
	case PASSWORDRESET:
		return "", true
	case VERIFYIDENTITY:
		return "", true
	default:
		return "", false
	}
}

func IsIdentificationType(identification_type string) (string, bool) {
	switch identification_type {
	case EMAIL:
		return "", true
	case PHONE:
		return "", true
	default:
		return "", false
	}
}

type OtpData struct {
	Otp string `json:"otp,omitempty" validate:"required"`
}

func NewOtpData() *OtpData {
	return new(OtpData)
}

func (data *OtpData) Valid() (string, error) {
	var validate *validator.Validate = validator.New()
	err := validate.Struct(data)
	if err != nil {
		return "", err
	}
	return data.Otp, nil
}
