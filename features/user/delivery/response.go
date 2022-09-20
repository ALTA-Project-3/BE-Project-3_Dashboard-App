package delivery

import "project/dashboard/features/user"

type Response struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Team     string `json:"team"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

func toRespon(data user.Core) Response {
	return Response{
		ID:       data.ID,
		Fullname: data.Fullname,
		Email:    data.Email,
		Password: data.Password,
		Team:     data.Team,
		Role:     data.Role,
		Status:   data.Status,
	}
}

func toResponList(data []user.Core) []Response {
	var dataAll []Response
	for key := range data {
		dataAll = append(dataAll, toRespon(data[key]))
	}

	return dataAll
}
