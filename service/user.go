package service

import (
	"context"
	"crypto/tls"
	"github.com/Nerzal/gocloak/v11"
	"sentinel/config"
	"sentinel/logger"
	"sentinel/model"
)

type UserServices interface {
	CreateUser(user model.User) error
}

func (suw *ServicesUtilitiesWrapper) CreateUser(user model.User) error {
	var (
		realm         = config.Config.Keycloak.Realm
		basePath      = config.Config.Keycloak.BasePath
		adminUsername = config.Config.Keycloak.Username
		adminPassword = config.Config.Keycloak.Password
	)

	ctx := context.Background()
	client := gocloak.NewClient(basePath, gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms"))
	restyClient := client.RestyClient()

	restyClient.SetDebug(true)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	token, err := client.LoginAdmin(ctx, adminUsername, adminPassword, realm)
	if err != nil {
		logger.LogFatalMessageInRed("Login failed: ", err)
		return err
	}

	var groups = []string{"Users"}
	var credentials = []gocloak.CredentialRepresentation{
		{
			// TODO: Here you can also handle hash and salt. See how
			Temporary: gocloak.BoolP(false),
			Value:     gocloak.StringP(user.Password),
		},
	}

	gocloakUser := gocloak.User{
		FirstName:   gocloak.StringP(user.FirstName),
		LastName:    gocloak.StringP(user.LastName),
		Email:       gocloak.StringP(user.Email),
		Enabled:     gocloak.BoolP(user.Enabled),
		Username:    gocloak.StringP(user.Username),
		Groups:      &groups,
		Credentials: &credentials,
	}

	_, err = client.CreateUser(ctx, token.AccessToken, realm, gocloakUser)
	if err != nil {
		logger.LogMessageInRed("Cannot create user. The reason was: " + err.Error())
		return err
	}

	return nil
}
