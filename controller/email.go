package controller

import (
	"net/http"
	"sentinel/model"
)

type EmailController interface {
	SendVerificationEmail(w http.ResponseWriter, r *http.Request)
}

type PostEmail struct {
	ReceiverEmail string `json:"receiverEmail" validate:"required"`
	Username      string `json:"username" validate:"required"`
}

func (ctrl *HTTPController) SendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	var postEmail PostEmail

	ctrl.DecodeBody(r, &postEmail)

	email := model.Email{
		ReceiverEmail: postEmail.ReceiverEmail,
		Username:      postEmail.Username,
	}

	err := ctrl.ServicesWrapper.SendVerificationEmail(email)
	if err != nil {
		w.WriteHeader(400)
	}
}
