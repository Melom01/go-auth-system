package service

import (
	"context"
	"crypto/tls"
	"github.com/Nerzal/gocloak/v11"
	"sentinel/apperror"
	"sentinel/config"
	"sentinel/model"
)

type UserServices interface {
	CreateUser(user model.User)
}

func (suw *ServicesUtilitiesWrapper) CreateUser(user model.User) {
	var (
		realm         = config.Config.Keycloak.Realm
		basePath      = config.Config.Keycloak.BasePath
		adminUsername = config.Config.Keycloak.Username
		adminPassword = config.Config.Keycloak.Password
	)

	ctx := context.Background()
	client := gocloak.NewClient(basePath, gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms"))
	restyClient := client.RestyClient()

	// TODO: set debug dynamically from config.json file
	restyClient.SetDebug(true)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	token, err := client.LoginAdmin(ctx, adminUsername, adminPassword, realm)
	if err != nil {
		apperror.ThrowError(apperror.ErrServerError("Unable to create Keycloak admin connection"))
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
		apperror.ThrowError(apperror.ErrUserAlreadyExist(err.Error()))
	}
}
