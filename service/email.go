package service

import (
	"sentinel/config"
	"sentinel/logger"
	"sentinel/model"
	"sentinel/utils"
)

type EmailServices interface {
	SendVerificationEmail(email model.VerificationEmail) error
}

func (suw *ServicesUtilitiesWrapper) SendVerificationEmail(email model.VerificationEmail) error {
	var (
		OTPLength = config.Config.Emailer.OTPLength
		sender    = config.Config.Emailer.Sender
		subject   = "Verify your email address"
	)

	OTPCode, err := utils.GenerateOTP(OTPLength)
	if err != nil {
		logger.LogMessageInRed("Cannot generate OTP code. The reason was: " + err.Error())
		return err
	}

	data := model.OTPStorage{
		Username: email.Username,
		Email:    email.Email,
		OtpCode:  OTPCode,
	}

	suw.StoreOTPUserData(data)

	body := utils.CreateEmailFromTemplate("verification_email.html", data)

	err = suw.Emailer.SendEmail(sender, model.VerificationEmail{
		Email:   email.Email,
		Subject: subject,
		Body:    body,
	})

	if err != nil {
		logger.LogMessageInYellow("An error occurred while sending the email. The error was: " + err.Error())
		return err
	}

	return nil
}
