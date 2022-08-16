package controller

import (
	"net/http"
	"sentinel/model"
)

type UserController interface {
	CheckIfUserAlreadyExist(w http.ResponseWriter, r *http.Request)
	CreateUser(_ http.ResponseWriter, r *http.Request)
}

func (ctrl *HTTPController) CheckIfUserAlreadyExist(w http.ResponseWriter, r *http.Request) {
	var (
		username = r.URL.Query().Get("username")
		email    = r.URL.Query().Get("email")
	)

	response := ctrl.ServicesWrapper.CheckIfUserAlreadyExist(username, email)

	SetJsonHeadersAndEncode(w, response)
}

type PostUser struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Enabled   *bool  `json:"enabled" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

func (ctrl *HTTPController) CreateUser(_ http.ResponseWriter, r *http.Request) {
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

	ctrl.ServicesWrapper.CreateUser(user)
}
