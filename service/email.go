package service

import (
	"sentinel/config"
	"sentinel/logger"
	"sentinel/model"
	"sentinel/utils"
)

type EmailServices interface {
	SendVerificationEmail(email model.Email) error
}

func (suw *ServicesUtilitiesWrapper) SendVerificationEmail(email model.Email) error {
	type Data struct {
		Username      string
		ReceiverEmail string
		OtpCode       string
	}

	data := Data{
		Username:      email.Username,
		ReceiverEmail: email.ReceiverEmail,
		OtpCode:       "123456",
	}

	body := utils.CreateEmailFromTemplate("verification_email.html", data)
	subject := "TEST SUBJECT"

	err := suw.Emailer.SendEmail(config.Config.Emailer.Sender, model.Email{
		ReceiverEmail: email.ReceiverEmail,
		Subject:       subject,
		Body:          body,
	})

	if err != nil {
		logger.LogMessageInYellow("An error occurred while sending the email. The error was: " + err.Error())
		return err
	}

	return nil
}
