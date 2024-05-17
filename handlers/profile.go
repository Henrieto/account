package handlers

import (
	"context"
	"net/http"

	"github.com/henrieto/account/auth"
	jwt_auth "github.com/henrieto/account/auth/jwt"
	"github.com/henrieto/account/config"
	"github.com/henrieto/account/models"
	"github.com/henrieto/account/models/repository"
	"github.com/henrieto/account/notification"
	"github.com/henrieto/account/utils"
	db_utils "github.com/henrieto/account/utils/db"
	"github.com/henrieto/account/validators"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	Background = context.Background()
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewSignupData()
	// bind the request data to the validator object
	err := utils.BindJson(r, data_validator)
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// check if the data is valid
	user, err := data_validator.Valid()
	// if the user is not valid return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	//add the user data to the database
	user, err = repository.User.Create(Background, user)
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " could not create user"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = "user created successfully"
	//  set the data for the response
	response.Data = user
	// send the response
	response.Send(w, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewLoginData()
	// bind the request data to the validator
	err := utils.BindJson(r, data_validator)
	// if there was an error , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// check if the data is valid
	email, password, err := data_validator.Valid()
	// if the user is not valid return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// fetch the user from the database
	user, err := repository.User.GetByEmail(Background, email)
	// if there was an error ,  return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " user does not exists"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// authenticate the user
	authenticated := auth.ComparePassword([]byte(user.PasswordHash), []byte(password))
	// if user authentication failed
	if !authenticated {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = "your credentials is invalid"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// fetch user permissions
	user_permissions, err := repository.Permission.GetAllUserPermissions(Background, user.ID)
	// if there was an error ,  return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " coul not retrieve user permissions"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// fetch user permissions
	user_group_permissions, err := repository.Permission.GetAllGroupPermissions(Background, user.GroupID.Int32)
	// if there was an error ,  return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " coul not retrieve user group permissions"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// add db user permissions to user object
	user.Permissions = append(user.Permissions, db_utils.TranslatePermission(user_permissions)...)
	// add db user group permissions to user object
	user.Permissions = append(user.Permissions, db_utils.TranslatePermission(user_group_permissions)...)
	// if there was an error ,  return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " coul not retrieve user permissions"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// initialize a jwt claim
	claims := jwt_auth.NewClaims()
	// add the user object in to the claims
	claims.Object = user
	// generate the jwt token
	token, err := claims.GenerateJwtToken(config.SECRET)
	// if there was an error , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = "user login failed"
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = "you are successfully logged in"
	//  set the data for the response
	response.Data = map[string]any{
		// jwt token for jwt authentication
		"token": token,
		// user data , retrieved from the database
		"user": user,
	}
	// send the response
	response.Send(w, http.StatusOK)
}

func Profile(w http.ResponseWriter, r *http.Request) {}

func VerifyIdentityRequest(w http.ResponseWriter, r *http.Request) {
	// initialize a data validator
	data_validator := validators.NewerifyIdentityData()
	// bind the request data to the validator
	err := utils.BindJson(r, data_validator)
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	verify_identity_data, err := data_validator.Valid()
	// if the data is not valid , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid"
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// add the verify identity data to the database
	verify_identity_data, err = repository.VerifyIdentityData.Create(Background, verify_identity_data)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " verify identity data failed "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// check if the user with the data exists
	// check if the verification identity type is email
	if verify_identity_data.IdentificationType == validators.EMAIL {
		// check if the email exists
		email_exists := repository.User.EmailExists(Background, verify_identity_data.IdentificationValue)
		// if email does not exists , return a failed response
		if !email_exists {
			// create a response
			response := utils.NewResponse()
			// set the response status e.g success
			response.Status = "failed"
			// set the message for the response
			response.Msg = " user with the email , does not exists "
			// send the response
			response.Send(w, http.StatusOK)
			return
		}
	}
	// check if the verification identity type is a phone (phone number)
	if verify_identity_data.IdentificationType == validators.PHONE {
		// check if the phone number exists
		phone_exists := repository.User.PhoneExists(Background, verify_identity_data.IdentificationValue)
		// if phone number does not exists , return a failed response
		if !phone_exists {
			// create a response
			response := utils.NewResponse()
			// set the response status e.g success
			response.Status = "failed"
			// set the message for the response
			response.Msg = " user with the phone number , does not exists "
			// send the response
			response.Send(w, http.StatusOK)
			return
		}
	}
	// notify the anonymouse user with otp
	err = notification.SendOTP(verify_identity_data.IdentificationType, verify_identity_data.IdentificationValue, verify_identity_data.Otp)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " verify identity data failed "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// return a successful response
	// create a response
	response := utils.NewResponse()
	// set the response status e.g success
	response.Status = "success"
	// set the message for the response
	response.Msg = " group created successfully"
	//  set the data for the response
	response.Data = verify_identity_data.RandomString
	// send the response
	response.Send(w, http.StatusOK)
}

func VerifyIdentity(w http.ResponseWriter, r *http.Request) {
	// retrieve the verify identity randon string from
	// the url
	random_string := r.PathValue("random_string")
	// check if the random string is empty
	if random_string == "" {
		// if empty , return a failed response
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " token is empty "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// fetch the verify identity data using the random string
	verify_identity_data, err := repository.VerifyIdentityData.Get(Background, random_string)
	// if an error occured , return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " token is not valid "
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// initialize a data validator
	data_validator := validators.NewOtpData()
	// bind the request data to the validator
	err = utils.BindJson(r, data_validator)
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid "
		// send the response
		response.Send(w, http.StatusBadRequest)
		return
	}
	// check if data is valid
	otp, err := data_validator.Valid()
	// if an error occured return a failed response
	if err != nil {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " data is invalid "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// check id the otp is valid
	otp_valid := verify_identity_data.Otp == otp
	// if not valid return a failed response
	if !otp_valid {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " otp is invalid "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// check if the otp has expired
	expired := auth.ExpiredOtp(verify_identity_data.Expiry.Time)
	// if the otp has expired , return a failed response
	if expired {
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "failed"
		// set the message for the response
		response.Msg = " otp has expired "
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
	// initialize a user variiable
	var user *models.User
	// check if the identifiaction type is email
	if verify_identity_data.IdentificationType == validators.EMAIL {
		// fetch the user using the email
		user, err = repository.User.GetByEmail(Background, verify_identity_data.IdentificationValue)
		// if an error occured , return a failed response
		if err != nil {
			// create a response
			response := utils.NewResponse()
			// set the response status e.g success
			response.Status = "failed"
			// set the message for the response
			response.Msg = " user with the email does not exists "
			// send the response
			response.Send(w, http.StatusOK)
			return
		}
	} else if verify_identity_data.IdentificationType == validators.PHONE {
		// check if the identifiaction type is phone
		// fetch the user using the phone number
		user, err = repository.User.GetByPhone(Background, verify_identity_data.IdentificationValue)
		// if an error occured , return a failed response
		if err != nil {
			// create a response
			response := utils.NewResponse()
			// set the response status e.g success
			response.Status = "failed"
			// set the message for the response
			response.Msg = " user with the email does not exists "
			// send the response
			response.Send(w, http.StatusOK)
			return
		}
	}

	// check the operation type for the identity verification
	if verify_identity_data.OperationType == validators.VERIFYIDENTITY {

		// set the user verified attribute to true
		user.Verified = pgtype.Bool{Bool: true}
		// update the database user data
		user, err = repository.User.Update(Background, user)
		// if an error occured , return a failed response
		if err != nil {
			// create a response
			response := utils.NewResponse()
			// set the response status e.g success
			response.Status = "failed"
			// set the message for the response
			response.Msg = " could not verify user"
			// send the response
			response.Send(w, http.StatusOK)
			return
		}
		// delete the user identity verification data
		err = repository.VerifyIdentityData.Delete(Background, verify_identity_data.ID)
		// if an error occured , return a failed response
		if err != nil {
			// create a response
			response := utils.NewResponse()
			// set the response status e.g success
			response.Status = "failed"
			// set the message for the response
			response.Msg = " could not delete identity verification data"
			// send the response
			response.Send(w, http.StatusOK)
			return
		}
		// return a successful response
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "success"
		// set the message for the response
		response.Msg = " user verified successfully"
		//  set the data for the response
		response.Data = user
		// send the response
		response.Send(w, http.StatusOK)
		return
	} else if verify_identity_data.OperationType == validators.PASSWORDRESET {
		// generate password reset token
		token := auth.GenerateRandomString(validators.RandomStringLength)
		// add the password reset token to the user auth id
		user.AuthID = pgtype.Text{String: token}
		// update the user database data
		_, err = repository.User.Update(Background, user)
		// if an error occured , return a failed response
		if err != nil {
			// create a response
			response := utils.NewResponse()
			// set the response status e.g success
			response.Status = "failed"
			// set the message for the response
			response.Msg = " could not update user"
			// send the response
			response.Send(w, http.StatusOK)
			return
		}
		// delete the user identity verification data
		err = repository.VerifyIdentityData.Delete(Background, verify_identity_data.ID)
		// if an error occured , return a failed response
		if err != nil {
			// create a response
			response := utils.NewResponse()
			// set the response status e.g success
			response.Status = "failed"
			// set the message for the response
			response.Msg = " could not delete identity verification data"
			// send the response
			response.Send(w, http.StatusOK)
			return
		}
		// return a successful response
		// create a response
		response := utils.NewResponse()
		// set the response status e.g success
		response.Status = "success"
		// set the message for the response
		response.Msg = " proceed to change password"
		//  set the data for the response
		response.Data = token
		// send the response
		response.Send(w, http.StatusOK)
		return
	}
}
