package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/savi2w/pupper/config"
	"github.com/savi2w/pupper/consts"
)

func SetMiddlewares(e *echo.Echo, cfg *config.Config) {
	// if !cfg.InternalConfig.RunningLocal {
	// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 		AllowOrigins: []string{"pupper-example.com"},
	// 	}))
	// }

	e.Pre(middleware.BodyLimit(consts.BodyLimit))
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())

	e.Use(middleware.ContextTimeout(consts.Timeout))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: consts.Timeout,
	}))
}
