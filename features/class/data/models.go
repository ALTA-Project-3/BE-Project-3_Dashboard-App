package data

import (
	"project/dashboard/features/class"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name      string
	StartDate string
	EndDate   string
	UserID    uint
	Mentee    []Mentee
}

type Mentee struct {
	gorm.Model
	Name        string
	Address     string
	Homeaddress string
	Email       string
	Gender      string
	Telegram    string
	Phone       uint64
	ClassID     uint
}

func CoreToModel(data class.CoreClass) Class {
	return Class{
		Name:      data.Name,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		UserID:    data.UserID,
	}
}

func ModelToCore(data Class) class.CoreClass {
	return class.CoreClass{
		ID:        data.ID,
		Name:      data.Name,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		UserID:    data.UserID,
	}
}

func ListToCore(data []Class) []class.CoreClass {
	var list []class.CoreClass
	for _, v := range data {
		list = append(list, ModelToCore(v))
	}

	return list
}

// type User struct {
// 	gorm.Model
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }
