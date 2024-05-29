package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/DanielVieirass/um_help/config"
	"github.com/DanielVieirass/um_help/consts"
	"github.com/DanielVieirass/um_help/util/cryptoutil"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

const TokenType = "Bearer"

func AcessJWSMiddleware(cfg *config.Config, logger *zerolog.Logger, crypt *cryptoutil.Cryptoutil) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			authorization := ctx.Request().Header.Get("Authorization")
			if len(authorization) == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
			}

			jws := strings.Split(authorization, " ")
			if len(jws) != 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			jwsType := jws[0]
			if jwsType != TokenType {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token type")
			}

			var claims cryptoutil.TokenClaims

			jwsValue := jws[1]
			err := crypt.VerifyJWS(jwsValue, &claims)
			if err != nil {
				logger.Warn().Err(err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			if claims.Issuer != cfg.InternalConfig.ServiceName {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid issuer")
			}

			if claims.ExpirationTime < time.Now().Unix() {
				return echo.NewHTTPError(http.StatusUnauthorized, "Expired token")
			}

			if claims.Type != consts.AcessTokenType {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token type ''"+claims.Type+"")
			}

			return next(ctx)
		}
	}
}

func RefreshJWSMiddleware(cfg *config.Config, logger *zerolog.Logger, crypt *cryptoutil.Cryptoutil) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			authorization := ctx.Request().Header.Get("Authorization")
			if len(authorization) == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
			}

			jws := strings.Split(authorization, " ")
			if len(jws) != 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			jwsType := jws[0]
			if jwsType != TokenType {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token type")
			}

			var claims cryptoutil.RefreshTokenClaims
			jwsValue := jws[1]
			err := crypt.VerifyJWS(jwsValue, &claims)
			if err != nil {
				logger.Warn().Err(err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			if claims.Type != consts.RefreshTokenType {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token type ''"+claims.Type+"")
			}

			ctx.Set("sign_id", claims.SignId)

			return next(ctx)
		}
	}
}
