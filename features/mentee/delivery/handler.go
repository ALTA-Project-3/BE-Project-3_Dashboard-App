package delivery

import (
	"project/dashboard/features/mentee"
	"project/dashboard/middlewares"
	"project/dashboard/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type menteeDelivery struct {
	menteeUsecase mentee.UsecaseInterface
}

func New(e *echo.Echo, data mentee.UsecaseInterface) {
	handler := &menteeDelivery{
		menteeUsecase: data,
	}

	e.POST("/user/mentee", handler.PostDataMentee, middlewares.JWTMiddleware())
	e.GET("/user/mentee", handler.GetMenteeList, middlewares.JWTMiddleware())
	e.PUT("/user/mentee/:id", handler.UpdateAMentee, middlewares.JWTMiddleware())

}

func (delivery *menteeDelivery) PostDataMentee(c echo.Context) error {

	var dataReq ReqCore
	err := c.Bind(&dataReq)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	idToken, _ := middlewares.ExtractToken(c)

	row := delivery.menteeUsecase.PostData(dataReq.toCore(), idToken)
	if row == -2 {
		return c.JSON(400, helper.FailedResponseHelper("data request harus diisi semua"))
	} else if row == -3 {
		return c.JSON(400, helper.FailedResponseHelper("gender must be Male or Female"))
	} else if row == -1 {
		return c.JSON(400, helper.FailedResponseHelper("failed insert data mentee"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success insert data mentee"))
}

func (delivery *menteeDelivery) GetMenteeList(c echo.Context) error {

	param := c.QueryParam("page")
	if param == "" {
		param = "0"
	}
	page, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param page must be number"))
	}
	class := c.QueryParam("class")
	category := c.QueryParam("category")
	status := c.QueryParam("status")

	idToken, _ := middlewares.ExtractToken(c)

	data, errGet := delivery.menteeUsecase.GetData(page, idToken)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("failed to get all mentee"))
	}

	dataRes := toResponList(data, class, category, status)
	return c.JSON(200, helper.SuccessDataResponseHelper("succes get all mentee", dataRes))
}

func (delivery *menteeDelivery) UpdateAMentee(c echo.Context) error {
	var req ReqCore
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("gagal bind"))
	}
	idmentee, errs := strconv.Atoi(c.Param("id"))
	if errs != nil {
		c.JSON(400, helper.FailedResponseHelper("Param harus nomor"))
	}

	msg, errr := delivery.menteeUsecase.UpdateData(req.toCore(), idmentee)
	if errr != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}
