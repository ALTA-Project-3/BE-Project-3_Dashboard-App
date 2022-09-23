package delivery

import (
	"project/dashboard/features/class"
	"project/dashboard/middlewares"
	"project/dashboard/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	From class.UsecaseInterface
}

func New(e *echo.Echo, data class.UsecaseInterface) {
	handler := &Delivery{
		From: data,
	}
	// e.GET("/logan", handler.Token)
	e.POST("/user/class", handler.AddClass, middlewares.JWTMiddleware())
	e.GET("/user/class", handler.GetAllClass, middlewares.JWTMiddleware())
	e.PUT("/user/class/:id", handler.EditClass, middlewares.JWTMiddleware())
	e.DELETE("/user/class/:id", handler.DeleteAClass, middlewares.JWTMiddleware())
}

func (user *Delivery) AddClass(c echo.Context) error {
	userid, _ := middlewares.ExtractToken(c)
	var req RequestClass
	errb := c.Bind(&req)
	if errb != nil {
		return c.JSON(400, helper.FailedResponseHelper("Wrong Data"))
	}
	if req.Name == "" {
		return c.JSON(400, helper.FailedResponseHelper("Nama kelas wajib di isi"))
	}
	core := req.ReqToCore(uint(userid))

	msg, err := user.From.CreateClass(core)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (user *Delivery) GetAllClass(c echo.Context) error {
	// userid, _ := middlewares.ExtractToken(c)
	param := c.QueryParam("page")
	if param == "" {
		param = "0"
	}
	page, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Harus Pakai Angka"))
	}

	listcore, err := user.From.GetAllClass(page)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Terjadi Kesalahan"))
	}

	listresponse := ListCoreResponse(listcore)
	return c.JSON(200, helper.SuccessDataResponseHelper("Sukses Mendapatkan data", listresponse))
}

func (user *Delivery) EditClass(c echo.Context) error {
	userid, _ := middlewares.ExtractToken(c)

	classid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	var request RequestClass
	errb := c.Bind(&request)
	if errb != nil {
		return c.JSON(400, helper.FailedResponseHelper("Bad Request"))
	}
	core := request.ReqToCore(uint(userid))
	msg, errs := user.From.EditAClass(core, uint(classid))
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (user *Delivery) DeleteAClass(c echo.Context) error {
	userid, _ := middlewares.ExtractToken(c)

	classid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	msg, errs := user.From.DeleteAClass(uint(userid), uint(classid))
	if errs != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

// func (user *Delivery) Token(c echo.Context) error {
// 	var login data.User
// 	errb := c.Bind(&login)
// 	if errb != nil {
// 		c.JSON(400, helper.FailedResponseHelper("Gagal"))
// 	}
// 	token, err := user.From.GetToken(login.Email, login.Password)
// 	if err != nil {
// 		return c.JSON(400, helper.FailedResponseHelper("gagal"))
// 	}

// 	return c.JSON(200, helper.SuccessDataResponseHelper("sukses", token))
// }
