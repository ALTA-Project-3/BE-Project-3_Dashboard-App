package delivery

import (
	"project/dashboard/features/mentee"
	"project/dashboard/middlewares"
	"project/dashboard/utils/helper"

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
