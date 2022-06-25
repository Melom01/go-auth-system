package config

import (
	"encoding/json"
	"io/ioutil"
	"sentinel/logger"
)

var Config Configuration

type Configuration struct {
	Server   ServerConfig
	Database DatabaseConfig
	Emailer  EmailerConfig
}

type ServerConfig struct {
	Title    string
	Version  string
	Hostname string
	Port     string
}

type DatabaseConfig struct {
	Username string
	Password string
	Hostname string
	Name     string
	Port     string
}

type EmailerConfig struct {
	Sender      string
	EmailTitle  string
	Password    string
	Encoding    string
	Encryption  string
	Host        string
	Port        int
	OTPLifeSpan int
}

func SetupConfig() {
	readConfig()
}

func readConfig() {
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		logger.LogFatalMessageInRed("Unable to read configuration file: ", err)
	}

	if err = json.Unmarshal(raw, &Config); err != nil {
		logger.LogFatalMessageInRed("Unable to parse configuration file: ", err)
	}
}
