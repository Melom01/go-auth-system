package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"sentinel/config"
	"sentinel/controller"
	"sentinel/database"
	"sentinel/logger"
	"sentinel/routing"
	"sentinel/service"
)

func main() {
	config.SetupConfig()

	var (
		db       = database.SetupPostgresDatabase()
		suw      = &service.ServicesUtilitiesWrapper{Database: db}
		ctrl     = controller.HTTPController{ServicesWrapper: suw}
		router   = mux.NewRouter()
		hostname = config.Config.Server.Hostname
		port     = config.Config.Server.Port
	)

	routing.SetupRouter(router, &ctrl)
	logger.LogMessageInGreen("Starting server at: " + port)

	if err := http.ListenAndServe(hostname+":"+port, router); err != nil {
		logger.LogFatalMessageInRed("The server cannot listen and serve: ", err)
	}
}
