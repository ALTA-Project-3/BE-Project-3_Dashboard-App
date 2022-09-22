package delivery

import (
	"project/dashboard/features/logs"
	"time"
)

type ResponseIdMentee struct {
	Name     string
	Class    string
	Major    string
	Graduate string
	Phone    string
	Telegram string
	// Discord  string
	Email string
}

type LogResponse struct {
	User       string
	Status     string
	Feedback   string
	Url        string
	Created_at time.Time
}

func CoreToLogRes(core logs.CoreLogs, name string) LogResponse {
	return LogResponse{
		User:       name,
		Status:     core.Status,
		Feedback:   core.Feedback,
		Url:        core.Url,
		Created_at: core.Created_at,
	}
}

func CoreToLogReslist(core []logs.CoreLogs, name []logs.User) []LogResponse {
	var list []LogResponse
	for k, v := range core {
		list = append(list, CoreToLogRes(v, name[k].Fullname))
	}

	return list
}

func CoreToRes(mentee logs.Mentee, education logs.Education) ResponseIdMentee {
	return ResponseIdMentee{
		Name:     mentee.Name,
		Class:    mentee.Class,
		Major:    education.Major,
		Graduate: education.Graduate,
		Phone:    mentee.Phone,
		Telegram: mentee.Telegram,
		Email:    mentee.Email,
	}
}
