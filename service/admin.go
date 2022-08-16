package service

import (
	"context"
	"crypto/tls"
	"github.com/Nerzal/gocloak/v11"
	"sentinel/apperror"
	"sentinel/config"
)

func GetGoCloakClient() gocloak.GoCloak {
	basePath := config.Config.Keycloak.BasePath

	return gocloak.NewClient(basePath, gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms"))
}

func ConnectToKeycloak() *gocloak.JWT {
	// TODO: Make this a public struct to make it accessible in other points
	// You can provide this variables with dependency injection (WIRE library)
	var (
		client        = GetGoCloakClient()
		ctx           = context.Background()
		realm         = config.Config.Keycloak.Realm
		adminUsername = config.Config.Keycloak.Username
		adminPassword = config.Config.Keycloak.Password
	)

	restyClient := client.RestyClient()

	// TODO: set debug & TLS skip insecure dynamically from config.json file
	restyClient.SetDebug(true)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	token, err := client.LoginAdmin(ctx, adminUsername, adminPassword, realm)
	if err != nil {
		apperror.ThrowError(apperror.ErrServerError("Unable to create Keycloak admin connection"))
		return nil
	}

	return token
}
