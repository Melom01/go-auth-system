package controller

import (
	"net/http"
	"sentinel/model"
)

type EmailController interface {
	SendVerificationEmail(w http.ResponseWriter, r *http.Request)
}

type PostVerificationEmail struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
}

func (ctrl *HTTPController) SendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	var postEmail PostVerificationEmail

	ctrl.DecodeBody(r, &postEmail)

	email := model.VerificationEmail{
		Email:    postEmail.Email,
		Username: postEmail.Username,
	}

	err := ctrl.ServicesWrapper.SendVerificationEmail(email)
	if err != nil {
		w.WriteHeader(400)
	}
}
