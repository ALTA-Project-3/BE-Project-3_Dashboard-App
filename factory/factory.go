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

	classData "project/dashboard/features/class/data"
	classDelivery "project/dashboard/features/class/delivery"
	classUsecase "project/dashboard/features/class/usecase"

	menteeData "project/dashboard/features/mentee/data"
	menteeDelivery "project/dashboard/features/mentee/delivery"
	menteeUsecase "project/dashboard/features/mentee/usecase"

	logsData "project/dashboard/features/logs/data"
	logsDelivery "project/dashboard/features/logs/delivery"
	logsUsecase "project/dashboard/features/logs/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

	classDataFactory := classData.New(db)
	classUsecaseFactory := classUsecase.New(classDataFactory)
	classDelivery.New(e, classUsecaseFactory)

	menteeDataFactory := menteeData.New(db)
	menteeUsecaseFactory := menteeUsecase.New(menteeDataFactory)
	menteeDelivery.New(e, menteeUsecaseFactory)

	logsDataFactory := logsData.New(db)
	logsUsecaseFactory := logsUsecase.New(logsDataFactory)
	logsDelivery.New(e, logsUsecaseFactory)
}
