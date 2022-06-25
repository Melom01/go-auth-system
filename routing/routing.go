package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"sentinel/controller"
	"sentinel/logger"
	"sentinel/middleware"
)

func SetupRouter(router *mux.Router, ctrl controller.Controller) {
	router.Use(middleware.LoggingMiddleware, middleware.RecoveryPanicMiddleware)

	setupPingRouter(router)
}

func setupPingRouter(router *mux.Router) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Server ON"))
		if err != nil {
			logger.LogMessageInRed("The server is not available.")
		}
	})
}
