package notification

import (
	"errors"

	"github.com/henrieto/account/config"
	"github.com/henrieto/account/notification/notifiers"
)

const (
	// channels
	PHONE = "phone"
	EMAIL = "email"
)

func SendOTP(channel, id string, token string) error {
	switch channel {
	case PHONE:
		return nil
	case EMAIL:
		payload := notifiers.EmailPayload{
			From: config.SMTP_EMAILSENDER,
			To: []string{
				id,
			},
			Msg: &notifiers.Msg{
				Subject: " verify identity",
				Body:    token,
			},
		}
		return SendMail(payload)
	default:
		return errors.New("channel not supported ")
	}
}

func SendMail(payload notifiers.EmailPayload) error {
	err := Manager.Notify("email", payload)
	if err != nil {
		return err
	}
	return nil
}
