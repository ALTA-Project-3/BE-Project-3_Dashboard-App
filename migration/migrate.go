package migration

import (
	classModel "project/dashboard/features/class/data"
	userModel "project/dashboard/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&classModel.Class{})
}
