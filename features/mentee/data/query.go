package data

import (
	"errors"
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

func (repo *dataMentee) SelectData(page, idToken int) ([]mentee.Join, error) {

	var cek User
	repo.db.First(&cek, "id = ?", idToken)
	if cek.ID != uint(idToken) {
		return nil, errors.New("not access")
	}

	perPage := 5
	offset := ((page - 1) * perPage)
	queryBuider := repo.db.Limit(perPage).Offset(offset)

	var data []Mentee
	if page != 0 {
		tx := queryBuider.Find(&data)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if page == 0 {
		tx := repo.db.Find(&data)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	var res []mentee.Join
	for key := range data {
		var classData Class
		repo.db.First(&classData, "id = ? ", data[key].ClassID)
		var tipe Education
		repo.db.First(&tipe, "mentee_id = ? ", data[key].ID)
		res = append(res, mentee.Join{
			ID:       data[key].ID,
			Name:     data[key].Name,
			Class:    classData.Name,
			Status:   data[key].StatusMentee,
			Category: tipe.Type,
			Gender:   data[key].Gender,
		})
	}

	return res, nil

	// var data []mentee.Join
	// if category != "" || class != "" || status != "" {
	// 	if page != 0 {
	// 		// tx := queryBuider.Model(&Class{}).Select("mentees.id, mentees.name, classes.name, mentees.status_mentee, educations.type, mentees.gender").Joins("left join mentees on mentees.class_id = classes.id").Joins("left join educations on educations.mentee_id = mentees.id").Where(where).Scan(&data)
	// 		// if tx.Error != nil {
	// 		// 	return nil, tx.Error
	// 		// }
	// 		// return data, nil
	// 	} else if page == 0 {
	// 		// tx := repo.db.Model(&Class{}).Select("mentees.id, mentees.name, classes.name, mentees.status_mentee, educations.type, mentees.gender").Joins("inner join mentees on mentees.class_id = classes.id").Joins("inner join educations on educations.mentee_id = mentees.id").Where(where).Scan(&data)
	// 		// if tx.Error != nil {
	// 		// 	return nil, tx.Error
	// 		// }
	// 		// return data, nil
	// 	}
	// } else {
	// 	if page != 0 {
	// 		tx := queryBuider.Unscoped().Model(&Class{}).Select("classes.name, mentees.id, mentees.name, mentees.status_mentee, mentees.gender, educations.type").Joins("left join mentees on mentees.class_id = classes.id").Joins("right join educations on educations.mentee_id = mentees.id").Scan(&data)
	// 		if tx.Error != nil {
	// 			return nil, tx.Error
	// 		}
	// 		return data, nil
	// 	} else if page == 0 {
	// 		tx := repo.db.Model(&Class{}).Select("mentees.id, mentees.name, classes.name, mentees.status_mentee, educations.type, mentees.gender").Joins("inner join mentees on mentees.class_id = classes.id").Joins("inner join educations on educations.mentee_id = mentees.id").Scan(&data)
	// 		if tx.Error != nil {
	// 			return nil, tx.Error
	// 		}
	// 		return data, nil
	// 	}
	// }

}
