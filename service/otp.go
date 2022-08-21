package service

import (
	"sentinel/model"
	"time"
)

func (suw *ServicesUtilitiesWrapper) StoreOTPUserData(data model.OTPStorage) {
	OTPUser := model.OTPUser{
		Username:  data.Username,
		Email:     data.Email,
		OTPCode:   data.OtpCode,
		Timestamp: time.Now(),
	}

	if suw.Database.CheckIfOTPUserDataAlreadyExist(data.Username) {
		suw.Database.UpdateOTPUserData(data.Username, OTPUser)
	} else {
		suw.Database.SaveOTPUserData(OTPUser)
	}
}
