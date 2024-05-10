package user

import (
	"net/http"

	"github.com/DanielVieirass/um_help/service"
	"github.com/DanielVieirass/um_help/util/resutil"
	"github.com/DanielVieirass/um_help/validation"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type Controller struct {
	logger  *zerolog.Logger
	resutil *resutil.ResUtil
	svc     *service.Service
}

func New(resutil *resutil.ResUtil, logger *zerolog.Logger, svc *service.Service) *Controller {
	return &Controller{
		logger:  logger,
		resutil: resutil,
		svc:     svc,
	}
}

func (ctrl *Controller) HandleNewUser(ctx echo.Context) error {
	req, err := validation.VerifyNewUserRequest(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusBadRequest))
	}

	if err := ctrl.svc.User.NewUser(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(ctrl.resutil.Wrap(nil, err, http.StatusInternalServerError))
	}

	return ctx.JSON(ctrl.resutil.Wrap(nil, nil, http.StatusCreated))
}
