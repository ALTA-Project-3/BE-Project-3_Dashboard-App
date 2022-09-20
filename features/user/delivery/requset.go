package delivery

import "project/dashboard/features/user"

type Request struct {
	ID       uint   `json:"id" form:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Team     string `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func (req *Request) toCoreReq() user.Core {
	return user.Core{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: req.Password,
		Team:     req.Team,
		Role:     req.Role,
		Status:   req.Status,
	}
}
