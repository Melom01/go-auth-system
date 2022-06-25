package utils

import (
	"bytes"
	"html/template"
	"sentinel/logger"
)

func CreateEmailFromTemplate(templateName string, data interface{}) string {
	var buff bytes.Buffer

	tmpl, err := template.New(templateName).ParseFiles(EmailTemplatePrefix + templateName)
	if err != nil {
		logger.LogMessageInRed("Failed to parse email template. The error was: " + err.Error())
		return buff.String()
	}

	err = tmpl.ExecuteTemplate(&buff, templateName, data)
	if err != nil {
		logger.LogMessageInRed("Cannot apply data to the email template. The error was: " + err.Error())
		return buff.String()
	}

	return buff.String()
}
