package delivery

import (
	"project/dashboard/features/user"
	"time"
)

type Response struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Team     string `json:"team"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

type ResponseDashboard struct {
	ID              uint               `json:"id"`
	Fullname        string             `json:"fullname"`
	Role            string             `json:"role"`
	Menteeactive    int                `json:"menteeActive"`
	MenteePlacement int                `json:"menteePlacement"`
	MenteeFeedback  int                `json:"menteeFeedback"`
	ActiveinMonth   []ResMonthActive   `json:"ActiveInMonth"`
	GraduateinMonth []ResMonthGraduate `json:"graduateInMonth"`
}

type ResMonthActive struct {
	Month string
	Count int
}

type ResMonthGraduate struct {
	Month string
	Count int
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

func toProfile(data user.Core, count user.DashBoard) ResponseDashboard {

	var Count1 []ResMonthActive
	for _, v := range count.ActiveInMonth {
		var m = time.Month(v.Month)
		Count1 = append(Count1, ResMonthActive{
			Month: m.String(),
			Count: int(v.Count),
		})
	}

	var Count2 []ResMonthGraduate
	for _, v1 := range count.GraduateInMonth {
		var m2 = time.Month(v1.Month)
		Count2 = append(Count2, ResMonthGraduate{
			Month: m2.String(),
			Count: int(v1.Count),
		})
	}

	return ResponseDashboard{
		ID:              data.ID,
		Fullname:        data.Fullname,
		Role:            data.Role,
		Menteeactive:    int(count.Active),
		MenteePlacement: int(count.Placement),
		MenteeFeedback:  int(count.FeedBack),
		ActiveinMonth:   Count1,
		GraduateinMonth: Count2,
	}
}
