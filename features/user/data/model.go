package data

import (
	"project/dashboard/features/user"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Password string
	Team     string
	Role     string
	Status   string
	Class    []Class
}

type Mentee struct {
	StatusMentee string
}

type Class struct {
	gorm.Model
	Name      string
	StartDate time.Time
	EndDate   time.Time
	UserID    uint
}

type Logs struct {
	gorm.Model
	UserID   uint
	MenteeID uint
	Status   string
	Feedback string
	Url      string
}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:       data.ID,
		Fullname: data.Fullname,
		Email:    data.Email,
		Password: data.Password,
		Team:     data.Team,
		Role:     data.Role,
		Status:   data.Status,
	}
}

func toCoreList(data []User) []user.Core {
	var dataAll []user.Core
	for key := range data {
		dataAll = append(dataAll, data[key].toCore())
	}

	return dataAll
}

func fromCore(data user.Core) User {
	return User{
		Fullname: data.Fullname,
		Email:    data.Email,
		Password: data.Password,
		Team:     data.Team,
		Role:     data.Role,
		Status:   data.Status,
	}
}
