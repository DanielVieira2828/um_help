package middleware

import (
	"github.com/DanielVieirass/um_help/util/cryptoutil"
	"github.com/labstack/echo/v4"
)

func JWSMiddleware(cryptoutil *cryptoutil.Cryptoutil) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return next(ctx)
		}
	}
}
