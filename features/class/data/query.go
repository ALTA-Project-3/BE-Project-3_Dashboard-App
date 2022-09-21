package data

import (
	"errors"
	"project/dashboard/features/class"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) class.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) InsertClass(core class.CoreClass) (string, error) {
	model := CoreToModel(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Terjadi Kesalahan", tx.Error
	}

	return "Success Add Class", nil

}

func (storage *Storage) SelectAllClass() ([]class.CoreClass, error) {
	var model []Class
	tx := storage.query.Find(&model)
	if tx.Error != nil {
		return []class.CoreClass{}, tx.Error
	}
	listcoreclass := ListToCore(model)

	return listcoreclass, nil
}

func (storage *Storage) GetClassUserid(classid uint) uint {
	var model Class
	tx := storage.query.Find(&model, "id = ?", classid)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return 0
	}

	return model.UserID
}

func (storage *Storage) UpdateAClass(core class.CoreClass, classid uint) (string, error) {
	update := CoreToModel(core)
	tx := storage.query.Model(&update).Where("id = ?", classid).Updates(update)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Gagal Update", tx.Error
	}

	return "Sukses Update", nil

}

func (storage *Storage) DeleteAClass(classid uint) (string, error) {
	tx := storage.query.Unscoped().Where("id = ? ", classid).Delete(&Class{})
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Gagal Menghapus Data", errors.New("error")
	}

	return "Sukses Menghapus Data", nil
}

// func (storage *Storage) GetToken(email, password string) (string, error) {
// 	var datauser User
// 	tx := storage.query.Where("email = ? and password = ?", email, password).First(&datauser)
// 	if tx.Error != nil {
// 		return "", tx.Error
// 	}

// 	token, err := middlewares.CreateToken(int(datauser.ID))
// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// }
