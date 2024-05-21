package auth

import (
	"net/http"

	"github.com/DanielVieirass/um_help/service"
	"github.com/DanielVieirass/um_help/util/resutil"
	"github.com/DanielVieirass/um_help/validation"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	resutil *resutil.ResUtil
	svc     *service.Service
}

func New(svc *service.Service, resutil *resutil.ResUtil) *Controller {
	return &Controller{
		resutil: resutil,
		svc:     svc,
	}
}

func (ctrl *Controller) HandleLogin(ctx echo.Context) error {
	req, err := validation.VerifyLoginRequest(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	resp, err := ctrl.svc.Auth.Login(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusUnauthorized))
	}

	return ctx.JSON(ctrl.resutil.Wrap(resp, nil, http.StatusOK))
}
