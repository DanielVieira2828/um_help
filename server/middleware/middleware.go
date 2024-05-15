package middleware

import (
	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/consts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMiddlewares(e *echo.Echo, cfg *config.Config) {
	e.Pre(middleware.BodyLimit(consts.BodyLimit))
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())

	if !cfg.InternalConfig.RunningLocal {
		// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// 	AllowOrigins: []string{"example.com"},
		// }))

		e.Use(middleware.ContextTimeout(consts.Timeout))
		e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: consts.Timeout,
		}))
	}
}
