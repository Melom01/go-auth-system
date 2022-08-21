package service

import (
	"context"
	"github.com/Nerzal/gocloak/v11"
	"sentinel/apperror"
	"sentinel/config"
	"sentinel/model"
)

type UserServices interface {
	CheckIfUserAlreadyExist(username string, email string) string
	CreateUser(user model.User)
}

func (suw *ServicesUtilitiesWrapper) CheckIfUserAlreadyExist(username string, email string) string {
	var (
		usernameFilter = gocloak.GetUsersParams{Username: gocloak.StringP(username)}
		emailFilter    = gocloak.GetUsersParams{Email: gocloak.StringP(email)}
	)

	usernameList := GetUsersByParams(usernameFilter)
	emailList := GetUsersByParams(emailFilter)

	if len(usernameList) > 0 && len(emailList) > 0 {
		return "USERNAME_AND_EMAIL_EXIST"
	} else if len(usernameList) > 0 {
		return "USERNAME_EXIST"
	} else if len(emailList) > 0 {
		return "EMAIL_EXIST"
	}

	// TODO -> Handle response
	return ""
}

func (suw *ServicesUtilitiesWrapper) CreateUser(user model.User) {
	// TODO: Make this a public struct to make it accessible in other points
	var (
		ctx    = context.Background()
		client = GetGoCloakClient()
		realm  = config.Config.Keycloak.Realm
		token  = ConnectToKeycloak().AccessToken
	)

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

	_, err := client.CreateUser(ctx, token, realm, gocloakUser)
	if err != nil {
		apperror.ThrowError(apperror.ErrUserAlreadyExist(err.Error()))
	}

	suw.Database.DeleteOTPUserData(user.Username)
}
