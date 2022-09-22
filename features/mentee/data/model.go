package data

import (
	"project/dashboard/features/mentee"
	"time"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name          string
	Address       string
	HomeAddress   string
	Email         string
	Gender        string
	Telegram      string
	Phone         string
	EmergencyData Emergency
	Education     Education
	ClassID       uint
	Class         Class `gorm:"foreignKey:ClassID;references:ID"`
	StatusMentee  string
}

type Emergency struct {
	MenteeID uint `gorm:"primary_key"`
	Name     string
	Phone    string
	Status   string
}

type Education struct {
	MenteeID uint `gorm:"primary_key"`
	Type     string
	Major    string
	Graduate string
}

type Class struct {
	gorm.Model
	Name      string
	StartDate time.Time
	EndDate   time.Time
	UserID    uint
	Mentee    []Mentee `gorm:"foreignKey:ClassID"`
}

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Password string
	Team     string
	Role     string
	Status   string
}

func FromCore(data mentee.MenteeCore) (Mentee, Emergency, Education) {
	var dataMentee = Mentee{
		Name:         data.Name,
		Address:      data.Address,
		HomeAddress:  data.Homeaddress,
		Email:        data.Email,
		Gender:       data.Gender,
		Telegram:     data.Telegram,
		Phone:        data.Phone,
		ClassID:      data.ClassID,
		StatusMentee: data.StatusMentee,
	}

	var dataEmergency = Emergency{
		Name:   data.EmergencyData.Name,
		Phone:  data.EmergencyData.Phone,
		Status: data.EmergencyData.Status,
	}

	var dataEducation = Education{
		Type:     data.Education.Type,
		Major:    data.Education.Major,
		Graduate: data.Education.Graduate,
	}

	return dataMentee, dataEmergency, dataEducation
}
