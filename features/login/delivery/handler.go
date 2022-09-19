package delivery

import (
	"project/dashboard/features/login"
	"project/dashboard/utils/helper"

	"github.com/labstack/echo/v4"
)

type LoginDelivery struct {
	loginUsecase login.UsecaseInterface
}

func New(e *echo.Echo, usecase login.UsecaseInterface) {

	handler := LoginDelivery{
		loginUsecase: usecase,
	}

	e.POST("/login", handler.Login)

}

func (h *LoginDelivery) Login(c echo.Context) error {

	var req LoginRequest
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("wrong request"))
	}

	str := h.loginUsecase.LoginAuthorized(req.Email, req.Password)
	if str == "please input email and password" || str == "email not found" || str == "wrong password" {
		return c.JSON(400, helper.FailedResponseHelper(str))
	} else if str == "failed to created token" {
		return c.JSON(500, helper.FailedResponseHelper(str))
	} else {
		return c.JSON(200, helper.SuccessDataResponseHelper("Login Success", str))
	}

}
