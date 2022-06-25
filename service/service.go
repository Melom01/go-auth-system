package service

import "sentinel/database"

type ServicesWrapper interface {
	DB() database.DBUtilitiesWrapper
}

type ServicesUtilitiesWrapper struct {
	Database database.DBUtilitiesWrapper
}

func (suw *ServicesUtilitiesWrapper) DB() database.DBUtilitiesWrapper {
	return suw.Database
}
