package controller

import (
	"net/http"
	"sentinel/model"
)

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type PostUser struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Enabled   *bool  `json:"enabled" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

func (ctrl *HTTPController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var postUser PostUser

	ctrl.DecodeBody(r, &postUser)

	user := model.User{
		FirstName: postUser.FirstName,
		LastName:  postUser.LastName,
		Email:     postUser.Email,
		Enabled:   *postUser.Enabled,
		Username:  postUser.Username,
		Password:  postUser.Password,
	}

	err := ctrl.ServicesWrapper.CreateUser(user)
	if err != nil {
		// TODO: return 409 or 500 NOT statically
		// See how to write header dynamically based on the error that you get
		w.WriteHeader(409)
	}
}
