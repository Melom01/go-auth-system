package controller

import "net/http"

type EmailController interface {
	SendVerificationEmail(w http.ResponseWriter, r *http.Request)
}

func (ctrl *HTTPController) SendVerificationEmail(_ http.ResponseWriter, _ *http.Request) {
	ctrl.ServicesWrapper.SendVerificationEmail()
}
