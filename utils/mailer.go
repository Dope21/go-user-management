package utils

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"user-management/constants/configs"
	msg "user-management/constants/messages"
)

func SendEmailConfirmation(sendTo string) error {
	cfg := configs.LoadConfig()

	_, err := mail.ParseAddress(sendTo)
	if err != nil {
		return err
	}

	token, err := GenerateConfirmEmailToken(sendTo)
	if err != nil {
		return err
	}

  body := fmt.Sprintf(msg.EmailLink, token)
  message := fmt.Sprintf(msg.EmailHeader, sendTo, msg.EmailTitle, body)

	auth := smtp.PlainAuth("", cfg.MailUser, cfg.MailPass, cfg.MailHost)

	err = smtp.SendMail(cfg.MailHost+":"+cfg.MailPort, auth, cfg.MailUser, []string{sendTo}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}