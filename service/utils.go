package service

import (
	"context"
	"github.com/Nerzal/gocloak/v11"
	"sentinel/apperror"
	"sentinel/config"
)

func GetUsersByParams(params gocloak.GetUsersParams) []*gocloak.User {
	// TODO: Make this a public struct to make it accessible in other points
	var (
		ctx    = context.Background()
		client = GetGoCloakClient()
		realm  = config.Config.Keycloak.Realm
		token  = ConnectToKeycloak().AccessToken
	)

	users, err := client.GetUsers(ctx, token, realm, params)
	if err != nil {
		apperror.ThrowError(apperror.ErrGocloak("Unable to search for users"))
		return nil
	}

	return users
}
