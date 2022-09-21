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

	e.GET("/users", handler.GetAllUser, middlewares.JWTMiddleware())
	e.PUT("/users", handler.PutDataUser, middlewares.JWTMiddleware())
	e.DELETE("/admin/:id", handler.DeleteDataUser, middlewares.JWTMiddleware())
	e.POST("/admin", handler.PostDataUser, middlewares.JWTMiddleware())
	e.GET("/user/profile", handler.GetProfileUser, middlewares.JWTMiddleware())

}

func (delivery *userDelivery) GetAllUser(c echo.Context) error {
	param := c.QueryParam("page")
	if param == "" {
		param = "0"
	}
	page, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("query param must be number"))
	}

	idToken, _ := middlewares.ExtractToken(c)

	data, errGet := delivery.userUsecase.GetAll(page, idToken)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get all"))
	} else if len(data) == 0 {
		return c.JSON(200, helper.SuccessResponseHelper("user data is empty"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get all data", toResponList(data)))
}

func (delivery *userDelivery) PutDataUser(c echo.Context) error {

	idToken, role := middlewares.ExtractToken(c)

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

	if role != "Admin" {
		if update.Role != "" || update.Status != "" || update.Team != "" || update.ID != uint(idToken) {
			return c.JSON(400, helper.FailedResponseHelper("not have access edit role, status and team"))
		}
	}

	updateData.ID = update.ID
	row := delivery.userUsecase.PutData(updateData)
	if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("failed to update data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success update data"))
}

func (delivery *userDelivery) DeleteDataUser(c echo.Context) error {

	_, role := middlewares.ExtractToken(c)
	if role != "Admin" {
		return c.JSON(400, helper.FailedResponseHelper("not have access"))
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row := delivery.userUsecase.DeleteData(id)
	if row == -1 || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("failed to delete data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete data"))
}

func (delivery *userDelivery) PostDataUser(c echo.Context) error {

	var data Request
	err := c.Bind(&data)
	if err != nil || data.ID != 0 {
		return c.JSON(400, helper.FailedResponseHelper("error bind and id cannot be filled"))
	}

	_, role := middlewares.ExtractToken(c)
	if role != "Admin" {
		return c.JSON(400, helper.FailedResponseHelper("not have access"))
	}

	row := delivery.userUsecase.PostData(data.toCoreReq())
	if row == -1 || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("failed post data"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *userDelivery) GetProfileUser(c echo.Context) error {

	idToken, _ := middlewares.ExtractToken(c)

	data, err := delivery.userUsecase.GetProfile(idToken)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed get profile"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get profile", toProfile(data)))
}
