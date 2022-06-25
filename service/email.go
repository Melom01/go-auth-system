package service

import (
	"sentinel/config"
	"sentinel/logger"
	"sentinel/model"
	"sentinel/utils"
)

type EmailServices interface {
	SendVerificationEmail()
}

func (suw *ServicesUtilitiesWrapper) SendVerificationEmail() {
	type Data struct {
		Username       string
		RecipientEmail string
		OtpCode        string
	}

	data := Data{
		Username:       "usernameGeneric",
		RecipientEmail: "mounir.progit@gmail.com",
		OtpCode:        "123456",
	}

	text := utils.CreateEmailFromTemplate("verification_email.html", data)

	err := suw.Emailer.SendEmail(config.Config.Emailer.Sender, model.Email{
		Subject: "TEST SUBJECT",
		Text:    text,
		To:      "mounir.progit@gmail.com",
	})
	if err != nil {
		logger.LogFatalMessageInRed("An error occurred while sending the email. The error was: ", err)
		return
	}
}
