package database

import (
	"sentinel/apperror"
	"sentinel/model"
)

type UserDatabase interface {
	SaveOTPUserData(user model.OTPUser)
	UpdateOTPUserData(username string, user model.OTPUser)
	DeleteOTPUserData(username string)
	CheckIfOTPUserDataAlreadyExist(username string) bool
}

func (db *PostgresDatabase) SaveOTPUserData(user model.OTPUser) {
	err := db.Session.Create(&user).Error

	if err != nil {
		apperror.ThrowError(apperror.ErrUnableToSaveUserData(err.Error()))
		return
	}
}

func (db *PostgresDatabase) UpdateOTPUserData(username string, user model.OTPUser) {
	err := db.Session.Model(model.OTPUser{}).Where("username = ?", username).Updates(user).Error

	if err != nil {
		apperror.ThrowError(apperror.ErrUnableToUpdateUserData(err.Error()))
		return
	}
}

func (db *PostgresDatabase) DeleteOTPUserData(username string) {
	err := db.Session.Delete(model.OTPUser{}, "username = ?", username).Error

	if err != nil {
		apperror.ThrowError(apperror.ErrUnableToDeleteUserData(err.Error()))
		return
	}
}

func (db *PostgresDatabase) CheckIfOTPUserDataAlreadyExist(username string) bool {
	var user model.OTPUser

	query := db.Session.Where("username = ?", username).First(&user)

	if query.Error != nil {
		return false
	}

	return true
}
