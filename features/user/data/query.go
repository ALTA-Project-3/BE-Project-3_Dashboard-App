package data

import (
	"errors"
	"project/dashboard/features/user"

	"gorm.io/gorm"
)

type dataUser struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &dataUser{
		db: db,
	}
}

func (repo *dataUser) SelectAll(page, token int) ([]user.Core, error) {

	perPage := 5
	offset := ((page - 1) * perPage)
	queryBuider := repo.db.Limit(perPage).Offset(offset)

	var data []User

	if page > 0 {
		tx1 := queryBuider.Find(&data).Order("fullname ASC")
		if tx1.Error != nil {
			return nil, tx1.Error
		}
	} else {
		tx2 := repo.db.Find(&data).Order("fullname ASC")
		if tx2.Error != nil {
			return nil, tx2.Error
		}
	}

	for _, v := range data {
		if v.ID == uint(token) {
			return toCoreList(data), nil
		}
	}

	return nil, errors.New("not have access")
}

func (repo *dataUser) UpdateData(data user.Core) int {

	newData := fromCore(data)

	tx := repo.db.Model(&User{}).Where("id = ? ", int(data.ID)).Updates(newData)
	if tx.Error != nil {
		return -1
	}

	return int(tx.RowsAffected)

}

func (repo *dataUser) DelData(id int) int {

	tx := repo.db.Unscoped().Where("id = ? ", id).Delete(&User{})
	if tx.Error != nil {
		return -1
	}

	return int(tx.RowsAffected)
}

func (repo *dataUser) InsertData(data user.Core) int {

	var dataInsert = fromCore(data)
	tx := repo.db.Create(&dataInsert)
	if tx.Error != nil {
		return -1
	}

	return int(tx.RowsAffected)
}

func (repo *dataUser) SelectProfile(id int) (user.Core, error) {

	var data user.Core
	tx := repo.db.First(&data, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	return data, nil
}
