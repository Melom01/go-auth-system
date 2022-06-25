package controller

import "sentinel/service"

type Controller interface {
}

type HTTPController struct {
	ServicesWrapper service.ServicesWrapper
}
