package tests

import (
	"testing"

	"github.com/henrieto/account/notification/notifiers"
)

func TestSmtpDefaultNotifier(t *testing.T) {
	// initialize the notifier
	notifier := notifiers.SmtpNotifier{
		Username:       "evihsltd@gmail.com",
		Password:       "bdizhilunqhmmqmm",
		Host:           "smtp.gmail.com",
		Server_Address: "smtp.gmail.com:587",
	}
	// create the payload
	payload := notifiers.NewSmtpPayload("evihsltd@gmail.com", "kalukennedyh@gmail.com")
	// set the payload message
	payload.SetMsg("RENTIT TEST EMAIL", "HELLO IT WORKED")
	// send the message
	err := notifier.Send(payload)
	// if an error occured ; test failed
	if err != nil {
		t.Error("email notification failed")
	}
}
