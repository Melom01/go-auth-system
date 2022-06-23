package routing

import (
	"github.com/gorilla/mux"
	"sentinel/middleware"
)

func SetupRouter(router *mux.Router) {
	router.Use(middleware.LoggingMiddleware, middleware.RecoveryPanicMiddleware)
}
