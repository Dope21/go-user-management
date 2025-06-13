package utils

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"user-management/constants/configs"
	msg "user-management/constants/messages"
	"user-management/models"
)

func SendEmailConfirmation(user *models.User) error {
	cfg := configs.LoadConfig()

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}

	token, err := GenerateConfirmEmailToken(user.ID)
	if err != nil {
		return err
	}

  body := fmt.Sprintf(msg.EmailLink, cfg.AppURL, token)
  message := fmt.Sprintf(msg.EmailHeader, user.Email, msg.EmailTitle, body)

	auth := smtp.PlainAuth("", cfg.MailUser, cfg.MailPass, cfg.MailHost)

	err = smtp.SendMail(
		cfg.MailHost+":"+cfg.MailPort, 
		auth, 
		cfg.MailUser, 
		[]string{user.Email}, 
		[]byte(message),
	)
	if err != nil {
		return err
	}

	return nil
}