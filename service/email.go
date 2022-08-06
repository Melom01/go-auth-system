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
	var (
		otpLength = config.Config.Emailer.OTPLength
		sender    = config.Config.Emailer.Sender
		subject   = "Verify your email address"
	)

	type Data struct {
		Username      string
		ReceiverEmail string
		OtpCode       string
	}

	otpCode, err := utils.GenerateOTP(otpLength)
	if err != nil {
		logger.LogMessageInRed("Cannot generate OTP code. The reason was: " + err.Error())
		return err
	}

	data := Data{
		Username:      email.Username,
		ReceiverEmail: email.ReceiverEmail,
		OtpCode:       otpCode,
	}

	body := utils.CreateEmailFromTemplate("verification_email.html", data)

	err = suw.Emailer.SendEmail(sender, model.Email{
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
