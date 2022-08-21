package emailer

import (
	"github.com/badoux/checkmail"
	mail "github.com/xhit/go-simple-mail/v2"
	"sentinel/config"
	"sentinel/logger"
	"sentinel/model"
)

type Emailer interface {
	SendEmail(sender string, emailData model.VerificationEmail) error
}

type SMTPEmailer struct {
	server *mail.SMTPServer
}

func SetupEmailer() *SMTPEmailer {
	server := mail.NewSMTPClient()

	server.Host = config.Config.Emailer.Host
	server.Port = config.Config.Emailer.Port
	server.Username = config.Config.Emailer.Sender
	server.Password = config.Config.Emailer.Password

	switch config.Config.Emailer.Encryption {
	case "NONE":
		server.Encryption = mail.EncryptionNone
	case "SSL":
		server.Encryption = mail.EncryptionSSL
	case "TLS":
		server.Encryption = mail.EncryptionTLS
	default:
		server.Encryption = mail.EncryptionNone
	}

	return &SMTPEmailer{
		server: server,
	}
}

func validateEmail(email string) error {
	err := checkmail.ValidateFormat(email)

	if err != nil {
		logger.LogMessageInYellow("Invalid email format. The email provided was: " + email)
		return err
	}

	err = checkmail.ValidateHost(email)

	if err != nil {
		logger.LogMessageInYellow("Invalid email host. The email provided was: " + email)
		return err
	}

	return nil
}

func (emailer *SMTPEmailer) SendEmail(sender string, emailData model.VerificationEmail) error {
	email := mail.NewMSG()

	err := validateEmail(emailData.Email)
	if err != nil {
		return err
	}

	email.SetFrom(sender).
		AddTo(emailData.Email).
		SetSubject(emailData.Subject).
		SetBody(mail.TextHTML, emailData.Body).
		AddHeader("Content-Transfer-Encoding", "quoted-printable")

	client, err := emailer.server.Connect()
	if err != nil {
		logger.LogMessageInRed("Cannot connect to the mail client: " + err.Error())
		return err
	}

	err = email.Send(client)
	if err != nil {
		logger.LogMessageInRed("Error while sending email to: " + emailData.Subject + ". The error was: " + err.Error())
		return err
	}

	return nil
}
