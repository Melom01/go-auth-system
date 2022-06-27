package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"sentinel/config"
	"sentinel/controller"
	"sentinel/database"
	"sentinel/emailer"
	"sentinel/logger"
	"sentinel/routing"
	"sentinel/service"
)

func Sum(x, y int) int {
	z := x + y
	return z
}

func main() {
	config.SetupConfig()

	var (
		db       = database.SetupPostgresDatabase()
		email    = emailer.SetupEmailer()
		suw      = &service.ServicesUtilitiesWrapper{Database: db, Emailer: email}
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
