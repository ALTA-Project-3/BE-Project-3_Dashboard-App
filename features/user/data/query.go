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
		tx1 := queryBuider.Unscoped().Find(&data).Order("fullname ASC")
		if tx1.Error != nil {
			return nil, tx1.Error
		}
	} else {
		tx2 := repo.db.Unscoped().Find(&data).Order("fullname ASC")
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

	repo.db.Model(&User{}).Where("id = ? ", id).Update("status", "Deleted")
	tx := repo.db.Where("id = ? ", id).Delete(&User{})
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

func (repo *dataUser) SelectProfile(id int) (user.Core, user.DashBoard, error) {

	var data User
	tx := repo.db.First(&data, id)
	if tx.Error != nil {
		return user.Core{}, user.DashBoard{}, tx.Error
	}

	var count user.DashBoard
	Act := "Active"
	Plc := "Placement"
	Grd := "Graduate"
	txCount := repo.db.Model(&Mentee{}).Where("status_mentee = ?", Act).Count(&count.Active)
	if txCount.Error != nil {
		return user.Core{}, user.DashBoard{}, txCount.Error
	}
	txCount2 := repo.db.Model(&Mentee{}).Where("status_mentee = ?", Plc).Count(&count.Placement)
	if txCount2.Error != nil {
		return user.Core{}, user.DashBoard{}, txCount.Error
	}
	txCount3 := repo.db.Model(&Logs{}).Count(&count.FeedBack)
	if txCount3.Error != nil {
		return user.Core{}, user.DashBoard{}, txCount.Error
	}

	txGroup := repo.db.Raw("SELECT MONTH(classes.start_date) AS month, COUNT(mentees.id) AS count from classes left join mentees on mentees.class_id = classes.id where mentees.status_mentee = ? AND mentees.deleted_at = NULL group by MONTH(classes.start_date)", Act).Scan(&count.ActiveInMonth)
	if txGroup.Error != nil {
		return user.Core{}, user.DashBoard{}, txGroup.Error
	}
	txGroup2 := repo.db.Raw("SELECT MONTH(classes.end_date) AS month, COUNT(mentees.id) AS count from classes left join mentees on mentees.class_id = classes.id where mentees.status_mentee = ? AND mentees.deleted_at = NULL group by MONTH(classes.end_date)", Grd).Scan(&count.GraduateInMonth)
	if txGroup2.Error != nil {
		return user.Core{}, user.DashBoard{}, txGroup2.Error
	}

	return data.toCore(), count, nil
}
