package data

import (
	"project/dashboard/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string
}

func toCore(user User) login.Core {

	var core = login.Core{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	return core
}
