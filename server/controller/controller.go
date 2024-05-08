package controller

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/pupper/server/controller/health"
	"github.com/savi2w/pupper/service"
	"github.com/savi2w/pupper/util/resutil"
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
