package delivery

import (
	"fmt"
	"net/http"
	"project/dashboard/config"
	"project/dashboard/features/logs"
	"project/dashboard/middlewares"
	"project/dashboard/utils/helper"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	From logs.UsecaseInterface
}

func New(e *echo.Echo, data logs.UsecaseInterface) {
	handler := &Delivery{
		From: data,
	}

	e.POST("/user/mentee/log/:id", handler.AddFeedback, middlewares.JWTMiddleware())
}

func (user *Delivery) AddFeedback(c echo.Context) error {
	userid, _ := middlewares.ExtractToken(c)

	menteeid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	var req Request
	erbind := c.Bind(&req)

	if erbind != nil || req.Status == "" || req.Feedback == "" {
		return c.JSON(400, helper.FailedResponseHelper("Terjadi Kesalahan"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("url")
	if fileErr != http.ErrMissingFile || fileErr == nil {
		ext, errs := helper.CheckFileType(fileInfo.Filename)
		if errs != nil {
			return c.JSON(400, helper.FailedResponseHelper("gagal membaca file exetension"))
		}

		if ext == "jpg" || ext == "png" || ext == "jpeg" {
			err_size := helper.CheckFileSize(fileInfo.Size, config.FileImageType)
			if err_size != nil {
				return c.JSON(400, helper.FailedResponseHelper("image size error"))
			}

			waktu := fmt.Sprintf("%v", time.Now())
			imageName := strconv.Itoa(userid) + "_" + c.Param("id") + waktu + "." + ext

			imageaddress, errupload := helper.UploadFileToS3(config.DirImage, imageName, config.FileImageType, fileData)
			if errupload != nil {
				return c.JSON(400, helper.FailedResponseHelper("failed to upload file"))
			}

			req.Url = imageaddress
		} else if ext == "pdf" {
			err_size := helper.CheckFileSize(fileInfo.Size, config.FilePdfType)
			if err_size != nil {
				return c.JSON(400, helper.FailedResponseHelper("pdf size error"))
			}

			waktu := fmt.Sprintf("%v", time.Now())
			pdfName := strconv.Itoa(userid) + "_" + c.Param("id") + waktu + "." + ext

			pdfaddress, errupload := helper.UploadFileToS3(config.DirPdf, pdfName, config.FilePdfType, fileData)
			if errupload != nil {
				return c.JSON(400, helper.FailedResponseHelper("failed to upload file"))
			}

			req.Url = pdfaddress
		} else {
			return c.JSON(400, helper.FailedResponseHelper("hanya menerima pdf, jpg, jpeg, png"))
		}
	}

	core := req.ReqToCore(uint(userid), uint(menteeid))
	msg, errr := user.From.AddLog(core)
	if errr != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}
