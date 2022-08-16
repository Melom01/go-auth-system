package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"sentinel/controller"
)

func setupUserRouter(router *mux.Router, ctrl controller.Controller) {
	router.HandleFunc("/v1/search-user", ctrl.CheckIfUserAlreadyExist).Methods(http.MethodGet)
	router.HandleFunc("/v1/create-user", ctrl.CreateUser).Methods(http.MethodPost)
}
