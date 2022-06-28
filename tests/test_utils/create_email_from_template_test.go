package test_utils

import (
	"bytes"
	"html/template"
	"testing"
)

func TestCreateEmailFromTemplate(t *testing.T) {
	type Data struct {
		Username       string
		RecipientEmail string
		OtpCode        string
	}

	var buff bytes.Buffer

	templatePrefix := "mock_data/"
	templateName := "verification_email_test.html"
	data := Data{
		Username:       "username",
		RecipientEmail: "recipient@mail.com",
		OtpCode:        "123456",
	}

	tmpl, err := template.New(templateName).ParseFiles(templatePrefix + templateName)
	if err != nil {
		t.Errorf("Failed to parse email template. The error was: %v", err.Error())
	}

	err = tmpl.ExecuteTemplate(&buff, templateName, data)
	if err != nil {
		t.Errorf("Cannot apply data to the email template. The error was: %v", err.Error())
	}
}
