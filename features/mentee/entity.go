package mentee

type MenteeCore struct {
	ID            uint
	Name          string
	Address       string
	Homeaddress   string
	Email         string
	Gender        string
	Telegram      string
	Phone         string
	EmergencyData EmergencyCore
	Education     EducationCore
	ClassID       uint
	Class         string
	StatusMentee  string
}

type EmergencyCore struct {
	ID     uint
	Name   string
	Phone  string
	Status string
}

type EducationCore struct {
	ID       uint
	Type     string
	Major    string
	Graduate string
}

type Join struct {
	ID       uint
	Name     string
	Class    string
	Status   string
	Category string
	Gender   string
}

type DataInterface interface {
	InsertData(data MenteeCore, idToken int) int
	SelectData(page, idToken int) ([]Join, error)
}

type UsecaseInterface interface {
	PostData(data MenteeCore, idToken int) int
	GetData(page, idToken int) ([]Join, error)
}
