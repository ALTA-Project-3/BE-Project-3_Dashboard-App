package delivery

import "project/dashboard/features/class"

type RequestClass struct {
	Name      string `json:"name" form:"name"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
}

func (data *RequestClass) ReqToCore(userid uint) class.CoreClass {
	return class.CoreClass{
		Name:      data.Name,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		UserID:    userid,
	}
}
