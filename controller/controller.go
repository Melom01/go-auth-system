package controller

import "sentinel/service"

type Controller interface {
	EmailController
}

type HTTPController struct {
	ServicesWrapper service.ServicesWrapper
}
