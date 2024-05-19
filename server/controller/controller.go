package controller

import (
	"github.com/DanielVieirass/um_help/server/controller/health"
	"github.com/DanielVieirass/um_help/server/controller/user"

	"github.com/DanielVieirass/um_help/service"
	"github.com/DanielVieirass/um_help/util/resutil"
	"github.com/rs/zerolog"

	"github.com/DanielVieirass/um_help/server/controller/auth"
)

type Controller struct {
	AuthController   *auth.Controller
	HealthController *health.Controller
	UserController   *user.Controller
}

func New(svc *service.Service, logger *zerolog.Logger) *Controller {
	resutil := resutil.New(logger)

	return &Controller{
		AuthController:   auth.New(svc, resutil),
		HealthController: health.New(resutil),
		UserController:   user.New(svc, resutil),
	}
}
