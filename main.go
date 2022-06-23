package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"sentinel/config"
	"sentinel/logger"
	"sentinel/routing"
)

func main() {
	config.SetupConfig()

	var (
		router   = mux.NewRouter()
		hostname = config.Config.Server.Hostname
		port     = config.Config.Server.Port
	)

	routing.SetupRouter(router)
	logger.LogMessageInGreen("Starting server at: " + port)

	if err := http.ListenAndServe(hostname+":"+port, router); err != nil {
		logger.LogFatalMessageInRed("The server cannot listen and serve: ", err)
	}
}
