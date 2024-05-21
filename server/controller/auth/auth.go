package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/savi2w/nano-go/service"
	"github.com/savi2w/nano-go/util/resutil"
	"github.com/savi2w/nano-go/validation"
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