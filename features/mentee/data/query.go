package data

import (
	"project/dashboard/features/mentee"

	"gorm.io/gorm"
)

type dataMentee struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.DataInterface {
	return &dataMentee{
		db: db,
	}
}

func (repo *dataMentee) InsertData(data mentee.MenteeCore, idToken int) int {

	var id int
	tx := repo.db.Model(&Class{}).Select("id").Where("name = ? ", data.Class).Scan(&id)
	if tx.Error != nil {
		return -1
	}

	var cek User
	repo.db.First(&cek, "id = ?", idToken)
	if cek.ID != uint(idToken) {
		return -1
	}

	data.ClassID = uint(id)
	data.StatusMentee = "Active"
	mentee, emergency, education := FromCore(data)
	txMente := repo.db.Create(&mentee)
	if txMente.Error != nil {
		return -1
	}
	txEmergency := repo.db.Create(&emergency)
	if txEmergency.Error != nil {
		return -1
	}
	txEducation := repo.db.Create(&education)
	if txEducation.Error != nil {
		return -1
	}

	return 1
}
