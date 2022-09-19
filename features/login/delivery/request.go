package delivery

import "project/dashboard/features/login"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(data LoginRequest) login.Core {
	return login.Core{
		Email:    data.Email,
		Password: data.Password,
	}
}
