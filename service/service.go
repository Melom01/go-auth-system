package service

import (
	"sentinel/database"
	"sentinel/emailer"
)

type ServicesWrapper interface {
	EmailServices

	GetDB() database.DBUtilitiesWrapper
	GetEmailer() emailer.Emailer
}

type ServicesUtilitiesWrapper struct {
	Database database.DBUtilitiesWrapper
	Emailer  emailer.Emailer
}

func (suw *ServicesUtilitiesWrapper) GetDB() database.DBUtilitiesWrapper {
	return suw.Database
}

func (suw *ServicesUtilitiesWrapper) GetEmailer() emailer.Emailer {
	return suw.Emailer
}
