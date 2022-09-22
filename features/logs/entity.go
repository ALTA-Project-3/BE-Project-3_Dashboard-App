package logs

import "time"

type CoreLogs struct {
	ID         uint
	UserID     uint
	MenteeID   uint
	Status     string
	Feedback   string
	Url        string
	Created_at time.Time
}

type DataInterface interface {
	InsertLog(core CoreLogs) (string, error)
}

type UsecaseInterface interface {
	AddLog(core CoreLogs) (string, error)
}
