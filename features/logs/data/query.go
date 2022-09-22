package data

import (
	"project/dashboard/features/logs"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) logs.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) InsertLog(core logs.CoreLogs) (string, error) {
	model := CoreToLogs(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Terjadi Kesalahan Pada Database", tx.Error
	}
	status := Mentee{
		StatusMentee: model.Status,
	}
	tx = storage.query.Model(&status).Where("id = ?", model.MenteeID).Updates(status)

	return "Sukses Ke Database", nil
}

// model := models.CoreToModel(core)
// tx := storage.query.Model(&model).Where("id = ?", idproduct).Updates(model)
// if tx.Error != nil || tx.RowsAffected != 1 {
// 	return "Gagal Update", tx.Error
// }
