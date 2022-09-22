package delivery

import (
	"project/dashboard/features/class"
	"time"
)

type RequestClass struct {
	Name      string `json:"name" form:"name"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
}

// var layoutFormat = "January"
var layout2 = "2006-01-02"

func (data *RequestClass) ReqToCore(userid uint) class.CoreClass {
	start, _ := time.Parse(layout2, data.StartDate)
	end, _ := time.Parse(layout2, data.EndDate)
	return class.CoreClass{
		Name:      data.Name,
		StartDate: start,
		EndDate:   end,
		UserID:    userid,
	}
}
