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
	setupPublicRouter(router.PathPrefix("/api").Subrouter(), ctrl)
}

func setupPublicRouter(router *mux.Router, ctrl controller.Controller) {
	setupEmailRouter(router.PathPrefix("/email").Subrouter(), ctrl)
	setupUserRouter(router.PathPrefix("/user").Subrouter(), ctrl)
}

func setupPingRouter(router *mux.Router) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Server ON"))
		if err != nil {
			logger.LogMessageInRed("Server not available.")
		}
	})
}
