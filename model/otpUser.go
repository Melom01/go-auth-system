package model

import "time"

type OTPUser struct {
	ID        string    `json:"id" gorm:"primaryKey;not null"`
	Username  string    `json:"username" gorm:"not null;default:null"`
	OTPCode   string    `json:"OTPCode" gorm:"not null;default:null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
