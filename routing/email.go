package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"sentinel/controller"
)

func setupEmailRouter(router *mux.Router, ctrl controller.Controller) {
	router.HandleFunc("/v1/send-verification-email", ctrl.SendVerificationEmail).Methods(http.MethodPost)
}
