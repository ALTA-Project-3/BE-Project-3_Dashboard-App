package data

import (
	class "project/dashboard/features/class/data"
	"project/dashboard/features/logs"
	mentee "project/dashboard/features/mentee/data"
	user "project/dashboard/features/user/data"

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

func (storage *Storage) SelectAMenteeLogs(menteeid uint) ([]logs.CoreLogs, string, error) {
	var listlog []Logs
	tx := storage.query.Find(&listlog, "mentee_id = ?", menteeid)
	if tx.Error != nil {
		return nil, "Error Saat Mengambil Logs", tx.Error
	}

	listcore := ListLogToCore(listlog)
	return listcore, "Sukses Pengambilan Listlogs", nil
}
func (storage *Storage) SelectUserList(data []logs.CoreLogs) ([]logs.User, string, error) {
	var listuser []logs.User
	listlog := ListCoreToLogs(data)
	for _, v := range listlog {
		var user user.User
		tx := storage.query.Find(&user, "id = ?", v.UserID)
		if tx.Error != nil {
			return nil, "error pada pengambilan user", tx.Error
		}
		listuser = append(listuser, logs.User{
			Fullname: user.Fullname,
		})
	}

	return listuser, "Sukses Mengambil data User", nil
}
func (storage *Storage) SelectAMenteePrivateData(menteeid uint) (logs.Mentee, string, error) {
	var model mentee.Mentee
	tx := storage.query.Find(&model, "id = ?", menteeid)
	if tx.Error != nil {
		return logs.Mentee{}, "Gagal Mendapat data mentee", tx.Error
	}
	var class class.Class
	storage.query.Find(&class, "id = ?", model.ClassID)

	DataMentee := MentTologsMentee(model, class.Name)
	return DataMentee, "Sukses Mendapat Data Mentee", nil
}

func (storage *Storage) SelectAMenteeEducation(menteeid uint) (logs.Education, string, error) {
	var model mentee.Education
	tx := storage.query.Find(&model, "mentee_id = ?", menteeid)
	if tx.Error != nil {
		return logs.Education{}, "Gagal Mendapat data mentee", tx.Error
	}
	edumentee := MentToEdumentee(model)

	return edumentee, "Sukses mendapatkan Education mentee", nil
}

// model := models.CoreToModel(core)
// tx := storage.query.Model(&model).Where("id = ?", idproduct).Updates(model)
// if tx.Error != nil || tx.RowsAffected != 1 {
// 	return "Gagal Update", tx.Error
// }
