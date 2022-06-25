package service

import (
	"sentinel/config"
	"sentinel/logger"
	"sentinel/model"
)

type EmailServices interface {
	SendVerificationEmail()
}

func (suw *ServicesUtilitiesWrapper) SendVerificationEmail() {
	err := suw.Emailer.SendEmail(config.Config.Emailer.Sender, model.Email{
		Subject: "TEST SUBJECT",
		Text:    "Hello from Golang",
		To:      "mounir.progit@gmail.com",
	})
	if err != nil {
		logger.LogFatalMessageInRed("An error occurred while sending the email. The error was: ", err)
		return
	}
}
