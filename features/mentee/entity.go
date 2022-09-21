package mentee

type MenteeCore struct {
	ID            uint
	Name          string
	Address       string
	Homeaddress   string
	Email         string
	Gender        string
	Telegram      string
	Phone         uint64
	EmergencyData EmergencyCore
	Education     EducationCore
	ClassID       uint
	Class         string
	StatusMentee  string
}

type EmergencyCore struct {
	ID     uint
	Name   string
	Phone  uint64
	Status string
}

type EducationCore struct {
	ID       uint
	Type     string
	Major    string
	Graduate string
}

type DataInterface interface {
	InsertData(data MenteeCore, idToken int) int
}

type UsecaseInterface interface {
	PostData(data MenteeCore, idToken int) int
}
