package usecase

import (
	"errors"
	"project/dashboard/features/class"
)

type Service struct {
	do class.DataInterface
}

func New(data class.DataInterface) class.UsecaseInterface {
	return &Service{
		do: data,
	}
}

func (Service *Service) CreateClass(core class.CoreClass) (string, error) {
	msg, err := Service.do.InsertClass(core)
	return msg, err
}

func (Service *Service) GetAllClass(page int) ([]class.CoreClass, error) {
	listcore, err := Service.do.SelectAllClass(page)
	return listcore, err
}

func (Service *Service) EditAClass(core class.CoreClass, classid uint) (string, error) {
	userid := Service.do.GetClassUserid(classid)
	if userid == 0 {
		return "data tidak ditemukan", errors.New("error")
	} else if core.UserID != userid {
		return "Hanya Pemilik Class Yang Bisa Mengedit Class", errors.New("error")
	}

	msg, err := Service.do.UpdateAClass(core, classid)
	return msg, err
}

func (Service *Service) DeleteAClass(userid, classid uint) (string, error) {
	useridfromdata := Service.do.GetClassUserid(classid)
	if useridfromdata == 0 {
		return "data tidak ditemukan", errors.New("error")
	} else if useridfromdata != userid {
		return "Hanya Pemilik Class Yang Bisa Menghapus Class", errors.New("error")
	}

	msg, err := Service.do.DeleteAClass(classid)
	return msg, err
}

// func (Service *Service) GetToken(email, password string) (string, error) {
// 	str, err := Service.do.GetToken(email, password)
// 	return str, err
// }
