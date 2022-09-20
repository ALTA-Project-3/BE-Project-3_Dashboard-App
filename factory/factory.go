package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userData "project/dashboard/features/user/data"
	userDelivery "project/dashboard/features/user/delivery"
	userUsecase "project/dashboard/features/user/usecase"

	loginData "project/dashboard/features/login/data"
	loginDelivery "project/dashboard/features/login/delivery"
	loginUsecase "project/dashboard/features/login/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)
}
