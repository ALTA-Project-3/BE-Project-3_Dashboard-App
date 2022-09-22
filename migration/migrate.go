package migration

import (
	classModel "project/dashboard/features/class/data"
	logsModel "project/dashboard/features/logs/data"
	menteeModel "project/dashboard/features/mentee/data"
	userModel "project/dashboard/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&classModel.Class{})
	db.AutoMigrate(&menteeModel.Mentee{})
	db.AutoMigrate(&menteeModel.Emergency{})
	db.AutoMigrate(&menteeModel.Education{})
	db.AutoMigrate(&logsModel.Logs{})
}
