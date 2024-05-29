package router

import (
	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/server/controller"
	"github.com/DanielVieirass/um_help/server/middleware"
	"github.com/DanielVieirass/um_help/util/cryptoutil"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func Register(cfg *config.Config, logger *zerolog.Logger, svr *echo.Echo, cryptoutil *cryptoutil.Cryptoutil, ctrl *controller.Controller) {
	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	root.POST("/user", ctrl.UserController.HandleNewUser)

	root.POST("/login", ctrl.AuthController.HandleLogin)

	root.POST("/refresh", ctrl.AuthController.HandleLogin, middleware.RefreshJWSMiddleware(cfg, logger, cryptoutil))
}
