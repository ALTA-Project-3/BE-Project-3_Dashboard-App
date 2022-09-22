package data

import (
	"project/dashboard/features/logs"
	mentee "project/dashboard/features/mentee/data"

	"gorm.io/gorm"
)

type Logs struct {
	gorm.Model
	UserID   uint
	MenteeID uint
	Status   string
	Feedback string
	Url      string
}

type Mentee struct {
	StatusMentee string
}

func CoreToLogs(core logs.CoreLogs) Logs {
	return Logs{
		UserID:   core.UserID,
		MenteeID: core.MenteeID,
		Status:   core.Status,
		Feedback: core.Feedback,
		Url:      core.Url,
	}
}

func ListCoreToLogs(core []logs.CoreLogs) []Logs {
	var models []Logs
	for _, v := range core {
		models = append(models, CoreToLogs(v))
	}

	return models
}

func LogToCore(model Logs) logs.CoreLogs {
	return logs.CoreLogs{
		ID:         model.ID,
		UserID:     model.UserID,
		MenteeID:   model.MenteeID,
		Status:     model.Status,
		Feedback:   model.Feedback,
		Url:        model.Url,
		Created_at: model.CreatedAt,
	}
}

func ListLogToCore(model []Logs) []logs.CoreLogs {
	var listcore []logs.CoreLogs
	for _, v := range model {
		listcore = append(listcore, LogToCore(v))
	}

	return listcore
}

func MentTologsMentee(data mentee.Mentee, classname string) logs.Mentee {
	return logs.Mentee{
		Name:     data.Name,
		Email:    data.Email,
		Telegram: data.Telegram,
		Phone:    data.Phone,
		Class:    classname,
	}
}

func MentToEdumentee(data mentee.Education) logs.Education {
	return logs.Education{
		Major:    data.Major,
		Graduate: data.Graduate,
	}
}
