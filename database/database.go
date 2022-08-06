package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sentinel/config"
	"sentinel/logger"
	"sentinel/model"
)

type DBUtilitiesWrapper interface {
	// TODO: add here ORM calls
}

type PostgresDatabase struct {
	Session *gorm.DB
}

func SetupPostgresDatabase() *PostgresDatabase {
	var psqlDatabase PostgresDatabase
	var err error
	var (
		dbUsername = config.Config.Database.Username
		dbPassword = config.Config.Database.Password
		dbHostname = config.Config.Database.Hostname
		dbName     = config.Config.Database.Name
		dbPort     = config.Config.Database.Port
		dbDSN      = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHostname, dbPort, dbName)
	)

	psqlDatabase.Session, err = gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		logger.LogFatalMessageInRed("Error during database initialization: ", err)
		return nil
	}

	err = psqlDatabase.Session.AutoMigrate(&model.OTPUser{})
	if err != nil {
		logger.LogFatalMessageInRed("Cannot execute database migrations: ", err)
		return nil
	}

	logger.LogMessageInGreen("Connected successfully to the database")
	return &psqlDatabase
}
