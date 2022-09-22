package logs

import (
	"time"
)

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
	SelectAMenteeLogs(menteeid uint) ([]CoreLogs, string, error)
	SelectUserList(data []CoreLogs) ([]User, string, error)
	SelectAMenteePrivateData(menteeid uint) (Mentee, string, error)
	SelectAMenteeEducation(menteeid uint) (Education, string, error)
}

type UsecaseInterface interface {
	AddLog(core CoreLogs) (string, error)
	DetailMentee(menteeid uint) ([]CoreLogs, []User, Mentee, Education, string, error)
}

type Mentee struct {
	Name     string
	Email    string
	Telegram string
	Phone    string
	Class    string
}

// type Mentee struct {
// 	Name         string
// 	Email        string
// 	Telegram     string
// 	Phone        string
// 	StatusMentee string
// }

type Education struct {
	Major    string
	Graduate string
}

type User struct {
	Fullname string
}
