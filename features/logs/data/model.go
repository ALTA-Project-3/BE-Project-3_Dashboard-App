package data

import (
	"project/dashboard/features/logs"

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
