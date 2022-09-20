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

	perPage := 8
	offset := ((page - 1) * perPage)
	queryBuider := repo.db.Limit(perPage).Offset(offset)

	var data []User
	tx := queryBuider.Find(&data).Order("fullname ASC")
	if tx.Error != nil {
		return nil, tx.Error
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

	tx := repo.db.Model(&User{}).Where("id = ? ", data.ID).Updates(&newData)
	if tx.Error != nil {
		return -1
	}

	return 1

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
