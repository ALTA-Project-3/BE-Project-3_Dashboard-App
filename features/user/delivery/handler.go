package delivery

import (
	"project/dashboard/features/user"
	"project/dashboard/middlewares"
	"project/dashboard/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, data user.UsecaseInterface) {
	handler := &userDelivery{
		userUsecase: data,
	}

	e.GET("/user", handler.GetAllUser, middlewares.JWTMiddleware())
	e.PUT("/user", handler.PutDataUser, middlewares.JWTMiddleware())

}

func (delivery *userDelivery) GetAllUser(c echo.Context) error {
	param := c.QueryParam("page")
	page, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("query param must be number"))
	}

	idToken := middlewares.ExtractToken(c)

	data, errGet := delivery.userUsecase.GetAll(page, idToken)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get all"))
	} else if len(data) == 0 {
		return c.JSON(200, helper.SuccessResponseHelper("user data is empty"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get all data", toResponList(data)))
}

func (delivery *userDelivery) PutDataUser(c echo.Context) error {

	idToken := middlewares.ExtractToken(c)

	var update Request
	err := c.Bind(&update)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	var updateData user.Core
	if update.Fullname != "" {
		updateData.Fullname = update.Fullname
	}
	if update.Email != "" {
		updateData.Email = update.Email
	}
	if update.Password != "" {
		updateData.Password = update.Password
	}
	if update.Role != "" {
		updateData.Role = update.Role
	}
	if update.Status != "" {
		updateData.Status = update.Status
	}
	if update.Team != "" {
		updateData.Team = update.Team
	}

	if idToken != 1 {
		if update.Role != "" || update.Status != "" || update.Team != "" {
			return c.JSON(400, helper.FailedResponseHelper("not have access edit role, status and team"))
		}
	}

	row := delivery.userUsecase.PutData(updateData)
	if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("failed to update data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success update data"))
}