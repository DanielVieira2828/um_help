package controller

import (
	"github.com/DanielVieirass/um_help/server/controller/health"
	"github.com/DanielVieirass/um_help/service"
	"github.com/DanielVieirass/um_help/util/resutil"
	"github.com/rs/zerolog"
)

type Controller struct {
	HealthController *health.Controller
}

func New(svc *service.Service, logger *zerolog.Logger) *Controller {
	resutil := resutil.New(logger)

	return &Controller{
		HealthController: health.New(resutil),
	}
}
