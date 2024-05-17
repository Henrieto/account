package notifiers

import (
	"errors"
	"fmt"
	"net/smtp"
)

type SmtpNotifier struct {
	Identity       string
	Username       string
	Password       string
	Server_Address string
	Host           string
}

func (esmtp *SmtpNotifier) Auth() (auth smtp.Auth) {
	auth = smtp.PlainAuth(
		esmtp.Identity,
		esmtp.Username,
		esmtp.Password,
		esmtp.Host,
	)
	return
}

func (esmpt *SmtpNotifier) _payload(payload any) (*SmptpPayload, error) {
	switch payload := payload.(type) {
	case *SmptpPayload:
		return payload, nil
	case *EmailPayload:
		return &SmptpPayload{
			From: payload.From,
			To:   payload.To,
			msg: &SmtpMsg{
				Subject: payload.Msg.Subject,
				Body:    payload.Msg.Body,
			},
		}, nil
	default:
		return nil, errors.New("not payload")
	}
}

func (esmtp *SmtpNotifier) _send(payload *SmptpPayload) error {
	err := smtp.SendMail(
		esmtp.Server_Address,
		esmtp.Auth(),
		payload.From,
		payload.To,
		payload.Msg(),
	)

	if err != nil {
		return err
	}
	return nil
}

func (esmtp SmtpNotifier) Send(payload any) error {
	_payload, err := esmtp._payload(payload)
	if err != nil {
		return err
	}
	err = esmtp._send(_payload)
	if err != nil {
		return err
	}
	return nil
}

type SmtpMsg struct {
	Subject string
	Body    string
}

func (msg *SmtpMsg) String() string {
	return fmt.Sprintf("Subject: %v\n%v ", msg.Subject, msg.Body)
}

type SmptpPayload struct {
	From string
	To   []string
	msg  *SmtpMsg
}

func NewSmtpPayload(from string, to ...string) *SmptpPayload {
	return &SmptpPayload{
		From: from,
		To:   to,
	}
}

func (pl *SmptpPayload) SetMsg(subject, body string) {
	pl.msg = &SmtpMsg{Subject: subject, Body: body}
}

func (pl *SmptpPayload) Msg() []byte {
	if pl.msg == nil {
		panic(" Message is need ")
	}
	return []byte(pl.msg.String())
}
