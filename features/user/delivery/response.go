package delivery

import "project/dashboard/features/user"

type Response struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Team     string `json:"team"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

type ResponseDashboard struct {
	ID               uint   `json:"id"`
	Fullname         string `json:"fullname"`
	Role             string `json:"role"`
	ActiveinMonth    int    `json:"registerInMonth"`
	PlacementinMonth int    `json:"placementInMonth"`
	GraduateinMonth  int    `json:"graduateInMonth"`
	Menteeactive     int    `json:"menteeActive"`
	MenteePlacement  int    `json:"menteePlacement"`
	MenteeFeedback   int    `json:"menteeFeedback"`
}

func toRespon(data user.Core) Response {
	return Response{
		ID:       data.ID,
		Fullname: data.Fullname,
		Email:    data.Email,
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

func toProfile(data user.Core) ResponseDashboard {
	return ResponseDashboard{
		ID:       data.ID,
		Fullname: data.Fullname,
		Role:     data.Role,
	}
}
