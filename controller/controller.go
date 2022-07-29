package controller

import "sentinel/service"

type Controller interface {
	EmailController
	UserController
}

type HTTPController struct {
	ServicesWrapper service.ServicesWrapper
}
