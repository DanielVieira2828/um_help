package router

import (
	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/server/controller"
	"github.com/labstack/echo/v4"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)

	root.POST("/user", ctrl.UserController.HandleNewUser)

	root.POST("/login", ctrl.AuthController.HandleLogin)
}
