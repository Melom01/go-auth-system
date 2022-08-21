package model

import "time"

type OTPUser struct {
	Username  string    `json:"username" gorm:"primaryKey;not null;default:null"`
	Email     string    `json:"email" gorm:"not null;default:null"`
	OTPCode   string    `json:"otp_code" gorm:"not null;default:null"`
	Timestamp time.Time `json:"timestamp" gorm:"not null;default:null"`
}
